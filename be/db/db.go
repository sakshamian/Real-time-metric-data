package db

import (
	"fmt"
	"log"
	"os"
	"udp-be/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := os.Getenv("DB_DSN")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	DB.AutoMigrate(&models.Message{})

	fmt.Println("Successfully connected to the database!")
}
