# Requirements Document

## Introduction

Esta aplicación en Go extraerá esquemas de bases de datos de forma paralela desde múltiples fuentes configuradas en un archivo YAML. La aplicación implementará un patrón strategy para soportar diferentes motores de bases de datos (inicialmente Firebird, con planes para SQL Server y otros). Los esquemas extraídos se exportarán como archivos YAML estructurados con nombres en inglés y excluyendo tablas del sistema.

## Requirements

### Requirement 1

**User Story:** Como desarrollador, quiero configurar múltiples conexiones de bases de datos en un archivo config.yaml, para poder extraer esquemas de diferentes bases de datos de forma centralizada.

#### Acceptance Criteria

1. WHEN la aplicación inicia THEN SHALL leer un archivo config.yaml del directorio actual
2. WHEN el archivo config.yaml contiene cadenas de conexión THEN SHALL parsear cada entrada con formato "nombre_db: cadena_conexion"
3. IF el archivo config.yaml no existe THEN SHALL mostrar un error descriptivo y terminar
4. WHEN las cadenas de conexión tienen formato "usuario:password@host:puerto/ruta" THEN SHALL extraer correctamente cada componente

### Requirement 2

**User Story:** Como desarrollador, quiero que la aplicación se conecte a múltiples bases de datos en paralelo, para optimizar el tiempo de extracción cuando tengo muchas bases de datos.

#### Acceptance Criteria

1. WHEN la aplicación procesa las conexiones THEN SHALL crear conexiones paralelas para cada base de datos
2. WHEN una conexión falla THEN SHALL continuar procesando las demás conexiones
3. WHEN todas las conexiones terminan THEN SHALL reportar el estado de cada una
4. IF una base de datos no responde THEN SHALL aplicar timeout y continuar con las demás

### Requirement 3

**User Story:** Como desarrollador, quiero extraer la estructura completa de cada base de datos (tablas, columnas, índices, stored procedures), para tener un inventario completo del esquema.

#### Acceptance Criteria

1. WHEN se conecta a una base de datos THEN SHALL extraer información de tablas de usuario únicamente
2. WHEN extrae tablas THEN SHALL obtener nombres, columnas con tipos de datos, constraints y índices
3. WHEN extrae stored procedures THEN SHALL obtener nombres, parámetros y definiciones
4. WHEN extrae índices THEN SHALL obtener nombres, columnas involucradas y tipo de índice
5. IF existen características específicas del motor THEN SHALL incluirlas en sección separada

### Requirement 4

**User Story:** Como desarrollador, quiero que la aplicación use el patrón strategy para diferentes motores de bases de datos, para poder agregar fácilmente soporte para nuevos motores sin modificar el código principal.

#### Acceptance Criteria

1. WHEN la aplicación identifica el tipo de base de datos THEN SHALL usar la estrategia correspondiente
2. WHEN se implementa soporte para Firebird THEN SHALL extraer generators y otras características específicas
3. WHEN se necesite agregar SQL Server THEN SHALL poder implementar nueva estrategia sin cambiar código existente
4. IF el motor de base de datos no es soportado THEN SHALL mostrar error descriptivo

### Requirement 5

**User Story:** Como desarrollador, quiero que los esquemas extraídos se guarden como archivos YAML con nombres en inglés, para tener un formato estándar y legible que pueda usar en otras herramientas.

#### Acceptance Criteria

1. WHEN la extracción termina exitosamente THEN SHALL generar archivo YAML con nombre "nombre_db.yaml"
2. WHEN estructura el YAML THEN SHALL usar nombres en inglés (tables, columns, indexes, procedures)
3. WHEN incluye información de columnas THEN SHALL especificar name, type, nullable, default_value
4. WHEN incluye características específicas del motor THEN SHALL agruparlas en sección "engine_specific"
5. IF la extracción falla THEN SHALL generar archivo de error con detalles del problema

### Requirement 6

**User Story:** Como desarrollador, quiero que la aplicación filtre automáticamente las tablas del sistema, para obtener solo las tablas creadas por usuarios en los archivos de salida.

#### Acceptance Criteria

1. WHEN extrae tablas de Firebird THEN SHALL excluir tablas que empiecen con "RDB$" y "MON$"
2. WHEN extrae de cualquier motor THEN SHALL aplicar filtros específicos para tablas del sistema
3. WHEN genera el archivo YAML THEN SHALL incluir solo tablas, vistas y objetos de usuario
4. IF no hay tablas de usuario THEN SHALL generar archivo YAML vacío con mensaje informativo