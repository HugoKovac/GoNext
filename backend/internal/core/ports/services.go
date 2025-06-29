package ports

import "GoNext/base/internal/core/domain"

type UserService interface {
	Register(user domain.User) (*domain.User, error)
	GetById(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(userID string, email string, oldPassword string, newPassword string) (domain.User, error)
	Delete(id string) error
}

type AuthService interface {
    Authenticate(username string, password string) (string, error)
    ValidateToken(tokenString string) (string, error)
}
