package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	"learn-n-grow.dev/m/auth"
)

func Connect() (*gorm.DB, error) {
	godotenv.Load()

	var USERNAME = os.Getenv("DB_USER")
	var PASSWORD = os.Getenv("DB_PASSWORD")
	var DATABASE = os.Getenv("DB_DB")
	url := fmt.Sprintf("postgresql://%s:%s@0.0.0.0:5432/%s", USERNAME, PASSWORD, DATABASE)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("db bad")
	}
	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&auth.User{})
	fmt.Println("migrated")
}
