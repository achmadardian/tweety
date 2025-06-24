package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
