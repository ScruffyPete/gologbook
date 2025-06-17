package in_memory

import (
	"context"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type UserRepository struct {
	users map[string]*domain.User
}

func NewUserRepository(users []*domain.User) *UserRepository {
	data := make(map[string]*domain.User)

	for _, user := range users {
		data[user.Email] = user
	}

	return &UserRepository{users: data}
}

func (repo *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if _, ok := repo.users[user.Email]; ok {
		return nil, domain.NewErrUserAlreadyExists(user.Email)
	}
	repo.users[user.Email] = user
	return user, nil
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, ok := repo.users[email]
	if !ok {
		return nil, domain.NewErrUserDoesNotExist(email)
	}
	return user, nil
}
