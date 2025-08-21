package output_test

import (
	"database-schema-extractor/internal/output"
	"database-schema-extractor/internal/schema"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestYAMLGenerator_GenerateOutput_Success(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "yaml_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test schema
	testSchema := schema.NewDatabaseSchema("test_db")
	testSchema.Tables = []schema.Table{
		{
			Name: "users",
			Columns: []schema.Column{
				{Name: "id", Type: "INTEGER", Nullable: false},
				{Name: "name", Type: "VARCHAR", Nullable: true},
			},
			PrimaryKey: []string{"id"},
		},
	}

	// Create successful extraction result
	result := output.NewExtractionResult("test_db", testSchema)

	// Generate output
	generator := output.NewYAMLGeneratorWithDir(tempDir)
	err = generator.GenerateOutput(result)
	if err != nil {
		t.Fatalf("Failed to generate output: %v", err)
	}

	// Verify file was created
	expectedFile := filepath.Join(tempDir, "test_db.yaml")
	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatalf("Expected output file %s was not created", expectedFile)
	}

	// Verify file content
	data, err := os.ReadFile(expectedFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	var parsedSchema schema.DatabaseSchema
	if err := yaml.Unmarshal(data, &parsedSchema); err != nil {
		t.Fatalf("Failed to parse generated YAML: %v", err)
	}

	if parsedSchema.DatabaseName != "test_db" {
		t.Errorf("Expected database name 'test_db', got %s", parsedSchema.DatabaseName)
	}

	if len(parsedSchema.Tables) != 1 {
		t.Errorf("Expected 1 table, got %d", len(parsedSchema.Tables))
	}
}

func TestYAMLGenerator_GenerateOutput_Error(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "yaml_error_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create error extraction result
	result := output.NewExtractionError("failed_db", errors.New("connection failed"))

	// Generate output
	generator := output.NewYAMLGeneratorWithDir(tempDir)
	err = generator.GenerateOutput(result)
	if err != nil {
		t.Fatalf("Failed to generate error output: %v", err)
	}

	// Verify error file was created
	expectedFile := filepath.Join(tempDir, "failed_db.error.yaml")
	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatalf("Expected error file %s was not created", expectedFile)
	}

	// Verify error file content
	data, err := os.ReadFile(expectedFile)
	if err != nil {
		t.Fatalf("Failed to read error file: %v", err)
	}

	var errorInfo map[string]interface{}
	if err := yaml.Unmarshal(data, &errorInfo); err != nil {
		t.Fatalf("Failed to parse error YAML: %v", err)
	}

	if errorInfo["database_name"] != "failed_db" {
		t.Errorf("Expected database name 'failed_db', got %v", errorInfo["database_name"])
	}

	if errorInfo["error"] != "connection failed" {
		t.Errorf("Expected error 'connection failed', got %v", errorInfo["error"])
	}
}

func TestGenerateSchemaYAML(t *testing.T) {
	testSchema := schema.NewDatabaseSchema("test_db")
	testSchema.Tables = []schema.Table{
		{
			Name: "test_table",
			Columns: []schema.Column{
				{Name: "id", Type: "INTEGER", Nullable: false},
			},
		},
	}

	yamlData, err := output.GenerateSchemaYAML(testSchema)
	if err != nil {
		t.Fatalf("Failed to generate YAML: %v", err)
	}

	if len(yamlData) == 0 {
		t.Error("Generated YAML data is empty")
	}

	// Verify it can be parsed back
	var parsedSchema schema.DatabaseSchema
	if err := yaml.Unmarshal(yamlData, &parsedSchema); err != nil {
		t.Fatalf("Generated YAML is not valid: %v", err)
	}

	if parsedSchema.DatabaseName != "test_db" {
		t.Errorf("Expected database name 'test_db', got %s", parsedSchema.DatabaseName)
	}
}