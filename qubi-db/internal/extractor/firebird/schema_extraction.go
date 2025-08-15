package firebird

import (
	"database-schema-extractor/internal/schema"
	"database/sql"
	"fmt"
	"strings"
)

// extractTables extracts all user tables from the database
func (fe *FirebirdExtractor) extractTables() ([]schema.Table, error) {
	rows, err := fe.db.Query(queryTables)
	if err != nil {
		return nil, fmt.Errorf("failed to query tables: %w", err)
	}
	defer rows.Close()

	var tables []schema.Table
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, fmt.Errorf("failed to scan table name: %w", err)
		}

		table, err := fe.extractTableDetails(tableName)
		if err != nil {
			return nil, fmt.Errorf("failed to extract details for table %s: %w", tableName, err)
		}

		tables = append(tables, *table)
	}

	return tables, nil
}

// extractTableDetails extracts detailed information for a specific table
func (fe *FirebirdExtractor) extractTableDetails(tableName string) (*schema.Table, error) {
	table := &schema.Table{
		Name:        tableName,
		Columns:     []schema.Column{},
		PrimaryKey:  []string{},
		ForeignKeys: []schema.ForeignKey{},
	}

	// Extract columns
	columns, err := fe.extractTableColumns(tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to extract columns: %w", err)
	}
	table.Columns = columns

	// Extract primary key
	primaryKey, err := fe.extractPrimaryKey(tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to extract primary key: %w", err)
	}
	table.PrimaryKey = primaryKey

	// Extract foreign keys
	foreignKeys, err := fe.extractForeignKeys(tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to extract foreign keys: %w", err)
	}
	table.ForeignKeys = foreignKeys

	return table, nil
}

// extractTableColumns extracts columns for a specific table
func (fe *FirebirdExtractor) extractTableColumns(tableName string) ([]schema.Column, error) {
	rows, err := fe.db.Query(queryColumns, tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to query columns: %w", err)
	}
	defer rows.Close()

	var columns []schema.Column
	for rows.Next() {
		var columnName string
		var fieldType int
		var fieldLength, fieldPrecision, fieldScale sql.NullInt64
		var nullFlag sql.NullInt64
		var defaultValue sql.NullString
		var fieldSubType sql.NullInt64

		err := rows.Scan(&columnName, &fieldType, &fieldLength, &fieldPrecision, 
			&fieldScale, &nullFlag, &defaultValue, &fieldSubType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan column: %w", err)
		}

		column := schema.Column{
			Name:     columnName,
			Type:     mapFirebirdType(fieldType, fieldSubType.Int64),
			Nullable: nullFlag.Int64 != 1, // In Firebird, 1 means NOT NULL
		}

		// Set size/precision/scale based on field type
		if fieldLength.Valid && fieldLength.Int64 > 0 {
			size := int(fieldLength.Int64)
			column.Size = &size
		}
		if fieldPrecision.Valid && fieldPrecision.Int64 > 0 {
			precision := int(fieldPrecision.Int64)
			column.Precision = &precision
		}
		if fieldScale.Valid {
			scale := int(fieldScale.Int64)
			column.Scale = &scale
		}

		// Set default value
		if defaultValue.Valid && defaultValue.String != "" {
			// Clean up the default value (remove DEFAULT keyword)
			cleanDefault := strings.TrimSpace(strings.TrimPrefix(defaultValue.String, "DEFAULT"))
			column.DefaultValue = &cleanDefault
		}

		columns = append(columns, column)
	}

	return columns, nil
}

// extractPrimaryKey extracts primary key columns for a table
func (fe *FirebirdExtractor) extractPrimaryKey(tableName string) ([]string, error) {
	rows, err := fe.db.Query(queryPrimaryKeys, tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to query primary key: %w", err)
	}
	defer rows.Close()

	var primaryKey []string
	for rows.Next() {
		var constraintName, columnName string
		if err := rows.Scan(&constraintName, &columnName); err != nil {
			return nil, fmt.Errorf("failed to scan primary key: %w", err)
		}
		primaryKey = append(primaryKey, columnName)
	}

	return primaryKey, nil
}

// extractForeignKeys extracts foreign key constraints for a table
func (fe *FirebirdExtractor) extractForeignKeys(tableName string) ([]schema.ForeignKey, error) {
	rows, err := fe.db.Query(queryForeignKeys, tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to query foreign keys: %w", err)
	}
	defer rows.Close()

	foreignKeyMap := make(map[string]*schema.ForeignKey)
	
	for rows.Next() {
		var constraintName, columnName, referencedTable, referencedColumn string

		err := rows.Scan(&constraintName, &columnName, &referencedTable, &referencedColumn)
		if err != nil {
			return nil, fmt.Errorf("failed to scan foreign key: %w", err)
		}

		if fk, exists := foreignKeyMap[constraintName]; exists {
			fk.Columns = append(fk.Columns, columnName)
			fk.ReferencedColumns = append(fk.ReferencedColumns, referencedColumn)
		} else {
			fk := &schema.ForeignKey{
				Name:              constraintName,
				Columns:           []string{columnName},
				ReferencedTable:   referencedTable,
				ReferencedColumns: []string{referencedColumn},
				// OnDelete and OnUpdate rules not available in older Firebird versions
				OnDelete:          "",
				OnUpdate:          "",
			}
			
			foreignKeyMap[constraintName] = fk
		}
	}

	var foreignKeys []schema.ForeignKey
	for _, fk := range foreignKeyMap {
		foreignKeys = append(foreignKeys, *fk)
	}

	return foreignKeys, nil
}

