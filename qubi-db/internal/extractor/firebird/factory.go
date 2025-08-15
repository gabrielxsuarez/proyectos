package firebird

import (
	"database-schema-extractor/internal/extractor"
	"strings"
)

// FirebirdExtractorFactory creates Firebird extractors
type FirebirdExtractorFactory struct{}

// NewFirebirdExtractorFactory creates a new Firebird extractor factory
func NewFirebirdExtractorFactory() *FirebirdExtractorFactory {
	return &FirebirdExtractorFactory{}
}

// CreateExtractor creates a new Firebird extractor
func (fef *FirebirdExtractorFactory) CreateExtractor(connectionString string) (extractor.DatabaseExtractor, error) {
	return NewFirebirdExtractor(), nil
}

// SupportsConnectionString checks if the connection string is for Firebird
func (fef *FirebirdExtractorFactory) SupportsConnectionString(connectionString string) bool {
	// Check for Firebird-specific patterns in connection string
	// Firebird connection strings typically contain .fdb or .gdb extensions
	// or specific port 3050
	return strings.Contains(connectionString, ".fdb") ||
		strings.Contains(connectionString, ".gdb") ||
		strings.Contains(connectionString, ":3050/")
}