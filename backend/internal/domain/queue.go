package domain

type MessageType string

const (
	MESSAGE_TYPE_NEW_ENTRY MessageType = "new_entry"
)

type Message struct {
	Type    MessageType    `json:"type"`
	Payload map[string]any `json:"payload"`
}

type Queue interface {
	Push(item Message) error
	Pop() (Message, error)
	IsEmpty() (bool, error)
}
