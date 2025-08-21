package firebird

import (
	"database-schema-extractor/internal/schema"
	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

// FirebirdExtractor implements DatabaseExtractor for Firebird databases
type FirebirdExtractor struct {
	connectionString string
	db               *sql.DB
}

// NewFirebirdExtractor creates a new Firebird extractor
func NewFirebirdExtractor() *FirebirdExtractor {
	return &FirebirdExtractor{}
}

// Connect establishes connection to the Firebird database
func (fe *FirebirdExtractor) Connect(connectionString string) error {
	fe.connectionString = connectionString
	
	db, err := sql.Open("firebirdsql", connectionString)
	if err != nil {
		return fmt.Errorf("failed to open Firebird connection: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return fmt.Errorf("failed to ping Firebird database: %w", err)
	}

	fe.db = db
	return nil
}

// ExtractSchema extracts the complete database schema from Firebird
func (fe *FirebirdExtractor) ExtractSchema() (*schema.DatabaseSchema, error) {
	if fe.db == nil {
		return nil, fmt.Errorf("database connection not established")
	}

	// Extract database name from connection string (simplified)
	dbName := extractDatabaseName(fe.connectionString)
	dbSchema := schema.NewDatabaseSchema(dbName)

	// Extract tables
	tables, err := fe.extractTables()
	if err != nil {
		return nil, fmt.Errorf("failed to extract tables: %w", err)
	}
	dbSchema.Tables = tables

	// Extract views
	views, err := fe.extractViews()
	if err != nil {
		return nil, fmt.Errorf("failed to extract views: %w", err)
	}
	dbSchema.Views = views

	// Extract procedures
	procedures, err := fe.extractProcedures()
	if err != nil {
		return nil, fmt.Errorf("failed to extract procedures: %w", err)
	}
	dbSchema.Procedures = procedures

	// Extract functions
	functions, err := fe.extractFunctions()
	if err != nil {
		return nil, fmt.Errorf("failed to extract functions: %w", err)
	}
	dbSchema.Functions = functions

	// Extract indexes
	indexes, err := fe.extractIndexes()
	if err != nil {
		return nil, fmt.Errorf("failed to extract indexes: %w", err)
	}
	dbSchema.Indexes = indexes

	// Extract Firebird-specific features
	engineSpecific, err := fe.extractEngineSpecific()
	if err != nil {
		return nil, fmt.Errorf("failed to extract engine-specific features: %w", err)
	}
	dbSchema.EngineSpecific = engineSpecific

	return dbSchema, nil
}

// Close closes the database connection
func (fe *FirebirdExtractor) Close() error {
	if fe.db != nil {
		return fe.db.Close()
	}
	return nil
}

// GetDriverName returns the driver name
func (fe *FirebirdExtractor) GetDriverName() string {
	return "firebirdsql"
}

// extractDatabaseName extracts database name from connection string
func extractDatabaseName(connectionString string) string {
	// Simple extraction - in real implementation, this could be more sophisticated
	// For now, we'll use a placeholder
	return "firebird_database"
}