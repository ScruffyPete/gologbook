//go:build integration

package postgres

import (
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	t.Run("creates a user", func(t *testing.T) {
		user := domain.MakeUser("test@example.com", "password")
		db, _ := testutil.NewTestDB(nil, nil, nil)
		defer db.Close()

		repo := NewUserRepository(db)
		repo_user, err := repo.CreateUser(user)

		assert.Nil(t, err)
		assert.Equal(t, user.ID, repo_user.ID)
		assert.Equal(t, user.CreatedAt, repo_user.CreatedAt)
		assert.Equal(t, user.Email, repo_user.Email)
		assert.Equal(t, user.Password, repo_user.Password)
	})

	t.Run("returns an error if the user already exists", func(t *testing.T) {
		user := domain.MakeUser("test@example.com", "password")
		db, _ := testutil.NewTestDB([]*domain.User{user}, nil, nil)
		defer db.Close()

		repo := NewUserRepository(db)
		_, err := repo.CreateUser(user)

		assert.NotNil(t, err)
	})
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	t.Run("returns a user", func(t *testing.T) {
		user := domain.MakeUser("test@example.com", "password")
		db, _ := testutil.NewTestDB([]*domain.User{user}, nil, nil)
		defer db.Close()

		repo := NewUserRepository(db)
		repo_user, err := repo.GetUserByEmail(user.Email)

		assert.Nil(t, err)
		assert.Equal(t, user.ID, repo_user.ID)
		assert.Equal(t, user.CreatedAt, repo_user.CreatedAt)
		assert.Equal(t, user.Email, repo_user.Email)
		assert.Equal(t, user.Password, repo_user.Password)
	})

	t.Run("returns an error if the user does not exist", func(t *testing.T) {
		email := "test@example.com"
		db, _ := testutil.NewTestDB(nil, nil, nil)
		defer db.Close()

		repo := NewUserRepository(db)
		_, err := repo.GetUserByEmail(email)

		assert.NotNil(t, err)
		assert.True(t, domain.NewErrUserDoesNotExist(email).Is(err))
	})

	t.Run("returns an error if the query fails", func(t *testing.T) {
		email := "test@example.com"
		db, _ := testutil.NewTestDB(nil, nil, nil)
		db.Close()

		repo := NewUserRepository(db)
		_, err := repo.GetUserByEmail(email)

		assert.NotNil(t, err)
	})
}
