-- Script SQL para insertar datos de prueba
-- Ejecutar con: psql -U amestris -d amestris_db -f scripts/seed.sql
-- O desde Docker: docker exec -i amestris-postgres psql -U amestris -d amestris_db < scripts/seed.sql

-- Insertar materiales
INSERT INTO materials (created_at, updated_at, name, type, description, stock, unit, price) VALUES
(NOW(), NOW(), 'Hierro', 'metal', 'Metal común utilizado en transmutaciones básicas', 1000.0, 'kg', 5.0),
(NOW(), NOW(), 'Acero', 'metal', 'Aleación de hierro y carbono', 500.0, 'kg', 10.0),
(NOW(), NOW(), 'Oro', 'metal', 'Metal precioso', 50.0, 'kg', 500.0),
(NOW(), NOW(), 'Carbón', 'mineral', 'Combustible y material de transmutación', 2000.0, 'kg', 2.0),
(NOW(), NOW(), 'Agua', 'organic', 'Elemento básico para transmutaciones', 5000.0, 'L', 0.5),
(NOW(), NOW(), 'Tierra', 'mineral', 'Material base para construcciones', 10000.0, 'kg', 1.0);

-- Insertar alquimistas
-- Nota: Las contraseñas están hasheadas con bcrypt. La contraseña es "password123" para todos
INSERT INTO alchemists (created_at, updated_at, name, email, password, rank, specialty, role, certified) VALUES
(NOW(), NOW(), 'Edward Elric', 'edward@amestris.gov', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'state', 'combat', 'alchemist', true),
(NOW(), NOW(), 'Alphonse Elric', 'alphonse@amestris.gov', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'state', 'research', 'alchemist', true),
(NOW(), NOW(), 'Roy Mustang', 'mustang@amestris.gov', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'national', 'combat', 'supervisor', true);

-- Insertar misiones
INSERT INTO missions (created_at, updated_at, title, description, status, alchemist_id, requested_at, approved_at, completed_at) VALUES
(NOW(), NOW(), 'Investigación de Transmutación Humana', 'Investigar y documentar los peligros de la transmutación humana prohibida', 'in_progress', 2, NOW() - INTERVAL '10 days', NOW() - INTERVAL '9 days', NULL),
(NOW(), NOW(), 'Protección de la Capital', 'Mantener la seguridad de Central durante el período de transición', 'pending', 1, NOW() - INTERVAL '2 days', NULL, NULL),
(NOW(), NOW(), 'Desarrollo de Nuevos Materiales', 'Investigar materiales alquímicos más eficientes para transmutaciones', 'completed', 2, NOW() - INTERVAL '30 days', NOW() - INTERVAL '30 days', NOW() - INTERVAL '5 days');

-- Insertar transmutaciones
INSERT INTO transmutations (created_at, updated_at, alchemist_id, status, description, cost, result, approved_at, completed_at, supervisor_id) VALUES
(NOW(), NOW(), 1, 'completed', 'Transmutación de hierro en acero para armamento', 500.0, 'Transmutación exitosa: 100kg de acero producido', NOW() - INTERVAL '7 days', NOW() - INTERVAL '6 days', 3),
(NOW(), NOW(), 2, 'pending', 'Transmutación de agua en hielo para experimentos', 50.0, NULL, NULL, NULL, NULL),
(NOW(), NOW(), 1, 'approved', 'Reparación de infraestructura usando transmutación de tierra', 200.0, NULL, NOW() - INTERVAL '1 day', NULL, 3);

-- Insertar materiales de transmutación
INSERT INTO transmutation_materials (created_at, updated_at, transmutation_id, material_id, quantity, is_input) VALUES
-- Transmutación 1 (hierro -> acero)
(NOW(), NOW(), 1, 1, 100.0, true),  -- Hierro (entrada)
(NOW(), NOW(), 1, 3, 1.0, true),   -- Oro (entrada)
(NOW(), NOW(), 1, 2, 100.0, false), -- Acero (salida)
-- Transmutación 2 (agua -> hielo)
(NOW(), NOW(), 2, 5, 10.0, true),  -- Agua (entrada)
-- Transmutación 3 (tierra)
(NOW(), NOW(), 3, 6, 50.0, true);  -- Tierra (entrada)

-- Insertar auditorías
INSERT INTO audits (created_at, updated_at, type, severity, description, alchemist_id, details, resolved, resolved_at, resolved_by) VALUES
(NOW(), NOW(), 'material_usage', 'medium', 'Uso excesivo de hierro detectado en el último mes', 1, '{"material": "Hierro", "usage": 500, "threshold": 300}', false, NULL, NULL),
(NOW(), NOW(), 'mission_check', 'low', 'Misión pendiente por más de 7 días', 1, '{"mission_id": 2, "days_pending": 2}', false, NULL, NULL),
(NOW(), NOW(), 'system', 'high', 'Verificación de seguridad del sistema completada', NULL, '{"check_type": "security", "status": "passed"}', true, NOW() - INTERVAL '1 day', 3);

