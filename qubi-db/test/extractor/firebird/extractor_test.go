package firebird_test

import (
	"database-schema-extractor/internal/extractor/firebird"
	"testing"
)

func TestFirebirdExtractor_GetDriverName(t *testing.T) {
	extractor := firebird.NewFirebirdExtractor()
	
	if extractor.GetDriverName() != "firebirdsql" {
		t.Errorf("Expected driver name 'firebirdsql', got %s", extractor.GetDriverName())
	}
}

func TestFirebirdExtractor_Connect_InvalidConnection(t *testing.T) {
	extractor := firebird.NewFirebirdExtractor()
	
	// Test with invalid connection string
	err := extractor.Connect("invalid_connection_string")
	if err == nil {
		t.Error("Expected error for invalid connection string")
	}
	
	// Ensure Close doesn't panic even if connection failed
	closeErr := extractor.Close()
	if closeErr != nil {
		t.Errorf("Close should not return error when connection was never established: %v", closeErr)
	}
}

func TestFirebirdExtractor_ExtractSchema_NoConnection(t *testing.T) {
	extractor := firebird.NewFirebirdExtractor()
	
	// Try to extract schema without connecting
	_, err := extractor.ExtractSchema()
	if err == nil {
		t.Error("Expected error when extracting schema without connection")
	}
	
	expectedMsg := "database connection not established"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

// Note: Real connection tests would require a running Firebird database
// These would be integration tests rather than unit tests
func TestFirebirdExtractor_Integration_Placeholder(t *testing.T) {
	t.Skip("Integration tests require running Firebird database")
	
	// This is where we would test with real connection strings like:
	// "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\ifarmacia.fdb"
	
	// extractor := firebird.NewFirebirdExtractor()
	// err := extractor.Connect("sysdba:masterkey@localhost:3050/test.fdb")
	// if err != nil {
	//     t.Fatalf("Failed to connect: %v", err)
	// }
	// defer extractor.Close()
	//
	// schema, err := extractor.ExtractSchema()
	// if err != nil {
	//     t.Fatalf("Failed to extract schema: %v", err)
	// }
	//
	// if schema == nil {
	//     t.Error("Schema should not be nil")
	// }
}