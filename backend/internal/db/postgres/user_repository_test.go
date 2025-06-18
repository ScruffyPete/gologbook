//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	t.Run("creates a user", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		user := domain.NewUser("test@example.com", "password")
		db := testutil.NewTestDB(t, ctx, nil, nil, nil, nil)

		repo := NewUserRepository(db)
		repo_user, err := repo.CreateUser(ctx, user)

		assert.Nil(t, err)
		assert.Equal(t, user.ID, repo_user.ID)
		assert.Equal(t, user.CreatedAt, repo_user.CreatedAt)
		assert.Equal(t, user.Email, repo_user.Email)
		assert.Equal(t, user.Password, repo_user.Password)
	})

	t.Run("returns an error if the user already exists", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		user := domain.NewUser("test@example.com", "password")
		db := testutil.NewTestDB(t, ctx, []*domain.User{user}, nil, nil, nil)

		repo := NewUserRepository(db)
		_, err := repo.CreateUser(ctx, user)

		assert.NotNil(t, err)
	})
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	t.Run("returns a user", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		user := domain.NewUser("test@example.com", "password")
		db := testutil.NewTestDB(t, ctx, []*domain.User{user}, nil, nil, nil)

		repo := NewUserRepository(db)
		repo_user, err := repo.GetUserByEmail(ctx, user.Email)

		assert.Nil(t, err)
		assert.Equal(t, user.ID, repo_user.ID)
		assert.Equal(t, user.CreatedAt, repo_user.CreatedAt)
		assert.Equal(t, user.Email, repo_user.Email)
		assert.Equal(t, user.Password, repo_user.Password)
	})

	t.Run("returns an error if the query fails", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		email := "test@example.com"
		db := testutil.NewTestDB(t, ctx, nil, nil, nil, nil)

		repo := NewUserRepository(db)
		_, err := repo.GetUserByEmail(ctx, email)

		assert.NotNil(t, err)
	})
}
