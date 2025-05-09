package queue

import (
	"context"
	"errors"
	"sync"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type InMemoryQueue struct {
	mu    sync.Mutex
	items []domain.Message
}

func NewInMemoryQueue() *InMemoryQueue {
	return &InMemoryQueue{items: []domain.Message{}}
}

func (q *InMemoryQueue) Push(ctx context.Context, item domain.Message) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
	return nil
}

func (q *InMemoryQueue) Pop() (domain.Message, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return domain.Message{}, errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
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
