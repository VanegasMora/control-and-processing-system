@echo off
REM Script para configurar variables de entorno para PostgreSQL local
REM Ejecutar: setup-local-postgres.bat

set POSTGRES_HOST=localhost
set POSTGRES_USER=postgres
set POSTGRES_PASSWORD=tu_password
set POSTGRES_DB=amestris_db
set JWT_SECRET=amestris-alchemy-secret-key-change-in-production

echo Variables de entorno configuradas para PostgreSQL local
echo POSTGRES_HOST=%POSTGRES_HOST%
echo POSTGRES_USER=%POSTGRES_USER%
echo POSTGRES_DB=%POSTGRES_DB%
echo.
echo Ahora puedes ejecutar: go run main.go

