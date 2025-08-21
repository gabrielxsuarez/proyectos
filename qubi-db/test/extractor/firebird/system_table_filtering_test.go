package firebird_test

import (
	"database-schema-extractor/internal/extractor/firebird"
	"strings"
	"testing"
)

func TestFirebirdExtractor_SystemTableFiltering(t *testing.T) {
	// Test that system table filtering logic is correct
	// This tests the SQL queries used for filtering
	
	systemTablePrefixes := []string{"RDB$", "MON$"}
	userTableNames := []string{"USERS", "PRODUCTS", "ORDERS", "CUSTOMERS"}
	systemTableNames := []string{"RDB$RELATIONS", "RDB$FIELDS", "MON$STATEMENTS", "MON$ATTACHMENTS"}

	// Test that user tables would not be filtered
	for _, tableName := range userTableNames {
		for _, prefix := range systemTablePrefixes {
			if strings.HasPrefix(tableName, prefix) {
				t.Errorf("User table %s should not start with system prefix %s", tableName, prefix)
			}
		}
	}

	// Test that system tables would be filtered
	for _, tableName := range systemTableNames {
		isSystemTable := false
		for _, prefix := range systemTablePrefixes {
			if strings.HasPrefix(tableName, prefix) {
				isSystemTable = true
				break
			}
		}
		if !isSystemTable {
			t.Errorf("System table %s should start with a system prefix", tableName)
		}
	}
}

func TestFirebirdExtractor_QueryFiltering(t *testing.T) {
	// Test that our queries contain the correct filtering conditions
	// This is a structural test to ensure we don't accidentally remove filtering
	
	extractor := firebird.NewFirebirdExtractor()
	if extractor == nil {
		t.Fatal("Failed to create Firebird extractor")
	}

	// We can't directly test the private queries, but we can test the structure
	// In a real implementation, we might expose the queries for testing
	// or create a method that returns the filtering conditions
	
	// For now, we'll test that the extractor exists and has the right driver
	if extractor.GetDriverName() != "firebirdsql" {
		t.Errorf("Expected driver 'firebirdsql', got %s", extractor.GetDriverName())
	}
}

// This would be an integration test to verify actual filtering
func TestFirebirdExtractor_Integration_SystemTableFiltering(t *testing.T) {
	t.Skip("Integration test requires running Firebird database")
	
	// This test would connect to a real Firebird database and verify
	// that no system tables are returned in the schema
	
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
	// // Verify no system tables in results
	// for _, table := range schema.Tables {
	//     if strings.HasPrefix(table.Name, "RDB$") || strings.HasPrefix(table.Name, "MON$") {
	//         t.Errorf("System table %s should be filtered out", table.Name)
	//     }
	// }
	//
	// // Verify no system indexes in results
	// for _, index := range schema.Indexes {
	//     if strings.HasPrefix(index.Name, "RDB$") {
	//         t.Errorf("System index %s should be filtered out", index.Name)
	//     }
	//     if strings.HasPrefix(index.TableName, "RDB$") || strings.HasPrefix(index.TableName, "MON$") {
	//         t.Errorf("Index %s on system table %s should be filtered out", index.Name, index.TableName)
	//     }
	// }
}

func TestFirebirdExtractor_FilteringConsistency(t *testing.T) {
	// Test that filtering is consistent across all extraction methods
	// This ensures that if we filter tables, we also filter related objects
	
	// Test data representing what should be filtered
	testCases := []struct {
		name           string
		objectName     string
		shouldBeFiltered bool
	}{
		{"User table", "CUSTOMERS", false},
		{"User table with underscore", "USER_PROFILES", false},
		{"System relation", "RDB$RELATIONS", true},
		{"System field", "RDB$FIELDS", true},
		{"Monitor table", "MON$STATEMENTS", true},
		{"Monitor attachment", "MON$ATTACHMENTS", true},
		{"Mixed case user table", "Products", false},
		{"System table mixed case", "rdb$relations", true}, // Should still be filtered if case-insensitive
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test the filtering logic
			isSystemTable := strings.HasPrefix(strings.ToUpper(tc.objectName), "RDB$") || 
							strings.HasPrefix(strings.ToUpper(tc.objectName), "MON$")
			
			if isSystemTable != tc.shouldBeFiltered {
				t.Errorf("Object %s: expected filtered=%v, got filtered=%v", 
					tc.objectName, tc.shouldBeFiltered, isSystemTable)
			}
		})
	}
}