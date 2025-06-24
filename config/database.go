package config

import (
	"fmt"
	"os"
	"time"

	"votes/models"
	z "votes/utils/logger"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	// Seed superadmin if not exists
	exists, err := checkSuperadmin(db)
	if err != nil {
		z.Log.Fatal().
			Str("event", "database.check_superadmin").
			Err(err).
			Msg("failed to check if superadmin exists")
	}

	if !exists {
		if err := seedSuperadmin(db); err != nil {
			z.Log.Fatal().
				Str("event", "database.seed_super_admin").
				Err(err).
				Msg("failed to seed superadmin")
		}
	}

	return &Database{DB: db}
}

func checkSuperadmin(db *gorm.DB) (bool, error) {
	var count int64
	if err := db.Model(&models.User{}).Where("role_id = ?", 1).Count(&count).Error; err != nil {
		return false, fmt.Errorf("check superadmin existence: %w", err)
	}

	return count > 0, nil
}

func seedSuperadmin(db *gorm.DB) error {
	password := os.Getenv("SUPERADMIN_PASSWORD")
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hashing password: %w", err)
	}

	lastName := "admin"
	superadmin := models.User{
		ID:        uuid.New(),
		FirstName: "super",
		LastName:  &lastName,
		Username:  "superadmin",
		Email:     "superadmin@example.com",
		Password:  string(hashedPass),
		RoleID:    1,
	}

	if err := db.Create(&superadmin).Error; err != nil {
		return fmt.Errorf("seeding superadmin: %w", err)
	}

	return nil
}
