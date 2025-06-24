package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  *string
	Username  string
	Email     string
	Password  string
	RoleID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Role      Role `gorm:"foreignKey:RoleID;references:ID"`
}
