//go:build integration

package queue

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisQueue_Push(t *testing.T) {
	queue, _ := NewRedisQueue()
	err := queue.PushPendingProject(context.Background(), "test_key", uuid.NewString())
	assert.NoError(t, err)
}

func TestRedisQueue_Subscribe(t *testing.T) {
	queue, err := NewRedisQueue()
	assert.Nil(t, err)

	projectID := uuid.NewString()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msgCh := queue.SubscribeForDocumentTokens(ctx, projectID)
	prefix := os.Getenv("REDIS_LLM_STREAM_CHANNEL_PREFIX")
	streamName := fmt.Sprintf("%s:%s", prefix, projectID)

	_, err = queue.client.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]any{"token": "first message"},
	}).Result()
	assert.NoError(t, err)

	_, err = queue.client.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]any{"token": "second message"},
	}).Result()
	assert.NoError(t, err)

	var received []string
	for len(received) < 2 {
		select {
		case <-ctx.Done():
			t.Fatal("timeout waiting for Redis messages")
		case msg, ok := <-msgCh:
			if !ok {
				t.Fatal("channel closed before expected")
			}
			received = append(received, msg)
		}
	}
	assert.Equal(t, 2, len(received))
	assert.Equal(t, "first message", received[0])
	assert.Equal(t, "second message", received[1])
}
