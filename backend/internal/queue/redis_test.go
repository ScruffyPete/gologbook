//go:build integration

package queue

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
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

	channelName := "test:document:stream"
	ctx := context.Background()

	msgCh := queue.SubscribeForDocumentTokens(ctx, channelName)

	go func() {
		time.Sleep(100 * time.Microsecond)
		queue.client.Publish(ctx, channelName, "first message")
		queue.client.Publish(ctx, channelName, "second message")
		queue.client.Publish(ctx, channelName, "[[STOP]]")
	}()

	var received []string
	for msg := range msgCh {
		received = append(received, msg)
	}

	assert.Equal(t, 2, len(received))
	assert.Equal(t, "first message", received[0])
	assert.Equal(t, "second message", received[1])
}
