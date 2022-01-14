package config

import (
	"d-crud/models"
	"d-crud/utils"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("host=%s port=5432 user=docker password=docker dbname=docker sslmode=disable", host)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.User{})
	utils.DB = db
	fmt.Println("Sukses Connect ke DB")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbPsql, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbPsql.Close()
}