// mapFirebirdType maps Firebird field types to standard type names
func mapFirebirdType(fieldType int, subType int64) string {
	switch fieldType {
	case 7: // SMALLINT
		return "SMALLINT"
	case 8: // INTEGER
		return "INTEGER"
	case 10: // FLOAT
		return "FLOAT"
	case 12: // DATE
		return "DATE"
	case 13: // TIME
		return "TIME"
	case 14: // CHAR
		return "CHAR"
	case 16: // BIGINT
		return "BIGINT"
	case 27: // DOUBLE PRECISION
		return "DOUBLE PRECISION"
	case 35: // TIMESTAMP
		return "TIMESTAMP"
	case 37: // VARCHAR
		return "VARCHAR"
	case 261: // BLOB
		if subType == 1 {
			return "BLOB SUB_TYPE TEXT"
		}
		return "BLOB"
	default:
		return fmt.Sprintf("UNKNOWN_TYPE_%d", fieldType)
	}
}

// extractViews extracts all user views from the database
func (fe *FirebirdExtractor) extractViews() ([]schema.View, error) {
	rows, err := fe.db.Query(queryViews)
	if err != nil {
		return nil, fmt.Errorf("failed to query views: %w", err)
	}
	defer rows.Close()

	var views []schema.View
	for rows.Next() {
		var viewName string
		var viewDefinition sql.NullString

		if err := rows.Scan(&viewName, &viewDefinition); err != nil {
			return nil, fmt.Errorf("failed to scan view: %w", err)
		}

		view := schema.View{
			Name:       viewName,
			Columns:    []schema.Column{}, // Would need separate query to get view columns
			Definition: viewDefinition.String,
		}

		views = append(views, view)
	}

	return views, nil
}

// extractProcedures extracts all stored procedures from the database
func (fe *FirebirdExtractor) extractProcedures() ([]schema.StoredProcedure, error) {
	rows, err := fe.db.Query(queryProcedures)
	if err != nil {
		return nil, fmt.Errorf("failed to query procedures: %w", err)
	}
	defer rows.Close()

	var procedures []schema.StoredProcedure
	for rows.Next() {
		var procedureName string
		var procedureSource sql.NullString

		if err := rows.Scan(&procedureName, &procedureSource); err != nil {
			return nil, fmt.Errorf("failed to scan procedure: %w", err)
		}

		// Extract parameters for this procedure
		parameters, err := fe.extractProcedureParameters(procedureName)
		if err != nil {
			return nil, fmt.Errorf("failed to extract parameters for procedure %s: %w", procedureName, err)
		}

		procedure := schema.StoredProcedure{
			Name:       procedureName,
			Parameters: parameters,
			Definition: procedureSource.String,
		}

		procedures = append(procedures, procedure)
	}

	return procedures, nil
}

// extractProcedureParameters extracts parameters for a specific procedure
func (fe *FirebirdExtractor) extractProcedureParameters(procedureName string) ([]schema.Parameter, error) {
	rows, err := fe.db.Query(queryProcedureParameters, procedureName)
	if err != nil {
		return nil, fmt.Errorf("failed to query procedure parameters: %w", err)
	}
	defer rows.Close()

	var parameters []schema.Parameter
	for rows.Next() {
		var paramName string
		var paramType int
		var fieldType int
		var fieldLength, fieldPrecision, fieldScale sql.NullInt64

		err := rows.Scan(&paramName, &paramType, &fieldType, &fieldLength, &fieldPrecision, &fieldScale)
		if err != nil {
			return nil, fmt.Errorf("failed to scan procedure parameter: %w", err)
		}

		direction := "IN"
		if paramType == 1 {
			direction = "OUT"
		}

		parameter := schema.Parameter{
			Name:      paramName,
			Type:      mapFirebirdType(fieldType, 0),
			Direction: direction,
		}

		parameters = append(parameters, parameter)
	}

	return parameters, nil
}

