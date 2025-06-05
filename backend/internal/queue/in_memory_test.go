package queue

import (
	"context"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	queue := NewInMemoryQueue()
	entry := domain.NewEntry(uuid.NewString(), "test")
	key := "test_zset"

	queue.PushPendingProject(context.Background(), key, entry.ProjectID)

	timestamp, ok := queue.pendingProjects[key][entry.ProjectID]

	assert.True(t, ok)
	assert.NotNil(t, timestamp)
}

func TestSubscribe(t *testing.T) {
	queue := NewInMemoryQueue()

	channelName := "test:document:stream"
	ch := make(chan string, 3)
	ch <- "first message"
	ch <- "second message"
	ch <- "[[STOP]]"
	close(ch)

	queue.documentStream[channelName] = ch

	ctx := context.Background()
	msgCh := queue.SubscribeForDocumentTokens(ctx, channelName)

	var received []string
	for msg := range msgCh {
		received = append(received, msg)
	}

	assert.Equal(t, 2, len(received))
	assert.Equal(t, "first message", received[0])
	assert.Equal(t, "second message", received[1])
}
