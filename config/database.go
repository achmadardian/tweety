package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	WriteConnection *gorm.DB
	ReadConnecion   *gorm.DB
}

func InitDB() *Database {
	dsn := "root:ardian@tcp(127.0.0.1:3306)/votes?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error initialize database: ", err)
	}
	log.Print("success initalize database")

	return &Database{
		WriteConnection: db,
		ReadConnecion:   db,
	}
}
