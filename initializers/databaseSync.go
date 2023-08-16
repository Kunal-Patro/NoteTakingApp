package initializers

import (
	"github.com/Kunal-Patro/NoteTakingApp/internal/logger"
	"github.com/Kunal-Patro/NoteTakingApp/models"
)

func MigrateDatabase() {
	logger.SetLevel("debug")
	log := logger.WithService("migration-service")

	log.Debug("Datbase migration started.")

	err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		log.WithError(err).Error("UUID-OSSP extension creation failed")
		return
	}
	err = DB.AutoMigrate(models.TABLES...)
	if err != nil {
		log.WithError(err).Error("Database migration failed.")
		return
	}

	log.Debug("Database migration completed.")
}
