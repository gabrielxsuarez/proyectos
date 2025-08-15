package firebird_test

import (
	"database-schema-extractor/internal/extractor/firebird"
	"testing"
)

func TestFirebirdExtractorFactory_SupportsConnectionString(t *testing.T) {
	factory := firebird.NewFirebirdExtractorFactory()

	tests := []struct {
		name             string
		connectionString string
		expected         bool
	}{
		{
			name:             "Firebird .fdb file",
			connectionString: "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\ifarmacia.fdb",
			expected:         true,
		},
		{
			name:             "Firebird .gdb file",
			connectionString: "sysdba:masterkey@localhost:3050/database.gdb",
			expected:         true,
		},
		{
			name:             "Firebird port 3050",
			connectionString: "sysdba:masterkey@localhost:3050/database",
			expected:         true,
		},
		{
			name:             "SQL Server connection",
			connectionString: "server=localhost;database=test;user=sa;password=pass",
			expected:         false,
		},
		{
			name:             "MySQL connection",
			connectionString: "user:password@tcp(localhost:3306)/database",
			expected:         false,
		},
		{
			name:             "PostgreSQL connection",
			connectionString: "postgres://user:password@localhost:5432/database",
			expected:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := factory.SupportsConnectionString(tt.connectionString)
			if result != tt.expected {
				t.Errorf("SupportsConnectionString(%s) = %v, want %v", 
					tt.connectionString, result, tt.expected)
			}
		})
	}
}

func TestFirebirdExtractorFactory_CreateExtractor(t *testing.T) {
	factory := firebird.NewFirebirdExtractorFactory()
	
	extractor, err := factory.CreateExtractor("test_connection_string")
	if err != nil {
		t.Fatalf("CreateExtractor failed: %v", err)
	}
	
	if extractor == nil {
		t.Error("CreateExtractor returned nil extractor")
	}
	
	if extractor.GetDriverName() != "firebirdsql" {
		t.Errorf("Expected driver name 'firebirdsql', got %s", extractor.GetDriverName())
	}
}