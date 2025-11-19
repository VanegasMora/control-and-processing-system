# 📚 Documentación de la API - Sistema de Alquimia Estatal

## Base URL
```
http://localhost:8000
```

## Autenticación

La mayoría de los endpoints requieren autenticación mediante JWT. Incluye el token en el header:
```
Authorization: Bearer {token}
```

---

## 🔐 Autenticación

### POST /api/auth/login
Iniciar sesión en el sistema.

**Request Body:**
```json
{
  "email": "edward@amestris.gov",
  "password": "password123"
}
```

**Response 200:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user_id": 1,
  "email": "edward@amestris.gov",
  "role": "alchemist",
  "name": "Edward Elric",
  "rank": "state",
  "specialty": "combat"
}
```

### POST /api/auth/register
Registrar un nuevo alquimista.

**Request Body:**
```json
{
  "name": "Edward Elric",
  "email": "edward@amestris.gov",
  "password": "password123",
  "rank": "state",
  "specialty": "combat"
}
```

**Response 200:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user_id": 1,
  "email": "edward@amestris.gov",
  "role": "alchemist",
  "name": "Edward Elric",
  "rank": "state",
  "specialty": "combat"
}
```

### GET /api/auth/profile
Obtener el perfil del usuario autenticado.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "id": 1,
  "name": "Edward Elric",
  "email": "edward@amestris.gov",
  "rank": "state",
  "specialty": "combat",
  "role": "alchemist",
  "certified": true,
  "created_at": "2025-11-19T00:00:00Z"
}
```

---

## 👥 Alquimistas

### GET /api/alchemists
Listar todos los alquimistas.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
[
  {
    "id": 1,
    "name": "Edward Elric",
    "email": "edward@amestris.gov",
    "rank": "state",
    "specialty": "combat",
    "role": "alchemist",
    "certified": true,
    "created_at": "2025-11-19T00:00:00Z"
  }
]
```

### GET /api/alchemists/{id}
Obtener un alquimista por ID.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "id": 1,
  "name": "Edward Elric",
  "email": "edward@amestris.gov",
  "rank": "state",
  "specialty": "combat",
  "role": "alchemist",
  "certified": true,
  "created_at": "2025-11-19T00:00:00Z"
}
```

### POST /api/alchemists
Crear un nuevo alquimista (requiere permisos de supervisor).

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "name": "Edward Elric",
  "email": "edward@amestris.gov",
  "password": "password123",
  "rank": "state",
  "specialty": "combat",
  "role": "alchemist",
  "certified": true
}
```

**Response 201:**
```json
{
  "id": 1,
  "name": "Edward Elric",
  "email": "edward@amestris.gov",
  "rank": "state",
  "specialty": "combat",
  "role": "alchemist",
  "certified": true,
  "created_at": "2025-11-19T00:00:00Z"
}
```

### PUT /api/alchemists/{id}
Actualizar un alquimista.

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "name": "Edward Elric",
  "rank": "national",
  "specialty": "combat",
  "certified": true
}
```

**Response 200:**
```json
{
  "id": 1,
  "name": "Edward Elric",
  "email": "edward@amestris.gov",
  "rank": "national",
  "specialty": "combat",
  "role": "alchemist",
  "certified": true,
  "created_at": "2025-11-19T00:00:00Z"
}
```

### DELETE /api/alchemists/{id}
Eliminar un alquimista (soft delete).

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "message": "Alchemist deleted successfully"
}
```

---

## 🎯 Misiones

### GET /api/missions
Listar todas las misiones.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
[
  {
    "id": 1,
    "title": "Investigación de Transmutación Humana",
    "description": "Investigar y documentar los peligros...",
    "status": "in_progress",
    "alchemist_id": 1,
    "alchemist": {
      "id": 1,
      "name": "Edward Elric",
      "email": "edward@amestris.gov"
    },
    "requested_at": "2025-11-09T00:00:00Z",
    "approved_at": "2025-11-10T00:00:00Z",
    "completed_at": null,
    "supervisor_id": 3
  }
]
```

### GET /api/missions/{id}
Obtener una misión por ID.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "id": 1,
  "title": "Investigación de Transmutación Humana",
  "description": "Investigar y documentar los peligros...",
  "status": "in_progress",
  "alchemist_id": 1,
  "requested_at": "2025-11-09T00:00:00Z",
  "approved_at": "2025-11-10T00:00:00Z",
  "completed_at": null,
  "supervisor_id": 3
}
```

### POST /api/missions
Crear una nueva misión.

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "title": "Nueva Misión de Investigación",
  "description": "Descripción detallada de la misión"
}
```

**Response 201:**
```json
{
  "id": 1,
  "title": "Nueva Misión de Investigación",
  "description": "Descripción detallada de la misión",
  "status": "pending",
  "alchemist_id": 1,
  "requested_at": "2025-11-19T00:00:00Z",
  "approved_at": null,
  "completed_at": null,
  "supervisor_id": null
}
```

### PUT /api/missions/{id}
Actualizar una misión.

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "title": "Título Actualizado",
  "description": "Descripción actualizada"
}
```

