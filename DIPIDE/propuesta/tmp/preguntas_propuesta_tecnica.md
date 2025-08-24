# PREGUNTAS PARA PROPUESTA TÉCNICA - DIPIDE

Este documento contiene preguntas específicas para completar cada sección de la propuesta técnica según los requerimientos informales del proyecto DIPIDE.

---

## SECCIÓN 1: OBJETIVOS GENERALES

### DATOS DE SOPORTE (CRUD)

#### OBJETIVO GENERAL 01.01 ~ GESTIÓN DE DATOS DE ENTIDADES INFORMÁTICAS
**Pregunta 1.1.1:** ¿Cuál será la estrategia específica para el manejo de importación/exportación CSV? ¿Qué separadores de campo utilizaremos (tabulador, coma, punto y coma)?
**Respuesta:** Se implementará un sistema configurable de importación/exportación CSV usando tabulador como separador principal (para evitar conflictos con descripciones que contengan comas). Se desarrollará un endpoint REST en Go Fiber que procese archivos CSV con validación de estructura y encoding UTF-8. La funcionalidad incluirá preview de datos antes de importar y logs detallados de errores. Dado que DIPIDE actualmente utiliza excel, los separadores seran los que provea excel por defecto. De momento no se conocen las estructuras de estos csv.

**Pregunta 1.1.2:** Para el ABM semi-automatizado por lotes, ¿qué tamaño máximo de lote consideraremos? ¿Habrá validación de integridad referencial durante la importación?
**Respuesta:** Tamaño máximo de lote: desconocido pero la aplicacion debe poder manejar algunas decenas de miles sin complejizar la arquitectura. Se implementarán transacciones PostgreSQL para garantizar atomicidad y validación completa de llaves foráneas antes de confirmar cambios. Sistema de rollback automático ante errores con reportes detallados.

**Pregunta 1.1.3:** ¿Cómo manejaremos los conflictos durante la importación (IDs duplicados, referencias inexistentes)?
**Respuesta:** El sistema debe priorizar la consistencia de la base de datos, por lo que si la consistencia peligra es preferible fallar la importacion y que el usuario arregle lo que deba arreglar

