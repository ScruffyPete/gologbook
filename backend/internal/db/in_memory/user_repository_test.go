package in_memory

import (
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	repo := NewUserRepository(nil)

	user := domain.NewUser("test@example.com", "password")
	createdUser, err := repo.CreateUser(user)

	assert.Nil(t, err)
	assert.Equal(t, user, createdUser)
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	t.Run("returns a user", func(t *testing.T) {
		user := domain.NewUser("test@example.com", "password")
		repo := NewUserRepository([]*domain.User{user})

		foundUser, err := repo.GetUserByEmail(user.Email)

		assert.Nil(t, err)
		assert.Equal(t, user, foundUser)
	})

	t.Run("returns an error if the user does not exist", func(t *testing.T) {
		repo := NewUserRepository(nil)

		_, err := repo.GetUserByEmail("test@example.com")

		assert.NotNil(t, err)
	})
}
