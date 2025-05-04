package domain

import "fmt"

type ErrProjectDoesNotExist struct {
	ID string
}

func NewErrProjectDoesNotExist(id string) *ErrProjectDoesNotExist {
	return &ErrProjectDoesNotExist{ID: id}
}

func (e *ErrProjectDoesNotExist) Error() string {
	return fmt.Sprintf("project with id %s does not exist", e.ID)
}

func (e *ErrProjectDoesNotExist) Is(target error) bool {
	if target == nil {
		return false
	}

	return target.Error() == e.Error()
}

type ErrUserDoesNotExist struct {
	Email string
}

func NewErrUserDoesNotExist(email string) *ErrUserDoesNotExist {
	return &ErrUserDoesNotExist{Email: email}
}

func (e *ErrUserDoesNotExist) Error() string {
	return fmt.Sprintf("user with email %s does not exist", e.Email)
}

func (e *ErrUserDoesNotExist) Is(target error) bool {
	if target == nil {
		return false
	}

	return target.Error() == e.Error()
}

type ErrUserAlreadyExists struct {
	Email string
}

func NewErrUserAlreadyExists(email string) *ErrUserAlreadyExists {
	return &ErrUserAlreadyExists{Email: email}
}

func (e *ErrUserAlreadyExists) Error() string {
	return fmt.Sprintf("user with email %s already exists", e.Email)
}

func (e *ErrUserAlreadyExists) Is(target error) bool {
	if target == nil {
		return false
	}

	return target.Error() == e.Error()
}

type ErrInvalidPassword struct {
}

func NewErrInvalidPassword() *ErrInvalidPassword {
	return &ErrInvalidPassword{}
}

func (e *ErrInvalidPassword) Error() string {
	return "invalid password"
}

func (e *ErrInvalidPassword) Is(target error) bool {
	if target == nil {
		return false
	}

	return target.Error() == e.Error()
}
