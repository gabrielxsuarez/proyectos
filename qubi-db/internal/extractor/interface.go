// Package extractor provides interfaces and implementations for database schema extraction.
//
// This package implements the Strategy pattern to support multiple database engines.
// Each database engine (Firebird, SQL Server, etc.) implements the DatabaseExtractor
// interface to provide engine-specific schema extraction logic.
//
// Example usage:
//
//	factory := firebird.NewFirebirdExtractorFactory()
//	if factory.SupportsConnectionString(connStr) {
//		extractor, err := factory.CreateExtractor(connStr)
//		if err != nil {
//			return err
//		}
//		defer extractor.Close()
//		
//		if err := extractor.Connect(connStr); err != nil {
//			return err
//		}
//		
//		schema, err := extractor.ExtractSchema()
//		if err != nil {
//			return err
//		}
//		// Use schema...
//	}
package extractor

import "database-schema-extractor/internal/schema"

// DatabaseExtractor defines the interface for database schema extraction strategies.
//
// This interface implements the Strategy pattern, allowing different database engines
// to provide their own extraction logic while maintaining a consistent API.
//
// Implementations should:
//   - Handle connection management (Connect/Close)
//   - Extract complete schema information including tables, views, procedures, etc.
//   - Filter out system tables and objects
//   - Include engine-specific features in the EngineSpecific field
//   - Handle errors gracefully and provide meaningful error messages
type DatabaseExtractor interface {
	// Connect establishes connection to the database using the provided connection string.
	//
	// The connection string format is engine-specific:
	//   - Firebird: "user:password@host:port/database_path"
	//   - SQL Server: "server=host;database=db;user=user;password=pass"
	//
	// Returns an error if the connection cannot be established.
	Connect(connectionString string) error
	
	// ExtractSchema extracts the complete database schema including:
	//   - Tables with columns, primary keys, foreign keys
	//   - Views with their definitions
	//   - Stored procedures and functions
	//   - Indexes
	//   - Engine-specific features (generators, domains, etc.)
	//
	// Only user-created objects are included; system tables and objects are filtered out.
	//
	// Returns the complete schema or an error if extraction fails.
	// The connection must be established before calling this method.
	ExtractSchema() (*schema.DatabaseSchema, error)
	
	// Close closes the database connection and releases associated resources.
	//
	// This method should be called when the extractor is no longer needed,
	// typically in a defer statement after successful connection.
	//
	// Returns an error if the connection cannot be closed properly.
	Close() error
	
	// GetDriverName returns the name of the database driver used by this extractor.
	//
	// This is used for logging and debugging purposes.
	// Examples: "firebirdsql", "sqlserver", "postgres"
	GetDriverName() string
}

// ExtractorFactory creates database extractors based on connection string patterns.
//
// Factories implement the Factory pattern to create appropriate extractors
// for different database engines. The orchestrator uses factories to determine
// which extractor to use for each connection string.
//
// Example implementation:
//
//	type FirebirdExtractorFactory struct{}
//	
//	func (f *FirebirdExtractorFactory) SupportsConnectionString(connStr string) bool {
//		return strings.Contains(connStr, ".fdb") || strings.Contains(connStr, ":3050/")
//	}
//	
//	func (f *FirebirdExtractorFactory) CreateExtractor(connStr string) (DatabaseExtractor, error) {
//		return NewFirebirdExtractor(), nil
//	}
type ExtractorFactory interface {
	// CreateExtractor creates a new DatabaseExtractor instance for the given connection string.
	//
	// The factory should create an appropriate extractor but not establish the connection.
	// The caller is responsible for calling Connect() on the returned extractor.
	//
	// Returns a new extractor instance or an error if creation fails.
	CreateExtractor(connectionString string) (DatabaseExtractor, error)
	
	// SupportsConnectionString determines if this factory can handle the given connection string.
	//
	// Factories should examine the connection string format to determine compatibility.
	// This method should be fast as it may be called multiple times during orchestration.
	//
	// Returns true if this factory can create an extractor for the connection string.
	SupportsConnectionString(connectionString string) bool
}