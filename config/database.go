package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	WriteConnection *gorm.DB
	ReadConnection  *gorm.DB
}

func InitDB() *Database {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Build the DSN (Data Source Name) string
	Dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error initialize database: ", err)
	}
	log.Print("success initialize database")

	return &Database{
		WriteConnection: db,
		ReadConnection:  db,
	}
}
