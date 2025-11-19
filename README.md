# 🧪 Sistema de Gestión de Alquimia Estatal - Amestris

Sistema full-stack moderno para la gestión integral de alquimistas estatales, misiones, transmutaciones y auditorías del Departamento de Alquimia de Amestris.

## 📋 Tabla de Contenidos

- [Descripción](#descripción)
- [Tecnologías](#tecnologías)
- [Instalación](#instalación)
- [Usuarios de Prueba](#usuarios-de-prueba)
- [Funcionalidades](#funcionalidades)
- [API Endpoints](#api-endpoints)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Comandos Útiles](#comandos-útiles)
- [Troubleshooting](#troubleshooting)

## 🎯 Descripción

Este sistema permite gestionar de manera eficiente:

- **Alquimistas**: Registro, autenticación y gestión de perfiles
- **Misiones**: Creación, aprobación y seguimiento de misiones estatales
- **Transmutaciones**: Solicitud y procesamiento de transmutaciones alquímicas
- **Materiales**: Inventario y gestión de materiales alquímicos
- **Auditorías**: Sistema de auditoría automática y manual
- **Notificaciones en Tiempo Real**: Actualizaciones instantáneas vía WebSocket

## 🛠 Tecnologías

### Backend
- **Go 1.23+** - Lenguaje principal
- **Gorilla Mux** - Router HTTP
- **GORM** - ORM para PostgreSQL
- **JWT** - Autenticación
- **WebSocket** - Comunicación en tiempo real

### Frontend
- **Next.js 14** - Framework React
- **TypeScript** - Tipado estático
- **Tailwind CSS** - Estilos modernos con tema oscuro
- **Axios** - Cliente HTTP

### Base de Datos
- **PostgreSQL 15** - Base de datos relacional

### Cola de Mensajes
- **Redis 7** - Sistema de cola de mensajes para procesamiento asíncrono

### Infraestructura
- **Docker & Docker Compose** - Contenedorización completa

## 🚀 Instalación

### Prerrequisitos

- Docker y Docker Compose instalados
- Go 1.23+ (para desarrollo local)
- Node.js 18+ y npm (para desarrollo local)

### Instalación con Docker (Recomendado)

1. **Clonar o descargar el proyecto**
```bash
cd Final
```

2. **Construir y levantar los servicios**
```bash
docker-compose up --build
```

3. **Poblar la base de datos con datos de prueba**
```bash
POSTGRES_HOST=localhost POSTGRES_USER=amestris POSTGRES_PASSWORD=amestris123 POSTGRES_DB=amestris_db go run scripts/seed.go
```

4. **Acceder a la aplicación**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8000
   - PostgreSQL: localhost:5432
   - Redis: localhost:6379

### Instalación Local (Desarrollo)

#### Backend

```bash
# Instalar dependencias
go mod download

# Configurar variables de entorno
export POSTGRES_HOST=localhost
export POSTGRES_USER=amestris
export POSTGRES_PASSWORD=amestris123
export POSTGRES_DB=amestris_db
export JWT_SECRET=amestris-alchemy-secret-key-change-in-production

# Ejecutar
go run main.go
```

#### Frontend

```bash
cd frontend

# Instalar dependencias
npm install

# Ejecutar en modo desarrollo
npm run dev
```

## 👤 Usuarios de Prueba

Todos los usuarios de prueba tienen la misma contraseña: **`password123`**

### Alquimistas

| Nombre | Email | Rango | Especialidad | Rol |
|--------|-------|-------|--------------|-----|
| Edward Elric | `edward@amestris.gov` | Estatal | Combate | Alquimista |
| Alphonse Elric | `alphonse@amestris.gov` | Estatal | Investigación | Alquimista |
| Riza Hawkeye | `hawkeye@amestris.gov` | Estatal | Combate | Alquimista |
| Winry Rockbell | `winry@amestris.gov` | Aprendiz | Industrial | Alquimista |
| Shou Tucker | `tucker@amestris.gov` | Estatal | Investigación | Alquimista |
| Izumi Curtis | `izumi@amestris.gov` | Nacional | Combate | Alquimista |

### Supervisores

| Nombre | Email | Rango | Especialidad | Rol |
|--------|-------|-------|--------------|-----|
| Roy Mustang | `mustang@amestris.gov` | Nacional | Combate | Supervisor |
| Alex Louis Armstrong | `armstrong@amestris.gov` | Nacional | Combate | Supervisor |

## ✨ Funcionalidades

### Para Alquimistas

#### Panel de Alquimista
- **Dashboard personalizado** con vista de misiones y transmutaciones
- **Crear nuevas misiones** con título y descripción
- **Solicitar transmutaciones** seleccionando materiales y cantidades
- **Ver estado en tiempo real** de misiones y transmutaciones
- **Notificaciones automáticas** cuando cambia el estado de sus solicitudes

#### Gestión de Misiones
- Crear misiones con título y descripción detallada
- Ver todas las misiones asignadas
- Seguimiento del estado: Pendiente → Aprobada → En Progreso → Completada
- Historial completo de actividades

#### Gestión de Transmutaciones
- Crear solicitudes de transmutación
- Seleccionar múltiples materiales de entrada
- Especificar cantidades requeridas
- Ver costo estimado de la transmutación
- Seguimiento del estado: Pendiente → Aprobada → Completada/Rechazada

### Para Supervisores

#### Panel de Supervisor
- **Dashboard con estadísticas**:
  - Misiones pendientes
  - Transmutaciones pendientes
  - Auditorías sin resolver
  - Costo total de transmutaciones

#### Aprobación de Solicitudes
- **Aprobar o rechazar misiones** pendientes
- **Aprobar o rechazar transmutaciones** pendientes
- Agregar comentarios y resultados a las transmutaciones

#### Gestión de Auditorías
- Ver todas las auditorías del sistema
- Filtrar por severidad (Baja, Media, Alta, Crítica)
- Resolver auditorías con acciones correctivas
- Ver detalles de cada auditoría

#### Monitoreo
- Monitoreo de uso de materiales
- Detección automática de anomalías
- Alertas de seguridad

### Sistema de Auditoría Automática

El sistema genera auditorías automáticamente cuando detecta:
- Uso excesivo de materiales
- Misiones pendientes por mucho tiempo
- Transmutaciones que violan protocolos
- Stock de materiales por debajo del mínimo
- Actividades sospechosas

### Notificaciones en Tiempo Real

- Actualizaciones instantáneas vía WebSocket
- Notificaciones cuando cambia el estado de misiones
- Alertas de nuevas transmutaciones
- Avisos de auditorías creadas

## 🔌 API Endpoints

### Autenticación

| Método | Endpoint | Descripción | Autenticación |
|--------|----------|-------------|---------------|
| POST | `/api/auth/login` | Iniciar sesión | No |
| POST | `/api/auth/register` | Registrarse | No |
| GET | `/api/auth/profile` | Obtener perfil | Sí |

### Alquimistas

| Método | Endpoint | Descripción | Autenticación |
|--------|----------|-------------|---------------|
| GET | `/api/alchemists` | Listar alquimistas | Sí |
| GET | `/api/alchemists/{id}` | Obtener alquimista | Sí |
| POST | `/api/alchemists` | Crear alquimista | Sí |
| PUT | `/api/alchemists/{id}` | Actualizar alquimista | Sí |
| DELETE | `/api/alchemists/{id}` | Eliminar alquimista | Sí |

### Misiones

| Método | Endpoint | Descripción | Autenticación |
|--------|----------|-------------|---------------|
| GET | `/api/missions` | Listar misiones | Sí |
| GET | `/api/missions/{id}` | Obtener misión | Sí |
| POST | `/api/missions` | Crear misión | Sí |
| PUT | `/api/missions/{id}` | Actualizar misión | Sí |
| PUT | `/api/missions/{id}/status` | Actualizar estado | Sí |
| DELETE | `/api/missions/{id}` | Eliminar misión | Sí |

### Materiales

| Método | Endpoint | Descripción | Autenticación |
|--------|----------|-------------|---------------|
| GET | `/api/materials` | Listar materiales | Sí |
| GET | `/api/materials/{id}` | Obtener material | Sí |
| POST | `/api/materials` | Crear material | Sí |
| PUT | `/api/materials/{id}` | Actualizar material | Sí |
| DELETE | `/api/materials/{id}` | Eliminar material | Sí |

### Transmutaciones

| Método | Endpoint | Descripción | Autenticación |
|--------|----------|-------------|---------------|
| GET | `/api/transmutations` | Listar transmutaciones | Sí |
| GET | `/api/transmutations/{id}` | Obtener transmutación | Sí |
| POST | `/api/transmutations` | Crear transmutación | Sí |
| PUT | `/api/transmutations/{id}` | Actualizar transmutación | Sí |
| PUT | `/api/transmutations/{id}/status` | Actualizar estado | Sí |
| DELETE | `/api/transmutations/{id}` | Eliminar transmutación | Sí |

### Auditorías

| Método | Endpoint | Descripción | Autenticación |
|--------|----------|-------------|---------------|
| GET | `/api/audits` | Listar auditorías | Sí |
| GET | `/api/audits/{id}` | Obtener auditoría | Sí |
| POST | `/api/audits` | Crear auditoría | Sí |
| PUT | `/api/audits/{id}` | Actualizar auditoría | Sí |
| PUT | `/api/audits/{id}/resolve` | Resolver auditoría | Sí |
| DELETE | `/api/audits/{id}` | Eliminar auditoría | Sí |

### WebSocket

| Endpoint | Descripción | Autenticación |
|----------|-------------|---------------|
| `/api/ws?token={jwt}` | Conexión WebSocket para notificaciones | Token en query |

## 📁 Estructura del Proyecto

```
Final/
├── api/                    # DTOs y estructuras de respuesta API
│   ├── alchemist.go
│   ├── auth.go
│   ├── material.go
│   ├── mission.go
│   └── transmutation.go
├── auth/                   # Autenticación JWT y middleware
│   ├── jwt.go
│   └── middleware.go
├── config/                 # Configuración
│   ├── config.go
│   └── config.json
├── logger/                 # Sistema de logging
│   └── logger.go
├── models/                 # Modelos de base de datos (GORM)
│   ├── alchemist.go
│   ├── audit.go
│   ├── material.go
│   ├── mission.go
│   └── transmutation.go
├── repository/             # Repositorios de acceso a datos
│   ├── alchemist_repository.go
│   ├── audit_repository.go
│   ├── material_repository.go
│   ├── mission_repository.go
│   └── transmutation_repository.go
├── server/                 # Handlers, router y lógica del servidor
│   ├── alchemist_handlers.go
│   ├── async_processor.go
│   ├── audit_handlers.go
│   ├── auth_handlers.go
│   ├── material_handlers.go
│   ├── mission_handlers.go
│   ├── router.go
│   ├── server.go
│   ├── task_queue.go
│   ├── transmutation_handlers.go
│   └── websocket.go
├── scripts/                # Scripts de utilidad
│   ├── seed.go            # Script para poblar base de datos
│   └── seed_demo.go
├── frontend/               # Aplicación Next.js
│   ├── app/               # Páginas y layouts
│   │   ├── dashboard/
│   │   ├── globals.css
│   │   ├── layout.tsx
│   │   └── page.tsx
│   ├── components/        # Componentes React
│   │   ├── AlchemistDashboard.tsx
│   │   ├── AuditList.tsx
│   │   ├── CreateMissionModal.tsx
│   │   ├── CreateTransmutationModal.tsx
│   │   ├── DashboardLayout.tsx
│   │   ├── DashboardStats.tsx
│   │   ├── LoginForm.tsx
│   │   ├── MissionList.tsx
│   │   ├── SupervisorDashboard.tsx
│   │   └── TransmutationList.tsx
│   ├── lib/               # Utilidades y API client
│   │   └── api.ts
│   └── package.json
├── docker-compose.yml      # Configuración Docker Compose
├── Dockerfile             # Dockerfile del backend
├── go.mod                 # Dependencias Go
└── main.go               # Punto de entrada
```

## 🛠 Comandos Útiles

### Docker

```bash
# Iniciar todos los servicios
docker-compose up

# Iniciar en segundo plano
docker-compose up -d

# Detener todos los servicios
docker-compose down

# Ver logs del backend
docker-compose logs backend

# Ver logs del frontend
docker-compose logs frontend

# Ver logs de PostgreSQL
docker-compose logs postgres

# Reiniciar un servicio específico
docker-compose restart backend

# Reconstruir contenedores
docker-compose up --build

# Eliminar volúmenes (⚠️ borra la base de datos)
docker-compose down -v
```

### Base de Datos

```bash
# Poblar base de datos con datos de prueba
POSTGRES_HOST=localhost POSTGRES_USER=amestris POSTGRES_PASSWORD=amestris123 POSTGRES_DB=amestris_db go run scripts/seed.go

# Conectar a PostgreSQL desde Docker
docker-compose exec postgres psql -U amestris -d amestris_db

# Ver tablas en la base de datos
docker-compose exec postgres psql -U amestris -d amestris_db -c "\dt"
```

### Desarrollo

```bash
# Backend - Ejecutar en modo desarrollo
go run main.go

# Frontend - Ejecutar en modo desarrollo
cd frontend && npm run dev

# Frontend - Construir para producción
cd frontend && npm run build

# Frontend - Ejecutar producción
cd frontend && npm start
```

## 🔧 Troubleshooting

### Problema: No puedo iniciar sesión con usuarios existentes

**Solución:**
```bash
# Ejecutar el script de seed para restaurar usuarios
POSTGRES_HOST=localhost POSTGRES_USER=amestris POSTGRES_PASSWORD=amestris123 POSTGRES_DB=amestris_db go run scripts/seed.go
```

### Problema: Los contenedores no inician

**Solución:**
```bash
# Verificar que los puertos no estén en uso
lsof -i :3000  # Frontend
lsof -i :8000  # Backend
lsof -i :5432  # PostgreSQL

# Detener y eliminar contenedores
docker-compose down

# Reconstruir desde cero
docker-compose up --build
```

### Problema: Error de conexión a la base de datos

**Solución:**
```bash
# Verificar que PostgreSQL esté corriendo
docker-compose ps

# Ver logs de PostgreSQL
docker-compose logs postgres

# Reiniciar PostgreSQL
docker-compose restart postgres
```

### Problema: El frontend no se conecta al backend

**Solución:**
1. Verificar que el backend esté corriendo en `http://localhost:8000`
2. Verificar la variable de entorno `NEXT_PUBLIC_API_URL` en `docker-compose.yml`
3. Revisar la consola del navegador para errores CORS

### Problema: No aparecen materiales al crear transmutación

**Solución:**
```bash
# Ejecutar el script de seed para crear materiales
POSTGRES_HOST=localhost POSTGRES_USER=amestris POSTGRES_PASSWORD=amestris123 POSTGRES_DB=amestris_db go run scripts/seed.go
```

### Problema: WebSocket no funciona

**Solución:**
1. Verificar que el token JWT esté presente en localStorage
2. Verificar que la URL del WebSocket sea correcta: `ws://localhost:8000/api/ws?token={token}`
3. Revisar los logs del backend para errores de conexión

## 📝 Notas Adicionales

### Tema Oscuro

El frontend utiliza un tema oscuro moderno con:
- Fondos oscuros con gradientes sutiles
- Efectos de glassmorphism
- Colores de acento (cyan, purple, pink)
- Transiciones suaves
- Diseño responsive

### Seguridad

- Todas las rutas API (excepto login/register) requieren autenticación JWT
- Las contraseñas se almacenan con hash bcrypt
- CORS configurado para desarrollo
- Validación de datos en backend

### Materiales Disponibles

El sistema viene con 10 materiales predefinidos:
- Hierro, Acero, Oro, Plata (Metales)
- Carbón, Tierra, Fósforo (Minerales)
- Agua (Orgánico)
- Cristal Alquímico, Mercurio (Sintéticos)

## 📞 Soporte

Para problemas o preguntas:
1. Revisar los logs: `docker-compose logs`
2. Verificar que todos los servicios estén corriendo: `docker-compose ps`
3. Ejecutar el script de seed si faltan datos: `go run scripts/seed.go`

## 📚 Documentación de la API

### Documentación Completa
Ver el archivo [API_DOCUMENTATION.md](./API_DOCUMENTATION.md) para documentación detallada de todos los endpoints.

### Colección de Postman
Importa el archivo `postman_collection.json` en Postman para probar todos los endpoints fácilmente.

**Características de la colección:**
- Variables de entorno configuradas
- Autenticación automática (el token se guarda automáticamente después del login)
- Ejemplos de request/response para cada endpoint
- Organizada por categorías

## ✅ Cumplimiento de Requerimientos

### Backend (Go)
- ✅ API REST implementada con Gorilla Mux (sin frameworks como Gin, Fiber, Echo)
- ✅ Go 1.23+ utilizado
- ✅ PostgreSQL como base de datos
- ✅ GORM como ORM
- ✅ Autenticación JWT implementada
- ✅ Endpoints CRUD completos para todas las entidades:
  - ✅ Alquimistas
  - ✅ Misiones
  - ✅ Materiales
  - ✅ Transmutaciones
  - ✅ Auditorías
- ✅ Sistema de procesamiento asíncrono:
  - ✅ Cola de tareas en memoria (TaskQueue)
  - ✅ Redis disponible para escalabilidad
  - ✅ Procesamiento de transmutaciones en segundo plano
  - ✅ Verificaciones automáticas diarias
  - ✅ Generación automática de auditorías

### Frontend (Next.js/React)
- ✅ Next.js 14 con TypeScript
- ✅ Interfaz responsive y moderna
- ✅ Login y registro de usuarios
- ✅ Paneles diferenciados:
  - ✅ Panel de Alquimista
  - ✅ Panel de Supervisor
- ✅ Visualizaciones de datos (estadísticas, listas)
- ✅ Notificaciones en tiempo real vía WebSocket
- ✅ Uso correcto de hooks (useState, useEffect)
- ✅ Código completamente tipado (sin `any` injustificados)

### Infraestructura (Docker)
- ✅ Docker Compose configurado
- ✅ Backend contenerizado
- ✅ Frontend contenerizado
- ✅ PostgreSQL contenerizado
- ✅ Redis contenerizado
- ✅ Script único para levantar todo el entorno

### Entregables
- ✅ Repositorio con estructura clara
- ✅ docker-compose.yml funcional
- ✅ Script de inicialización de base de datos (seed.go)
- ✅ Documentación de API (API_DOCUMENTATION.md)
- ✅ Colección de Postman (postman_collection.json)
- ✅ README.md completo con manual de despliegue

### Sistema Asíncrono
- ✅ Procesamiento de transmutaciones en segundo plano
- ✅ Verificaciones automáticas de uso de materiales
- ✅ Detección de misiones no cerradas
- ✅ Generación automática de auditorías
- ✅ Notificaciones WebSocket en tiempo real

### Características Adicionales
- ✅ Sistema de auditoría automática completo
- ✅ Verificaciones diarias programadas
- ✅ WebSocket para actualizaciones en tiempo real
- ✅ Interfaz moderna con tema oscuro
- ✅ Diseño responsive

## 📄 Licencia

Este proyecto es parte de un sistema de gestión académico.

---

**Desarrollado con ❤️ para el Departamento de Alquimia Estatal de Amestris**

