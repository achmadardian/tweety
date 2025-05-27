package requests

import "github.com/google/uuid"

type RegisterRequest struct {
	ID        uuid.UUID `json:"-"`
	FirstName string    `json:"first_name" binding:"required,max=50"`
	LastName  string    `json:"last_name" binding:"required,max=50"`
	Username  string    `json:"username" binding:"required,max=20"`
	Email     string    `json:"email" binding:"required,email,max=254"`
	Password  string    `json:"password" binding:"required,min=8"`
}
