package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ConfigManager handles loading and parsing of configuration files
type ConfigManager struct {
	configPath string
}

// NewConfigManager creates a new ConfigManager with default config path
func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		configPath: "config.yaml",
	}
}

// NewConfigManagerWithPath creates a new ConfigManager with custom config path
func NewConfigManagerWithPath(path string) *ConfigManager {
	return &ConfigManager{
		configPath: path,
	}
}

// LoadConfig loads and parses the configuration file
func (cm *ConfigManager) LoadConfig() (*Config, error) {
	// Check if file exists
	if _, err := os.Stat(cm.configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w: %s", ErrConfigFileNotFound, cm.configPath)
	}

	// Read file content
	data, err := os.ReadFile(cm.configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidConfigFormat, err.Error())
	}

	// Validate configuration
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return &config, nil
}

// GetConfigPath returns the current config file path
func (cm *ConfigManager) GetConfigPath() string {
	return cm.configPath
}