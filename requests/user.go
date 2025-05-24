package requests

import (
	"github.com/google/uuid"
)

type CreatUserRequest struct {
	Id       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *UserUpdateRequest) IsEmpty() bool {
	return u.Name == "" && u.Email == "" && u.Password == ""
}
