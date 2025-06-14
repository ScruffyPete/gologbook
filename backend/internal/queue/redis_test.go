//go:build integration

package queue

import (
	"context"
	"fmt"
	"os"
	"testing"

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
	ctx := context.Background()
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

	_, err = queue.client.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]any{"token": "[[STOP]]"},
	}).Result()
	assert.NoError(t, err)

	var received []string
	for msg := range msgCh {
		received = append(received, msg)
	}
	assert.Equal(t, 2, len(received))
	assert.Equal(t, "first message", received[0])
	assert.Equal(t, "second message", received[1])
}
