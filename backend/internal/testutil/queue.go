package testutil

import (
	"errors"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type FailingQueue struct{}

var ErrQueueFailed = errors.New("failed to push message")

func (q *FailingQueue) Push(msg domain.Message) error {
	return ErrQueueFailed
}

func (q *FailingQueue) Pop() (domain.Message, error) {
	return domain.Message{}, ErrQueueFailed
}

func (q *FailingQueue) IsEmpty() (bool, error) {
	return true, ErrQueueFailed
}
