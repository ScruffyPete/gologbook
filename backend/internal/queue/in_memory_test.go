package queue

import (
	"context"
	"fmt"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	queue := NewInMemoryQueue()
	entry := domain.NewEntry(uuid.NewString(), "test")
	payload := map[string]any{"entry": entry}
	message := domain.Message{
		Type:    domain.MESSAGE_TYPE_NEW_ENTRY,
		Payload: payload,
	}
	key := fmt.Sprintf("project:%s", entry.ProjectID)

	queue.Push(context.Background(), key, &message)

	poppedMessage, err := queue.Pop(key)
	assert.NoError(t, err)
	assert.Equal(t, &message, poppedMessage)
}

func TestPop(t *testing.T) {
	t.Run("should pop message from queue", func(t *testing.T) {
		queue := NewInMemoryQueue()
		entry := domain.NewEntry(uuid.NewString(), "test")
		payload := map[string]any{"entry": entry}
		message := domain.Message{
			Type:    domain.MESSAGE_TYPE_NEW_ENTRY,
			Payload: payload,
		}
		key := fmt.Sprintf("project:%s", entry.ProjectID)

		queue.Push(context.Background(), key, &message)
		poppedMessage, err := queue.Pop(key)
		assert.NoError(t, err)
		assert.Equal(t, &message, poppedMessage)
	})

	t.Run("should return error if queue is empty", func(t *testing.T) {
		queue := NewInMemoryQueue()
		_, err := queue.Pop("new queue")
		assert.Error(t, err)
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("should return true if queue is empty", func(t *testing.T) {
		queue := NewInMemoryQueue()
		empty, err := queue.IsEmpty()
		assert.NoError(t, err)
		assert.True(t, empty)
	})

	t.Run("should return false if queue is not empty", func(t *testing.T) {
		queue := NewInMemoryQueue()
		entry := domain.NewEntry(uuid.NewString(), "test")
		payload := map[string]any{"entry": entry}
		message := domain.Message{
			Type:    domain.MESSAGE_TYPE_NEW_ENTRY,
			Payload: payload,
		}
		key := fmt.Sprintf("project:%s", entry.ProjectID)
		queue.Push(context.Background(), key, &message)
		empty, err := queue.IsEmpty()
		assert.NoError(t, err)
		assert.False(t, empty)
	})
}
