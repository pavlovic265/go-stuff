package config

import (
	"go-bookstore/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	log.Println(">>> Connect to db")
	// d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=postgres password= dbname=gorm port=5432 sslmode=disable",
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
