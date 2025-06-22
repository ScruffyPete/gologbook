package domain

import "context"

type MessageType string

const (
	MESSAGE_TYPE_NEW_ENTRY MessageType = "new_entry"
)

type Message struct {
	Type    MessageType    `json:"type"`
	Payload map[string]any `json:"payload"`
}

type Queue interface {
	PushPendingProject(ctx context.Context, key string, projectID string) error
	SubscribeForDocumentTokens(ctx context.Context, projectID string) <-chan string
	ClearProjectStream(ctx context.Context, projectID string) error
	Close() error
}
