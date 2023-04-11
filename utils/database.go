package utils

import (
	"fmt"
	"github.com/Ruchanov/Golang_2023/assignment3/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	// Set up database connection
	dsn := "host = localhost port = 5432 user = postgres dbname = bookstore password = Ayef1407_ sslmode = disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	err = migrateDB(db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
}
func migrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		return fmt.Errorf("failed to migrate books table: %v", err)
	}

	return nil
}
