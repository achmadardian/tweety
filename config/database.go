package config

import (
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
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error initialize database: ", err)
	}
	log.Print("success initialize database")

	return &Database{
		WriteConnection: db,
		ReadConnection:  db,
	}
}
