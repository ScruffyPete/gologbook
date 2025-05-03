package postgres

import (
	"database/sql"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	_, err := repo.db.Exec(
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

func (repo *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	row, err := repo.db.Query(
		"SELECT id, created_at, email, password FROM users WHERE email = $1",
		email,
	)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if !row.Next() {
		return nil, domain.NewErrUserDoesNotExist(email)
	}

	var user domain.User
	if err := row.Scan(&user.ID, &user.CreatedAt, &user.Email, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}
