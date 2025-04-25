package response

import "votes/models"

type UserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponseUpdate struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (u *UserResponseUpdate) Map(user *models.User) {
	u.Name = user.Name
	u.Email = user.Email
}
