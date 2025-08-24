# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**DIPIDE** (Dirección Provincial de Impresión y Digitalización del Estado) - Digital Transformation Project for modular information systems to optimize print work order management processes.

### Project Context
- **Organization**: DIPIDE - Provincial government printing and digitalization services (La Plata, Buenos Aires)
- **Development Team**: 3-person team (Claudio Paz - Full Stack, Gabriel Suárez - Backend/Security, Adriana Aguirre - Functional Analyst/UX-UI)
- **PMO Coordinator**: Adrián Sergio Buratovich (Serian IO) - Project intermediary and change management
- **Development Timeline**: 6 months development + 3 months implementation + ongoing support

## Technical Architecture

### Core Technology Stack (MANDATORY REQUIREMENTS)
- **Frontend**: Flutter/Dart (cross-platform: Desktop, Web, Mobile)
- **Backend**: Web server with HTTPS (REST API)
- **Database**: PostgreSQL (relational SQL database)
- **Deployment**: Physical server at DIPIDE headquarters (La Plata)
- **Web Services**: REST-XML (primary for digital documents), REST-JSON (alternative)
- **Interfaces**: Web Responsive (terminals and mobile devices)

### Modular Architecture (5 Required Modules)
1. **Main Module**: Centralized security/access control, configuration management, role-based navigation
2. **Specialized Management Modules**: Inherited access, functional unit task management
3. **Specialized Task Modules**: Specific process tasks, multi-level access control
4. **Specific Action Modules**: Real-time processing, flexible data configuration
5. **Data Administration Modules**: CRUD interfaces, information manipulation

## Core Business Process

**Primary Process**: GO.001.001 - WORK ORDER MANAGEMENT (PRINTING JOBS)

### Process Flow (8 Stages)
1. **TRAMO 01** - Start/Internal Feasibility
2. **TRAMO 02** - Prerequisites  
3. **TRAMO 03** - Programming
4. **TRAMO 04** - Coordination and Management
5. **TRAMO 05** - Monitoring and Control
6. **TRAMO 06** - Completion
7. **TRAMO 07** - [Reserved]
8. **TRAMO 08** - Incidents

### Functional Units
- **COMERCIALIZACIÓN**: Responsible for start, prerequisites
- **PRODUCCIÓN**: Responsible for programming, management, control
- **ALMACÉN**: Responsible for monitoring, control, completion
- **TRANSPORTE**: Informed in all processes

## Development Methodology

### Sprint Planning (6 months / 12 sprints)
- **Methodology**: AGILE/SCRUM 
- **Sprint Duration**: 2 weeks each
- **Backlog Review**: First 3 months compressed to avoid unforeseen issues
- **Implementation**: Starts in month 4 (mid-way)

### 5 Critical Milestones

#### **MILESTONE 1: Database Administration** (Sprints 1-2)
- **Semi-automated CRUD** via CSV import/export
- **Configurable CSV format** (preferably tab-separated)
- **Batch operations** for high/low/modify records
- **Full table export** for backup purposes
- **No dedicated DBA** - tool must be self-sufficient

#### **MILESTONE 2: Movements/Traceability** (Sprints 3-4)
- **Process traceability** via file numbers and work orders
- **Management/control states** (differentiated for files vs orders)
- **Complete persistence** - NO movement deletion allowed
- **Chronological stock control** with baseline + movements
- **Audit trail** for all operations

#### **MILESTONE 3: Inventory States** (Sprints 5-6)
- **Material/Supply differentiation** (Materials = high priority, Supplies = medium)
- **State management**: available/reserved/in-process
- **Stock control** with starting point and chronological movements
- **Movement types**: configurable ingress/egress operations

#### **MILESTONE 4: Replenishment Fields** (Sprints 7-8) - **DEFERRED**
- **Stock minimums/maximums** (postponed to later phases)
- **Provisional orders and reserves**
- **Integration** with functional units

#### **MILESTONE 5: Modular Interfaces** (Sprints 9-12)
- **Web responsive interfaces** (mobile NOT initially required)
- **Role-based organizational modules**
- **Architecture prepared** for future web services
- **User training** (critical due to low technical level)

## Key Business Rules

### Data Management
- **Master Data**: Products, suppliers, organizational units, user roles
- **Movement Data**: All transactions with complete audit trail
- **Relational Data**: Cross-reference tables (suppliers-products, user-roles, etc.)
- **No data deletion**: All operations must be traceable

### User Access & Security
- **Internal authentication**: Simple credential system (no Active Directory integration required)
- **Role-based access**: Users → Organizational Units → Activities → Interfaces
- **Geographic context**: DIPIDE La Plata (single location initially)
- **Multi-level permissions**: Inherited from main module

### Stock Management
- **Control points**: Periodic stock counts establish baseline
- **Movement tracking**: Chronological from last control point
- **State differentiation**: Available, reserved, in-process, etc.
- **Adjustment movements**: For discrepancies (no direct deletion)

## Data Integration & Formats

### CSV Management
- **Format**: Configurable separators (tab preferred over comma)
- **Import/Export**: Full table operations
- **Batch processing**: High/low/modify operations with ID references
- **Client responsibility**: Data standardization and backup

### Future Web Services
- **REST Architecture**: XML for formal documents, JSON for data exchange
- **Service Layer**: Prepared in "Specific Actions" modules
- **No SOA contract required**: Simple REST over HTTP
- **Document generation**: XML-CSS for standardized reports

## Infrastructure Constraints

### Server Environment
- **Single physical server** at DIPIDE (no backup/development environments)
- **Limited technical support**: Shared IT resources (not 24/7 dedicated)
- **Client responsibility**: Infrastructure, backups, and maintenance
- **Scalability requirement**: Must handle future high demand

