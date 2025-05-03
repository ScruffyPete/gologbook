package in_memory

import (
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

func (repo *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	repo.users[user.Email] = user
	return user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	user, ok := repo.users[email]
	if !ok {
		return nil, domain.NewErrUserDoesNotExist(email)
	}
	return user, nil
}
