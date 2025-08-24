# RESUMEN EJECUTIVO - PROYECTO DIPIDE
## Sistema de Soluciones Informáticas Modulares

---

## 1. DESCRIPCIÓN GENERAL DEL PROYECTO

### Organización Cliente
**DIPIDE** - Dirección Provincial de Impresión y Digitalización del Estado
- Ubicación: La Plata, Provincia de Buenos Aires
- Sector: Gobierno Provincial - Servicios de impresión y digitalización

### Contexto del Proyecto
Adrián Sergio Buratovich de **Serian IO** propone un modelo de **Oficina PMO Externa** para coordinar la transformación digital de DIPIDE a través del desarrollo de **soluciones informáticas modulares** que optimicen los procesos de gestión de órdenes de trabajo de impresión.

---

## 2. ALCANCE Y OBJETIVOS

### Objetivo Principal
Desarrollar un **sistema modular de gestión de inventarios y órdenes de trabajo** para optimizar el proceso completo "desde la solicitud de trabajos hasta la entrega de trabajos terminados" del Departamento de Logística/Depósito de DIPIDE.

### Proceso Principal a Optimizar
**GO.001.001 - GESTIÓN DE ÓRDENES DE TRABAJO (TRABAJOS DE IMPRESIÓN)**

**Tramos del Proceso:**
1. **TRAMO 01** - Inicio / Factibilidad Interna
2. **TRAMO 02** - Prerrequisitos
3. **TRAMO 03** - Programación
4. **TRAMO 04** - Coordinación y Gestión
5. **TRAMO 05** - Seguimiento, Monitoreo y Control
6. **TRAMO 06** - Finalización
7. **TRAMO 07** - [Reservado]
8. **TRAMO 08** - Incidencias

### Unidades Funcionales Involucradas
- **COMERCIALIZACIÓN**: Responsable inicio, prerrequisitos
- **PRODUCCIÓN**: Responsable programación, gestión, control
- **ALMACÉN**: Responsable seguimiento, control, finalización
- **TRANSPORTE**: Informado en todos los procesos

---

## 3. OBJETIVOS GENERALES DE INTEGRACIÓN

### 3.1 DATOS DE SOPORTE (CRUD)
- Gestión de entidades informáticas (maestros)
- Gestión de registros relacionales (tablas de vinculación)

### 3.2 TRAZABILIDAD DE PROCESOS
- Trazabilidad de expedientes
- Trazabilidad de órdenes de trabajo

### 3.3 ESTADOS DE PROCESOS
- Estados de gestión y control de expedientes
- Estados de gestión y control de órdenes

### 3.4 INVENTARIOS
- **Insumos**: Movimientos y estados
- **Materiales**: Movimientos y estados
- **Trabajos terminados**: Movimientos y estados

### 3.5 REPOSICIÓN DE INVENTARIO
- Planes de reposición de insumos
- Planes de reposición de materiales
- Planes de reposición de trabajos terminados

### 3.6 USABILIDAD DE RECURSOS
- Disponibilidad de recursos informáticos
- Control de acceso por credenciales y roles
- Navegabilidad adaptada a unidades funcionales

---

## 4. REQUISITOS TÉCNICOS OBLIGATORIOS
**(Los 10 requisitos que deben detallarse en la Propuesta Técnica - Anexo I, Parte 2)**

### 4.1 ENTORNO DE DESARROLLO (REQUISITO 01)
- **Tecnología**: Flutter/Dart
- **Plataformas**: Escritorio, Web, Móviles
- **Sistemas**: Windows, macOS, iOS, Linux, Android

### 4.2 BASE DE DATOS (REQUISITO 02)
- **Tipo**: Relacional SQL
- **Compatibilidad**: PostgreSQL
- **Escalabilidad**: Alta demanda futura

### 4.3 SERVIDOR WEB (REQUISITO 03)
- **Ubicación**: Servidor físico sede central DIPIDE (La Plata)
- **Protocolo**: HTTPS obligatorio
- **Redes**: Privadas y públicas
- **Escalabilidad**: Alta demanda futura

### 4.4 SERVICIOS WEB (REQUISITO 04)
- **Primario**: REST-XML (para documentos digitales XML-CSS)
- **Alternativo**: REST-JSON
- **Escalabilidad**: Alta demanda futura

