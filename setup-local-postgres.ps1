# Script para configurar variables de entorno para PostgreSQL local
# Ejecutar en PowerShell: .\setup-local-postgres.ps1

$env:POSTGRES_HOST = "localhost"
$env:POSTGRES_USER = "postgres"  # O tu usuario de PostgreSQL
$env:POSTGRES_PASSWORD = "tu_password"  # Cambia esto por tu contraseña
$env:POSTGRES_DB = "amestris_db"
$env:JWT_SECRET = "amestris-alchemy-secret-key-change-in-production"

Write-Host "Variables de entorno configuradas para PostgreSQL local"
Write-Host "POSTGRES_HOST: $env:POSTGRES_HOST"
Write-Host "POSTGRES_USER: $env:POSTGRES_USER"
Write-Host "POSTGRES_DB: $env:POSTGRES_DB"
Write-Host ""
Write-Host "Ahora puedes ejecutar: go run main.go"

