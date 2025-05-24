package requests

import (
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	Id       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type UserRequestUpdate struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	UpdatedAt time.Time
}

func (u *UserRequestUpdate) IsEmpty() bool {
	return u.Name == "" && u.Email == "" && u.Password == ""
}
