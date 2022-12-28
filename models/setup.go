package models

import (
	"os"
	"fmt"
	
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
)


var Database *gorm.DB

func ConnectDatabase() {
	godotenv.Load(".env")

	database, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URL")), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database! %s"))
	}

	err = database.AutoMigrate(&Url{})

	if err != nil {
		return
	}

	Database = database
}
