package initializers

import (
	"fmt"
	"time"

	"github.com/Kunal-Patro/NoteTakingApp/models"
)

func MigrateDatabase() {
	fmt.Printf("[INFO][%v] Database migration started... \n",
		time.Now().In(time.Local).Format("2006-01-02 15:04:05"))

	err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		fmt.Printf("[ERROR][%v] UUID-OSSP extension creation failed. \n",
			time.Now().In(time.Local).Format("2006-01-02 15:04:05"))
		return
	}
	err = DB.AutoMigrate(models.TABLES...)
	if err != nil {
		fmt.Printf("[ERROR][%v] Database migration failed. \n",
			time.Now().In(time.Local).Format("2006-01-02 15:04:05"))
		return
	}

	fmt.Printf("[INFO][%v] Database migration done. \n",
		time.Now().In(time.Local).Format("2006-01-02 15:04:05"))
}
