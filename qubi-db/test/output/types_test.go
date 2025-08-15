package output_test

import (
	"database-schema-extractor/internal/output"
	"database-schema-extractor/internal/schema"
	"errors"
	"testing"
	"time"
)

func TestNewExtractionResult(t *testing.T) {
	testSchema := schema.NewDatabaseSchema("test_db")
	result := output.NewExtractionResult("test_db", testSchema)

	if result.DatabaseName != "test_db" {
		t.Errorf("Expected database name 'test_db', got %s", result.DatabaseName)
	}

	if result.Schema != testSchema {
		t.Error("Schema should match the provided schema")
	}

	if result.Error != nil {
		t.Errorf("Error should be nil for successful result, got %v", result.Error)
	}

	if result.Timestamp.IsZero() {
		t.Error("Timestamp should be set")
	}

	if !result.IsSuccess() {
		t.Error("Result should be successful")
	}

	if result.IsError() {
		t.Error("Result should not be an error")
	}
}

func TestNewExtractionError(t *testing.T) {
	testError := errors.New("connection failed")
	result := output.NewExtractionError("failed_db", testError)

	if result.DatabaseName != "failed_db" {
		t.Errorf("Expected database name 'failed_db', got %s", result.DatabaseName)
	}

	if result.Schema != nil {
		t.Error("Schema should be nil for error result")
	}

	if result.Error != testError {
		t.Errorf("Error should match the provided error, got %v", result.Error)
	}

	if result.Timestamp.IsZero() {
		t.Error("Timestamp should be set")
	}

	if result.IsSuccess() {
		t.Error("Result should not be successful")
	}

	if !result.IsError() {
		t.Error("Result should be an error")
	}
}

func TestExtractionResult_IsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		result   output.ExtractionResult
		expected bool
	}{
		{
			name: "successful result",
			result: output.ExtractionResult{
				DatabaseName: "test",
				Schema:       schema.NewDatabaseSchema("test"),
				Error:        nil,
				Timestamp:    time.Now(),
			},
			expected: true,
		},
		{
			name: "error result",
			result: output.ExtractionResult{
				DatabaseName: "test",
				Schema:       nil,
				Error:        errors.New("failed"),
				Timestamp:    time.Now(),
			},
			expected: false,
		},
		{
			name: "nil schema but no error",
			result: output.ExtractionResult{
				DatabaseName: "test",
				Schema:       nil,
				Error:        nil,
				Timestamp:    time.Now(),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result.IsSuccess() != tt.expected {
				t.Errorf("IsSuccess() = %v, want %v", tt.result.IsSuccess(), tt.expected)
			}
		})
	}
}

func TestExtractionResult_IsError(t *testing.T) {
	tests := []struct {
		name     string
		result   output.ExtractionResult
		expected bool
	}{
		{
			name: "successful result",
			result: output.ExtractionResult{
				DatabaseName: "test",
				Schema:       schema.NewDatabaseSchema("test"),
				Error:        nil,
				Timestamp:    time.Now(),
			},
			expected: false,
		},
		{
			name: "error result",
			result: output.ExtractionResult{
				DatabaseName: "test",
				Schema:       nil,
				Error:        errors.New("failed"),
				Timestamp:    time.Now(),
			},
			expected: true,
		},
		{
			name: "nil error",
			result: output.ExtractionResult{
				DatabaseName: "test",
				Schema:       schema.NewDatabaseSchema("test"),
				Error:        nil,
				Timestamp:    time.Now(),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result.IsError() != tt.expected {
				t.Errorf("IsError() = %v, want %v", tt.result.IsError(), tt.expected)
			}
		})
	}
}