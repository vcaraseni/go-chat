package database

import (
	"fmt"
	"log"

	"github.com/vcaraseni/go-chat/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	cfg := config.AppConfig.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Connection Error to the DB: %v", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
}

func Migrate() {
	DB.AutoMigrate(&Message{})
	fmt.Println("Migration has been processed")
}
