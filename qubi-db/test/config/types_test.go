package config_test

import (
	"database-schema-extractor/internal/config"
	"testing"
)

func TestConfig_GetConnections(t *testing.T) {
	cfg := &config.Config{
		Databases: map[string]string{
			"db1": "conn1",
			"db2": "conn2",
		},
	}

	connections := cfg.GetConnections()

	if len(connections) != 2 {
		t.Errorf("Expected 2 connections, got %d", len(connections))
	}

	// Check that all connections are present
	found := make(map[string]bool)
	for _, conn := range connections {
		found[conn.Name] = true
		if conn.ConnectionString == "" {
			t.Errorf("Connection string for %s is empty", conn.Name)
		}
	}

	if !found["db1"] || !found["db2"] {
		t.Error("Not all expected connections were found")
	}
}

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *config.Config
		wantErr bool
		errType error
	}{
		{
			name: "valid config",
			config: &config.Config{
				Databases: map[string]string{
					"db1": "connection1",
					"db2": "connection2",
				},
			},
			wantErr: false,
		},
		{
			name: "no databases",
			config: &config.Config{
				Databases: map[string]string{},
			},
			wantErr: true,
			errType: config.ErrNoDatabasesConfigured,
		},
		{
			name: "empty database name",
			config: &config.Config{
				Databases: map[string]string{
					"":    "connection1",
					"db2": "connection2",
				},
			},
			wantErr: true,
			errType: config.ErrEmptyDatabaseName,
		},
		{
			name: "empty connection string",
			config: &config.Config{
				Databases: map[string]string{
					"db1": "",
					"db2": "connection2",
				},
			},
			wantErr: true,
			errType: config.ErrEmptyConnectionString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errType != nil && err != tt.errType {
				t.Errorf("Config.Validate() error = %v, want %v", err, tt.errType)
			}
		})
	}
}