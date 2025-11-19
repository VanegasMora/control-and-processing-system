package server

import (
	"backend-avanzada/config"
	"backend-avanzada/logger"
	"backend-avanzada/models"
	"backend-avanzada/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	DB                      *gorm.DB
	Config                  *config.Config
	Handler                 http.Handler
	logger                  *logger.Logger
	taskQueue               *TaskQueue
	AlchemistRepository     *repository.AlchemistRepository
	MissionRepository       *repository.MissionRepository
	MaterialRepository      *repository.MaterialRepository
	TransmutationRepository *repository.TransmutationRepository
	AuditRepository         *repository.AuditRepository
	websocketClients        map[uint]*WebSocketClient
	websocketMutex          sync.RWMutex
}

func NewServer() *Server {
	s := &Server{
		logger:    logger.NewLogger(),
		taskQueue: NewTaskQueue(),
	}
	var config config.Config
	configFile, err := os.ReadFile("config/config.json")
	if err != nil {
		s.logger.Fatal(err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		s.logger.Fatal(err)
	}
	s.Config = &config
	return s
}

func (s *Server) StartServer() {
	fmt.Println("Inicializando base de datos...")
	s.initDB()
	fmt.Println("Inicializando mux...")
	srv := &http.Server{
		Addr:    s.Config.Address,
		Handler: s.router(),
	}
	fmt.Println("Escuchando en el puerto ", s.Config.Address)
	if err := srv.ListenAndServe(); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) initDB() {
	switch s.Config.Database {
	case "sqlite":
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			s.logger.Fatal(err)
		}
		s.DB = db
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			s.logger.Fatal(err)
		}
		s.DB = db
	}
	fmt.Println("Aplicando migraciones...")
	s.DB.AutoMigrate(
		&models.Alchemist{},
		&models.Mission{},
		&models.Material{},
		&models.Transmutation{},
		&models.TransmutationMaterial{},
		&models.Audit{},
	)

	s.AlchemistRepository = repository.NewAlchemistRepository(s.DB)
	s.MissionRepository = repository.NewMissionRepository(s.DB)
	s.MaterialRepository = repository.NewMaterialRepository(s.DB)
	s.TransmutationRepository = repository.NewTransmutationRepository(s.DB)
	s.AuditRepository = repository.NewAuditRepository(s.DB)

	s.websocketClients = make(map[uint]*WebSocketClient)

	// Iniciar procesador de tareas asíncronas
	go s.startAsyncProcessor()
}
