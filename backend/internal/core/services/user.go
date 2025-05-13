package services

import (
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (s *UserService) Register(user *domain.User) (*domain.User, error) {
	// fmt.Println("Register: ", user)
	// s.UserRepository.Create(user)
	// return nil
	// Check if user already exists
    existingUser, _ := s.UserRepository.FindByEmail(user.Email)
    if existingUser != nil {
        return nil, errors.New("user with this email already exists")
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    user.Password = string(hashedPassword)

    // Set timestamps
    now := time.Now()
    user.CreatedAt = now
    user.UpdatedAt = now

    return s.UserRepository.Create(user)
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
