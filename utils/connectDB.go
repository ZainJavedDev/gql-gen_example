package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbHost := GoDotEnvVariable("DB_HOST")
	dbUser := GoDotEnvVariable("DB_USER")
	dbName := GoDotEnvVariable("DB_NAME")
	dbPassword := GoDotEnvVariable("DB_PASSWORD")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", dbHost, dbUser, dbName, dbPassword, "5432")

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
