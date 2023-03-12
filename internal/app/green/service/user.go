package service

import (
	"green-echo/internal/app/green/domain"
)

type UserService interface {
	GetById(userID uint) (*domain.User, error)
}

type userService struct {
	repository domain.UserRepository
}

func NewUserService(repository domain.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) GetById(userID uint) (*domain.User, error) {
	user, err := s.repository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
