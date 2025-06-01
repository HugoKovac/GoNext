package services

import (
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/mocks"
	"GoNext/base/pkg/jwt"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewAuthService(t *testing.T){
	mockUserRepo := mocks.NewMockUserRepository(t)
	authService := NewAuthService(mockUserRepo, "secret")
	assert.NotNil(t, authService)
	assert.IsType(t, &AuthService{}, authService)
}

func TestAuthenticate(t *testing.T) {
	mockUserRepo := mocks.NewMockUserRepository(t)

	pass := "pass"
	bcryptPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	user := domain.User{
		Id: "1",
		Email: "test@email.com",
		Password: string(bcryptPass),
		CreatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		UpdatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		Role: "user",

	}

	mockUserRepo.On("FindByEmail", user.Email).Return(&user,nil)
	
	authService := NewAuthService(mockUserRepo, "secret")
	assert.Nil(t, err)
	jwt, err := authService.Authenticate(user.Email, pass)

	assert.NotEmpty(t, jwt)
	assert.Nil(t, err)
}

func TestAuthenticateDoesntExists(t *testing.T) {
	mockUserRepo := mocks.NewMockUserRepository(t)

	user := domain.User{
		Id: "1",
		Email: "test@email.com",
		Password: "pass",
		CreatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		UpdatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		Role: "user",

	}

	mockUserRepo.On("FindByEmail", user.Email).Return(nil,errors.New("user does not exist"))
	
	authService := NewAuthService(mockUserRepo, "secret")
	jwt, err := authService.Authenticate(user.Email, user.Password)

	assert.Empty(t, jwt)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "user does not exist")
}

func TestAuthenticateWrongPassword(t *testing.T) {
	mockUserRepo := mocks.NewMockUserRepository(t)

	user := domain.User{
		Id: "1",
		Email: "test@email.com",
		Password: "pass",
		CreatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		UpdatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		Role: "user",

	}

	mockUserRepo.On("FindByEmail", user.Email).Return(&user,nil)
	
	authService := NewAuthService(mockUserRepo, "secret")
	jwt, err := authService.Authenticate(user.Email, user.Password)

	assert.Empty(t, jwt)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "invalid credentials")
}

func TestValidateToken(t *testing.T){
	mockUserRepo := mocks.NewMockUserRepository(t)

	pass := "pass"
	bcryptPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	user := domain.User{
		Id: "1",
		Email: "test@email.com",
		Password: string(bcryptPass),
		CreatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		UpdatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		Role: "user",

	}

	mockUserRepo.On("FindByEmail", user.Email).Return(&user,nil)
	
	authService := NewAuthService(mockUserRepo, "secret")
	assert.Nil(t, err)
	jwt, err := authService.Authenticate(user.Email, pass)

	assert.NotEmpty(t, jwt)
	assert.Nil(t, err)

	rtn, err := authService.ValidateToken(jwt)
	assert.NotEmpty(t, rtn)
	assert.Nil(t, err)
}

func TestValidateTokenIvalid(t *testing.T){
	mockUserRepo := mocks.NewMockUserRepository(t)

	pass := "pass"
	bcryptPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	user := domain.User{
		Id: "1",
		Email: "test@email.com",
		Password: string(bcryptPass),
		CreatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		UpdatedAt: time.Date(2025,05,4,5,13,24,0,time.Now().Local().Location()),
		Role: "user",

	}
	
	authService := NewAuthService(mockUserRepo, "secret")
	assert.Nil(t, err)

	jwt, err := jwt.GenerateToken(user.Id, "tt", user.Role)
	assert.Nil(t, err)

	rtn, err := authService.ValidateToken(jwt)
	assert.Empty(t, rtn)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "token signature is invalid: signature is invalid")
}