**Response 200:**
```json
{
  "id": 1,
  "title": "Título Actualizado",
  "description": "Descripción actualizada",
  "status": "pending",
  "alchemist_id": 1,
  "requested_at": "2025-11-19T00:00:00Z"
}
```

### PUT /api/missions/{id}/status
Actualizar el estado de una misión.

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "status": "approved"
}
```

**Estados válidos:** `pending`, `approved`, `in_progress`, `completed`, `cancelled`

**Response 200:**
```json
{
  "id": 1,
  "title": "Investigación de Transmutación Humana",
  "status": "approved",
  "approved_at": "2025-11-19T00:00:00Z"
}
```

### DELETE /api/missions/{id}
Eliminar una misión.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "message": "Mission deleted successfully"
}
```

---

## ⚗️ Materiales

### GET /api/materials
Listar todos los materiales.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
[
  {
    "id": 1,
    "name": "Hierro",
    "type": "metal",
    "description": "Metal común utilizado en transmutaciones básicas",
    "stock": 5000.0,
    "unit": "kg",
    "price": 5.0
  }
]
```

### GET /api/materials/{id}
Obtener un material por ID.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "id": 1,
  "name": "Hierro",
  "type": "metal",
  "description": "Metal común utilizado en transmutaciones básicas",
  "stock": 5000.0,
  "unit": "kg",
  "price": 5.0
}
```

### POST /api/materials
Crear un nuevo material.

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "name": "Hierro",
  "type": "metal",
  "description": "Metal común utilizado en transmutaciones básicas",
  "stock": 5000.0,
  "unit": "kg",
  "price": 5.0
}
```

**Tipos válidos:** `metal`, `mineral`, `organic`, `synthetic`

**Response 201:**
```json
{
  "id": 1,
  "name": "Hierro",
  "type": "metal",
  "description": "Metal común utilizado en transmutaciones básicas",
  "stock": 5000.0,
  "unit": "kg",
  "price": 5.0
}
```

### PUT /api/materials/{id}
Actualizar un material.

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "stock": 6000.0,
  "price": 5.5
}
```

**Response 200:**
```json
{
  "id": 1,
  "name": "Hierro",
  "type": "metal",
  "description": "Metal común utilizado en transmutaciones básicas",
  "stock": 6000.0,
  "unit": "kg",
  "price": 5.5
}
```

### DELETE /api/materials/{id}
Eliminar un material.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "message": "Material deleted successfully"
}
```

---

## 🔬 Transmutaciones

### GET /api/transmutations
Listar todas las transmutaciones.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
[
  {
    "id": 1,
    "alchemist_id": 1,
    "alchemist": {
      "id": 1,
      "name": "Edward Elric",
      "email": "edward@amestris.gov"
    },
    "status": "completed",
    "input_materials": [
      {
        "id": 1,
        "material": {
          "id": 1,
          "name": "Hierro",
          "stock": 5000.0
        },
        "quantity": 500.0,
        "is_input": true
      }
    ],
    "output_materials": [],
    "description": "Transmutación de hierro en acero",
    "cost": 2500.0,
    "result": "Transmutación exitosa",
    "supervisor_id": 3,
    "approved_at": "2025-11-15T00:00:00Z",
    "completed_at": "2025-11-16T00:00:00Z",
    "created_at": "2025-11-15T00:00:00Z"
  }
]
```

### GET /api/transmutations/{id}
Obtener una transmutación por ID.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "id": 1,
  "alchemist_id": 1,
  "status": "completed",
  "input_materials": [...],
  "description": "Transmutación de hierro en acero",
  "cost": 2500.0,
  "result": "Transmutación exitosa",
  "created_at": "2025-11-15T00:00:00Z"
}
```

### POST /api/transmutations
Crear una nueva transmutación.

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "description": "Transmutación de hierro en acero para armamento",
  "input_materials": [
    {
      "material_id": 1,
      "quantity": 500.0
    },
    {
      "material_id": 5,
      "quantity": 50.0
    }
  ],
  "output_materials": [
    {
      "material_id": 2,
      "quantity": 500.0
    }
  ]
}
```

**Response 201:**
```json
{
  "id": 1,
  "alchemist_id": 1,
  "status": "pending",
  "input_materials": [...],
  "description": "Transmutación de hierro en acero para armamento",
  "cost": 2500.0,
  "created_at": "2025-11-19T00:00:00Z"
}
```

### PUT /api/transmutations/{id}
Actualizar una transmutación.

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "description": "Descripción actualizada"
}
```

**Response 200:**
```json
{
  "id": 1,
  "description": "Descripción actualizada",
  "status": "pending"
}
```

### PUT /api/transmutations/{id}/status
Actualizar el estado de una transmutación (solo supervisores).

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "status": "approved",
  "result": "Transmutación aprobada para procesamiento"
}
```

