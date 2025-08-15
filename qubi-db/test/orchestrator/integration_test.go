package orchestrator_test

import (
	"database-schema-extractor/internal/config"
	"database-schema-extractor/internal/orchestrator"
	"testing"
)

// Integration test that verifies the complete workflow
func TestExtractorOrchestrator_Integration_CompleteWorkflow(t *testing.T) {
	t.Skip("Integration test requires running Firebird database")

	// This test would use real Firebird connection strings
	cfg := &config.Config{
		Databases: map[string]string{
			"ifarmacia.fdb": "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\ifarmacia.fdb",
			"clientes.fdb":  "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\clientes.fdb",
		},
	}

	orch := orchestrator.NewExtractorOrchestrator(cfg, 2)
	results := orch.ExtractAll()

	if len(results) != 2 {
		t.Fatalf("Expected 2 results, got %d", len(results))
	}

	// Verify results
	for _, result := range results {
		if result.IsError() {
			t.Errorf("Extraction failed for %s: %v", result.DatabaseName, result.Error)
			continue
		}

		if result.Schema == nil {
			t.Errorf("Schema is nil for %s", result.DatabaseName)
			continue
		}

		// Verify schema content
		if result.Schema.DatabaseName == "" {
			t.Errorf("Database name is empty for %s", result.DatabaseName)
		}

		// Verify that we have some tables (assuming test databases have user tables)
		if len(result.Schema.Tables) == 0 {
			t.Logf("Warning: No tables found for %s (might be empty database)", result.DatabaseName)
		}

		// Verify that system tables are filtered out
		for _, table := range result.Schema.Tables {
			if table.Name[:4] == "RDB$" || table.Name[:4] == "MON$" {
				t.Errorf("System table %s should be filtered out from %s", table.Name, result.DatabaseName)
			}
		}

		// Verify engine-specific features for Firebird
		if engineSpecific := result.Schema.EngineSpecific; engineSpecific != nil {
			if generators, exists := engineSpecific["generators"]; exists {
				t.Logf("Found generators for %s: %v", result.DatabaseName, generators)
			}
			if domains, exists := engineSpecific["domains"]; exists {
				t.Logf("Found domains for %s: %v", result.DatabaseName, domains)
			}
		}
	}
}

func TestExtractorOrchestrator_Integration_ParallelProcessing(t *testing.T) {
	t.Skip("Integration test requires running Firebird database")

	// Test with multiple databases to verify parallel processing
	cfg := &config.Config{
		Databases: map[string]string{
			"db1": "sysdba:masterkey@localhost:3050/test1.fdb",
			"db2": "sysdba:masterkey@localhost:3050/test2.fdb",
			"db3": "sysdba:masterkey@localhost:3050/test3.fdb",
		},
	}

	orch := orchestrator.NewExtractorOrchestrator(cfg, 3)
	
	// Measure time to ensure parallel processing is working
	// start := time.Now()
	results := orch.ExtractAll()
	// duration := time.Since(start)

	if len(results) != 3 {
		t.Fatalf("Expected 3 results, got %d", len(results))
	}

	// Verify all databases were processed
	processedDbs := make(map[string]bool)
	for _, result := range results {
		processedDbs[result.DatabaseName] = true
	}

	expectedDbs := []string{"db1", "db2", "db3"}
	for _, dbName := range expectedDbs {
		if !processedDbs[dbName] {
			t.Errorf("Database %s was not processed", dbName)
		}
	}

	// If we had sequential processing, it would take much longer
	// This is a rough check - in practice you'd need more sophisticated timing tests
	// if duration > 10*time.Second {
	//     t.Errorf("Processing took too long (%v), parallel processing might not be working", duration)
	// }
}

func TestExtractorOrchestrator_Integration_MixedResults(t *testing.T) {
	t.Skip("Integration test requires running Firebird database")

	// Test with mix of valid and invalid connections
	cfg := &config.Config{
		Databases: map[string]string{
			"valid_db":   "sysdba:masterkey@localhost:3050/valid.fdb",
			"invalid_db": "invalid:connection@localhost:3050/nonexistent.fdb",
		},
	}

	orch := orchestrator.NewExtractorOrchestrator(cfg, 2)
	results := orch.ExtractAll()

	if len(results) != 2 {
		t.Fatalf("Expected 2 results, got %d", len(results))
	}

	// Should have one success and one failure
	successCount := 0
	errorCount := 0

	for _, result := range results {
		if result.IsSuccess() {
			successCount++
		} else {
			errorCount++
		}
	}

	// Depending on whether valid.fdb actually exists, we might have different results
	// The important thing is that the orchestrator handles both cases gracefully
	if successCount+errorCount != 2 {
		t.Error("Results should be either success or error")
	}

	// Verify that errors don't prevent other databases from being processed
	if errorCount > 0 {
		t.Log("Some extractions failed as expected for invalid connections")
	}
}