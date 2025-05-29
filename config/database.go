package config

import (
	"fmt"
	"os"
	"time"
	z "votes/utils/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func InitDB() *Database {
	dbHost := os.Getenv("APP_DB_HOST")
	dbUser := os.Getenv("APP_DB_USER")
	dbPass := os.Getenv("APP_DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("APP_DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		z.Log.Fatal().
			Str("event", "database.connect").
			Err(err).
			Msg("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		z.Log.Fatal().
			Str("event", "database.get_sql").
			Err(err).
			Msg("failed to unwrap sql.DB from GORM")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Database{DB: db}
}
