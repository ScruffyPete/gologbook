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
	Push(ctx context.Context, key string, projectID string) error
	Close() error
}
