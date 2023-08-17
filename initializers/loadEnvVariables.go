package initializers

import (
	"github.com/Kunal-Patro/NoteTakingApp/internal/logger"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		logger.WithService("environment_loader-service").WithError(err).Error("Cannot load .env file")
	}
}
