package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/auth"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/config"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/database"
	"gorm.io/gorm"
)

type App struct {
	cfg    config.Config
	db     *gorm.DB
	router *gin.Engine

	jwtService *auth.JWTService
}

func New(cfg config.Config) (*App, error) {
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSAllowedOrigins,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	a := &App{
		cfg:    cfg,
		db:     db,
		router: router,
	}

	a.registerRoutes()

	return a, nil
}

func (a *App) Run() error {
	addr := ":" + a.cfg.AppPort
	log.Printf("server running on %s", addr)

	return a.router.Run(addr)
}

func (a *App) Close() {
	sqlDB, err := a.db.DB()
	if err != nil {
		log.Printf("could not get sql db: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("could not close database connection: %v", err)
	}
}
