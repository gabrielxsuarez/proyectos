# Índice de Documentación - DIPIDE

## Proyecto de Transformación Digital - Dirección Provincial de Impresión y Digitalización del Estado

---

## DOCUMENTOS PRINCIPALES

### 1. **CLAUDE.md**

**Resumen:** Guía técnica completa para Claude Code sobre el proyecto DIPIDE. Especifica arquitectura modular obligatoria con Flutter/Dart + PostgreSQL, define el proceso principal GO.001.001 de gestión de órdenes de trabajo con 8 tramos, establece 5 hitos críticos de desarrollo, metodología AGILE con 12 sprints en 6 meses, y detalla reglas de negocio para inventarios con persistencia completa. Incluye información del equipo de desarrollo (Claudio Paz, Gabriel Suárez, Adriana Aguirre) y PMO (Adrián Buratovich).

### 2. **resumen.md**

**Resumen:** Resumen ejecutivo completo del proyecto DIPIDE con objetivos generales de integración, requisitos técnicos obligatorios, arquitectura modular de 5 capas, cronograma detallado de 12 sprints organizados en 5 hitos principales. Define estructura de propuesta técnica con 6 objetivos generales y 10 requisitos específicos. Incluye factores críticos de éxito, riesgos principales, aspectos técnicos clarificados en reuniones, y contactos del equipo PMO y desarrollo.

---

## PROPUESTAS TÉCNICAS Y COMERCIALES

### 3. **propuesta/propuesta_comercial.md**

**Resumen:** Propuesta comercial del equipo de 3 desarrolladores con dos modalidades de contratación. Opción 1 (equipo completo): USD 4000/mes desarrollo, USD 3500/mes implementación, USD 1500/mes soporte. Opción 2 (individual): USD 3000/mes base + especializaciones opcionales. Incluye perfiles del equipo: Claudio Paz (18 años, Full Stack, Accenture), Gabriel Suárez (20 años, arquitectura/seguridad, Banco Hipotecario), Adriana Aguirre (7 años, análisis funcional/UX-UI, Dicsys).

### 4. **propuesta/propuesta_tecnica.md**

**Resumen:** Propuesta técnica detallada con stack Go Fiber + Flutter Web + PostgreSQL para cumplir 6 objetivos generales y 10 requisitos técnicos obligatorios. Incluye sistema semi-automatizado de CRUD mediante CSV, trazabilidad completa con auditoría automática, máquina de estados configurable, control cronológico de inventarios, sistema RBAC con contexto organizacional, y arquitectura preparada para web services REST-XML/JSON. Presenta demo funcional en dipide.gabrielsuarez.ar como concepto visual.

### 5. **propuesta/DER.md**

**Resumen:** Diagrama entidad-relación completo del sistema DIPIDE con 18 entidades principales implementando trazabilidad de expedientes y órdenes de trabajo, gestión de inventarios con persistencia completa, sistema RBAC contextual, auditoría automática, y configuración flexible. Incluye entidades clave: EXPEDIENTE, ORDEN_TRABAJO, MOVIMIENTO_WORKFLOW, MOVIMIENTO_STOCK, sistema de usuarios/roles, alertas automáticas, y estructura preparada para reportes configurables y escalabilidad futura.

### 6. **propuesta/tmp/preguntas_propuesta_tecnica.md**

**Resumen:** Documento de 23 categorías con 100+ preguntas técnicas específicas para desarrollar la propuesta completa DIPIDE. Cubre objetivos generales, requisitos técnicos, implementación, capacitación, migración, integración, operación, validación técnica, rendimiento, seguridad y contingencias. Incluye respuestas técnicas detalladas sobre stack tecnológico, volúmenes de datos, políticas de backup, autenticación, y recomendaciones de arquitectura simplificada con Go Fiber + PostgreSQL + Flutter.

---

## REUNIONES Y ACLARACIONES

### 7. **reuniones/reunion-20250823.md**

**Resumen:** Transcripción completa de reunión técnica de clarificación con 5 hitos críticos: (1) Administración BD con ABM semi-automatizado vía CSV, (2) Movimientos con trazabilidad de expedientes/órdenes y estados diferenciados, (3) Estados de inventario con diferenciación materiales/insumos y control cronológico de stock, (4) Campos de reposición diferidos, (5) Interfaces modulares con roles organizacionales. Incluye discusión sobre infraestructura (servidor único sin backup), contexto histórico de problemas anteriores, y definición de responsabilidades entre PMO y equipo técnico.

---

## DOCUMENTACIÓN METODOLÓGICA SERIAN IO

### 8. **documentacion/Desarrollador XXXXX - Propuesta Comercial y Técnica - DIPIDE (Soluciones Informáticas Modulares) - 20250812-19450001.md**

