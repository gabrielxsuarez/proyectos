package main

import (
	"database-schema-extractor/internal/config"
	"database-schema-extractor/internal/orchestrator"
	"database-schema-extractor/internal/output"
	"os"
	"path/filepath"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

// End-to-end integration test
func TestEndToEndIntegration(t *testing.T) {
	t.Skip("Integration test requires running Firebird database")

	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "integration_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test config file
	testConfig := map[string]string{
		"ifarmacia.fdb": "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\ifarmacia.fdb",
		"clientes.fdb":  "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\clientes.fdb",
	}

	configPath := filepath.Join(tempDir, "config.yaml")
	configData, err := yaml.Marshal(testConfig)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, configData, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Load configuration
	configManager := config.NewConfigManager()
	cfg, err := configManager.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Run extraction
	orch := orchestrator.NewExtractorOrchestrator(cfg, 2)
	results := orch.ExtractAll()

	if len(results) != 2 {
		t.Fatalf("Expected 2 results, got %d", len(results))
	}

	// Generate output files
	generator := output.NewYAMLGenerator()
	for _, result := range results {
		if err := generator.GenerateOutput(result); err != nil {
			t.Errorf("Failed to generate output for %s: %v", result.DatabaseName, err)
		}
	}

	// Verify output files were created
	expectedFiles := []string{"ifarmacia.fdb.yaml", "clientes.fdb.yaml"}
	for _, filename := range expectedFiles {
		filePath := filepath.Join(tempDir, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Expected output file %s was not created", filename)
			continue
		}

		// Verify file content is valid YAML
		data, err := os.ReadFile(filePath)
		if err != nil {
			t.Errorf("Failed to read output file %s: %v", filename, err)
			continue
		}

		var schema map[string]interface{}
		if err := yaml.Unmarshal(data, &schema); err != nil {
			t.Errorf("Output file %s contains invalid YAML: %v", filename, err)
			continue
		}

		// Verify basic schema structure
		if dbName, exists := schema["database_name"]; !exists || dbName == "" {
			t.Errorf("Output file %s missing or empty database_name", filename)
		}

		if tables, exists := schema["tables"]; exists {
			if tableList, ok := tables.([]interface{}); ok {
				t.Logf("File %s contains %d tables", filename, len(tableList))
			}
		}
	}
}

func TestEndToEndWithErrors(t *testing.T) {
	// Test end-to-end workflow with some invalid connections
	tempDir, err := os.MkdirTemp("", "integration_error_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create config with mix of valid and invalid connections
	testConfig := map[string]string{
		"invalid_db":    "invalid:connection@localhost:3050/nonexistent.fdb",
		"unsupported":   "mysql://user:pass@localhost:3306/db",
		"another_invalid": "bad_connection_string",
	}

	configPath := filepath.Join(tempDir, "config.yaml")
	configData, err := yaml.Marshal(testConfig)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, configData, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Load configuration
	configManager := config.NewConfigManager()
	cfg, err := configManager.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Run extraction
	orch := orchestrator.NewExtractorOrchestrator(cfg, 3)
	orch.SetTimeout(5 * time.Second) // Short timeout for faster test
	results := orch.ExtractAll()

	if len(results) != 3 {
		t.Fatalf("Expected 3 results, got %d", len(results))
	}

	// All should fail
	for _, result := range results {
		if result.IsSuccess() {
			t.Errorf("Expected extraction to fail for %s", result.DatabaseName)
		}
	}

	// Generate output files (should create error files)
	generator := output.NewYAMLGenerator()
	for _, result := range results {
		if err := generator.GenerateOutput(result); err != nil {
			t.Errorf("Failed to generate output for %s: %v", result.DatabaseName, err)
		}
	}

	// Verify error files were created
	expectedErrorFiles := []string{
		"invalid_db.error.yaml",
		"unsupported.error.yaml", 
		"another_invalid.error.yaml",
	}

	for _, filename := range expectedErrorFiles {
		filePath := filepath.Join(tempDir, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Expected error file %s was not created", filename)
			continue
		}

		// Verify error file content
		data, err := os.ReadFile(filePath)
		if err != nil {
			t.Errorf("Failed to read error file %s: %v", filename, err)
			continue
		}

		var errorInfo map[string]interface{}
		if err := yaml.Unmarshal(data, &errorInfo); err != nil {
			t.Errorf("Error file %s contains invalid YAML: %v", filename, err)
			continue
		}

		// Verify error file structure
		if dbName, exists := errorInfo["database_name"]; !exists || dbName == "" {
			t.Errorf("Error file %s missing database_name", filename)
		}

		if errorMsg, exists := errorInfo["error"]; !exists || errorMsg == "" {
			t.Errorf("Error file %s missing error message", filename)
		}

		if timestamp, exists := errorInfo["timestamp"]; !exists || timestamp == "" {
			t.Errorf("Error file %s missing timestamp", filename)
		}
	}
}

func TestConfigurationValidation(t *testing.T) {
	// Test various configuration scenarios
	tests := []struct {
		name        string
		config      map[string]string
		expectError bool
	}{
		{
			name: "valid config",
			config: map[string]string{
				"db1": "connection1",
				"db2": "connection2",
			},
			expectError: false,
		},
		{
			name:        "empty config",
			config:      map[string]string{},
			expectError: true,
		},
		{
			name: "empty connection string",
			config: map[string]string{
				"db1": "",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "config_test")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tempDir)

			configPath := filepath.Join(tempDir, "config.yaml")
			configData, err := yaml.Marshal(tt.config)
			if err != nil {
				t.Fatalf("Failed to marshal config: %v", err)
			}

			if err := os.WriteFile(configPath, configData, 0644); err != nil {
				t.Fatalf("Failed to write config file: %v", err)
			}

			configManager := config.NewConfigManagerWithPath(configPath)
			_, err = configManager.LoadConfig()

			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}