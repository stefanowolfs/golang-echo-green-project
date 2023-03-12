package handler

import "green-echo/internal/app/green/domain"

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (j *UserResponse) ToUser() *domain.User {
	return &domain.User{
		ID:   j.ID,
		Name: j.Name,
	}
}

func (j *UserResponse) FromUser(user *domain.User) {
	j.ID = user.ID
	j.Name = user.Name
}
