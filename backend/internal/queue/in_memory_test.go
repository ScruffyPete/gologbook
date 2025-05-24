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

	queue.Push(context.Background(), key, entry.ProjectID)

	timestamp, ok := queue.items[key][entry.ProjectID]

	assert.True(t, ok)
	assert.NotNil(t, timestamp)
}
