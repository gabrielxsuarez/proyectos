package output

import (
	"database-schema-extractor/internal/schema"
	"time"
)

// ExtractionResult represents the result of a database schema extraction
type ExtractionResult struct {
	DatabaseName string
	Schema       *schema.DatabaseSchema
	Error        error
	Timestamp    time.Time
}

// NewExtractionResult creates a new successful extraction result
func NewExtractionResult(databaseName string, schema *schema.DatabaseSchema) ExtractionResult {
	return ExtractionResult{
		DatabaseName: databaseName,
		Schema:       schema,
		Error:        nil,
		Timestamp:    time.Now(),
	}
}

// NewExtractionError creates a new failed extraction result
func NewExtractionError(databaseName string, err error) ExtractionResult {
	return ExtractionResult{
		DatabaseName: databaseName,
		Schema:       nil,
		Error:        err,
		Timestamp:    time.Now(),
	}
}

// IsSuccess returns true if the extraction was successful
func (er ExtractionResult) IsSuccess() bool {
	return er.Error == nil && er.Schema != nil
}

// IsError returns true if the extraction failed
func (er ExtractionResult) IsError() bool {
	return er.Error != nil
}