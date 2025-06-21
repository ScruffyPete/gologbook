//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	uow, err := NewTestUnitOfWork()
	if err != nil {
		t.Fatalf("failed to create unit of work: %v", err)
	}
	defer uow.Close()

	user := domain.NewUser("test@example.com", "password")

	t.Run("creates a user", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var createdUser *domain.User
		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			var err error
			createdUser, err = repos.Users.CreateUser(ctx, user)
			return err
		})
		if err != nil {
			t.Fatalf("WithTx error: %v", err)
		}
		assert.Equal(t, user, createdUser)
	})

	t.Run("returns an error if the user already exists", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			if _, err := repos.Users.CreateUser(ctx, user); err != nil {
				return err
			}
			_, err = repos.Users.CreateUser(ctx, user)
			return err
		})
		assert.NotNil(t, err)
	})

	t.Run("get user by email", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var gotUser *domain.User
		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			if _, err := repos.Users.CreateUser(ctx, user); err != nil {
				return err
			}
			gotUser, err = repos.Users.GetUserByEmail(ctx, user.Email)
			return err
		})
		if err != nil {
			t.Fatalf("WithTx error: %v", err)
		}
		assert.Equal(t, user, gotUser)
	})
}
