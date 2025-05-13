package services

import (
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
)

type UserService struct {
	UserRepository ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (s *UserService) Register(user *domain.User) error {
	return nil
}

func (s *UserService) GetById(id string) error {
	return nil
}

func (s *UserService) GetByEmail(email string) error {
	return nil
}

func (s *UserService) Update(user *domain.User) error {
	return nil
}

func (s *UserService) Delete(id string) error {
	return nil
}
