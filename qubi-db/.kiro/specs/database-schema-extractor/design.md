# Design Document

## Overview

La aplicación Database Schema Extractor es una herramienta CLI en Go que extrae esquemas de bases de datos de forma paralela y los exporta como archivos YAML estructurados. Utiliza el patrón Strategy para soportar múltiples motores de bases de datos de manera extensible.

## Architecture

La aplicación sigue una arquitectura modular con separación clara de responsabilidades:

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Config        │    │   Extractor      │    │   Output        │
│   Manager       │───▶│   Orchestrator   │───▶│   Generator     │
└─────────────────┘    └──────────────────┘    └─────────────────┘
                              │
                              ▼
                       ┌──────────────────┐
                       │   DB Strategy    │
                       │   Interface      │
                       └──────────────────┘
                              │
                    ┌─────────┼─────────┐
                    ▼         ▼         ▼
              ┌──────────┐ ┌──────────┐ ┌──────────┐
              │ Firebird │ │SQL Server│ │  Future  │
              │ Strategy │ │ Strategy │ │Strategies│
              └──────────┘ └──────────┘ └──────────┘
```

## Components and Interfaces

### 1. Configuration Management

**ConfigManager**
- Responsabilidad: Leer config.yaml
- Métodos principales:
  - `LoadConfig() (*Config, error)`

**Estructuras:**
```go
type Config struct {
    Databases map[string]string `yaml:",inline"`
}

type DatabaseConnection struct {
    Name             string
    ConnectionString string
}
```

### 2. Database Strategy Pattern

**DatabaseExtractor Interface**
```go
type DatabaseExtractor interface {
    Connect(connectionString string) error
    ExtractSchema() (*DatabaseSchema, error)
    Close() error
    GetDriverName() string
}
```

**FirebirdExtractor**
- Implementa DatabaseExtractor
- Maneja conexiones específicas de Firebird
- Filtra tablas del sistema (RDB$*, MON$*)
- Extrae generators específicos de Firebird

**Futuras implementaciones:**
- SQLServerExtractor
- PostgreSQLExtractor
- MySQLExtractor

### 3. Schema Data Models

**DatabaseSchema**
```go
type DatabaseSchema struct {
    DatabaseName   string                 `yaml:"database_name"`
    Tables         []Table               `yaml:"tables"`
    Views          []View                `yaml:"views"`
    Procedures     []StoredProcedure     `yaml:"procedures"`
    Functions      []Function            `yaml:"functions"`
    Indexes        []Index               `yaml:"indexes"`
    EngineSpecific map[string]interface{} `yaml:"engine_specific,omitempty"`
}

type Table struct {
    Name        string   `yaml:"name"`
    Columns     []Column `yaml:"columns"`
    PrimaryKey  []string `yaml:"primary_key,omitempty"`
    ForeignKeys []ForeignKey `yaml:"foreign_keys,omitempty"`
}

type Column struct {
    Name         string      `yaml:"name"`
    Type         string      `yaml:"type"`
    Size         *int        `yaml:"size,omitempty"`
    Precision    *int        `yaml:"precision,omitempty"`
    Scale        *int        `yaml:"scale,omitempty"`
    Nullable     bool        `yaml:"nullable"`
    DefaultValue *string     `yaml:"default_value,omitempty"`
    AutoIncrement bool       `yaml:"auto_increment,omitempty"`
}
```

### 4. Extraction Orchestrator

**ExtractorOrchestrator**
- Coordina la extracción paralela
- Maneja timeouts y errores
- Implementa worker pool pattern

```go
type ExtractorOrchestrator struct {
    config      *Config
    maxWorkers  int
    timeout     time.Duration
}

func (eo *ExtractorOrchestrator) ExtractAll() []ExtractionResult
```

### 5. Output Generation

**YAMLGenerator**
- Convierte DatabaseSchema a YAML
- Maneja serialización con nombres en inglés
- Genera archivos de salida y error

## Data Models

### Firebird Specific Extensions

Para Firebird, la sección `engine_specific` incluirá:

```yaml
engine_specific:
  generators:
    - name: "GEN_CLIENTE_ID"
      current_value: 1000
    - name: "GEN_PRODUCTO_ID" 
      current_value: 500
  domains:
    - name: "D_BOOLEAN"
      type: "SMALLINT"
      check_constraint: "VALUE IN (0,1)"
```

### Connection String Usage

Las cadenas de conexión se usarán tal como están definidas en el config.yaml, sin parsing adicional. Cada implementación de DatabaseExtractor será responsable de interpretar el formato específico de su motor de base de datos.

Ejemplo para Firebird: `sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\ifarmacia.fdb`

## Error Handling

### Estrategia de Manejo de Errores

1. **Errores de Configuración**: Terminan la aplicación con mensaje descriptivo
2. **Errores de Conexión**: Se registran pero no detienen otras extracciones
3. **Errores de Extracción**: Se genera archivo de error específico
4. **Timeouts**: Se configuran por conexión (default: 30 segundos)

### Logging

- Uso de structured logging (logrus o zap)
- Niveles: ERROR, WARN, INFO, DEBUG
- Salida a stdout con formato JSON para facilitar parsing

## Testing Strategy

### Unit Tests

1. **ConfigManager Tests**
   - Parsing de config.yaml válido e inválido
   - Parsing de connection strings
   - Manejo de archivos faltantes

2. **DatabaseExtractor Tests**
   - Mocks para cada implementación de strategy
   - Validación de schema extraction
   - Manejo de errores de conexión

3. **YAMLGenerator Tests**
   - Serialización correcta de estructuras
   - Validación de formato de salida
   - Manejo de caracteres especiales

### Integration Tests

1. **End-to-End Tests**
   - Usar bases de datos de prueba en contenedores
   - Validar extracción completa con datos reales
   - Verificar archivos de salida generados

2. **Database-Specific Tests**
   - Tests específicos para Firebird con datos de prueba
   - Validación de filtrado de tablas del sistema
   - Extracción de características específicas (generators)

### Performance Tests

1. **Concurrency Tests**
   - Validar extracción paralela con múltiples DBs
   - Medir tiempo de ejecución vs extracción secuencial
   - Validar manejo de recursos y memory leaks

2. **Load Tests**
   - Probar con bases de datos grandes (1000+ tablas)
   - Validar comportamiento con conexiones lentas
   - Medir uso de memoria y CPU

### Test Data Setup

- Scripts SQL para crear estructuras de prueba en Firebird
- Datos de prueba que incluyan todos los tipos de objetos
- Configuraciones de prueba con diferentes escenarios