package output

import (
	"database-schema-extractor/internal/schema"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// YAMLGenerator handles generation of YAML output files
type YAMLGenerator struct {
	outputDir string
}

// NewYAMLGenerator creates a new YAMLGenerator with default output directory
func NewYAMLGenerator() *YAMLGenerator {
	return &YAMLGenerator{
		outputDir: ".",
	}
}

// NewYAMLGeneratorWithDir creates a new YAMLGenerator with custom output directory
func NewYAMLGeneratorWithDir(dir string) *YAMLGenerator {
	return &YAMLGenerator{
		outputDir: dir,
	}
}

// GenerateOutput generates YAML output file for extraction result
func (yg *YAMLGenerator) GenerateOutput(result ExtractionResult) error {
	if result.Error != nil {
		return yg.generateErrorFile(result)
	}

	return yg.generateSchemaFile(result)
}

// generateSchemaFile generates the main schema YAML file
func (yg *YAMLGenerator) generateSchemaFile(result ExtractionResult) error {
	filename := fmt.Sprintf("%s.yaml", result.DatabaseName)
	filepath := filepath.Join(yg.outputDir, filename)

	// Convert schema to YAML
	yamlData, err := yaml.Marshal(result.Schema)
	if err != nil {
		return fmt.Errorf("failed to marshal schema to YAML: %w", err)
	}

	// Write to file
	if err := os.WriteFile(filepath, yamlData, 0644); err != nil {
		return fmt.Errorf("failed to write YAML file %s: %w", filepath, err)
	}

	return nil
}

// generateErrorFile generates an error file for failed extractions
func (yg *YAMLGenerator) generateErrorFile(result ExtractionResult) error {
	filename := fmt.Sprintf("%s.error.yaml", result.DatabaseName)
	filepath := filepath.Join(yg.outputDir, filename)

	errorInfo := map[string]interface{}{
		"database_name": result.DatabaseName,
		"error":         result.Error.Error(),
		"timestamp":     result.Timestamp,
	}

	yamlData, err := yaml.Marshal(errorInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal error to YAML: %w", err)
	}

	if err := os.WriteFile(filepath, yamlData, 0644); err != nil {
		return fmt.Errorf("failed to write error file %s: %w", filepath, err)
	}

	return nil
}

// SetOutputDir sets the output directory for generated files
func (yg *YAMLGenerator) SetOutputDir(dir string) {
	yg.outputDir = dir
}

// GetOutputDir returns the current output directory
func (yg *YAMLGenerator) GetOutputDir() string {
	return yg.outputDir
}

// GenerateSchemaYAML converts a DatabaseSchema to YAML bytes
func GenerateSchemaYAML(schema *schema.DatabaseSchema) ([]byte, error) {
	return yaml.Marshal(schema)
}