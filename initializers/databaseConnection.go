package initializers

import (
	"os"

	"github.com/Kunal-Patro/NoteTakingApp/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// log.Fatal("Error connecting to database: ", err)
		logger.WithService("database-service").WithError(err).Error("Error connecting to database")
	}

}