### 4.5 INTERFACES DE USUARIO WEB RESPONSIVE (REQUISITO 05)
- **Tipo**: Web Responsive
- **Dispositivos**: Terminales de trabajo y móviles

### 4.6 MÓDULO PRINCIPAL (REQUISITO 06)
- Control centralizado de seguridad y acceso
- Gestión de configuraciones y navegabilidad por roles
- Persistencia de parámetros de configuración

### 4.7 MÓDULOS DE GESTIÓN ESPECIALIZADA (REQUISITO 07)
- Acceso heredado desde módulo principal
- Gestión de tareas especializadas por unidad funcional
- Configuración adaptable organizacional

### 4.8 MÓDULOS DE TAREAS ESPECIALIZADAS (REQUISITO 08)
- Realización de tareas específicas del proceso
- Control de acceso multi-nivel heredado
- Configuración por procesos y actividades

### 4.9 MÓDULOS DE ACCIONES ESPECÍFICAS (REQUISITO 09)
- Acciones puntuales y procesamiento en tiempo real
- Acceso a datos pre-procesados y procesados en tiempo real
- Configuración flexible de orígenes de datos

### 4.10 MÓDULOS DE ADMINISTRACIÓN DE DATOS (REQUISITO 10)
- Manipulación y procesamiento de información
- Interfaces CRUD sobre modelo de datos
- Cumplimiento de objetivos generales de integración

---

## 5. ARQUITECTURA MODULAR REQUERIDA
**(Detalles expandidos de los Requisitos 06-10)**

### 5.1 MÓDULO PRINCIPAL
- Control centralizado de seguridad y acceso
- Gestión de configuraciones
- Navegabilidad por roles organizacionales
- Configuración persistente local

### 5.2 MÓDULOS DE GESTIÓN ESPECIALIZADA
- Acceso heredado desde módulo principal
- Gestión de tareas especializadas
- Configuración adaptable por unidad funcional

### 5.3 MÓDULOS DE TAREAS ESPECIALIZADAS
- Realización de tareas específicas
- Control de acceso multi-nivel
- Configuración por procesos y actividades

### 5.4 MÓDULOS DE ACCIONES ESPECÍFICAS
- Acciones puntuales del sistema
- Procesamiento en tiempo real
- Configuración flexible de datos

### 5.5 MÓDULOS DE ADMINISTRACIÓN DE DATOS
- Manipulación y procesamiento de información
- Cumplimiento de objetivos generales
- Interfaces CRUD sobre modelo de datos

---

## 6. PLANIFICACIÓN POR FASES (ACTUALIZADA SEGÚN REUNIÓN 23/08/2025)

### FASE 01 - DESARROLLO
- **Duración**: 6 meses
- **Metodología**: AGILE/SCRUM
- **Organización**: 12 sprints (2 semanas cada uno)
- **Revisión de Backlog**: Primeros 3 meses comprimidos para evitar imprevistos
- **Implementación**: Inicia en el 4to mes (mediados de camino)

### FASE 02 - IMPLEMENTACIÓN
- **Duración**: 3 meses
- **Enfoque**: Implementación general e integración
- **Flexibilidad**: Manejo adaptable según avance del desarrollo

### FASE 03 - SOPORTE Y TRANSFERENCIA
- **Duración**: Mensual continuo
- **Actividades**: Soporte técnico, transferencia a operaciones

---

## 7. CRONOGRAMA DE SPRINTS (PROYECTO TDIG AR-20250801-04.02)

### PLANIFICACIÓN DETALLADA - 6 MESES / 12 SPRINTS

| Mes | Sprint | Enfoque | Prioridad |
|-----|--------|---------|-----------|
| **MES 1** | Sprint 01 | Módulos principales y configuración base | Alta |
| | Sprint 02 | Gestión de datos maestros | Alta |
| **MES 2** | Sprint 03 | Trazabilidad de procesos | Alta |
| | Sprint 04 | Estados de gestión | Alta |
| **MES 3** | Sprint 05 | Inventarios - Insumos | Alta |
| | Sprint 06 | Inventarios - Materiales | Media |
| **MES 4** | Sprint 07 | Inventarios - Trabajos terminados | Media |
| | Sprint 08 | Planes de reposición | Media |
| **MES 5** | Sprint 09 | Interfaces de usuario especializadas | Media |
| | Sprint 10 | Módulos de acciones específicas | Baja |
| **MES 6** | Sprint 11 | Integración y usabilidad | Baja |
| | Sprint 12 | Testing, documentación y transferencia | Baja |

