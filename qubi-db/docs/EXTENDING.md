# Extending Database Schema Extractor

This document explains how to add support for new database engines to the Database Schema Extractor.

## Architecture Overview

The application uses the Strategy pattern to support multiple database engines. Each engine implements the `DatabaseExtractor` interface and provides a corresponding `ExtractorFactory`.

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Orchestrator  │───▶│ ExtractorFactory │───▶│ DatabaseExtractor│
└─────────────────┘    └──────────────────┘    └─────────────────┘
                              │                         │
                              ▼                         ▼
                       ┌──────────────────┐    ┌─────────────────┐
                       │ SupportsConnection│    │ ExtractSchema   │
                       │ String()         │    │ Connect()       │
                       └──────────────────┘    │ Close()         │
                                               └─────────────────┘
```

## Adding a New Database Engine

### Step 1: Create the Extractor Implementation

Create a new package under `internal/extractor/` for your database engine:

```go
// internal/extractor/sqlserver/extractor.go
package sqlserver

import (
    "database-schema-extractor/internal/schema"
    "database/sql"
    "fmt"
    
    _ "github.com/denisenkom/go-mssqldb" // SQL Server driver
)

type SQLServerExtractor struct {
    connectionString string
    db               *sql.DB
}

func NewSQLServerExtractor() *SQLServerExtractor {
    return &SQLServerExtractor{}
}

func (se *SQLServerExtractor) Connect(connectionString string) error {
    se.connectionString = connectionString
    
    db, err := sql.Open("sqlserver", connectionString)
    if err != nil {
        return fmt.Errorf("failed to open SQL Server connection: %w", err)
    }
    
    if err := db.Ping(); err != nil {
        db.Close()
        return fmt.Errorf("failed to ping SQL Server database: %w", err)
    }
    
    se.db = db
    return nil
}

func (se *SQLServerExtractor) ExtractSchema() (*schema.DatabaseSchema, error) {
    if se.db == nil {
        return nil, fmt.Errorf("database connection not established")
    }
    
    dbName := extractDatabaseName(se.connectionString)
    dbSchema := schema.NewDatabaseSchema(dbName)
    
    // Extract tables
    tables, err := se.extractTables()
    if err != nil {
        return nil, fmt.Errorf("failed to extract tables: %w", err)
    }
    dbSchema.Tables = tables
    
    // Extract other objects...
    
    return dbSchema, nil
}

func (se *SQLServerExtractor) Close() error {
    if se.db != nil {
        return se.db.Close()
    }
    return nil
}

func (se *SQLServerExtractor) GetDriverName() string {
    return "sqlserver"
}
```

### Step 2: Implement Schema Extraction Methods

Create methods to extract different database objects:

```go
// internal/extractor/sqlserver/schema_extraction.go
package sqlserver

import (
    "database-schema-extractor/internal/schema"
    "database/sql"
)

func (se *SQLServerExtractor) extractTables() ([]schema.Table, error) {
    query := `
        SELECT TABLE_NAME 
        FROM INFORMATION_SCHEMA.TABLES 
        WHERE TABLE_TYPE = 'BASE TABLE' 
          AND TABLE_SCHEMA != 'sys'
          AND TABLE_SCHEMA != 'INFORMATION_SCHEMA'
        ORDER BY TABLE_NAME`
    
    rows, err := se.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var tables []schema.Table
    for rows.Next() {
        var tableName string
        if err := rows.Scan(&tableName); err != nil {
            return nil, err
        }
        
        table, err := se.extractTableDetails(tableName)
        if err != nil {
            return nil, err
        }
        
        tables = append(tables, *table)
    }
    
    return tables, nil
}

func (se *SQLServerExtractor) extractTableDetails(tableName string) (*schema.Table, error) {
    // Implementation for extracting table columns, constraints, etc.
    // Similar to Firebird implementation but with SQL Server specific queries
}
```

### Step 3: Create the Factory

```go
// internal/extractor/sqlserver/factory.go
package sqlserver

import (
    "database-schema-extractor/internal/extractor"
    "strings"
)

type SQLServerExtractorFactory struct{}

func NewSQLServerExtractorFactory() *SQLServerExtractorFactory {
    return &SQLServerExtractorFactory{}
}

func (f *SQLServerExtractorFactory) CreateExtractor(connectionString string) (extractor.DatabaseExtractor, error) {
    return NewSQLServerExtractor(), nil
}

func (f *SQLServerExtractorFactory) SupportsConnectionString(connectionString string) bool {
    // Check for SQL Server connection string patterns
    return strings.Contains(connectionString, "server=") ||
           strings.Contains(connectionString, "sqlserver://") ||
           strings.Contains(connectionString, "Data Source=")
}
```

### Step 4: Register the Factory

Add your factory to the orchestrator:

```go
// In cmd/main.go or where the orchestrator is configured
import "database-schema-extractor/internal/extractor/sqlserver"

func main() {
    // ... existing code ...
    
    orch := orchestrator.NewExtractorOrchestrator(cfg, 5)
    
    // Add SQL Server support
    orch.AddExtractorFactory(sqlserver.NewSQLServerExtractorFactory())
    
    // ... rest of the code ...
}
```

## Database-Specific Considerations

### Connection String Formats

Each database engine has its own connection string format. Document the expected format:

- **Firebird**: `user:password@host:port/database_path`
- **SQL Server**: `server=host;database=db;user=user;password=pass`
- **PostgreSQL**: `postgres://user:password@host:port/database`
- **MySQL**: `user:password@tcp(host:port)/database`

### System Table Filtering

