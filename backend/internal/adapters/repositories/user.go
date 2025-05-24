package repositories

import (
	"GoNext/base/ent"
	"GoNext/base/ent/user"
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) ports.UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) toDomainUser(entUser *ent.User) *domain.User {
	if entUser == nil {
		return nil
	}

	return &domain.User{
		Id: entUser.ID.String(),
		Email: entUser.Email,
		Password: entUser.Password,
		CreatedAt: entUser.CreatedAt,
		UpdatedAt: entUser.CreatedAt,
	}
}

func (r *UserRepository) Create(user domain.User) (*domain.User, error) {
	ctx := context.Background()
	dUser, err := r.client.User.Create().SetEmail(user.Email).SetPassword(user.Password).SetID(uuid.New()).SetCreatedAt(time.Now()).SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		fmt.Println("Creating User: ", err)
		return nil, err
	}
	fmt.Println(user.Email, " user created")
	
	return r.toDomainUser(dUser), nil
}

func (r *UserRepository) FindById(id string) (*domain.User, error) {
	ctx := context.Background()
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	u, err := r.client.User.Query().Where(user.ID(uuid)).Only(ctx)
	return r.toDomainUser(u), err
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	ctx := context.Background()
	u, err := r.client.User.Query().Where(user.Email(email)).Only(ctx)
	return r.toDomainUser(u), err
}

func (r *UserRepository) Update(user *domain.User) error {
	return nil
}

func (r *UserRepository) Delete(id string) error {
	return nil
}