---

## 8. ENTREGABLES POR SPRINT

### SPRINTS 1-2 (MES 1) - HITO 1: ADMINISTRACIÓN BD
- Arquitectura base del sistema modular
- Módulo principal con autenticación interna simple
- **Importación/exportación CSV configurable**
- **ABM semi-automatizado por lotes**
- Configuración inicial de base de datos

### SPRINTS 3-4 (MES 2) - HITO 2: MOVIMIENTOS
- **Sistema de trazabilidad por expedientes/órdenes**
- **Estados de gestión y control diferenciados**
- **Persistencia completa sin eliminación de movimientos**
- Interfaces de gestión especializada

### SPRINTS 5-6 (MES 3) - HITO 3: ESTADOS INVENTARIO
- **Módulos de inventario: Materiales (prioridad alta), Insumos (media)**
- **Estados diferenciados: disponible/reservado/en proceso**
- **Control de stock con punto de partida y movimientos cronológicos**
- Consultas y reportes básicos

### SPRINTS 7-8 (MES 4) - IMPLEMENTACIÓN INICIA
- **Inventario trabajos terminados (prioridad baja)**
- **Órdenes provisorias y reservas**
- Integración con unidades funcionales
- **(Hito 4: Reposición diferido para fases posteriores)**

### SPRINTS 9-10 (MES 5) - HITO 5: INTERFACES
- **Interfaces web responsive (NO móviles requeridas inicialmente)**
- **Módulos por roles organizacionales**
- **Arquitectura preparada para web services futuros**
- Optimización UX/UI

### SPRINTS 11-12 (MES 6) - FINALIZACIÓN
- Integración final de módulos
- Testing integral y corrección de bugs
- Documentación técnica y funcional
- **Capacitación usuarios finales (crítico por nivel técnico bajo)**

---

## 9. ESTRUCTURA DE LA PROPUESTA TÉCNICA

### 9.1 DOCUMENTO PRINCIPAL
Basado en el template "**Desarrollador XXXXX - Propuesta Comercial y Técnica - DIPIDE**":

#### SECCIÓN 1: INFORMACIÓN COMERCIAL
- Nombre completo del desarrollador
- Correo electrónico de contacto
- Propuesta comercial (costos mensuales USD)
  - Costo base general: USD ____ / Mes
  - Cargos adicionales previstos: USD ____ / Mes
  - Bonificaciones previstas: USD ____ / Mes
- Observaciones comerciales
- Confirmación de criterios de inclusión:
  - ✓ No ser empleado del Estado de Buenos Aires
  - ✓ No ser empleado de DIPIDE
  - ✓ No haber trabajado en proyectos Universidad de La Matanza para DIPIDE

#### SECCIÓN 2: PROPUESTA TÉCNICA (ANEXO I)
**"RECOPILACIÓN ANALÍTICA DE PROPUESTAS TÉCNICAS"**

### 9.2 ANEXO I - ESTRUCTURA TÉCNICA OBLIGATORIA

#### PARTE 1: OBJETIVOS GENERALES DE INTEGRACIÓN
Para cada uno de los 6 objetivos generales, especificar:
- Alcance técnico propuesto
- Metodología de implementación
- Tecnologías específicas
- Criterios de aceptación
- Estimación de esfuerzo (horas/días)

#### PARTE 2: REQUISITOS GENERALES DE INTEGRACIÓN
Para cada uno de los **10 requisitos técnicos obligatorios** (detallados en sección 4), especificar:

**Los 10 Requisitos Técnicos a detallar son:**
1. **REQUISITO 01** - Entorno de Desarrollo (Flutter/Dart)
2. **REQUISITO 02** - Base de Datos (Relacional SQL/PostgreSQL)
3. **REQUISITO 03** - Servidor Web (HTTPS, La Plata)
4. **REQUISITO 04** - Servicios Web (REST-XML/JSON)
5. **REQUISITO 05** - Interfaces Web Responsive
6. **REQUISITO 06** - Módulo Principal (Control centralizado)
7. **REQUISITO 07** - Módulos de Gestión Especializada
8. **REQUISITO 08** - Módulos de Tareas Especializadas
9. **REQUISITO 09** - Módulos de Acciones Específicas
10. **REQUISITO 10** - Módulos de Administración de Datos

