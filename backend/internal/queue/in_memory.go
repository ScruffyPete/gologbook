package queue

import (
	"context"
	"errors"
	"sync"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type InMemoryQueue struct {
	mu    sync.Mutex
	items map[string][]*domain.Message
}

func NewInMemoryQueue() *InMemoryQueue {
	return &InMemoryQueue{items: map[string][]*domain.Message{}}
}

func (q *InMemoryQueue) Push(ctx context.Context, key string, item *domain.Message) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items[key] = append(q.items[key], item)
	return nil
}

func (q *InMemoryQueue) Pop(key string) (*domain.Message, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items[key]) == 0 {
		return nil, errors.New("queue is empty")
	}
	item := q.items[key][0]
	q.items[key] = q.items[key][1:]
	return item, nil
}

func (q *InMemoryQueue) Close() error {
	return nil
}

func (q *InMemoryQueue) IsEmpty() (bool, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items) == 0, nil
}
