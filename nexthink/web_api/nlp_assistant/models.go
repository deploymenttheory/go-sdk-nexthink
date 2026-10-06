package nlp_assistant

import (
	"encoding/json"
	"fmt"
)

type ChatRequest struct {
	ID               string    `json:"id"`
	Source           string    `json:"source"`
	Messages         []Message `json:"messages"`
	ConversationMode string    `json:"conversation_mode"`
}
type Message struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}
type ChatResponse struct {
	Events []Event `json:"events"`
}

// Event preserves the original SSE data, including polymorphic artifacts.
type Event struct {
	Type string          `json:"type"`
	ID   string          `json:"id,omitempty"`
	Data json.RawMessage `json:"data"`
}
type EventData struct {
	Metadata  json.RawMessage `json:"metadata,omitempty"`
	Content   string          `json:"content,omitempty"`
	Message   string          `json:"message,omitempty"`
	Status    int             `json:"status,omitempty"`
	Artifacts json.RawMessage `json:"artifacts,omitempty"`
}

func (e Event) Decode() (*EventData, error) {
	var data EventData
	if err := json.Unmarshal(e.Data, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

type StreamError struct {
	Status  int
	Message string
}

func (e *StreamError) Error() string {
	return fmt.Sprintf("NLP assistance stream (%d): %s", e.Status, e.Message)
}
