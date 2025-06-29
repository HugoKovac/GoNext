package services

import (
	"errors"
	"testing"
	"time"

	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUserService(t *testing.T) {
	mockUserRepo := mocks.NewMockUserRepository(t)
	service := NewUserService(mockUserRepo)

	assert.NotNil(t, service)
	assert.IsType(t, &UserService{}, service)
}


func TestRegister(t *testing.T){
	mockUserRepo := mocks.NewMockUserRepository(t)

	user := domain.User{
		Id: "1",
		Email: "test@email.com",
		Password: "pass",
		CreatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		UpdatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		Role: "user",

	}

	mockUserRepo.On("FindByEmail", user.Email).Return(&user, nil)
	mockUserRepo.On("Create", mock.Anything).Return(&user, nil)

	service := NewUserService(mockUserRepo)
	rtnUser, err := service.Register(user)
	assert.Nil(t, err)
	assert.IsType(t, &domain.User{}, rtnUser)
	assert.Equal(t, user.CreatedAt, rtnUser.CreatedAt)
	assert.Equal(t, user.UpdatedAt, rtnUser.UpdatedAt)
	assert.Equal(t, user.Id, rtnUser.Id)
	assert.Equal(t, user.Email, rtnUser.Email)
	assert.Equal(t, user.Password, rtnUser.Password)
	assert.Equal(t, user.Role, rtnUser.Role)
}

func TestRegisterEmailAlreadyExists(t *testing.T){
	mockUserRepo := mocks.NewMockUserRepository(t)

	user := domain.User{
		Id: "1",
		Email: "user@email.com",
		Password: "pass",
		CreatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		UpdatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		Role: "user",

	}

	mockUserRepo.On("FindByEmail", user.Email).Return(nil, errors.New("Already Exists"))

	service := NewUserService(mockUserRepo)
	rtnUser, err := service.Register(user)
	assert.Nil(t, rtnUser)
	assert.NotNil(t, err)
}

func TestGetById(t *testing.T){
	mockUserRepo := mocks.NewMockUserRepository(t)
	user := domain.User{
		Id: "55",
		Email: "test55@email.com",
		Password: "pass",
		CreatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		UpdatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		Role: "user",

	}
	mockUserRepo.On("FindById", user.Id).Return(&user, nil)
	service := NewUserService(mockUserRepo)
	rtnUser, err := service.GetById(user.Id)
	assert.Nil(t, err)
	assert.IsType(t, &domain.User{}, rtnUser)
	assert.Equal(t, user.CreatedAt, rtnUser.CreatedAt)
	assert.Equal(t, user.UpdatedAt, rtnUser.UpdatedAt)
	assert.Equal(t, user.Id, rtnUser.Id)
	assert.Equal(t, user.Email, rtnUser.Email)
	assert.Equal(t, user.Password, rtnUser.Password)
	assert.Equal(t, user.Role, rtnUser.Role)

}

func TestGetByIdUserDoesntExists(t *testing.T){
	mockUserRepo := mocks.NewMockUserRepository(t)
	mockUserRepo.On("FindById", "44").Return(nil, errors.New("Doesn't exists"))
	service := NewUserService(mockUserRepo)
	rtnUser, err := service.GetById("44")
	assert.Nil(t, rtnUser)
	assert.NotNil(t, err)
}

func TestGetByEmail(t *testing.T){
	mockUserRepo := mocks.NewMockUserRepository(t)
	user := domain.User{
		Id: "55",
		Email: "test66@email.com",
		Password: "pass",
		CreatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		UpdatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		Role: "user",

	}
	mockUserRepo.On("FindByEmail", user.Email).Return(&user, nil)
	service := NewUserService(mockUserRepo)
	rtnUser, err := service.GetByEmail(user.Email)
	assert.Nil(t, err)
	assert.IsType(t, &domain.User{}, rtnUser)
	assert.Equal(t, user.CreatedAt, rtnUser.CreatedAt)
	assert.Equal(t, user.UpdatedAt, rtnUser.UpdatedAt)
	assert.Equal(t, user.Id, rtnUser.Id)
	assert.Equal(t, user.Email, rtnUser.Email)
	assert.Equal(t, user.Password, rtnUser.Password)
	assert.Equal(t, user.Role, rtnUser.Role)

}

func TestGetByEmailUserDoesntExists(t *testing.T){
	mockUserRepo := mocks.NewMockUserRepository(t)
	mockUserRepo.On("FindByEmail", "nexists@gmail.com").Return(nil, errors.New("Doesn't exists"))
	service := NewUserService(mockUserRepo)
	rtnUser, err := service.GetByEmail("nexists@gmail.com")
	assert.Nil(t, rtnUser)
	assert.NotNil(t, err)
}

// todo: add test for Update