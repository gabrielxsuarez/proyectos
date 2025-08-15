package orchestrator_test

import (
	"database-schema-extractor/internal/config"
	"database-schema-extractor/internal/orchestrator"
	"testing"
	"time"
)

func TestNewExtractorOrchestrator(t *testing.T) {
	cfg := &config.Config{
		Databases: map[string]string{
			"test1": "connection1",
			"test2": "connection2",
		},
	}

	orch := orchestrator.NewExtractorOrchestrator(cfg, 3)
	if orch == nil {
		t.Fatal("NewExtractorOrchestrator returned nil")
	}

	// Test that we can set timeout
	orch.SetTimeout(60 * time.Second)
}

func TestExtractorOrchestrator_ExtractAll_EmptyConfig(t *testing.T) {
	cfg := &config.Config{
		Databases: map[string]string{},
	}

	orch := orchestrator.NewExtractorOrchestrator(cfg, 1)
	results := orch.ExtractAll()

	if len(results) != 0 {
		t.Errorf("Expected 0 results for empty config, got %d", len(results))
	}
}

func TestExtractorOrchestrator_ExtractAll_UnsupportedDatabase(t *testing.T) {
	cfg := &config.Config{
		Databases: map[string]string{
			"unsupported": "mysql://user:pass@localhost:3306/db",
		},
	}

	orch := orchestrator.NewExtractorOrchestrator(cfg, 1)
	results := orch.ExtractAll()

	if len(results) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(results))
	}

	result := results[0]
	if result.IsSuccess() {
		t.Error("Expected extraction to fail for unsupported database")
	}

	if result.Error == nil {
		t.Error("Expected error for unsupported database")
	}
}

func TestExtractorOrchestrator_ExtractAll_InvalidFirebirdConnection(t *testing.T) {
	cfg := &config.Config{
		Databases: map[string]string{
			"invalid_firebird": "invalid:connection@localhost:3050/nonexistent.fdb",
		},
	}

	orch := orchestrator.NewExtractorOrchestrator(cfg, 1)
	results := orch.ExtractAll()

	if len(results) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(results))
	}

	result := results[0]
	if result.IsSuccess() {
		t.Error("Expected extraction to fail for invalid connection")
	}

	if result.DatabaseName != "invalid_firebird" {
		t.Errorf("Expected database name 'invalid_firebird', got %s", result.DatabaseName)
	}
}

func TestExtractorOrchestrator_WorkerCount(t *testing.T) {
	// Test that worker count is limited by number of databases
	cfg := &config.Config{
		Databases: map[string]string{
			"db1": "connection1",
			"db2": "connection2",
		},
	}

	// Request more workers than databases
	orch := orchestrator.NewExtractorOrchestrator(cfg, 10)
	results := orch.ExtractAll()

	// Should still process all databases
	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}
}

func TestExtractorOrchestrator_Timeout(t *testing.T) {
	cfg := &config.Config{
		Databases: map[string]string{
			"slow_db": "invalid:connection@localhost:3050/slow.fdb",
		},
	}

	orch := orchestrator.NewExtractorOrchestrator(cfg, 1)
	orch.SetTimeout(100 * time.Millisecond) // Very short timeout

	start := time.Now()
	results := orch.ExtractAll()
	duration := time.Since(start)

	if len(results) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(results))
	}

	result := results[0]
	if result.IsSuccess() {
		t.Error("Expected extraction to fail due to timeout or connection error")
	}

	// Should complete relatively quickly due to timeout or immediate connection failure
	if duration > 5*time.Second {
		t.Errorf("Extraction took too long: %v", duration)
	}
}