package main

import (
	"backend-avanzada/auth"
	"backend-avanzada/models"
	"fmt"
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

	fmt.Println("Conectado a la base de datos. Limpiando datos existentes...")

	// Limpiar datos existentes (en orden inverso de dependencias)
	db.Exec("DELETE FROM transmutation_materials")
	db.Exec("DELETE FROM audits")
	db.Exec("DELETE FROM transmutations")
	db.Exec("DELETE FROM missions")
	db.Exec("DELETE FROM materials")
	db.Exec("DELETE FROM alchemists")

	fmt.Println("Datos limpiados. Insertando nuevos datos de prueba...")

	// Crear materiales primero
	materials := []models.Material{
		{Name: "Hierro", Type: models.MaterialTypeMetal, Description: "Metal común utilizado en transmutaciones básicas", Stock: 5000.0, Unit: "kg", Price: 5.0},
		{Name: "Acero", Type: models.MaterialTypeMetal, Description: "Aleación de hierro y carbono", Stock: 2500.0, Unit: "kg", Price: 10.0},
		{Name: "Oro", Type: models.MaterialTypeMetal, Description: "Metal precioso para transmutaciones avanzadas", Stock: 100.0, Unit: "kg", Price: 500.0},
		{Name: "Plata", Type: models.MaterialTypeMetal, Description: "Metal precioso de uso alquímico", Stock: 200.0, Unit: "kg", Price: 300.0},
		{Name: "Carbón", Type: models.MaterialTypeMineral, Description: "Combustible y material de transmutación", Stock: 10000.0, Unit: "kg", Price: 2.0},
		{Name: "Agua", Type: models.MaterialTypeOrganic, Description: "Elemento básico para transmutaciones", Stock: 20000.0, Unit: "L", Price: 0.5},
		{Name: "Tierra", Type: models.MaterialTypeMineral, Description: "Material base para construcciones", Stock: 50000.0, Unit: "kg", Price: 1.0},
		{Name: "Cristal Alquímico", Type: models.MaterialTypeSynthetic, Description: "Cristal sintético para transmutaciones complejas", Stock: 150.0, Unit: "kg", Price: 1000.0},
		{Name: "Mercurio", Type: models.MaterialTypeMetal, Description: "Metal líquido para experimentos", Stock: 50.0, Unit: "L", Price: 800.0},
		{Name: "Fósforo", Type: models.MaterialTypeMineral, Description: "Mineral reactivo para transmutaciones", Stock: 300.0, Unit: "kg", Price: 150.0},
	}

	for i := range materials {
		if err := db.Create(&materials[i]).Error; err != nil {
			fmt.Printf("Error creando material %s: %v\n", materials[i].Name, err)
		} else {
			fmt.Printf("✓ Material creado: %s\n", materials[i].Name)
		}
	}

	// Crear alquimistas
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
			Name:      "Alex Louis Armstrong",
			Email:     "armstrong@amestris.gov",
			Password:  hashedPassword,
			Rank:      models.RankNational,
			Specialty: models.SpecialtyCombat,
			Role:      models.RoleSupervisor,
			Certified: true,
		},
		{
			Name:      "Shou Tucker",
			Email:     "tucker@amestris.gov",
			Password:  hashedPassword,
			Rank:      models.RankState,
			Specialty: models.SpecialtyResearch,
			Role:      models.RoleAlchemist,
			Certified: true,
		},
		{
			Name:      "Izumi Curtis",
			Email:     "izumi@amestris.gov",
			Password:  hashedPassword,
			Rank:      models.RankNational,
			Specialty: models.SpecialtyCombat,
			Role:      models.RoleAlchemist,
			Certified: true,
		},
	}

	for i := range alchemists {
		if err := db.Create(&alchemists[i]).Error; err != nil {
			fmt.Printf("Error creando alquimista %s: %v\n", alchemists[i].Name, err)
		} else {
			fmt.Printf("✓ Alquimista creado: %s (%s)\n", alchemists[i].Name, alchemists[i].Email)
		}
	}

	// Crear misiones
	now := time.Now()
	missions := []models.Mission{
		{
			Title:        "Investigación de Transmutación Humana",
			Description:  "Investigar y documentar los peligros de la transmutación humana prohibida. Revisar casos históricos y establecer protocolos de seguridad.",
			Status:       models.MissionStatusInProgress,
			AlchemistID:  alchemists[1].ID, // Alphonse
			RequestedAt:  now.AddDate(0, 0, -10),
			ApprovedAt:   &[]time.Time{now.AddDate(0, 0, -9)}[0],
			SupervisorID: &alchemists[2].ID, // Mustang
		},
		{
			Title:       "Protección de la Capital",
			Description: "Mantener la seguridad de Central durante el período de transición política. Patrullas regulares y respuesta a emergencias.",
			Status:      models.MissionStatusPending,
			AlchemistID: alchemists[0].ID, // Edward
			RequestedAt: now.AddDate(0, 0, -2),
		},
		{
			Title:        "Desarrollo de Nuevos Materiales",
			Description:  "Investigar materiales alquímicos más eficientes para transmutaciones. Enfoque en reducir costos y mejorar resultados.",
			Status:       models.MissionStatusCompleted,
			AlchemistID:  alchemists[1].ID, // Alphonse
			RequestedAt:  now.AddDate(0, -1, 0),
			ApprovedAt:   &[]time.Time{now.AddDate(0, -1, 0)}[0],
			CompletedAt:  &[]time.Time{now.AddDate(0, 0, -5)}[0],
			SupervisorID: &alchemists[2].ID, // Mustang
		},
		{
			Title:        "Reparación de Infraestructura en Resembool",
			Description:  "Restaurar edificios dañados usando transmutación de tierra y materiales locales.",
			Status:       models.MissionStatusInProgress,
			AlchemistID:  alchemists[0].ID, // Edward
			RequestedAt:  now.AddDate(0, 0, -7),
			ApprovedAt:   &[]time.Time{now.AddDate(0, 0, -6)}[0],
			SupervisorID: &alchemists[5].ID, // Armstrong
		},
		{
			Title:        "Capacitación de Aprendices",
			Description:  "Impartir clases sobre principios básicos de alquimia y seguridad en transmutaciones.",
			Status:       models.MissionStatusCompleted,
			AlchemistID:  alchemists[7].ID, // Izumi
			RequestedAt:  now.AddDate(0, 0, -20),
			ApprovedAt:   &[]time.Time{now.AddDate(0, 0, -19)}[0],
			CompletedAt:  &[]time.Time{now.AddDate(0, 0, -10)}[0],
			SupervisorID: &alchemists[5].ID, // Armstrong
		},
		{
			Title:       "Investigación de Automail",
			Description: "Mejorar la tecnología de automail usando transmutaciones avanzadas de metales.",
			Status:      models.MissionStatusPending,
			AlchemistID: alchemists[4].ID, // Winry
			RequestedAt: now.AddDate(0, 0, -1),
		},
		{
			Title:        "Auditoría de Materiales en Depósitos",
			Description:  "Verificar inventario y estado de materiales almacenados en los depósitos estatales.",
			Status:       models.MissionStatusInProgress,
			AlchemistID:  alchemists[3].ID, // Hawkeye
			RequestedAt:  now.AddDate(0, 0, -5),
			ApprovedAt:   &[]time.Time{now.AddDate(0, 0, -4)}[0],
			SupervisorID: &alchemists[2].ID, // Mustang
		},
		{
			Title:       "Desarrollo de Armas Alquímicas",
			Description: "Crear y probar nuevas armas basadas en transmutación para el ejército.",
			Status:      models.MissionStatusPending,
			AlchemistID: alchemists[6].ID, // Tucker
			RequestedAt: now.AddDate(0, 0, -3),
		},
	}

	for i := range missions {
		if err := db.Create(&missions[i]).Error; err != nil {
			fmt.Printf("Error creando misión %s: %v\n", missions[i].Title, err)
		} else {
			fmt.Printf("✓ Misión creada: %s\n", missions[i].Title)
		}
	}

	// Crear transmutaciones
	transmutations := []models.Transmutation{
		{
			AlchemistID:  alchemists[0].ID, // Edward
			Status:       models.TransmutationStatusCompleted,
			Description:  "Transmutación de hierro en acero para armamento militar",
			Cost:         2500.0,
			Result:       "Transmutación exitosa: 500kg de acero de alta calidad producido. Material verificado y aprobado para uso militar.",
			ApprovedAt:   &[]time.Time{now.AddDate(0, 0, -15)}[0],
			CompletedAt:  &[]time.Time{now.AddDate(0, 0, -14)}[0],
			SupervisorID: &alchemists[2].ID, // Mustang
		},
		{
			AlchemistID: alchemists[1].ID, // Alphonse
			Status:      models.TransmutationStatusPending,
			Description: "Transmutación de agua en hielo para experimentos de conservación",
			Cost:        250.0,
		},
		{
			AlchemistID:  alchemists[0].ID, // Edward
			Status:       models.TransmutationStatusApproved,
			Description:  "Reparación de infraestructura usando transmutación de tierra y roca",
			Cost:         1000.0,
			ApprovedAt:   &[]time.Time{now.AddDate(0, 0, -2)}[0],
			SupervisorID: &alchemists[5].ID, // Armstrong
		},
		{
			AlchemistID:  alchemists[3].ID, // Hawkeye
			Status:       models.TransmutationStatusCompleted,
			Description:  "Creación de municiones especiales mediante transmutación de metales",
			Cost:         1500.0,
			Result:       "1000 unidades de munición especial producidas. Calidad verificada.",
			ApprovedAt:   &[]time.Time{now.AddDate(0, 0, -8)}[0],
			CompletedAt:  &[]time.Time{now.AddDate(0, 0, -7)}[0],
			SupervisorID: &alchemists[2].ID, // Mustang
		},
		{
			AlchemistID: alchemists[4].ID, // Winry
			Status:      models.TransmutationStatusPending,
			Description: "Transmutación de metales para componentes de automail",
			Cost:        800.0,
		},
		{
			AlchemistID:  alchemists[7].ID, // Izumi
			Status:       models.TransmutationStatusCompleted,
			Description:  "Transmutación de materiales orgánicos para suministros médicos",
			Cost:         600.0,
			Result:       "Suministros médicos básicos producidos exitosamente.",
			ApprovedAt:   &[]time.Time{now.AddDate(0, 0, -12)}[0],
			CompletedAt:  &[]time.Time{now.AddDate(0, 0, -11)}[0],
			SupervisorID: &alchemists[5].ID, // Armstrong
		},
		{
			AlchemistID:  alchemists[1].ID, // Alphonse
			Status:       models.TransmutationStatusApproved,
			Description:  "Experimento de transmutación de cristales alquímicos",
			Cost:         5000.0,
			ApprovedAt:   &[]time.Time{now.AddDate(0, 0, -1)}[0],
			SupervisorID: &alchemists[2].ID, // Mustang
		},
		{
			AlchemistID:  alchemists[6].ID, // Tucker
			Status:       models.TransmutationStatusRejected,
			Description:  "Transmutación experimental con materiales orgánicos",
			Cost:         3000.0,
			Result:       "Transmutación rechazada por violación de protocolos de seguridad. Requiere revisión del Consejo.",
			ApprovedAt:   nil,
			SupervisorID: &alchemists[2].ID, // Mustang
		},
		{
			AlchemistID: alchemists[0].ID, // Edward
			Status:      models.TransmutationStatusPending,
			Description: "Transmutación de oro para financiamiento de investigación",
			Cost:        15000.0,
		},
		{
			AlchemistID: alchemists[3].ID, // Hawkeye
			Status:      models.TransmutationStatusPending,
			Description: "Transmutación de materiales para equipamiento táctico",
			Cost:        1200.0,
		},
	}

	for i := range transmutations {
		if err := db.Create(&transmutations[i]).Error; err != nil {
			fmt.Printf("Error creando transmutación: %v\n", err)
			continue
		}

		// Agregar materiales según el tipo de transmutación
		switch i {
		case 0: // Hierro -> Acero
			inputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[0].ID, Quantity: 500.0, IsInput: true},
				{TransmutationID: transmutations[i].ID, MaterialID: materials[4].ID, Quantity: 50.0, IsInput: true},
			}
			outputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[1].ID, Quantity: 500.0, IsInput: false},
			}
			for _, im := range inputMaterials {
				db.Create(&im)
			}
			for _, om := range outputMaterials {
				db.Create(&om)
			}
		case 1: // Agua -> Hielo
			inputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[5].ID, Quantity: 100.0, IsInput: true},
			}
			for _, im := range inputMaterials {
				db.Create(&im)
			}
		case 2: // Tierra para reparación
			inputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[6].ID, Quantity: 200.0, IsInput: true},
			}
			for _, im := range inputMaterials {
				db.Create(&im)
			}
		case 3: // Municiones
			inputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[1].ID, Quantity: 100.0, IsInput: true},
				{TransmutationID: transmutations[i].ID, MaterialID: materials[9].ID, Quantity: 10.0, IsInput: true},
			}
			for _, im := range inputMaterials {
				db.Create(&im)
			}
		case 4: // Automail
			inputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[1].ID, Quantity: 50.0, IsInput: true},
				{TransmutationID: transmutations[i].ID, MaterialID: materials[3].ID, Quantity: 5.0, IsInput: true},
			}
			for _, im := range inputMaterials {
				db.Create(&im)
			}
		case 5: // Suministros médicos
			inputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[5].ID, Quantity: 200.0, IsInput: true},
			}
			for _, im := range inputMaterials {
				db.Create(&im)
			}
		case 6: // Cristales alquímicos
			inputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[7].ID, Quantity: 5.0, IsInput: true},
				{TransmutationID: transmutations[i].ID, MaterialID: materials[2].ID, Quantity: 2.0, IsInput: true},
			}
			for _, im := range inputMaterials {
				db.Create(&im)
			}
		case 7: // Rechazada
			inputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[5].ID, Quantity: 50.0, IsInput: true},
			}
			for _, im := range inputMaterials {
				db.Create(&im)
			}
		case 8: // Oro
			inputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[0].ID, Quantity: 1000.0, IsInput: true},
				{TransmutationID: transmutations[i].ID, MaterialID: materials[2].ID, Quantity: 10.0, IsInput: true},
			}
			outputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[2].ID, Quantity: 30.0, IsInput: false},
			}
			for _, im := range inputMaterials {
				db.Create(&im)
			}
			for _, om := range outputMaterials {
				db.Create(&om)
			}
		case 9: // Equipamiento táctico
			inputMaterials := []models.TransmutationMaterial{
				{TransmutationID: transmutations[i].ID, MaterialID: materials[1].ID, Quantity: 80.0, IsInput: true},
				{TransmutationID: transmutations[i].ID, MaterialID: materials[7].ID, Quantity: 1.0, IsInput: true},
			}
			for _, im := range inputMaterials {
				db.Create(&im)
			}
		}

		fmt.Printf("✓ Transmutación creada: %s\n", transmutations[i].Description)
	}

	// Crear auditorías
	audits := []models.Audit{
		{
			Type:        models.AuditTypeMaterialUsage,
			Severity:    models.AuditSeverityMedium,
			Description: "Uso excesivo de hierro detectado en el último mes. Alquimista ha utilizado 1500kg cuando el límite mensual es 1000kg",
			AlchemistID: &alchemists[0].ID,
			Details:     `{"material": "Hierro", "usage": 1500, "threshold": 1000, "period": "monthly"}`,
			Resolved:    false,
		},
		{
			Type:        models.AuditTypeMissionCheck,
			Severity:    models.AuditSeverityLow,
			Description: "Misión pendiente por más de 7 días sin aprobación",
			AlchemistID: &alchemists[0].ID,
			Details:     `{"mission_id": 2, "days_pending": 2, "status": "pending"}`,
			Resolved:    false,
		},
		{
			Type:        models.AuditTypeSystem,
			Severity:    models.AuditSeverityHigh,
			Description: "Verificación de seguridad del sistema completada exitosamente",
			Details:     `{"check_type": "security", "status": "passed", "date": "` + now.Format(time.RFC3339) + `"}`,
			Resolved:    true,
			ResolvedAt:  &[]time.Time{now.AddDate(0, 0, -1)}[0],
			ResolvedBy:  &alchemists[2].ID,
		},
		{
			Type:        models.AuditTypeTransmutation,
			Severity:    models.AuditSeverityHigh,
			Description: "Transmutación rechazada por violación de protocolos. Requiere investigación adicional",
			AlchemistID: &alchemists[6].ID,
			Details:     `{"transmutation_id": 8, "reason": "protocol_violation", "violation_type": "organic_materials"}`,
			Resolved:    false,
		},
		{
			Type:        models.AuditTypeMaterialUsage,
			Severity:    models.AuditSeverityLow,
			Description: "Stock de oro por debajo del nivel mínimo recomendado",
			Details:     `{"material": "Oro", "current_stock": 100, "minimum_recommended": 150}`,
			Resolved:    false,
		},
		{
			Type:        models.AuditTypeMissionCheck,
			Severity:    models.AuditSeverityMedium,
			Description: "Misión en progreso por más de 30 días sin actualización",
			AlchemistID: &alchemists[1].ID,
			Details:     `{"mission_id": 1, "days_in_progress": 10, "last_update": "` + now.AddDate(0, 0, -10).Format(time.RFC3339) + `"}`,
			Resolved:    false,
		},
		{
			Type:        models.AuditTypeMaterialUsage,
			Severity:    models.AuditSeverityCritical,
			Description: "Uso crítico de cristal alquímico detectado. Verificar autorizaciones",
			AlchemistID: &alchemists[1].ID,
			Details:     `{"material": "Cristal Alquímico", "usage": 5, "threshold": 3, "requires_authorization": true}`,
			Resolved:    false,
		},
		{
			Type:        models.AuditTypeSystem,
			Severity:    models.AuditSeverityMedium,
			Description: "Revisión mensual de transmutaciones completada",
			Details:     `{"review_type": "monthly", "total_transmutations": 10, "pending": 4, "completed": 5, "rejected": 1}`,
			Resolved:    true,
			ResolvedAt:  &[]time.Time{now.AddDate(0, 0, -2)}[0],
			ResolvedBy:  &alchemists[5].ID,
		},
	}

	for i := range audits {
		if err := db.Create(&audits[i]).Error; err != nil {
			fmt.Printf("Error creando auditoría: %v\n", err)
		} else {
			fmt.Printf("✓ Auditoría creada: %s\n", audits[i].Description)
		}
	}

	fmt.Println("\n✓ Datos de prueba insertados correctamente!")
	fmt.Println("\n📋 Resumen de datos creados:")
	fmt.Printf("  - %d materiales\n", len(materials))
	fmt.Printf("  - %d alquimistas\n", len(alchemists))
	fmt.Printf("  - %d misiones\n", len(missions))
	fmt.Printf("  - %d transmutaciones\n", len(transmutations))
	fmt.Printf("  - %d auditorías\n", len(audits))
	fmt.Println("\n👤 Usuarios de prueba (todos con contraseña: password123):")
	fmt.Println("  - Edward Elric (edward@amestris.gov) - Alquimista Estatal")
	fmt.Println("  - Alphonse Elric (alphonse@amestris.gov) - Alquimista Estatal")
	fmt.Println("  - Roy Mustang (mustang@amestris.gov) - Supervisor Nacional")
	fmt.Println("  - Riza Hawkeye (hawkeye@amestris.gov) - Alquimista Estatal")
	fmt.Println("  - Winry Rockbell (winry@amestris.gov) - Aprendiz")
	fmt.Println("  - Alex Louis Armstrong (armstrong@amestris.gov) - Supervisor Nacional")
	fmt.Println("  - Shou Tucker (tucker@amestris.gov) - Alquimista Estatal")
	fmt.Println("  - Izumi Curtis (izumi@amestris.gov) - Alquimista Nacional")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