### Development Considerations
- **Documentation-only repository**: This repository contains project specifications and requirements
- **No source code yet**: Implementation will follow proposed technical architecture
- **Future development**: When code is added, follow Flutter/Go Fiber architecture specified in propuesta_tecnica.md

## Contact Information

**PMO Coordinator**: Adrián Sergio Buratovich  
**Email**: as.buratovich@gmail.com  
**Phone**: +54 9 11 4094-3221  
**LinkedIn**: http://www.linkedin.com/in/adrian-buratovich  
**Website**: http://www.serianweb.com.ar

**Development Team Lead**: Claudio Alejandro Paz  
**Email**: alejandro30.11.2007@gmail.com

**Backend/Security Specialist**: Gabriel Ernesto Suárez  
**Email**: gabrielxsuarez@gmail.com

**Functional Analyst**: Adriana Noelia Aguirre  
**Email**: adriana.aguirre88@gmail.com

## Repository Structure

This repository contains project documentation only:
- `index.md`: Comprehensive index of all project documentation with summaries
- `resumen.md`: Executive summary with technical requirements
- `propuesta/`: Commercial and technical proposals with team information
- `documentacion/`: Detailed project specifications and requirements (Serian IO methodology)
- `reuniones/`: Meeting notes and technical clarifications

### Document Format Priority

**IMPORTANT**: All PDF, DOCX, and audio files in this project have been transformed to Markdown (.md) format for better accessibility and version control. When seeking project context or information:

1. **ALWAYS prefer the .md versions** over original formats (PDF, DOCX, audio transcripts)
2. **Use `index.md`** as the primary navigation document for finding specific information
3. **Original files** (PDF, DOCX) are kept for reference but may be outdated compared to their .md transformations
4. **All audio meeting recordings** have been transcribed to .md files in the `reuniones/` directory

The `.md` files contain the most current and accessible version of all project documentation, optimized for development workflow integration and Claude Code analysis.

## Development Workflow

### When Development Begins
Future development should follow this structure based on the technical proposal:

```
project-root/
├── backend/          # Go Fiber API
│   ├── main.go
│   ├── routes/       # REST endpoints
│   ├── models/       # PostgreSQL models (see DER.md)
│   ├── middleware/   # Authentication, RBAC
│   └── services/     # Business logic
├── frontend/         # Flutter Web application
│   ├── lib/
│   │   ├── main.dart
│   │   ├── modules/  # 5 required modules (see architecture)
│   │   ├── services/ # API clients
│   │   └── models/   # Data models
└── database/
    ├── migrations/   # PostgreSQL schema
    └── seeds/        # Initial data
```

### Key Development Commands (Future)
Based on the proposed technical stack:

- **Backend (Go Fiber)**:
  - `go run main.go` - Start development server
  - `go test ./...` - Run all tests
  - `go build` - Build production binary

- **Frontend (Flutter Web)**:
  - `flutter run -d web` - Start development server
  - `flutter test` - Run unit tests
  - `flutter build web` - Build for production

- **Database (PostgreSQL)**:
  - Migrations and seeds should be implemented for the DER schema in `DER.md`

## Development Notes

- **Requirements gathering**: All technical specs are in the documentation folder
- **Client intermediation**: All technical questions must go through Adrián (PMO coordinator)  
- **Change management**: Significant organizational change management component
- **Training critical**: End users have low technical skills - requires extensive training
- **Architecture reference**: See `propuesta/propuesta_tecnica.md` for complete technical implementation details

## Specific Technical Considerations

### CSV Processing
- Tab-separated values preferred (avoid comma conflicts in descriptions)
- Must handle special characters in product descriptions  
- Export functionality for complete tables (backup purposes)
- Import functionality for batch updates using internal IDs

### Movement Persistence  
- **NEVER delete movements** - only add corrective/adjustment entries
- Complete audit trail with timestamps and user identification
- State changes tracked as separate movements
- Support for provisional orders and reservations

### User Interface Requirements
- Role-based navigation and functionality
- Context-aware filtering (by organizational unit, activity, etc.)
- Standard HTML form controls (no custom formatting initially)
- Responsive design for terminals and basic mobile access

### Future Extensibility
- Modular architecture to support additional organizational units
- Web service layer preparation for external system integration
- Document generation system for formal reports and orders
- Integration preparation for production and commercial modules

## Critical Implementation Requirements

### Database Design
- **Refer to `propuesta/DER.md`** for complete entity-relationship diagram
- **18 core entities** including EXPEDIENTE, ORDEN_TRABAJO, MOVIMIENTO_STOCK, USUARIO, etc.
- **Audit trail**: All changes automatically logged via PostgreSQL triggers
- **Soft deletes**: Never physically delete records, only logical deletion

### API Architecture
- **Go Fiber framework** with modular route structure
- **JWT authentication** with refresh tokens for security
- **RBAC middleware** validating permissions before each endpoint
- **OpenAPI 3.0** documentation automatically generated at `/docs`
- **JSON primary**, XML for formal documents

### Frontend Modules (Mandatory)
Implementation must include exactly these 5 modules:
1. **Main Module**: Authentication, navigation, role-based UI
2. **Specialized Management**: Per-unit functional interfaces
3. **Specialized Tasks**: Process-specific workflows  
4. **Specific Actions**: Reports, real-time operations
5. **Data Administration**: CRUD operations, CSV import/export

### Data Integrity Rules
- **NO movement deletion**: Only corrective adjustments allowed
- **Complete audit trail**: User, timestamp, before/after values for all changes
- **Transactional imports**: CSV operations must be atomic (all-or-nothing)
- **State machine validation**: Workflow transitions enforced at database level