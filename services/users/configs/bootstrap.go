package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/adhitamafikri/go-simple-pms/services/users/router"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	APP_ENV           string
	APP_PORT          int
	APP_HOST          string
	POSTGRES_DB       string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_PORT     int
	POSTGRES_SSL_MODE string

	IS_LOCAL bool
}

func loadAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load .env")
	}

	appEnv := os.Getenv("APP_ENV")
	appPortRaw := os.Getenv("APP_PORT")
	appPort, appPortErr := strconv.Atoi(appPortRaw)
	if appPortErr != nil {
		log.Fatal().Err(err).Msg("Failed to load APP_PORT from .env")
	}
	appHost := os.Getenv("APP_HOST")

	postgresDB := os.Getenv("POSTGRES_DB")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresPortRaw := os.Getenv("POSTGRES_PORT")
	postgresPort, postgresPortErr := strconv.Atoi(postgresPortRaw)
	if postgresPortErr != nil {
		log.Fatal().Err(err).Msg("Failed to load POSTGRES_PORT from .env")
	}
	postgresSSLMode := os.Getenv("POSTGRES_SSL_MODE")

	return &AppConfig{
		APP_ENV:           appEnv,
		APP_PORT:          appPort,
		APP_HOST:          appHost,
		POSTGRES_DB:       postgresDB,
		POSTGRES_USER:     postgresUser,
		POSTGRES_PASSWORD: postgresPassword,
		POSTGRES_PORT:     postgresPort,
		POSTGRES_SSL_MODE: postgresSSLMode,
		IS_LOCAL:          appEnv == "local",
	}
}

func loadRestRouter() gin.Engine {
	userRouter := router.NewRouter()
	engine := gin.Default()

	return *router.RegisterRoute(engine, userRouter)
}

func Bootstrap() {
	// Load Env
	appConfig := loadAppConfig()

	// Connect DB

	// Load Routers
	r := loadRestRouter()
	addr := fmt.Sprintf("%s:%d", appConfig.APP_HOST, appConfig.APP_PORT)
	r.Run(addr)
}