Each database has different system tables that should be filtered out:

```go
// Example for SQL Server
func (se *SQLServerExtractor) isSystemTable(tableName, schemaName string) bool {
    systemSchemas := []string{"sys", "INFORMATION_SCHEMA", "db_owner", "db_accessadmin"}
    for _, sysSchema := range systemSchemas {
        if schemaName == sysSchema {
            return true
        }
    }
    return false
}
```

### Engine-Specific Features

Use the `EngineSpecific` field to include database-specific features:

```go
// SQL Server specific features
engineSpecific := map[string]interface{}{
    "schemas": schemas,
    "user_defined_types": userTypes,
    "assemblies": assemblies,
}
dbSchema.EngineSpecific = engineSpecific
```

## Testing Your Implementation

### Unit Tests

Create unit tests for your extractor:

```go
// test/extractor/sqlserver/extractor_test.go
package sqlserver_test

import (
    "database-schema-extractor/internal/extractor/sqlserver"
    "testing"
)

func TestSQLServerExtractor_GetDriverName(t *testing.T) {
    extractor := sqlserver.NewSQLServerExtractor()
    if extractor.GetDriverName() != "sqlserver" {
        t.Errorf("Expected driver name 'sqlserver', got %s", extractor.GetDriverName())
    }
}

func TestSQLServerExtractorFactory_SupportsConnectionString(t *testing.T) {
    factory := sqlserver.NewSQLServerExtractorFactory()
    
    tests := []struct {
        connStr  string
        expected bool
    }{
        {"server=localhost;database=test;user=sa;password=pass", true},
        {"sqlserver://sa:pass@localhost:1433?database=test", true},
        {"user:pass@localhost:3050/test.fdb", false},
    }
    
    for _, tt := range tests {
        result := factory.SupportsConnectionString(tt.connStr)
        if result != tt.expected {
            t.Errorf("SupportsConnectionString(%s) = %v, want %v", 
                tt.connStr, result, tt.expected)
        }
    }
}
```

### Integration Tests

Create integration tests that use a real database:

```go
func TestSQLServerExtractor_Integration(t *testing.T) {
    t.Skip("Integration test requires running SQL Server database")
    
    extractor := sqlserver.NewSQLServerExtractor()
    err := extractor.Connect("server=localhost;database=test;user=sa;password=password")
    if err != nil {
        t.Fatalf("Failed to connect: %v", err)
    }
    defer extractor.Close()
    
    schema, err := extractor.ExtractSchema()
    if err != nil {
        t.Fatalf("Failed to extract schema: %v", err)
    }
    
    // Verify schema content
    if schema.DatabaseName == "" {
        t.Error("Database name should not be empty")
    }
}
```

## Common SQL Queries by Database

### Tables and Columns

**SQL Server:**
```sql
SELECT t.TABLE_NAME, c.COLUMN_NAME, c.DATA_TYPE, c.IS_NULLABLE
FROM INFORMATION_SCHEMA.TABLES t
JOIN INFORMATION_SCHEMA.COLUMNS c ON t.TABLE_NAME = c.TABLE_NAME
WHERE t.TABLE_TYPE = 'BASE TABLE' AND t.TABLE_SCHEMA != 'sys'
```

**PostgreSQL:**
```sql
SELECT t.tablename, c.column_name, c.data_type, c.is_nullable
FROM pg_tables t
JOIN information_schema.columns c ON t.tablename = c.table_name
WHERE t.schemaname = 'public'
```

**MySQL:**
```sql
SELECT t.TABLE_NAME, c.COLUMN_NAME, c.DATA_TYPE, c.IS_NULLABLE
FROM information_schema.TABLES t
JOIN information_schema.COLUMNS c ON t.TABLE_NAME = c.TABLE_NAME
WHERE t.TABLE_SCHEMA = DATABASE() AND t.TABLE_TYPE = 'BASE TABLE'
```

### Primary Keys

**SQL Server:**
```sql
SELECT tc.TABLE_NAME, kcu.COLUMN_NAME
FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc
JOIN INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu ON tc.CONSTRAINT_NAME = kcu.CONSTRAINT_NAME
WHERE tc.CONSTRAINT_TYPE = 'PRIMARY KEY'
```

### Foreign Keys

**SQL Server:**
```sql
SELECT 
    fk.name AS constraint_name,
    tp.name AS table_name,
    cp.name AS column_name,
    tr.name AS referenced_table,
    cr.name AS referenced_column
FROM sys.foreign_keys fk
JOIN sys.foreign_key_columns fkc ON fk.object_id = fkc.constraint_object_id
JOIN sys.tables tp ON fkc.parent_object_id = tp.object_id
JOIN sys.columns cp ON fkc.parent_object_id = cp.object_id AND fkc.parent_column_id = cp.column_id
JOIN sys.tables tr ON fkc.referenced_object_id = tr.object_id
JOIN sys.columns cr ON fkc.referenced_object_id = cr.object_id AND fkc.referenced_column_id = cr.column_id
```

## Best Practices

1. **Error Handling**: Always provide meaningful error messages with context
2. **Resource Management**: Use defer statements to ensure connections are closed
3. **Logging**: Use structured logging to help with debugging
4. **Testing**: Write both unit and integration tests
5. **Documentation**: Document connection string formats and engine-specific features
6. **Performance**: Use prepared statements for repeated queries
7. **Security**: Never log connection strings or passwords

## Example: Complete PostgreSQL Implementation

See the `examples/postgresql/` directory for a complete implementation example that demonstrates all the concepts covered in this guide.