**Resumen:** Plantilla oficial de propuesta comercial y técnica para desarrolladores. Define estructura de costos (base, cargos adicionales, bonificaciones), criterios de inclusión obligatorios (no ser empleado Estado/DIPIDE/proyectos UNLaM), y anexo de recopilación analítica de propuestas técnicas. Documento template del PMO Adrián Sergio Buratovich para estandarizar presentación de propuestas al proyecto de transformación digital DIPIDE.

### 9. **documentacion/Serian IO - DIPIDE (EPM - 00.02 - Resumen de Proyecto - Integración de Servicios de Apoyo EPM - PM) - 20250606-14300001.md**

**Resumen:** Propuesta de servicios PMO externos Serian-IO para DIPIDE. Define oficina de gestión de procesos y proyectos facilitando transformación digital, articulación entre unidades funcionales y prestadores de servicios. Establece grupos de trabajo para integrantes DIPIDE, equipos interdisciplinarios y desarrolladores de soluciones modulares con enfoque en gestión del cambio organizacional.

### 10. **documentacion/Serian IO - DIPIDE (EPM - 00.03 - Resumen de Proyecto - Integración de Servicios de Desarrollo Informático) - 20250731-19300001.md**

**Resumen:** Marco de integración de prestadores de servicios de desarrollo informático para mejora de procesos del Departamento de Logística y Depósito DIPIDE. Especifica requisitos técnicos mandatorios (Flutter/Dart, servidor web La Plata), modalidades de facturación, criterios de trabajo, y proceso estructurado de presentación de propuestas comerciales y técnicas con coordinación PMO.

### 11. **documentacion/Serian IO - DIPIDE (EPM - 01 - Criterios Generales de Gestión de Portafolios de Proyectos) - 20250731-19300001.md**

**Resumen:** Marco metodológico integral para gestión de portafolio de proyectos estratégicos DIPIDE con gestión del cambio y conocimiento. Define 4 áreas de gestión estratégica, planificación 18 meses, proceso GO.001.001 desde solicitud hasta entrega, 6 objetivos generales de integración de sistemas informáticos, y requisitos técnicos detallados para soluciones modulares escalables (Flutter/Dart, SQL, HTTPS, REST).

### 12. **documentacion/Serian IO - DIPIDE (EPM - 02 - Recopilación Analítica de Procesos) - 20250731-19300001.md**

**Resumen:** Modelo de recopilación analítica de procesos organizacionales según metodología Serian-IO (adaptación BPM). Clasifica procesos en 6 categorías: estratégicos, directivos, gestión, operativos, soporte, interorganizacionales. Programa de revisión con distribución responsabilidades RACI y backlog detallado para proceso de gestión de órdenes de trabajo con trazabilidad completa desde solicitud hasta entrega.

### 13. **documentacion/Serian IO - DIPIDE (EPM - 03 - Recopilación Analítica de Propuestas Técnicas) - 20250731-19300001.md**

**Resumen:** Marco metodológico para recopilación de propuestas técnicas de soluciones informáticas modulares. Articula proyectos de gestión del cambio con transformación digital, establece backlog general de integración de soluciones informáticas, describe programa de integración con distribución de responsabilidades entre unidades funcionales y coordinación entre grupos de trabajo organizacionales y equipos de proyectos técnicos.

### 14. **documentacion/Serian IO - DIPIDE (EPM - 07 - Reporte de seguimiento de backlog de integración - Soluciones informáticas modulares) - 20250818.md**

**Resumen:** Reporte de seguimiento del backlog de integración proyecto TDIG (Transformación Digital) con 6 objetivos generales de sistemas informáticos: datos de soporte, trazabilidad, estados de procesos, inventarios, reposición, usabilidad. Define 10 requisitos técnicos específicos (Flutter/Dart, SQL, servidor web, REST, interfaces modulares) con planificación por sprints y estados de desarrollo para versión 01 del sistema modular escalable.

---

## Información de Contacto

**Adrián Sergio Buratovich**  
📧 as.buratovich@gmail.com  
📞 +54 9 11 4094-3221  
🔗 [LinkedIn](http://www.linkedin.com/in/adrian-buratovich)  
🌐 [Sitio Web](http://www.serianweb.com.ar)

---

## Contexto del Proyecto

### Alcance Técnico
- **Frontend:** Flutter/Dart (multiplataforma)
- **Backend:** Servidor web con REST API (XML/JSON)
- **Base de datos:** SQL relacional (compatible PostgreSQL)
- **Despliegue:** Servidor HTTPS físico en La Plata, Buenos Aires
- **Arquitectura:** Sistema modular escalable

### Objetivo Principal
Desarrollo e implementación de soluciones informáticas modulares para la gestión integral de órdenes de trabajo, inventarios y procesos documentales del Departamento de Logística y Depósito de DIPIDE.

---

*Documentación generada: Agosto 2025*