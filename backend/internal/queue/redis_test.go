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
	err := queue.Push(context.Background(), domain.Message{
		Type: "test",
		Payload: map[string]any{
			"test": "test",
		},
	})

	assert.NoError(t, err)
}
