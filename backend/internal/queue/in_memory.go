package queue

import (
	"context"
	"errors"
	"sync"
	"time"
)

type InMemoryQueue struct {
	mu    sync.Mutex
	items map[string]map[string]float64
}

func NewInMemoryQueue() *InMemoryQueue {
	return &InMemoryQueue{items: map[string]map[string]float64{}}
}

func (q *InMemoryQueue) Push(ctx context.Context, key string, projectID string) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.items == nil {
		q.items = make(map[string]map[string]float64)
	}
	if q.items[key] == nil {
		q.items[key] = make(map[string]float64)
	}

	q.items[key][projectID] = float64(time.Now().UnixNano())
	return nil
}

func (q *InMemoryQueue) Pop(key string, projectID string) (float64, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	projects, ok := q.items[key]
	if !ok {
		return 0, errors.New("no such key")
	}

	item, ok := projects[projectID]
	if !ok {
		return 0, errors.New("no item for project")
	}

	delete(q.items[key], projectID) // Remove after pop
	return item, nil
}

func (q *InMemoryQueue) Close() error {
	return nil
}

// func (q *InMemoryQueue) IsEmpty() (bool, error) {
// 	q.mu.Lock()
// 	defer q.mu.Unlock()
// 	return len(q.items) == 0, nil
// }
