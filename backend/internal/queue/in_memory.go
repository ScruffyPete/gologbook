package queue

import (
	"context"
	"errors"
	"sync"
	"time"
)

type InMemoryQueue struct {
	mu              sync.Mutex
	pendingProjects map[string]map[string]float64
	documentStream  map[string]<-chan string
}

func NewInMemoryQueue() *InMemoryQueue {
	return &InMemoryQueue{
		pendingProjects: map[string]map[string]float64{},
		documentStream:  map[string]<-chan string{},
	}
}

func (q *InMemoryQueue) PushPendingProject(ctx context.Context, key string, projectID string) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.pendingProjects == nil {
		q.pendingProjects = make(map[string]map[string]float64)
	}
	if q.pendingProjects[key] == nil {
		q.pendingProjects[key] = make(map[string]float64)
	}

	q.pendingProjects[key][projectID] = float64(time.Now().UnixNano())
	return nil
}

func (q *InMemoryQueue) Pop(key string, projectID string) (float64, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	projects, ok := q.pendingProjects[key]
	if !ok {
		return 0, errors.New("no such key")
	}

	item, ok := projects[projectID]
	if !ok {
		return 0, errors.New("no item for project")
	}

	delete(q.pendingProjects[key], projectID) // Remove after pop
	return item, nil
}

func (q *InMemoryQueue) SubscribeForDocumentTokens(ctx context.Context, projectID string) <-chan string {
	out := make(chan string, 100)

	go func() {
		defer close(out)

		for msg := range q.documentStream[projectID] {
			if msg == "[[STOP]]" {
				return
			}
			out <- msg
		}
	}()

	return out
}

func (q *InMemoryQueue) Close() error {
	return nil
}

func (q *InMemoryQueue) SetDocumentStream(channelName string, documentStream <-chan string) {
	q.documentStream[channelName] = documentStream
}

// func (q *InMemoryQueue) IsEmpty() (bool, error) {
// 	q.mu.Lock()
// 	defer q.mu.Unlock()
// 	return len(q.items) == 0, nil
// }
