package postgres

import (
	"context"
	"database/sql"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type userRepository struct {
	tx *sql.Tx
}

func NewUserRepository(tx *sql.Tx) *userRepository {
	return &userRepository{tx: tx}
}

func (repo *userRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	_, err := repo.tx.ExecContext(
		ctx,
		"INSERT INTO users (id, created_at, email, password) VALUES ($1, $2, $3, $4)",
		user.ID,
		user.CreatedAt,
		user.Email,
		user.Password,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := repo.tx.QueryRowContext(
		ctx,
		"SELECT id, created_at, email, password FROM users WHERE email = $1",
		email,
	)

	var user domain.User
	if err := row.Scan(&user.ID, &user.CreatedAt, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewErrUserDoesNotExist(email)
		}
		return nil, err
	}
	return &user, nil
}
