package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	uow domain.UnitOfWork
}

func NewAuthService(uow domain.UnitOfWork) *AuthService {
	return &AuthService{uow: uow}
}

func (s *AuthService) SignUp(ctx context.Context, email string, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("generate password hash: %w", err)
	}
	user := domain.NewUser(email, string(hashedPassword))

	err = s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		_, err = repos.Users.CreateUser(ctx, user)
		return err
	})

	if err != nil {
		return fmt.Errorf("sign up user: %w", err)
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	var user *domain.User

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		user, err = repos.Users.GetUserByEmail(ctx, email)
		if err != nil {
			return err
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return domain.NewErrInvalidPassword()
		}
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("login user: %w", err)
	}

	tokenString, err := createToken(user)
	if err != nil {
		return "", fmt.Errorf("create token: %w", err)
	}

	return tokenString, nil
}

func createToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("generate password hash: %w", err)
	}
	return string(hashedPassword), nil
}