**Pregunta 1.1.4:** ¿Qué entidades maestras específicas identificamos hasta ahora? (productos, proveedores, usuarios, unidades organizacionales, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible.

#### OBJETIVO GENERAL 01.02 ~ GESTIÓN DE DATOS DE REGISTROS RELACIONALES
**Pregunta 1.2.1:** ¿Cuáles son las tablas de vinculación más críticas? (producto-proveedor, usuario-rol, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 1.2.2:** ¿Cómo garantizaremos la consistencia referencial cuando se actualicen registros relacionales por lotes?
**Respuesta:** Uso de transacciones PostgreSQL con nivel de aislamiento SERIALIZABLE. Implementación de triggers de base de datos para validación automática y procedures almacenados para operaciones complejas. Sistema de locks optimistas para prevenir condiciones de carrera.

### TRAZABILIDAD DE PROCESOS

#### OBJETIVO GENERAL 02.01 ~ TRAZABILIDAD DE GESTIÓN (EXPEDIENTES)
**Pregunta 2.1.1:** ¿Cuál es el formato específico de los números de expediente? ¿Hay algún estándar del GDE provincial que debamos seguir?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 2.1.2:** ¿Qué timestamps específicos necesitamos registrar? (creación, modificación, cambios de estado, etc.)
**Respuesta:** Se registrarán timestamps con zona horaria (timestamptz de PostgreSQL) para: fecha_creacion, fecha_actualizacion, nuevo_estado, cada movimiento de inventario, y acciones de usuario. Sistema de auditoría completo usando triggers automáticos y tablas de log con retención configurable.

**Pregunta 2.1.3:** ¿Habrá integración futura con el sistema GDE provincial o será completamente independiente?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

#### OBJETIVO GENERAL 02.02 ~ TRAZABILIDAD DE GESTIÓN (ÓRDENES)
**Pregunta 2.2.1:** ¿Cuál es la relación exacta entre expedientes y órdenes? ¿Un expediente puede tener múltiples órdenes?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 2.2.2:** ¿Cómo se generarán los números de orden de trabajo? ¿Secuencial, con prefijos, por período?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

### ESTADOS DE PROCESOS

#### OBJETIVO GENERAL 03.01/03.02 ~ ESTADOS DE GESTIÓN Y CONTROL
**Pregunta 3.1.1:** ¿Cuáles son los estados de gestión específicos para expedientes? (ej: Recibido, En Análisis, Aprobado, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 3.1.2:** ¿Cuáles son los estados de gestión específicos para órdenes? (ej: Provisoria, Confirmada, En Producción, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 3.1.3:** ¿Cuáles son los estados de control? (ej: Pendiente Aprobación, Aprobado, Rechazado, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 3.1.4:** ¿Habrá restricciones en las transiciones de estado? ¿Un diagrama de flujo de estados?
**Respuesta:** Se implementará una máquina de estados usando enum types de PostgreSQL con validación automática de transiciones permitidas mediante triggers. Los diagramas de estados se configurarán en JSON y se validarán en el backend Go Fiber antes de cada cambio de estado.

### GESTIÓN DE MOVIMIENTOS Y ESTADO DE INVENTARIO

#### OBJETIVO GENERAL 04.01/04.02/04.03 ~ INVENTARIOS
**Pregunta 4.1.1:** ¿Cuál es la diferencia práctica entre "insumos" y "materiales" en DIPIDE? ¿Ejemplos concretos?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 4.1.2:** ¿Qué tipos de movimiento específicos manejan actualmente? (ingreso por compra, egreso por producción, devolución, ajuste, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 4.1.3:** ¿Cómo manejan actualmente el control de stock? ¿Con qué frecuencia hacen inventarios físicos?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 4.1.4:** Para "trabajos terminados", ¿se considera inventario de productos listos para entrega o trabajos en diferentes etapas?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 4.1.5:** ¿Qué estados de inventario específicos necesitan? (disponible, reservado, en proceso, defectuoso, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible. Cada cambio de estado generará un registro de auditoría automático con timestamp y usuario responsable.

### PLANES REPOSICIÓN DE INVENTARIO

#### OBJETIVO GENERAL 05.01/05.02/05.03 ~ REPOSICIÓN
**Pregunta 5.1.1:** ¿Actualmente tienen criterios de stock mínimo/máximo definidos?
**Respuesta:** No se conoce, la aplicacion debe ser configurable por un administrador.

**Pregunta 5.1.2:** ¿Cómo se calculará el punto de reposición? ¿Basado en consumo histórico, demanda proyectada?
**Respuesta:** Se podrán definir alertas por cantidad de stock, lo del consumo historico y la demanda proyectada puede ser interesante pero lo dejaria como opcional.

**Pregunta 5.1.3:** ¿La reposición será automática (alertas) o manual?
**Respuesta:** Sistema híbrido: alertas automáticas mediante jobs programados en Go Fiber que revisen stock mínimo diariamente y generen notificaciones push/email. Las órdenes de reposición serán manuales con aprobación de supervisor pero con datos pre-calculados.

**Pregunta 5.1.4:** ¿Qué información de proveedores necesitamos para los planes de reposición?
**Respuesta:**No se conoce, la propuesta debe ser flexible.

### USABILIDAD DE RECURSOS INFORMÁTICOS

#### OBJETIVO GENERAL 06.01/06.02/06.03 ~ USABILIDAD
**Pregunta 6.1.1:** ¿Cuáles son los roles organizacionales específicos? (Jefe de Almacén, Operario, Administrador, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 6.1.2:** ¿Qué nivel de granularidad necesitamos en los permisos? (por funcionalidad, por datos, por unidad organizacional?)
**Respuesta:** Sistema de permisos basado en RBAC (Role-Based Access Control) con granularidad a nivel de recurso y operación. Se implementarán middleware de autorización en Go Fiber que validen permisos antes de cada endpoint. Los roles se configurarán en base de datos con herencia jerárquica.

**Pregunta 6.1.3:** ¿Habrá usuarios que trabajen en múltiples unidades funcionales?
**Respuesta:** No se conoce, la propuesta debe ser flexible, por lo que asumiremos que si.

---

## SECCIÓN 2: REQUISITOS GENERALES

### ENTORNO DE DESARROLLO

**Pregunta 7.1:** ¿Qué versión específica de Flutter/Dart planificamos utilizar? ¿Hay restricciones de compatibilidad?
**Respuesta:** Es requisito usar flutter, como no se especifica version usaremos la ultima.

**Pregunta 7.2:** ¿Necesitamos compatibilidad con versiones específicas de navegadores web?
**Respuesta:** Con que funcione en chrome es suficiente.

**Pregunta 7.3:** ¿Hay algún framework UI específico de Flutter que prefiera DIPIDE o que debamos evitar?
**Respuesta:** No

**Pregunta 7.4:** ¿El desarrollo inicial se enfocará en web responsive o también incluirá aplicaciones nativas desde el primer sprint?
**Respuesta:** Inicialmente quieren una pagina web responsive en flutter, de momento no se necesitan versiones de android y ios, sin embargo en el futuro inmediato a la implementacion de este proyecto, planean evolucionar la aplicacion con soporte mobile.

### BASE DE DATOS

**Pregunta 8.1:** ¿Hay alguna versión específica de PostgreSQL que debamos usar o evitar?
**Respuesta:** No se especifica version por lo que usaremos la ultima.

**Pregunta 8.2:** ¿Cuáles son las estimaciones de volumen de datos? (registros por tabla, crecimiento anual esperado)
**Respuesta:** No se conoce, pero no creo que el volumen sea un problema, de momento priorizaremos la simplicidad y si luego esto crece demasiado nos adaptaremos. Aunque si se pueden tener en cuenta la escalabilidad sin agregar complejidad lo haremos desde el primer momento.

**Pregunta 8.3:** ¿Necesitamos considerar particionamiento de tablas desde el diseño inicial?
**Respuesta:** No inicialmente. Se diseñarán tablas con posibilidad futura de particionamiento por fecha en tablas de movimientos e historiales. Se implementarán índices apropiados desde el inicio.

**Pregunta 8.4:** ¿Habrá necesidad de réplicas de lectura o backup automático desde la aplicación?
**Respuesta:** Backup automático mediante pg_dump programado con cron/systemd timers. No réplicas de lectura inicialmente pero arquitectura preparada para PostgreSQL streaming replication en el futuro.

**Pregunta 8.5:** ¿Qué estrategia de backup requiere DIPIDE? ¿Con qué frecuencia?
**Respuesta:** No se conoce, no hay servidor de respaldo, por lo que los backups automaticos estaran en el mismo servidor, deberiamos darle la posibilidad al usuario de hacer un restore y de descargar el backup.

### SERVIDOR WEB

**Pregunta 9.1:** ¿Cuáles son las especificaciones exactas del servidor físico? (CPU, RAM, almacenamiento)
**Respuesta:** No se conoce, pero esperamos que sea potente, en caso de que no lo sea la solucion será alojarlo en otro lugar.

**Pregunta 9.2:** ¿Qué sistema operativo ejecuta el servidor? ¿Windows Server, Linux?
**Respuesta:** No se conoce, esperamos que sea linux, para la propuesta tecnica asume eso pero indica que puede haber cambios si el servidor es windows.

**Pregunta 9.3:** ¿Hay restricciones de firewall o puertos específicos que debamos considerar?
**Respuesta:** Configuración estándar con Caddy como reverse proxy en puerto 443 (HTTPS) y 80 (HTTP redirect). Go Fiber backend en puerto interno 8080. Configuración firewall mediante ufw en caso de que el servidor sea ubuntu.

**Pregunta 9.4:** ¿Cuántos usuarios concurrentes se esperan en el sistema?
**Respuesta:** No se conoce, pero no muchos, decenas como mucho.

**Pregunta 9.5:** ¿Habrá acceso desde fuera de la red local de DIPIDE? ¿VPN, acceso público?
**Respuesta:** la pagina web será de acceso publico por internet.

**Pregunta 9.6:** ¿Qué tipo de certificado SSL/TLS se utilizará? ¿Autofirmado, CA corporativa?
**Respuesta:** Caddy con Let's Encrypt para certificados automáticos si hay acceso público, o certificados autofirmados para uso interno. Configuración automática según el tipo de despliegue.

### SERVICIOS WEB

**Pregunta 10.1:** ¿Para qué casos de uso específicos se utilizará REST-XML vs REST-JSON?
**Respuesta:** JSON como formato principal para todas las operaciones CRUD y comunicación frontend-backend. XML reservado para reportes formales y documentos oficiales que requieran estructura específica. Ambos soportados mediante content-type negotiation.

**Pregunta 10.2:** ¿Hay algún estándar XML específico que DIPIDE ya utilice o deba cumplir?
**Respuesta:** Si pero no se conocen los detalles.

**Pregunta 10.3:** ¿Los web services serán utilizados inicialmente o es preparación para futuras integraciones?
**Respuesta:** Preparación para futuras integraciones. Se desarrollará API REST completa desde el inicio para facilitar integraciones futuras con otros sistemas gubernamentales.

**Pregunta 10.4:** ¿Qué nivel de documentación de API necesitamos? (Swagger, OpenAPI)
**Respuesta:** Documentación OpenAPI 3.0 automática usando swagger middleware en Go Fiber. Documentación interactiva disponible en /docs endpoint con autenticación.

**Pregunta 10.5:** ¿Habrá autenticación específica para los web services (API keys, OAuth)?
**Respuesta:** JWT tokens para autenticación de API con refresh tokens. Soporte para API keys para integraciones de sistemas.

### TIPOS DE INTERFACES DE USUARIO

**Pregunta 11.1:** ¿Cuáles son los dispositivos específicos desde los que se accederá? (PCs de escritorio, tablets, smartphones)
**Respuesta:** Principalmente desde pcs de escritorios, pero que sea responsivo es un requisito.

**Pregunta 11.2:** ¿Qué resoluciones de pantalla debemos soportar como mínimo?
**Respuesta:** Diseño responsive con soporte para resoluciones desde 320px (móvil) hasta 4K. Breakpoints estándar: móvil (<768px), tablet (768-1024px), desktop (>1024px).

**Pregunta 11.3:** ¿Hay usuarios con necesidades de accesibilidad específicas?
**Respuesta:** No.

**Pregunta 11.4:** ¿Se requiere funcionalidad offline o siempre será con conexión?
**Respuesta:** Siempre será con conexion.

### INTERFACES DE USUARIO (MÓDULO PRINCIPAL)

**Pregunta 12.1:** ¿Cómo debe lucir la pantalla de login? ¿Logo de DIPIDE, información institucional?
**Respuesta:** Mientras se vea lindo no hay mayores requisitos.

**Pregunta 12.2:** ¿Qué información debe mostrar el dashboard principal? ¿KPIs, notificaciones, accesos directos?
**Respuesta:** En principio, no se tiene planificado hacer un dashboard con el cliente, sin embargo a mi personalmente me parece que será necesario, por lo que lo haremos de todas formas. Dado que no está planificado, tampoco se conoce que informacion se debe mostrar. Aunque seria interesante si fuera configurable por usuario, es decir que cada usuario seleccione la informacion que le interesa ver.

**Pregunta 12.3:** ¿Hay algún estándar de diseño visual que debamos seguir? ¿Colores institucionales, tipografías?
**Respuesta:** No, asi que usaremos los estilos por defecto de flutter.

**Pregunta 12.4:** ¿La configuración será por usuario o por rol organizacional?
**Respuesta:** Configuración híbrida: configuraciones base por rol organizacional con posibilidad de personalización individual. Almacenamiento en PostgreSQL con cache en browser para configuraciones de UI frecuentes.

### INTERFACES DE USUARIO (MÓDULOS DE GESTIÓN ESPECIALIZADA)

**Pregunta 13.1:** ¿Cuáles son los módulos de gestión especializada específicos? (Gestión de Almacén, Gestión de Producción, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 13.2:** ¿Cada unidad funcional tendrá su propia interfaz especializada o habrá solapamiento?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 13.3:** ¿Qué nivel de personalización por unidad funcional se requiere?
**Respuesta:** Sistema modular con componentes reutilizables pero configurables por unidad funcional. Personalización de campos visibles, validaciones específicas y flujos de trabajo mediante configuración en base de datos.

### INTERFACES DE USUARIO (MÓDULOS DE TAREAS ESPECIALIZADAS)

**Pregunta 14.1:** ¿Cuáles son las tareas especializadas más críticas? (Registro de orden provisoria, Actualización de inventario, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 14.2:** ¿Habrá wizards o formularios de múltiples pasos para tareas complejas?
**Respuesta:** Sí, implementación de wizard components reutilizables con navegación step-by-step, validación por pasos y persistencia temporal en browser storage para recuperación ante desconexión.

**Pregunta 14.3:** ¿Se necesitan validaciones en tiempo real durante la captura de datos?
**Respuesta:** Validaciones en tiempo real usando debounced API calls para validación de disponibilidad, duplicados y reglas de negocio. Validación client-side para formato y server-side para integridad de datos.

### INTERFACES DE USUARIO (MÓDULOS DE ACCIONES ESPECÍFICAS)

**Pregunta 15.1:** ¿Cuáles son las acciones específicas más comunes? (Consulta de stock, Búsqueda de expediente, etc.)
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 15.2:** ¿Se necesitan reportes en tiempo real o pueden ser procesados por lotes?
**Respuesta:** Reportes básicos en tiempo real para consultas simples, reportes complejos procesados en background con notificación al usuario. Sistema de cola usando Go routines para reportes grandes.

**Pregunta 15.3:** ¿Qué formato tendrán los reportes? (PDF, Excel, CSV, vista en pantalla)
**Respuesta:** Múltiples formatos: vista en pantalla (HTML), exportación CSV/Excel usando librerías Go, y PDF usando wkhtmltopdf para reportes formales. Generación asíncrona para archivos grandes.

### INTERFACES DE USUARIO (MÓDULOS DE ADMINISTRACIÓN DE DATOS)

**Pregunta 16.1:** ¿Quiénes tendrán acceso a los módulos de administración de datos?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 16.2:** ¿Se necesitan herramientas de auditoría para ver quién modificó qué datos y cuándo?
**Respuesta:** Sí, auditoría completa usando triggers PostgreSQL que registren automáticamente en tabla de auditoría todos los cambios con: usuario, timestamp, valores anteriores/nuevos, IP de origen.

**Pregunta 16.3:** ¿Habrá funciones de eliminación lógica o todo será persistencia completa?
**Respuesta:** Persistencia completa sin eliminación física. Eliminación lógica usando campos deleted_at y soft delete patterns. Funcionalidad de "papelera" para recuperación de registros eliminados lógicamente.

**Pregunta 16.4:** ¿Se necesitan herramientas de importación/exportación masiva de datos?
**Respuesta:** Sí, herramientas integradas de import/export CSV con mapeo de campos configurable, validación de datos previa, y procesamiento por lotes con barra de progreso en tiempo real.

---

## PREGUNTAS ADICIONALES SOBRE IMPLEMENTACIÓN

### CAPACITACIÓN Y ADOPCIÓN

**Pregunta 17.1:** ¿Cuál es el nivel técnico actual de los usuarios finales? ¿Experiencia con Excel, sistemas informáticos?
**Respuesta:** No tengo conocimiento de primera mano pero al parecer el conocimiento técnico es bajo.

**Pregunta 17.2:** ¿Cuántos usuarios serán capacitados en total? ¿Por roles?
**Respuesta:** No se conocen la cantidad, la capacitacion será general para todos por igual.

**Pregunta 17.3:** ¿Hay usuarios clave que puedan actuar como "multiplicadores" durante la capacitación?
**Respuesta:** No.

### MIGRACIÓN DE DATOS

**Pregunta 18.1:** ¿En qué formato están los datos actuales? ¿Excel distribuido, sistemas legacy?
**Respuesta:** Excel principalmente

**Pregunta 18.2:** ¿Cuáles son los datos más críticos que deben migrarse sin pérdida?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 18.3:** ¿Hay datos históricos que deban preservarse? ¿Cuántos años hacia atrás?
**Respuesta:** En la propuesta indica que no se preservaran datos historicos, pero somos flexibles si es un requisito.

**Pregunta 18.4:** ¿Se puede tener un período de funcionamiento paralelo (sistema viejo + nuevo)?
**Respuesta:** No, una vez implementado deberan usar el sistema nuevo.

### INTEGRACIÓN CON SISTEMAS EXISTENTES

**Pregunta 19.1:** ¿Hay otros sistemas informáticos en DIPIDE que eventualmente deban integrarse?
**Respuesta:** En esta etapa seguro que no.

**Pregunta 19.2:** ¿Hay intercambio de información con otras dependencias del gobierno provincial?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

**Pregunta 19.3:** ¿Se utilizan actualmente sistemas de gestión documental específicos?
**Respuesta:** No se conoce, la propuesta debe ser flexible.

### OPERACIÓN Y SOPORTE

**Pregunta 20.1:** ¿Quién será el administrador técnico del sistema en DIPIDE?
**Respuesta:** En una primera etapa los postulantes (gabriel, claudio y adriana) pero luego se capacitara a un referente de DIPIDE.

**Pregunta 20.2:** ¿Qué horarios de operación tendrá el sistema? ¿24/7 o horario administrativo?
**Respuesta:** Probablemente se utilizara en horario laboral, pero debe estar disponible 24/7. aunque se pueden hacer tareas de mantenimiento en algun horario acordado.

**Pregunta 20.3:** ¿Hay ventanas de mantenimiento disponibles para actualizaciones?
**Respuesta:** Será negociable.

**Pregunta 20.4:** ¿Cómo se escalará el soporte técnico? ¿Niveles de soporte?
**Respuesta:** En principio Adrian Buratovich será la primera linea de soporte y el equipo de desarrollo la segunda.

---

## PREGUNTAS DE VALIDACIÓN TÉCNICA

### RENDIMIENTO

**Pregunta 21.1:** ¿Cuáles son los tiempos de respuesta esperados para las consultas más comunes?
**Respuesta:** Objetivos de rendimiento: consultas simples <200ms, consultas complejas <2s, reportes <30s. Se implementarán índices PostgreSQL optimizados, cache Redis para consultas frecuentes, y paginación para listados grandes.

**Pregunta 21.2:** ¿Cuál será el volumen de transacciones diarias estimado?
**Respuesta:** No se conoce, pero no espero que sea mucho.

**Pregunta 21.3:** ¿Hay operaciones que requieran procesamiento en background?
**Respuesta:** Sí: generación de reportes grandes, procesamiento de importaciones masivas, cálculos de reposición de inventario. Se implementará usando Go routines con worker pools y sistema de cola con monitoreo de progreso.

### SEGURIDAD

**Pregunta 22.1:** ¿Hay políticas de seguridad específicas del estado provincial que debamos cumplir?
**Respuesta:** No se conoce, asume que no.

**Pregunta 22.2:** ¿Se requiere encriptación de datos en reposo?
**Respuesta:** Encriptación de datos sensibles usando AES-256 en PostgreSQL con campos encrypted. Encriptación automática de contraseñas usando bcrypt. Certificados TLS para datos en tránsito.

**Pregunta 22.3:** ¿Qué políticas de contraseñas se deben implementar?
**Respuesta:** No seremos muy estricto con esto porque el riesgo de ataque es bajo. Permitiremos reutilizar contraseñas y usar contraseñas debiles.

**Pregunta 22.4:** ¿Habrá auditorías de seguridad periódicas?
**Respuesta:** Manuales no. Podriamos incluir algun reporte para ver a demanda, aunque esto no seria parte de lo planificado, se cobraria por separado.

### CONTINGENCIAS

**Pregunta 23.1:** ¿Qué sucede si el servidor principal falla? ¿Hay plan de contingencia?
**Respuesta:** No, por eso debemos permitir que descargar manualmente el backup.

**Pregunta 23.2:** ¿Cómo se maneja la pérdida de conectividad de red?
**Respuesta:** Intentaremos que todas las operaciones queden en un estado consistente con capacidad de hacer rollback en caso de que una operacion se corte por la mitad.

**Pregunta 23.3:** ¿Hay procedimientos manuales de respaldo para operaciones críticas?
**Respuesta:** No es algo pedido por la DIPIDE, agregarlo como algo de valor adicional que recomendamos que se cotizara por separado.

---

## PROPUESTA DE ARQUITECTURA TÉCNICA SIMPLIFICADA

### STACK TECNOLÓGICO RECOMENDADO

**Backend:**
- **Go con Fiber Framework**: API REST rápida y eficiente
- **PostgreSQL 14+**: Base de datos con JSON/JSONB support
- **Podman**: Containerización para fácil despliegue y mantenimiento

**Frontend:**
- **Flutter**: Con soporte web en una primera etapa. Android/IOs mas adelante.

**Infraestructura:**
- **Caddy Server**: Reverse proxy con HTTPS automático
- **Systemd**: Gestión de servicios Linux
- **Scripts de deployment**: Automatización de instalación y updates

### VENTAJAS DE ESTA ARQUITECTURA

1. **Simplicidad**: Stack tecnológico maduro y bien documentado
2. **Performance**: Go es extremadamente eficiente para APIs REST
3. **Mantenibilidad**: Menos dependencias y menor complejidad técnica
4. **Escalabilidad**: Fácil escalado horizontal con contenedores
5. **Costos**: Menor complejidad = menores costos de desarrollo y mantenimiento

---

**NOTA:** No se menciona en ningun momento el ABM de usuarios/roles, sin embargo es un requisito que nadie de la DIPIDE toquetee la base de datos, por lo que debemos agregar algun panel de administracion. Asumiremos que DIPIDE no tiene active directory ni nada similar, pero seremos flexible a agregar autenticacion externa en el codigo aunque no sea la opcion primaria.

**NOTA IMPORTANTE:** Estas preguntas están diseñadas para obtener la información específica necesaria para elaborar una propuesta técnica completa y detallada. Es recomendable agrupar las preguntas por tema y realizarlas en sesiones de trabajo estructuradas con los stakeholders de DIPIDE, coordinadas a través de Adrián Sergio Buratovich (PMO).