**Para cada requisito, detallar:**
- Solución técnica propuesta
- Herramientas y frameworks específicos
- Configuraciones y ajustes necesarios
- Plan de testing y validación
- Documentación entregable

### 9.3 SECCIONES ADICIONALES RECOMENDADAS

#### EXPERIENCIA Y REFERENCIAS
- Portfolio de proyectos similares
- Tecnologías dominadas (Flutter/Dart, PostgreSQL, etc.)
- Experiencia en sistemas modulares
- Referencias de clientes anteriores

#### METODOLOGÍA DE TRABAJO
- Proceso de desarrollo AGILE/SCRUM
- Herramientas de gestión de proyectos
- Comunicación y reportes
- Gestión de cambios y requerimientos adicionales

#### PLAN DE IMPLEMENTACIÓN DETALLADO
- Cronograma específico por sprint
- Dependencias entre módulos
- Hitos de validación con cliente
- Plan de riesgos y contingencias

#### SOPORTE Y MANTENIMIENTO
- Modelo de soporte post-implementación
- Actualizaciones y mejoras
- Transferencia de conocimiento
- Documentación técnica y funcional

---

## 10. OBSERVACIONES CRÍTICAS (ACTUALIZADAS SEGÚN REUNIONES)

### 10.1 FACTORES DE ÉXITO
- **Cumplimiento estricto** de requisitos técnicos (Flutter/Dart + PostgreSQL)
- **Arquitectura modular** escalable y mantenible
- **Integración perfecta** con procesos organizacionales existentes
- **Capacitación efectiva** a usuarios finales
- **Intermediación efectiva** con Director General de DIPIDE
- **Gestión de cambio organizacional** acompañada de capacitación

### 10.2 RIESGOS PRINCIPALES
- Complejidad de integración con sistemas legacy (Excel distribuido y desordenado)
- Cambios en requerimientos durante desarrollo
- Coordinación entre múltiples unidades funcionales
- Adopción de tecnología por usuarios finales (algunos sin conocimiento básico de Excel)
- **Infraestructura limitada**: Un solo servidor físico sin backup
- **Recursos técnicos limitados**: Área técnica no dedicada 24/7

### 10.3 RECOMENDACIONES
- Implementar **prototipo funcional** en Sprint 2
- **Validaciones incrementales** con usuarios clave
- **Testing paralelo** durante desarrollo
- **Documentación continua** de decisiones técnicas
- **Gestión de importación/exportación CSV** desde el primer hito
- **Trazabilidad completa** sin eliminación de movimientos
- **Backup responsabilidad del cliente** claramente establecida

### 10.4 HITOS ESPECÍFICOS CONFIRMADOS (5 HITOS PRINCIPALES)
1. **Administración de base de datos** - Importación/exportación CSV, ABM semi-automatizado
2. **Gestión de movimientos** - Trazabilidad por expedientes/órdenes, estados de gestión
3. **Estados de inventario** - Diferenciación materiales/insumos, reservas y disponibilidad
4. **Campos de reposición** - Stock mínimos/máximos (diferido para fases posteriores)
5. **Interfaces modulares** - Arquitectura por roles organizacionales, preparación para web services

### 10.5 ASPECTOS TÉCNICOS CLARIFICADOS
- **Gestión de usuarios**: Sistema interno simple, sin integración Active Directory
- **Servidor**: Físico en DIPIDE La Plata, sin ambiente de desarrollo separado
- **Formatos**: CSV configurables, separadores definibles (tabulador preferido)
- **Dashboard tiempo real**: NO es requisito (puede ser opcional futuro)
- **Web services**: REST preparado para futuro, XML para documentos formales

---

**Contacto Coordinación PMO Externa:**
- **Adrián Sergio Buratovich**
- Email: as.buratovich@gmail.com
- Tel: +54 9 11 4094-3221
- LinkedIn: http://www.linkedin.com/in/adrian-buratovich
- Web: http://www.serianweb.com.ar