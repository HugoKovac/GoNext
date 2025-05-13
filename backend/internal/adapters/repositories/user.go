package repositories

import (
	"GoNext/base/ent"
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) ports.UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) Create(user *domain.User) error {
	return nil
}

func (r *UserRepository) FindById(id string) error {
	return nil
}

func (r *UserRepository) FindByEmail(email string) error {
	return nil
}

func (r *UserRepository) Update(user *domain.User) error {
	return nil
}

func (r *UserRepository) Delete(id string) error {
	return nil
}
