package ports

import "GoNext/base/internal/core/domain"

type UserService interface {
	Register(user *domain.User) error
	GetById(id string) error
	GetByEmail(email string) error
	Update(user *domain.User) error
	Delete(id string) error
}