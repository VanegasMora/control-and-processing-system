# Datos de Prueba

Este documento explica cómo insertar datos de prueba en la base de datos.

## Opción 1: Usando el script Go (Recomendado)

El script Go se conecta a la base de datos y crea todos los datos de prueba usando GORM.

### Desde Docker:

```bash
# Ejecutar el script dentro del contenedor del backend
docker exec -it amestris-backend go run /app/scripts/seed.go
```

### Desde tu máquina local:

1. Asegúrate de tener las variables de entorno configuradas:
```powershell
$env:POSTGRES_HOST="localhost"
$env:POSTGRES_USER="amestris"
$env:POSTGRES_PASSWORD="alchemy123"
$env:POSTGRES_DB="amestris_db"
```

2. Ejecuta el script:
```bash
go run scripts/seed.go
```

## Opción 2: Usando SQL

### Desde Docker:

```bash
docker exec -i amestris-postgres psql -U amestris -d amestris_db < scripts/seed.sql
```

### Desde tu máquina local:

```bash
psql -U amestris -d amestris_db -f scripts/seed.sql
```

## Datos de Prueba Incluidos

### Materiales (6)
- Hierro, Acero, Oro, Carbón, Agua, Tierra

### Alquimistas (3)
1. **Edward Elric** (edward@amestris.gov)
   - Rango: Estatal
   - Especialidad: Combate
   - Contraseña: `password123`

2. **Alphonse Elric** (alphonse@amestris.gov)
   - Rango: Estatal
   - Especialidad: Investigación
   - Contraseña: `password123`

3. **Roy Mustang** (mustang@amestris.gov)
   - Rango: Nacional
   - Especialidad: Combate
   - Rol: Supervisor
   - Contraseña: `password123`

### Misiones (3)
- Investigación de Transmutación Humana (en progreso)
- Protección de la Capital (pendiente)
- Desarrollo de Nuevos Materiales (completada)

### Transmutaciones (3)
- Transmutación de hierro en acero (completada)
- Transmutación de agua en hielo (pendiente)
- Reparación de infraestructura (aprobada)

### Auditorías (3)
- Uso excesivo de hierro (sin resolver)
- Misión pendiente por más de 7 días (sin resolver)
- Verificación de seguridad del sistema (resuelta)

## Notas

- Todas las contraseñas de prueba son: `password123`
- Los datos están diseñados para mostrar diferentes estados y relaciones
- Puedes ejecutar el script múltiples veces (usará `ON CONFLICT` o verificará existencia)

