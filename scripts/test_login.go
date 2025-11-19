package main

import (
	"backend-avanzada/auth"
	"backend-avanzada/repository"
	"fmt"
	"os"

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

	repo := repository.NewAlchemistRepository(db)

	// Probar con Edward
	email := "edward@amestris.gov"
	password := "password123"

	alchemist, err := repo.FindByEmail(email)
	if err != nil {
		fmt.Printf("Error buscando alquimista: %v\n", err)
		os.Exit(1)
	}

	if alchemist == nil {
		fmt.Printf("Alquimista no encontrado: %s\n", email)
		os.Exit(1)
	}

	fmt.Printf("Alquimista encontrado: %s\n", alchemist.Name)
	fmt.Printf("Password hash: %s\n", alchemist.Password)
	fmt.Printf("Password hash length: %d\n", len(alchemist.Password))

	// Verificar contraseña
	isValid := auth.CheckPasswordHash(password, alchemist.Password)
	fmt.Printf("Contraseña válida: %v\n", isValid)

	if !isValid {
		fmt.Println("\n⚠️  La contraseña no coincide. Regenerando hash...")
		newHash, err := auth.HashPassword(password)
		if err != nil {
			fmt.Printf("Error generando hash: %v\n", err)
			os.Exit(1)
		}
		alchemist.Password = newHash
		repo.Save(alchemist)
		fmt.Println("✓ Hash actualizado en la base de datos")
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
