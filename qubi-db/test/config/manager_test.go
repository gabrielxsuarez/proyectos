package config_test

import (
	"database-schema-extractor/internal/config"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConfigManager_LoadConfig(t *testing.T) {
	// Create temporary directory for test files
	tempDir, err := os.MkdirTemp("", "config_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name        string
		configData  string
		wantErr     bool
		expectedErr error
	}{
		{
			name: "valid config",
			configData: `ifarmacia.fdb: "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\ifarmacia.fdb"
clientes.fdb: "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\clientes.fdb"`,
			wantErr: false,
		},
		{
			name:        "invalid yaml",
			configData:  `invalid: yaml: content: [`,
			wantErr:     true,
			expectedErr: config.ErrInvalidConfigFormat,
		},
		{
			name:        "empty config",
			configData:  ``,
			wantErr:     true,
			expectedErr: config.ErrNoDatabasesConfigured,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test config file
			configPath := filepath.Join(tempDir, "test_config.yaml")
			err := os.WriteFile(configPath, []byte(tt.configData), 0644)
			if err != nil {
				t.Fatalf("Failed to write test config: %v", err)
			}

			// Test config loading
			manager := config.NewConfigManagerWithPath(configPath)
			cfg, err := manager.LoadConfig()

			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigManager.LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if tt.expectedErr != nil && !isErrorType(err, tt.expectedErr) {
					t.Errorf("ConfigManager.LoadConfig() error = %v, want error type %v", err, tt.expectedErr)
				}
				return
			}

			// Validate successful load
			if cfg == nil {
				t.Error("ConfigManager.LoadConfig() returned nil config")
				return
			}

			if len(cfg.Databases) == 0 {
				t.Error("ConfigManager.LoadConfig() returned empty databases")
			}
		})
	}
}

func TestConfigManager_LoadConfig_FileNotFound(t *testing.T) {
	manager := config.NewConfigManagerWithPath("nonexistent.yaml")
	_, err := manager.LoadConfig()

	if err == nil {
		t.Error("Expected error for nonexistent file")
		return
	}

	if !isErrorType(err, config.ErrConfigFileNotFound) {
		t.Errorf("Expected ErrConfigFileNotFound, got %v", err)
	}
}

func TestConfigManager_GetConfigPath(t *testing.T) {
	manager := config.NewConfigManager()
	if manager.GetConfigPath() != "config.yaml" {
		t.Errorf("Expected default path 'config.yaml', got %s", manager.GetConfigPath())
	}

	customPath := "custom/path/config.yaml"
	manager = config.NewConfigManagerWithPath(customPath)
	if manager.GetConfigPath() != customPath {
		t.Errorf("Expected custom path %s, got %s", customPath, manager.GetConfigPath())
	}
}

// Helper function to check if error is of specific type
func isErrorType(err, target error) bool {
	if err == nil || target == nil {
		return err == target
	}
	// Check if the error contains the target error message
	return strings.Contains(err.Error(), target.Error())
}