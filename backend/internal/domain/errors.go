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
