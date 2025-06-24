package responses

import "github.com/google/uuid"

type RegisterResponse struct {
	ID        uuid.UUID        `json:"id"`
	FirstName string           `json:"first_name"`
	LastName  *string          `json:"last_name"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Role      RoleUserResponse `json:"role"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
