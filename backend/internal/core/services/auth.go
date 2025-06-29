// internal/core/services/auth_service.go
package services

import (
	"errors"

	"GoNext/base/internal/core/ports"
	"GoNext/base/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    userRepo ports.UserRepository
    jwtSecret string
}

func NewAuthService(userRepo ports.UserRepository, jwtSecret string) ports.AuthService {
    return &AuthService{
        userRepo: userRepo,
        jwtSecret: jwtSecret,
    }
}

func (s *AuthService) Authenticate(username string, password string) (string, error) {
    // Get user by email
    user, err := s.userRepo.FindByEmail(username)
    if err != nil {
        return "", errors.New("user does not exist")
    }
    
    // Compare passwords
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    // Generate JWT token
    return jwt.GenerateToken(user.Id, s.jwtSecret, user.Role)
}

func (s *AuthService) ValidateToken(tokenString  string) (string, error) {
    return jwt.ValidateToken(tokenString, s.jwtSecret)
}
