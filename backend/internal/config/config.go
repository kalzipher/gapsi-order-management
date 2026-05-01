package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string

	CORSAllowedOrigins []string

	DatabaseURL string

	JWTAccessSecret  string
	JWTRefreshSecret string
	JWTAccessTTL     time.Duration
	JWTRefreshTTL    time.Duration

	AdminEmail    string
	AdminPassword string
	AdminName     string

	CSVPath string
}

func Load() Config {
	_ = godotenv.Load()

	accessTTLMinutes := getEnvAsInt("JWT_ACCESS_TTL_MINUTES", 15)
	refreshTTLHours := getEnvAsInt("JWT_REFRESH_TTL_HOURS", 168)

	return Config{
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),

		CORSAllowedOrigins: getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000"}),

		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/orders_db?sslmode=disable"),

		JWTAccessSecret:  getEnv("JWT_ACCESS_SECRET", "change_me_access_secret"),
		JWTRefreshSecret: getEnv("JWT_REFRESH_SECRET", "change_me_refresh_secret"),
		JWTAccessTTL:     time.Duration(accessTTLMinutes) * time.Minute,
		JWTRefreshTTL:    time.Duration(refreshTTLHours) * time.Hour,

		AdminEmail:    getEnv("ADMIN_EMAIL", "admin@test.com"),
		AdminPassword: getEnv("ADMIN_PASSWORD", "admin123"),
		AdminName:     getEnv("ADMIN_NAME", "Admin"),

		CSVPath: getEnv("CSV_PATH", "./data/orders_db.csv"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}

func getEnvAsSlice(key string, fallback []string) []string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	if len(result) == 0 {
		return fallback
	}

	return result
}
