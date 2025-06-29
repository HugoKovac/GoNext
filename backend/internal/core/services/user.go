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

func (s *UserService) Register(user domain.User) (*domain.User, error) {
	existingUser, err := s.UserRepository.FindByEmail(user.Email)
	if existingUser != nil || err == nil {
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

func (s *UserService) Update(userId string, email string, oldPassword string, newPassword string) (domain.User, error) {
	user, err := s.UserRepository.FindById(userId)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	if user == nil {
		return domain.User{}, errors.New("user not found")
	}
	if email != "" && email != user.Email {
		existingUser, err := s.UserRepository.FindByEmail(email)
		if existingUser != nil || err == nil {
			return domain.User{}, errors.New("user with this email already exists")
		}
		user.Email = email
	}
	if oldPassword != "" && newPassword != "" {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
		if err != nil {
			return domain.User{}, errors.New("old password is incorrect")
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			return domain.User{}, err
		}
		user.Password = string(hashedPassword)
	}
	user.UpdatedAt = time.Now()
	err = s.UserRepository.Update(user)
	if err != nil {
		return domain.User{}, errors.New("failed to update user")
	}
	return *user, nil
}

func (s *UserService) Delete(id string) error {
	//todo: To implement and test
	return nil
}
