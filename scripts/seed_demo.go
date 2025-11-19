package main

import (
	"backend-avanzada/auth"
	"backend-avanzada/models"
	"fmt"
	"math/rand"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Conectar a la base de datos
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("POSTGRES_HOST", "localhost"),
		getEnv("POSTGRES_USER", "amestris"),
		getEnv("POSTGRES_PASSWORD", "alchemy123"),
		getEnv("POSTGRES_DB", "amestris_db"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error conectando a la base de datos: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("🗄️  Conectado a la base de datos")
	fmt.Println("📦 Insertando datos de demostración...\n")

	// Limpiar datos existentes (opcional - comentar si quieres mantener datos)
	fmt.Println("🧹 Limpiando datos existentes...")
	db.Exec("TRUNCATE TABLE transmutation_materials CASCADE")
	db.Exec("TRUNCATE TABLE transmutations CASCADE")
	db.Exec("TRUNCATE TABLE missions CASCADE")
	db.Exec("TRUNCATE TABLE audits CASCADE")
	db.Exec("TRUNCATE TABLE materials CASCADE")
	db.Exec("TRUNCATE TABLE alchemists CASCADE")

	// Crear materiales primero
	fmt.Println("📦 Creando materiales...")
	materials := []models.Material{
		{Name: "Hierro", Type: models.MaterialTypeMetal, Description: "Metal común utilizado en transmutaciones básicas", Stock: 5000.0, Unit: "kg", Price: 5.0},
		{Name: "Acero", Type: models.MaterialTypeMetal, Description: "Aleación de hierro y carbono, más resistente", Stock: 2000.0, Unit: "kg", Price: 10.0},
		{Name: "Oro", Type: models.MaterialTypeMetal, Description: "Metal precioso para transmutaciones avanzadas", Stock: 100.0, Unit: "kg", Price: 500.0},
		{Name: "Plata", Type: models.MaterialTypeMetal, Description: "Metal precioso con propiedades alquímicas", Stock: 200.0, Unit: "kg", Price: 300.0},
		{Name: "Carbón", Type: models.MaterialTypeMineral, Description: "Combustible y material de transmutación", Stock: 10000.0, Unit: "kg", Price: 2.0},
		{Name: "Agua", Type: models.MaterialTypeOrganic, Description: "Elemento básico para transmutaciones", Stock: 50000.0, Unit: "L", Price: 0.5},
		{Name: "Tierra", Type: models.MaterialTypeMineral, Description: "Material base para construcciones", Stock: 50000.0, Unit: "kg", Price: 1.0},
		{Name: "Arena", Type: models.MaterialTypeMineral, Description: "Material para vidrio y cerámica", Stock: 30000.0, Unit: "kg", Price: 0.8},
		{Name: "Madera", Type: models.MaterialTypeOrganic, Description: "Material orgánico para estructuras", Stock: 20000.0, Unit: "kg", Price: 3.0},
		{Name: "Piedra", Type: models.MaterialTypeMineral, Description: "Material de construcción", Stock: 40000.0, Unit: "kg", Price: 1.5},
		{Name: "Cristal", Type: models.MaterialTypeSynthetic, Description: "Material sintético para experimentos", Stock: 500.0, Unit: "kg", Price: 50.0},
		{Name: "Mercurio", Type: models.MaterialTypeMetal, Description: "Metal líquido para transmutaciones especiales", Stock: 50.0, Unit: "L", Price: 200.0},
	}

	for i := range materials {
		if err := db.Create(&materials[i]).Error; err != nil {
			fmt.Printf("❌ Error creando material %s: %v\n", materials[i].Name, err)
		} else {
			fmt.Printf("  ✓ %s (Stock: %.1f %s)\n", materials[i].Name, materials[i].Stock, materials[i].Unit)
		}
	}

	// Crear alquimistas
	fmt.Println("\n👥 Creando alquimistas...")
	hashedPassword, _ := auth.HashPassword("password123")

	alchemists := []models.Alchemist{
		{
			Name:      "Edward Elric",
			Email:     "edward@amestris.gov",
			Password:  hashedPassword,
			Rank:      models.RankState,
			Specialty: models.SpecialtyCombat,
			Role:      models.RoleAlchemist,
			Certified: true,
		},
		{
			Name:      "Alphonse Elric",
			Email:     "alphonse@amestris.gov",
			Password:  hashedPassword,
			Rank:      models.RankState,
			Specialty: models.SpecialtyResearch,
			Role:      models.RoleAlchemist,
			Certified: true,
		},
		{
			Name:      "Roy Mustang",
			Email:     "mustang@amestris.gov",
			Password:  hashedPassword,
			Rank:      models.RankNational,
			Specialty: models.SpecialtyCombat,
			Role:      models.RoleSupervisor,
			Certified: true,
		},
		{
			Name:      "Riza Hawkeye",
			Email:     "hawkeye@amestris.gov",
			Password:  hashedPassword,
			Rank:      models.RankState,
			Specialty: models.SpecialtyCombat,
			Role:      models.RoleAlchemist,
			Certified: true,
		},
		{
			Name:      "Winry Rockbell",
			Email:     "winry@amestris.gov",
			Password:  hashedPassword,
			Rank:      models.RankApprentice,
			Specialty: models.SpecialtyIndustrial,
			Role:      models.RoleAlchemist,
			Certified: false,
		},
		{
			Name:      "Izumi Curtis",
			Email:     "izumi@amestris.gov",
			Password:  hashedPassword,
			Rank:      models.RankNational,
			Specialty: models.SpecialtyMedical,
			Role:      models.RoleAlchemist,
			Certified: true,
		},
		{
			Name:      "Maes Hughes",
			Email:     "hughes@amestris.gov",
			Password:  hashedPassword,
			Rank:      models.RankState,
			Specialty: models.SpecialtyResearch,
			Role:      models.RoleSupervisor,
			Certified: true,
		},
	}

	for i := range alchemists {
		if err := db.Create(&alchemists[i]).Error; err != nil {
			fmt.Printf("❌ Error creando alquimista %s: %v\n", alchemists[i].Name, err)
		} else {
			roleStr := "Alquimista"
			if alchemists[i].Role == models.RoleSupervisor {
				roleStr = "Supervisor"
			}
			fmt.Printf("  ✓ %s (%s) - %s\n", alchemists[i].Name, alchemists[i].Email, roleStr)
		}
	}

	// Crear misiones
	fmt.Println("\n📋 Creando misiones...")
	now := time.Now()
	rand.Seed(time.Now().UnixNano())

	missionTitles := []string{
		"Investigación de Transmutación Humana",
		"Protección de la Capital Central",
		"Desarrollo de Nuevos Materiales Alquímicos",
		"Reparación de Infraestructura en Lior",
		"Análisis de Transmutaciones Prohibidas",
		"Seguridad en la Frontera Este",
		"Investigación de Filosofales",
		"Mejora de Sistemas de Transmutación",
		"Auditoría de Materiales en Depósitos",
		"Capacitación de Nuevos Alquimistas",
		"Investigación de Transmutación Animal",
		"Protección de Laboratorios de Investigación",
	}

	missionDescriptions := []string{
		"Investigar y documentar los peligros de la transmutación humana prohibida",
		"Mantener la seguridad de Central durante el período de transición política",
		"Investigar materiales alquímicos más eficientes para transmutaciones",
		"Reparar estructuras dañadas usando transmutación de tierra y piedra",
		"Analizar casos históricos de transmutaciones prohibidas",
		"Reforzar las defensas en la frontera este del país",
		"Investigar la existencia y propiedades de las piedras filosofales",
		"Desarrollar sistemas más eficientes de transmutación",
		"Realizar inventario y auditoría de materiales almacenados",
		"Capacitar a nuevos alquimistas en técnicas avanzadas",
		"Investigar transmutaciones animales para aplicaciones médicas",
		"Asegurar la protección de laboratorios de investigación alquímica",
	}

	statuses := []models.MissionStatus{
		models.MissionStatusPending,
		models.MissionStatusApproved,
		models.MissionStatusInProgress,
		models.MissionStatusCompleted,
	}

	missions := []models.Mission{}
	for i := 0; i < len(missionTitles); i++ {
		status := statuses[rand.Intn(len(statuses))]
		alchemistID := alchemists[rand.Intn(len(alchemists))].ID

		mission := models.Mission{
			Title:       missionTitles[i],
			Description: missionDescriptions[i],
			Status:      status,
			AlchemistID: alchemistID,
			RequestedAt: now.AddDate(0, 0, -rand.Intn(30)),
		}

		if status == models.MissionStatusApproved || status == models.MissionStatusInProgress {
			approvedAt := mission.RequestedAt.AddDate(0, 0, rand.Intn(5))
			mission.ApprovedAt = &approvedAt
			if alchemists[2].ID != 0 { // Mustang como supervisor
				mission.SupervisorID = &alchemists[2].ID
			}
		}

		if status == models.MissionStatusCompleted {
			approvedAt := mission.RequestedAt.AddDate(0, 0, rand.Intn(5))
			mission.ApprovedAt = &approvedAt
			completedAt := approvedAt.AddDate(0, 0, rand.Intn(20)+5)
			mission.CompletedAt = &completedAt
			if alchemists[2].ID != 0 {
				mission.SupervisorID = &alchemists[2].ID
			}
		}

		if err := db.Create(&mission).Error; err != nil {
			fmt.Printf("❌ Error creando misión: %v\n", err)
		} else {
			fmt.Printf("  ✓ %s [%s]\n", mission.Title, mission.Status)
		}
		missions = append(missions, mission)
	}

	// Crear transmutaciones
	fmt.Println("\n⚗️  Creando transmutaciones...")

	transmutationDescriptions := []string{
		"Transmutación de hierro en acero para armamento",
		"Transmutación de agua en hielo para experimentos",
		"Reparación de infraestructura usando transmutación de tierra",
		"Creación de herramientas mediante transmutación de metal",
		"Transmutación de arena en vidrio",
		"Conversión de madera en carbón",
		"Transmutación de piedra en materiales de construcción",
		"Creación de aleaciones especiales",
		"Transmutación de materiales orgánicos",
		"Procesamiento de minerales mediante transmutación",
		"Transmutación de metales preciosos",
		"Creación de materiales sintéticos",
		"Transmutación de agua purificada",
		"Conversión de materiales de desecho",
		"Transmutación de cristales alquímicos",
	}

	transStatuses := []models.TransmutationStatus{
		models.TransmutationStatusPending,
		models.TransmutationStatusApproved,
		models.TransmutationStatusCompleted,
		models.TransmutationStatusFailed,
	}

	for i := 0; i < 15; i++ {
		status := transStatuses[rand.Intn(len(transStatuses))]
		alchemistID := alchemists[rand.Intn(len(alchemists))].ID

		// Calcular costo
		cost := float64(rand.Intn(1000) + 50)

		transmutation := models.Transmutation{
			AlchemistID: alchemistID,
			Status:      status,
			Description: transmutationDescriptions[i],
			Cost:        cost,
		}

		if status == models.TransmutationStatusApproved || status == models.TransmutationStatusCompleted {
			approvedAt := now.AddDate(0, 0, -rand.Intn(10))
			transmutation.ApprovedAt = &approvedAt
			if alchemists[2].ID != 0 {
				transmutation.SupervisorID = &alchemists[2].ID
			}
		}

		if status == models.TransmutationStatusCompleted {
			completedAt := now.AddDate(0, 0, -rand.Intn(5))
			transmutation.CompletedAt = &completedAt
			transmutation.Result = fmt.Sprintf("Transmutación completada exitosamente. Resultado: %.0f%% de eficiencia", float64(rand.Intn(30)+70))
		}

		if status == models.TransmutationStatusFailed {
			transmutation.Result = "Transmutación fallida: Error en el proceso de conversión"
		}

		if err := db.Create(&transmutation).Error; err != nil {
			fmt.Printf("❌ Error creando transmutación: %v\n", err)
			continue
		}

		// Agregar materiales de entrada
		inputCount := rand.Intn(3) + 1
		for j := 0; j < inputCount; j++ {
			material := materials[rand.Intn(len(materials))]
			quantity := float64(rand.Intn(100) + 10)

			tm := models.TransmutationMaterial{
				TransmutationID: transmutation.ID,
				MaterialID:      material.ID,
				Quantity:        quantity,
				IsInput:         true,
			}
			db.Create(&tm)
		}

		// Agregar materiales de salida (solo si está completada)
		if status == models.TransmutationStatusCompleted {
			outputCount := rand.Intn(2) + 1
			for j := 0; j < outputCount; j++ {
				material := materials[rand.Intn(len(materials))]
				quantity := float64(rand.Intn(50) + 5)

				tm := models.TransmutationMaterial{
					TransmutationID: transmutation.ID,
					MaterialID:      material.ID,
					Quantity:        quantity,
					IsInput:         false,
				}
				db.Create(&tm)
			}
		}

		fmt.Printf("  ✓ %s [%s] - Costo: %.2f\n", transmutation.Description, transmutation.Status, transmutation.Cost)
	}

	// Crear auditorías
	fmt.Println("\n🔍 Creando auditorías...")

	auditTypes := []models.AuditType{
		models.AuditTypeMaterialUsage,
		models.AuditTypeMissionCheck,
		models.AuditTypeTransmutation,
		models.AuditTypeSystem,
	}

	severities := []models.AuditSeverity{
		models.AuditSeverityLow,
		models.AuditSeverityMedium,
		models.AuditSeverityHigh,
		models.AuditSeverityCritical,
	}

	auditDescriptions := []string{
		"Uso excesivo de hierro detectado en el último mes",
		"Misión pendiente por más de 7 días sin actualización",
		"Transmutación con eficiencia por debajo del estándar",
		"Verificación de seguridad del sistema completada",
		"Materiales faltantes en inventario",
		"Transmutación no autorizada detectada",
		"Uso de materiales sin aprobación previa",
		"Auditoría rutinaria de procesos alquímicos",
		"Verificación de cumplimiento de protocolos",
		"Análisis de patrones de uso de materiales",
		"Revisión de transmutaciones recientes",
		"Verificación de certificaciones de alquimistas",
	}

	for i := 0; i < 12; i++ {
		auditType := auditTypes[rand.Intn(len(auditTypes))]
		severity := severities[rand.Intn(len(severities))]
		resolved := rand.Float32() < 0.4 // 40% resueltas

		var alchemistID *uint
		if rand.Float32() < 0.7 { // 70% tienen alquimista asociado
			id := alchemists[rand.Intn(len(alchemists))].ID
			alchemistID = &id
		}

		audit := models.Audit{
			Type:        auditType,
			Severity:    severity,
			Description: auditDescriptions[i],
			AlchemistID: alchemistID,
			Details:     fmt.Sprintf(`{"check_id": %d, "timestamp": "%s"}`, i+1, time.Now().Format(time.RFC3339)),
			Resolved:    resolved,
		}

		if resolved {
			resolvedAt := now.AddDate(0, 0, -rand.Intn(5))
			audit.ResolvedAt = &resolvedAt
			if alchemists[2].ID != 0 {
				audit.ResolvedBy = &alchemists[2].ID
			}
		}

		if err := db.Create(&audit).Error; err != nil {
			fmt.Printf("❌ Error creando auditoría: %v\n", err)
		} else {
			resolvedStr := "Pendiente"
			if audit.Resolved {
				resolvedStr = "Resuelta"
			}
			fmt.Printf("  ✓ %s [%s] - %s\n", audit.Description, audit.Severity, resolvedStr)
		}
	}

	fmt.Println("\n✅ Datos de demostración insertados correctamente!")
	fmt.Println("\n📊 Resumen:")
	fmt.Printf("  • %d materiales\n", len(materials))
	fmt.Printf("  • %d alquimistas\n", len(alchemists))
	fmt.Printf("  • %d misiones\n", len(missions))
	fmt.Printf("  • 15 transmutaciones\n")
	fmt.Printf("  • 12 auditorías\n")
	fmt.Println("\n🔑 Credenciales de acceso:")
	fmt.Println("  Todos los usuarios tienen la contraseña: password123")
	fmt.Println("\n  👤 Alquimistas:")
	for _, a := range alchemists {
		if a.Role == models.RoleAlchemist {
			fmt.Printf("    • %s (%s)\n", a.Name, a.Email)
		}
	}
	fmt.Println("\n  👔 Supervisores:")
	for _, a := range alchemists {
		if a.Role == models.RoleSupervisor {
			fmt.Printf("    • %s (%s)\n", a.Name, a.Email)
		}
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
