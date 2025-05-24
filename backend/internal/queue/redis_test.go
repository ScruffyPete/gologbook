//go:build integration

package queue

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRedisQueue_Push(t *testing.T) {
	queue, _ := NewRedisQueue()
	err := queue.Push(context.Background(), "test_key", uuid.NewString())
	assert.NoError(t, err)
}
