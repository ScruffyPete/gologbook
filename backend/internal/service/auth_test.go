package service

import (
	"context"
	"os"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_Register(t *testing.T) {
	t.Run("returns a user", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		authService := NewAuthService(uow)
		user, err := authService.Register(context.Background(), "test@example.com", "password")
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("returns an error if the user already exists", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		authService := NewAuthService(uow)
		user, err := authService.Register(context.Background(), "test@example.com", "password")
		assert.NoError(t, err)
		assert.NotNil(t, user)

		user, err = authService.Register(context.Background(), "test@example.com", "password")
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}

func TestAuthService_Login(t *testing.T) {
	t.Run("returns a token", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		authService := NewAuthService(uow)
		user, err := authService.Register(context.Background(), "test@example.com", "password")
		assert.NoError(t, err)
		assert.NotNil(t, user)

		token, err := authService.Login(context.Background(), "test@example.com", "password")
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("returns an error if the user does not exist", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		authService := NewAuthService(uow)
		token, err := authService.Login(context.Background(), "test@example.com", "password")
		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("returns an error if the password is incorrect", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		authService := NewAuthService(uow)
		user, err := authService.Register(context.Background(), "test@example.com", "password")
		assert.NoError(t, err)
		assert.NotNil(t, user)

		token, err := authService.Login(context.Background(), "test@example.com", "wrongpassword")
		assert.Error(t, err)
		assert.Empty(t, token)
	})
}

func TestCreateToken(t *testing.T) {
	user := domain.NewUser("test@example.com", "password")
	tokenString, err := createToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, user.Email, claims["username"])
}

func TestValidateToken(t *testing.T) {
	user := domain.NewUser("test@example.com", "password")
	tokenString, err := createToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := ValidateToken(tokenString)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, user.Email, claims["username"])
}
