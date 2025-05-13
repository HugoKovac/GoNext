package services

import (
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"fmt"
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
	fmt.Println("Register: ", user)
	s.UserRepository.Create(user)
	return nil
}

func (s *UserService) GetById(id string) (*domain.User, error) {
	user, err := s.UserRepository.FindById(id)
	return user, err
}

func (s *UserService) GetByEmail(email string) (*domain.User, error) {
	user, err := s.UserRepository.FindByEmail(email)
	return user, err
}

func (s *UserService) Update(user *domain.User) error {
	return nil
}

func (s *UserService) Delete(id string) error {
	return nil
}
