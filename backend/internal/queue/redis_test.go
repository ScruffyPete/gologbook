//go:build integration

package queue

import (
	"context"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestRedisQueue_Push(t *testing.T) {
	queue, _ := NewRedisQueue()
	message := domain.Message{
		Type: "test",
		Payload: map[string]any{
			"test": "test",
		},
	}
	err := queue.Push(context.Background(), "test_key", &message)

	assert.NoError(t, err)
}
