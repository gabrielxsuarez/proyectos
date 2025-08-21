package schema_test

import (
	"database-schema-extractor/internal/schema"
	"testing"
)

func TestNewDatabaseSchema(t *testing.T) {
	name := "test_database"
	dbSchema := schema.NewDatabaseSchema(name)

	if dbSchema.DatabaseName != name {
		t.Errorf("Expected database name %s, got %s", name, dbSchema.DatabaseName)
	}

	if dbSchema.Tables == nil {
		t.Error("Tables slice should be initialized")
	}

	if dbSchema.Views == nil {
		t.Error("Views slice should be initialized")
	}

	if dbSchema.Procedures == nil {
		t.Error("Procedures slice should be initialized")
	}

	if dbSchema.Functions == nil {
		t.Error("Functions slice should be initialized")
	}

	if dbSchema.Indexes == nil {
		t.Error("Indexes slice should be initialized")
	}

	if dbSchema.EngineSpecific == nil {
		t.Error("EngineSpecific map should be initialized")
	}

	// Test that slices are empty but not nil
	if len(dbSchema.Tables) != 0 {
		t.Error("Tables slice should be empty initially")
	}

	if len(dbSchema.Views) != 0 {
		t.Error("Views slice should be empty initially")
	}
}

func TestColumn_Structure(t *testing.T) {
	size := 255
	precision := 10
	scale := 2
	defaultValue := "test"

	column := schema.Column{
		Name:          "test_column",
		Type:          "VARCHAR",
		Size:          &size,
		Precision:     &precision,
		Scale:         &scale,
		Nullable:      true,
		DefaultValue:  &defaultValue,
		AutoIncrement: false,
	}

	if column.Name != "test_column" {
		t.Errorf("Expected column name 'test_column', got %s", column.Name)
	}

	if column.Type != "VARCHAR" {
		t.Errorf("Expected column type 'VARCHAR', got %s", column.Type)
	}

	if column.Size == nil || *column.Size != 255 {
		t.Errorf("Expected column size 255, got %v", column.Size)
	}

	if !column.Nullable {
		t.Error("Expected column to be nullable")
	}
}

func TestTable_Structure(t *testing.T) {
	table := schema.Table{
		Name: "test_table",
		Columns: []schema.Column{
			{Name: "id", Type: "INTEGER", Nullable: false},
			{Name: "name", Type: "VARCHAR", Nullable: true},
		},
		PrimaryKey: []string{"id"},
		ForeignKeys: []schema.ForeignKey{
			{
				Name:              "fk_test",
				Columns:           []string{"parent_id"},
				ReferencedTable:   "parent_table",
				ReferencedColumns: []string{"id"},
			},
		},
	}

	if table.Name != "test_table" {
		t.Errorf("Expected table name 'test_table', got %s", table.Name)
	}

	if len(table.Columns) != 2 {
		t.Errorf("Expected 2 columns, got %d", len(table.Columns))
	}

	if len(table.PrimaryKey) != 1 || table.PrimaryKey[0] != "id" {
		t.Errorf("Expected primary key ['id'], got %v", table.PrimaryKey)
	}

	if len(table.ForeignKeys) != 1 {
		t.Errorf("Expected 1 foreign key, got %d", len(table.ForeignKeys))
	}
}