**Estados válidos:** `pending`, `approved`, `completed`, `rejected`

**Response 200:**
```json
{
  "id": 1,
  "status": "approved",
  "result": "Transmutación aprobada para procesamiento",
  "approved_at": "2025-11-19T00:00:00Z",
  "supervisor_id": 3
}
```

### DELETE /api/transmutations/{id}
Eliminar una transmutación.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "message": "Transmutation deleted successfully"
}
```

---

## 🔍 Auditorías

### GET /api/audits
Listar todas las auditorías.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
[
  {
    "id": 1,
    "type": "material_usage",
    "severity": "high",
    "description": "Uso excesivo de hierro detectado",
    "alchemist_id": 1,
    "alchemist": {
      "id": 1,
      "name": "Edward Elric"
    },
    "details": "{\"material\": \"Hierro\", \"usage\": 1500}",
    "resolved": false,
    "resolved_at": null,
    "resolved_by": null,
    "created_at": "2025-11-19T00:00:00Z"
  }
]
```

### GET /api/audits/{id}
Obtener una auditoría por ID.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "id": 1,
  "type": "material_usage",
  "severity": "high",
  "description": "Uso excesivo de hierro detectado",
  "alchemist_id": 1,
  "details": "{\"material\": \"Hierro\", \"usage\": 1500}",
  "resolved": false,
  "created_at": "2025-11-19T00:00:00Z"
}
```

### POST /api/audits
Crear una nueva auditoría.

**Headers:**
```
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "type": "material_usage",
  "severity": "high",
  "description": "Uso excesivo de materiales detectado",
  "alchemist_id": 1,
  "details": "{\"material\": \"Hierro\", \"usage\": 1500}"
}
```

**Tipos válidos:** `material_usage`, `mission_check`, `transmutation`, `system`

**Severidades válidas:** `low`, `medium`, `high`, `critical`

**Response 201:**
```json
{
  "id": 1,
  "type": "material_usage",
  "severity": "high",
  "description": "Uso excesivo de materiales detectado",
  "alchemist_id": 1,
  "resolved": false,
  "created_at": "2025-11-19T00:00:00Z"
}
```

### PUT /api/audits/{id}/resolve
Resolver una auditoría (solo supervisores).

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "id": 1,
  "type": "material_usage",
  "severity": "high",
  "description": "Uso excesivo de materiales detectado",
  "resolved": true,
  "resolved_at": "2025-11-19T00:00:00Z",
  "resolved_by": 3
}
```

### DELETE /api/audits/{id}
Eliminar una auditoría.

**Headers:**
```
Authorization: Bearer {token}
```

**Response 200:**
```json
{
  "message": "Audit deleted successfully"
}
```

---

## 🔌 WebSocket

### GET /api/ws?token={jwt_token}
Establecer conexión WebSocket para notificaciones en tiempo real.

**Query Parameters:**
- `token`: JWT token de autenticación

**Mensajes recibidos:**
```json
{
  "type": "mission_status_changed",
  "data": {
    "mission_id": 1,
    "status": "approved"
  }
}
```

```json
{
  "type": "transmutation_status_changed",
  "data": {
    "transmutation_id": 1,
    "status": "completed"
  }
}
```

```json
{
  "type": "audit_created",
  "data": {
    "audit_id": 1,
    "severity": "high"
  }
}
```

---

## 📊 Códigos de Estado HTTP

- `200 OK` - Operación exitosa
- `201 Created` - Recurso creado exitosamente
- `400 Bad Request` - Solicitud inválida
- `401 Unauthorized` - No autenticado o token inválido
- `403 Forbidden` - No tiene permisos para esta operación
- `404 Not Found` - Recurso no encontrado
- `500 Internal Server Error` - Error del servidor

---

## 🔒 Autenticación y Autorización

- Todos los endpoints excepto `/api/auth/login` y `/api/auth/register` requieren autenticación JWT
- El token debe incluirse en el header `Authorization: Bearer {token}`
- Los supervisores tienen permisos adicionales para:
  - Aprobar/rechazar misiones
  - Aprobar/rechazar transmutaciones
  - Resolver auditorías
  - Crear alquimistas

---

## ⚠️ Errores

### Error de Autenticación
```json
{
  "error": "Unauthorized"
}
```

### Error de Validación
```json
{
  "error": "Invalid request body",
  "details": "Email is required"
}
```

### Error de Recurso No Encontrado
```json
{
  "error": "Resource not found"
}
```

---

## 📝 Notas

- Todas las fechas están en formato ISO 8601 (UTC)
- Los IDs son números enteros positivos
- Los estados y tipos tienen valores predefinidos (ver documentación de cada endpoint)
- El sistema procesa transmutaciones de forma asíncrona
- Las auditorías se generan automáticamente por el sistema

