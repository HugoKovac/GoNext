package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"GoNext/base/internal/core/mocks"
)

func TestNewUserService(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository(t)
	service := NewUserService(mockRepo)

	assert.NotNil(t, service)
	assert.IsType(t, &UserService{}, service)
}

