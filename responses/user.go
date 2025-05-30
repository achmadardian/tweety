package responses

import "github.com/google/uuid"

type Me struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  *string   `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}
