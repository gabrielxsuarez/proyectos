# Database Schema Extractor

Una aplicación CLI en Go que extrae esquemas de bases de datos de forma paralela y los exporta como archivos YAML estructurados.

## Características

- ✅ **Procesamiento paralelo**: Extrae esquemas de múltiples bases de datos simultáneamente
- ✅ **Patrón Strategy**: Arquitectura extensible para soportar diferentes motores de bases de datos
- ✅ **Soporte inicial para Firebird**: Con planes para SQL Server y otros motores
- ✅ **Filtrado automático**: Excluye tablas del sistema automáticamente
- ✅ **Formato YAML**: Salida estructurada con nombres en inglés
- ✅ **Características específicas**: Soporte para generators, domains y otras características específicas de cada motor

## Instalación

### Prerrequisitos

- Go 1.21 o superior
- Acceso a las bases de datos que deseas extraer

### Compilación

```bash
# Clonar el repositorio
git clone <repository-url>
cd database-schema-extractor

# Descargar dependencias
go mod download

# Compilar la aplicación
go build -o schema-extractor cmd/main.go
```

## Configuración

Crea un archivo `config.yaml` en el directorio donde ejecutarás la aplicación:

```yaml
# Ejemplo de configuración para bases de datos Firebird
ifarmacia.fdb: "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\ifarmacia.fdb"
clientes.fdb: "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\clientes.fdb"
inventario.fdb: "sysdba:masterkey@localhost:3050/C:\\AlfaBeta\\firebird\\inventario.fdb"
```

### Formato de Cadenas de Conexión

#### Firebird
```
usuario:contraseña@host:puerto/ruta_completa_base_datos
```

Ejemplos:
- `sysdba:masterkey@localhost:3050/C:\\database\\mydb.fdb`
- `admin:password@192.168.1.100:3050/opt/firebird/data/production.fdb`

## Uso

### Ejecución Básica

```bash
# Ejecutar desde el directorio que contiene config.yaml
./schema-extractor
```

### Ejemplo de Salida

La aplicación generará archivos YAML para cada base de datos:

```
ifarmacia.fdb.yaml
clientes.fdb.yaml
inventario.fdb.yaml
```

### Estructura de Archivos de Salida

```yaml
database_name: "ifarmacia"
tables:
  - name: "CLIENTES"
    columns:
      - name: "ID"
        type: "INTEGER"
        nullable: false
        auto_increment: false
      - name: "NOMBRE"
        type: "VARCHAR"
        size: 100
        nullable: true
        default_value: null
    primary_key:
      - "ID"
    foreign_keys:
      - name: "FK_CLIENTE_CIUDAD"
        columns: ["CIUDAD_ID"]
        referenced_table: "CIUDADES"
        referenced_columns: ["ID"]
        on_delete: "RESTRICT"

views:
  - name: "V_CLIENTES_ACTIVOS"
    definition: "SELECT * FROM CLIENTES WHERE ACTIVO = 1"
    columns: []

procedures:
  - name: "SP_INSERTAR_CLIENTE"
    parameters:
      - name: "P_NOMBRE"
        type: "VARCHAR"
        direction: "IN"
    definition: "BEGIN ... END"

indexes:
  - name: "IDX_CLIENTE_NOMBRE"
    table_name: "CLIENTES"
    columns: ["NOMBRE"]
    unique: false
    type: "BTREE"

engine_specific:
  generators:
    - name: "GEN_CLIENTE_ID"
      current_value: 1000
  domains:
    - name: "D_BOOLEAN"
      type: "SMALLINT"
      check_constraint: "VALUE IN (0,1)"
```

## Manejo de Errores

### Archivos de Error

Si la extracción de una base de datos falla, se generará un archivo de error:

```yaml
# ejemplo: problematica.fdb.error.yaml
database_name: "problematica.fdb"
error: "failed to connect to database: connection refused"
timestamp: "2024-01-15T10:30:45Z"
```

### Códigos de Salida

