# Sistema de Gestión de Alquimia Estatal - Amestris

Sistema full-stack para la gestión de alquimistas estatales, misiones, transmutaciones y auditorías del Departamento de Alquimia de Amestris.

## Arquitectura

- **Backend**: Go 1.21+ con Gorilla Mux
- **Frontend**: Next.js 14 con TypeScript y React
- **Base de Datos**: PostgreSQL
- **Cola de Mensajes**: Redis (para procesamiento asíncrono)
- **Autenticación**: JWT
- **Tiempo Real**: WebSocket

## Instalación y Ejecución

### Usando Docker Compose

```bash
# Construir y levantar todos los servicios
docker-compose up --build

# En modo detached
docker-compose up -d --build
```

Los servicios estarán disponibles en:

- Frontend: http://localhost:3000
- Backend API: http://localhost:8000
- PostgreSQL: localhost:5432
- Redis: localhost:6379

### Desarrollo Local

#### Backend

```bash
# Instalar dependencias
go mod download

# Configurar variables de entorno
export POSTGRES_HOST=localhost
export POSTGRES_USER=amestris
export POSTGRES_PASSWORD=alchemy123
export POSTGRES_DB=amestris_db
export JWT_SECRET=amestris-alchemy-secret-key

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

## Estructura del Proyecto

```
.
├── api/              # DTOs y estructuras de respuesta
├── auth/             # Autenticación JWT y middleware
├── config/           # Configuración
├── logger/           # Sistema de logging
├── models/           # Modelos de base de datos (GORM)
├── repository/       # Repositorios de acceso a datos
├── server/           # Handlers, router y lógica del servidor
├── frontend/         # Aplicación Next.js
│   ├── app/          # Páginas y layouts
│   ├── components/   # Componentes React
│   └── lib/          # Utilidades y API client
└── docker-compose.yml
```

## Características y funciones

### Para Alquimistas

- Registro e inicio de sesión
- Crear y gestionar misiones
- Solicitar transmutaciones
- Ver historial de actividades
- Notificaciones en tiempo real

### Para Supervisores

- Aprobar/rechazar misiones
- Aprobar/rechazar transmutaciones
- Ver y resolver auditorías
- Dashboard con estadísticas
- Monitoreo de uso de materiales

### Sistema Asíncrono

- Procesamiento de transmutaciones en segundo plano
- Verificaciones automáticas de uso de materiales
- Auditorías automáticas
- Notificaciones WebSocket

## API Endpoints

### Autenticación

- `POST /api/auth/login` - Iniciar sesión
- `POST /api/auth/register` - Registrarse
- `GET /api/auth/profile` - Obtener perfil

### Alquimistas

- `GET /api/alchemists` - Listar alquimistas
- `GET /api/alchemists/{id}` - Obtener alquimista
- `POST /api/alchemists` - Crear alquimista
- `PUT /api/alchemists/{id}` - Actualizar alquimista
- `DELETE /api/alchemists/{id}` - Eliminar alquimista

### Misiones

- `GET /api/missions` - Listar misiones
- `GET /api/missions/{id}` - Obtener misión
- `POST /api/missions` - Crear misión
- `PUT /api/missions/{id}/status` - Actualizar estado

### Materiales

- `GET /api/materials` - Listar materiales
- `POST /api/materials` - Crear material
- `PUT /api/materials/{id}` - Actualizar material

### Transmutaciones

- `GET /api/transmutations` - Listar transmutaciones
- `POST /api/transmutations` - Crear transmutación
- `PUT /api/transmutations/{id}/status` - Actualizar estado

### Auditorías

- `GET /api/audits` - Listar auditorías
- `PUT /api/audits/{id}/resolve` - Resolver auditoría

### WebSocket

- `GET /api/ws` - Conexión WebSocket para notificaciones

## Variables de Entorno

### Backend

- `POSTGRES_HOST` - Host de PostgreSQL
- `POSTGRES_USER` - Usuario de PostgreSQL
- `POSTGRES_PASSWORD` - Contraseña de PostgreSQL
- `POSTGRES_DB` - Nombre de la base de datos
- `JWT_SECRET` - Clave secreta para JWT

### Frontend

- `NEXT_PUBLIC_API_URL` - URL del backend API

## Desarrollo

### Ejecutar Tests

```bash
# Backend
go test ./...

# Frontend
cd frontend
npm test
```

### Linting

```bash
# Backend
golangci-lint run

# Frontend
cd frontend
npm run lint
```

## Yerson Andres Perez Cadena - Santiago Bejarano Morera - Jessica Rivera