// extractFunctions extracts all user-defined functions from the database
func (fe *FirebirdExtractor) extractFunctions() ([]schema.Function, error) {
	// Note: This query might not work on older Firebird versions
	rows, err := fe.db.Query(queryFunctions)
	if err != nil {
		// If functions table doesn't exist (older Firebird), return empty slice
		return []schema.Function{}, nil
	}
	defer rows.Close()

	var functions []schema.Function
	for rows.Next() {
		var functionName string
		var functionSource, returnType sql.NullString

		if err := rows.Scan(&functionName, &functionSource, &returnType); err != nil {
			return nil, fmt.Errorf("failed to scan function: %w", err)
		}

		function := schema.Function{
			Name:       functionName,
			Parameters: []schema.Parameter{}, // Would need separate query for function parameters
			ReturnType: returnType.String,
			Definition: functionSource.String,
		}

		functions = append(functions, function)
	}

	return functions, nil
}

// extractIndexes extracts all user indexes from the database
func (fe *FirebirdExtractor) extractIndexes() ([]schema.Index, error) {
	rows, err := fe.db.Query(queryIndexes)
	if err != nil {
		return nil, fmt.Errorf("failed to query indexes: %w", err)
	}
	defer rows.Close()

	indexMap := make(map[string]*schema.Index)

	for rows.Next() {
		var indexName, tableName, columnName string
		var uniqueFlag sql.NullInt64
		var fieldPosition int

		err := rows.Scan(&indexName, &tableName, &uniqueFlag, &columnName, &fieldPosition)
		if err != nil {
			return nil, fmt.Errorf("failed to scan index: %w", err)
		}

		if index, exists := indexMap[indexName]; exists {
			index.Columns = append(index.Columns, columnName)
		} else {
			index := &schema.Index{
				Name:      indexName,
				TableName: tableName,
				Columns:   []string{columnName},
				Unique:    uniqueFlag.Valid && uniqueFlag.Int64 == 1,
				Type:      "BTREE", // Firebird uses B-tree indexes by default
			}
			indexMap[indexName] = index
		}
	}

	var indexes []schema.Index
	for _, index := range indexMap {
		indexes = append(indexes, *index)
	}

	return indexes, nil
}

// extractEngineSpecific extracts Firebird-specific features
func (fe *FirebirdExtractor) extractEngineSpecific() (map[string]interface{}, error) {
	engineSpecific := make(map[string]interface{})

	// Extract generators
	generators, err := fe.extractGenerators()
	if err != nil {
		return nil, fmt.Errorf("failed to extract generators: %w", err)
	}
	if len(generators) > 0 {
		engineSpecific["generators"] = generators
	}

	// Extract domains
	domains, err := fe.extractDomains()
	if err != nil {
		return nil, fmt.Errorf("failed to extract domains: %w", err)
	}
	if len(domains) > 0 {
		engineSpecific["domains"] = domains
	}

	return engineSpecific, nil
}

// extractGenerators extracts all user-defined generators/sequences
func (fe *FirebirdExtractor) extractGenerators() ([]map[string]interface{}, error) {
	rows, err := fe.db.Query(queryGenerators)
	if err != nil {
		return nil, fmt.Errorf("failed to query generators: %w", err)
	}
	defer rows.Close()

	var generators []map[string]interface{}
	for rows.Next() {
		var generatorName string
		var generatorValue sql.NullInt64

		if err := rows.Scan(&generatorName, &generatorValue); err != nil {
			return nil, fmt.Errorf("failed to scan generator: %w", err)
		}

		generator := map[string]interface{}{
			"name": generatorName,
		}

		if generatorValue.Valid {
			generator["current_value"] = generatorValue.Int64
		}

		generators = append(generators, generator)
	}

	return generators, nil
}

// extractDomains extracts all user-defined domains
func (fe *FirebirdExtractor) extractDomains() ([]map[string]interface{}, error) {
	rows, err := fe.db.Query(queryDomains)
	if err != nil {
		return nil, fmt.Errorf("failed to query domains: %w", err)
	}
	defer rows.Close()

	var domains []map[string]interface{}
	for rows.Next() {
		var domainName string
		var fieldType int
		var fieldLength, fieldPrecision, fieldScale sql.NullInt64
		var checkConstraint sql.NullString

		err := rows.Scan(&domainName, &fieldType, &fieldLength, &fieldPrecision, &fieldScale, &checkConstraint)
		if err != nil {
			return nil, fmt.Errorf("failed to scan domain: %w", err)
		}

		domain := map[string]interface{}{
			"name": domainName,
			"type": mapFirebirdType(fieldType, 0),
		}

		if fieldLength.Valid && fieldLength.Int64 > 0 {
			domain["length"] = fieldLength.Int64
		}
		if fieldPrecision.Valid && fieldPrecision.Int64 > 0 {
			domain["precision"] = fieldPrecision.Int64
		}
		if fieldScale.Valid {
			domain["scale"] = fieldScale.Int64
		}
		if checkConstraint.Valid && checkConstraint.String != "" {
			domain["check_constraint"] = strings.TrimSpace(checkConstraint.String)
		}

		domains = append(domains, domain)
	}

	return domains, nil
}