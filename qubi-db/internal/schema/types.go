package schema

// DatabaseSchema represents the complete schema of a database
type DatabaseSchema struct {
	DatabaseName   string                 `yaml:"database_name"`
	Tables         []Table                `yaml:"tables"`
	Views          []View                 `yaml:"views"`
	Procedures     []StoredProcedure      `yaml:"procedures"`
	Functions      []Function             `yaml:"functions"`
	Indexes        []Index                `yaml:"indexes"`
	EngineSpecific map[string]interface{} `yaml:"engine_specific,omitempty"`
}

// Table represents a database table
type Table struct {
	Name        string       `yaml:"name"`
	Columns     []Column     `yaml:"columns"`
	PrimaryKey  []string     `yaml:"primary_key,omitempty"`
	ForeignKeys []ForeignKey `yaml:"foreign_keys,omitempty"`
}

// Column represents a table column
type Column struct {
	Name          string  `yaml:"name"`
	Type          string  `yaml:"type"`
	Size          *int    `yaml:"size,omitempty"`
	Precision     *int    `yaml:"precision,omitempty"`
	Scale         *int    `yaml:"scale,omitempty"`
	Nullable      bool    `yaml:"nullable"`
	DefaultValue  *string `yaml:"default_value,omitempty"`
	AutoIncrement bool    `yaml:"auto_increment,omitempty"`
}

// View represents a database view
type View struct {
	Name       string   `yaml:"name"`
	Columns    []Column `yaml:"columns"`
	Definition string   `yaml:"definition"`
}

// StoredProcedure represents a stored procedure
type StoredProcedure struct {
	Name       string      `yaml:"name"`
	Parameters []Parameter `yaml:"parameters"`
	Definition string      `yaml:"definition"`
}

// Function represents a database function
type Function struct {
	Name       string      `yaml:"name"`
	Parameters []Parameter `yaml:"parameters"`
	ReturnType string      `yaml:"return_type"`
	Definition string      `yaml:"definition"`
}

// Parameter represents a procedure or function parameter
type Parameter struct {
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Direction string `yaml:"direction"` // IN, OUT, INOUT
}

// Index represents a database index
type Index struct {
	Name      string   `yaml:"name"`
	TableName string   `yaml:"table_name"`
	Columns   []string `yaml:"columns"`
	Unique    bool     `yaml:"unique"`
	Type      string   `yaml:"type,omitempty"` // BTREE, HASH, etc.
}

// ForeignKey represents a foreign key constraint
type ForeignKey struct {
	Name            string   `yaml:"name"`
	Columns         []string `yaml:"columns"`
	ReferencedTable string   `yaml:"referenced_table"`
	ReferencedColumns []string `yaml:"referenced_columns"`
	OnDelete        string   `yaml:"on_delete,omitempty"`
	OnUpdate        string   `yaml:"on_update,omitempty"`
}

// NewDatabaseSchema creates a new DatabaseSchema with initialized slices
func NewDatabaseSchema(name string) *DatabaseSchema {
	return &DatabaseSchema{
		DatabaseName:   name,
		Tables:         make([]Table, 0),
		Views:          make([]View, 0),
		Procedures:     make([]StoredProcedure, 0),
		Functions:      make([]Function, 0),
		Indexes:        make([]Index, 0),
		EngineSpecific: make(map[string]interface{}),
	}
}