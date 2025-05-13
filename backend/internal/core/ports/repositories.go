package ports

import (
	"GoNext/base/internal/core/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindById(id string) error
	FindByEmail(email string) error
	Update(user *domain.User) error
	Delete(id string) error
}