- `0`: Éxito completo
- `1`: Error de configuración
- `2`: Error de extracción (todas las bases de datos fallaron)
- `3`: Error de generación de archivos de salida
- `4`: Éxito parcial (algunas bases de datos fallaron)

## Desarrollo

### Estructura del Proyecto

```
├── cmd/
│   └── main.go                 # Punto de entrada de la aplicación
├── internal/
│   ├── config/                 # Gestión de configuración
│   ├── extractor/              # Interfaces y estrategias de extracción
│   │   └── firebird/           # Implementación para Firebird
│   ├── orchestrator/           # Coordinador de procesamiento paralelo
│   ├── output/                 # Generación de archivos YAML
│   └── schema/                 # Modelos de datos del esquema
├── test/                       # Tests separados del código principal
└── config.yaml                 # Archivo de configuración
```

### Ejecutar Tests

```bash
# Ejecutar todos los tests unitarios
go run test/run_tests.go

# Ejecutar tests con cobertura
./test/coverage.sh

# Ejecutar tests de un paquete específico
go test -v ./test/config
```

### Agregar Soporte para Nuevos Motores de BD

1. Crear nueva implementación en `internal/extractor/`:

```go
// internal/extractor/sqlserver/extractor.go
type SQLServerExtractor struct {
    // implementación específica
}

func (se *SQLServerExtractor) Connect(connectionString string) error {
    // lógica de conexión para SQL Server
}

func (se *SQLServerExtractor) ExtractSchema() (*schema.DatabaseSchema, error) {
    // lógica de extracción para SQL Server
}
```

2. Crear factory correspondiente:

```go
// internal/extractor/sqlserver/factory.go
func (f *SQLServerExtractorFactory) SupportsConnectionString(connStr string) bool {
    return strings.Contains(connStr, "sqlserver://") || 
           strings.Contains(connStr, "server=")
}
```

3. Registrar en el orchestrator:

```go
// En main.go o donde se configure el orchestrator
orch.AddExtractorFactory(sqlserver.NewSQLServerExtractorFactory())
```

## Características Específicas por Motor

### Firebird

- **Generators**: Secuencias automáticas (equivalente a AUTO_INCREMENT)
- **Domains**: Tipos de datos personalizados con validaciones
- **Filtrado de tablas del sistema**: Excluye automáticamente tablas `RDB$*` y `MON$*`

### Futuras Implementaciones

- **SQL Server**: Soporte para schemas, funciones definidas por usuario, tipos personalizados
- **PostgreSQL**: Soporte para extensiones, tipos ENUM, arrays
- **MySQL**: Soporte para engines, particiones, triggers

## Troubleshooting

### Problemas Comunes

1. **"config.yaml file not found"**
   - Asegúrate de que el archivo `config.yaml` esté en el directorio actual
   - Verifica que el archivo tenga el formato YAML correcto

2. **"failed to connect to database"**
   - Verifica que el servidor de base de datos esté ejecutándose
   - Confirma que las credenciales y la ruta sean correctas
   - Para Firebird, asegúrate de que el puerto 3050 esté abierto

3. **"no suitable extractor found"**
   - Verifica que la cadena de conexión tenga el formato correcto para el motor soportado
   - Actualmente solo se soporta Firebird (archivos .fdb/.gdb o puerto 3050)

### Logging

La aplicación usa logging estructurado en formato JSON. Para ver logs más detallados:

```bash
# Ejecutar con nivel de log DEBUG (si se implementa)
LOG_LEVEL=debug ./schema-extractor
```

## Contribuir

1. Fork el repositorio
2. Crea una rama para tu feature (`git checkout -b feature/nueva-caracteristica`)
3. Commit tus cambios (`git commit -am 'Agregar nueva característica'`)
4. Push a la rama (`git push origin feature/nueva-caracteristica`)
5. Crea un Pull Request

## Licencia

[Especificar licencia aquí]

## Changelog

### v1.0.0
- Soporte inicial para Firebird
- Procesamiento paralelo
- Generación de archivos YAML
- Filtrado automático de tablas del sistema
- Extracción de generators y domains específicos de Firebird