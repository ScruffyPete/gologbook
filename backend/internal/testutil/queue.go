package testutil

import (
	"context"
	"errors"
)

type FailingQueue struct{}

var ErrQueueFailed = errors.New("failed to push message")

func (q *FailingQueue) PushPendingProject(ctx context.Context, key string, projectID string) error {
	return ErrQueueFailed
}

func (q *FailingQueue) SubscribeForDocumentTokens(ctx context.Context, channelName string) <-chan string {
	return make(<-chan string)
}

func (q *FailingQueue) Close() error {
	return nil
}
