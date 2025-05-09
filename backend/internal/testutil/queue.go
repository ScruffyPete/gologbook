package testutil

import (
	"context"
	"errors"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type FailingQueue struct{}

var ErrQueueFailed = errors.New("failed to push message")

func (q *FailingQueue) Push(ctx context.Context, msg domain.Message) error {
	return ErrQueueFailed
}

func (q *FailingQueue) Close() error {
	return nil
}
