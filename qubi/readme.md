# Guía de Estilos - Go

Esta guía define las convenciones de código para proyectos Go, enfocada en mantener consistencia y legibilidad.

## Estructura General

### Organización del Código
- Utilizar comentarios de sección con formato `/* ========== NOMBRE_SECCION ========== */`
- Agrupar funciones relacionadas bajo la misma sección
- Mantener el orden: MAIN, ESTRUCTURAS, funciones específicas por funcionalidad

### Ejemplo de Secciones
```go
/* ========== MAIN ========== */
/* ========== ESTRUCTURAS ========== */
/* ========== FUNCIONALIDAD_PRINCIPAL ========== */
/* ========== UTILIDADES ========== */
/* ========== HELPERS ========== */
```

## Convenciones de Nomenclatura

### Funciones
- **Funciones públicas**: PascalCase (ej: `MenuItem`)
- **Funciones privadas**: camelCase con nombres descriptivos en español

### Variables
- **Variables locales**: camelCase en español (ej: `menuPadre`, `nuevoElemento`)
- **Constantes implícitas**: camelCase descriptivo (ej: `esArchivo`, `esTxt`)

### Estructuras
- **Nombres de struct**: PascalCase descriptivo
- **Campos de struct**: PascalCase, preferiblemente en el idioma principal del proyecto

## Manejo de Errores

### Patrón Estándar
```go
if err != nil {
    return fmt.Errorf("descripción del error: %w", err)
}
```

### Mensajes de Error
- Usar `fmt.Errorf` con wrapping (`%w`) para preservar la cadena de errores
- Mensajes descriptivos en español
- Incluir contexto específico del error

### Ejemplos
```go
return fmt.Errorf("error al cargar el icono: %w", err)
return fmt.Errorf("error al leer el archivo %s: %w", archivo, err)
```

## Estilo de Código

### Condicionales
- Usar variables booleanas descriptivas para mejorar legibilidad:
```go
esArchivo := !archivo.IsDir()
esTxt := strings.HasSuffix(archivo.Name(), ".txt")
contieneMenu := strings.Contains(archivo.Name(), "menu")

if esArchivo && esTxt && contieneMenu {
    // lógica
}
```

### Bucles
- Preferir `range` sobre índices cuando sea posible
- Usar `for range` para canales:
```go
for range menuItem.ClickedCh {
    // lógica
}
```

## Gorrutinas y Concurrencia

### Patrón para Event Handlers
```go
go func() {
    for range item.ClickedCh {
        // lógica del evento
    }
}()
```

### Patrón para Eventos Únicos
```go
go func() {
    <-salirItem.ClickedCh
    systray.Quit()
}()
```

## Funciones Auxiliares

### Principio de Responsabilidad Única
- Cada función debe tener una responsabilidad específica
- Extraer lógica compleja en funciones separadas
- Usar nombres descriptivos que expliquen la acción

### Ejemplos de Buena Separación
```go
func procesarElemento(elemento Tipo) error
func validarDatos(datos []string) bool
func ejecutarAccion(accion string, parametros []string)
```

## Compatibilidad Multiplataforma

### Detección de SO
```go
if runtime.GOOS == "windows" {
    // lógica específica de Windows
} else {
    // lógica para otros sistemas
}
```

### Rutas de Archivos
- Usar rutas relativas desde el directorio del ejecutable
- Separar lógica específica del SO cuando sea necesario

## Logging y Debug

### Mensajes Informativos
```go
fmt.Printf("Ejecutando comando: %s con argumentos: %v\n", partes[0], partes[1:])
```

### Mensajes de Error
```go
fmt.Println("Error al ejecutar el comando:", err)
```

## Comentarios

### Comentarios de Sección
- Usar el formato establecido con `=` para delimitar secciones
- Nombres de sección en mayúsculas
- Espaciado consistente

### Comentarios de Código
- Evitar comentarios obvios
- Comentar lógica compleja o no intuitiva, sin embargo es preferible partirla en funciones mas pequeñas cuando sea posible
- Usar español para consistencia con el resto del código

## Imports

### Organización
- Imports estándar primero
- Imports de terceros después
- Agrupar imports relacionados
- Usar nombres descriptivos cuando sea necesario

### Ejemplo
```go
import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/usuario/paquete-externo"
)
```

## Principios Generales

1. **Consistencia**: Mantener el mismo estilo en todo el proyecto
2. **Legibilidad**: Priorizar código claro sobre código "inteligente"
3. **Simplicidad**: Evitar complejidad innecesaria
4. **Idioma**: Usar español para nombres de variables y funciones internas, excepto para términos técnicos establecidos en inglés
5. **Separación de responsabilidades**: Una función, una tarea
6. **Manejo robusto de errores**: Siempre verificar y propagar errores apropiadamente
7. **Uso de idiomas**: Preferir español en general, pero usar inglés para términos técnicos ampliamente aceptados (ej: "Factory" en lugar de "Fábrica", "Handler" en lugar de "Manejador")

---

*Esta guía debe seguirse para mantener la coherencia del código y facilitar el mantenimiento futuro de cualquier proyecto Go.*