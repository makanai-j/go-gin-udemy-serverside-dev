package infra

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	env := os.Getenv("ENV")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var (
		db  *gorm.DB
		err error
	)

	if env == "prod" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		log.Println("setup db for production. db is postgres")
	} else {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		log.Println("setup db for development or test. db is sqlite")
	}

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
