package ui_events

import "encoding/json"

type Message struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	OccurredOn string          `json:"occurredOn"`
	Payload    json.RawMessage `json:"payload"`
}
type Link struct {
	Href string `json:"href"`
}
type Links struct {
	Next *Link `json:"next,omitempty"`
}
type MessagesResponse struct {
	Messages []Message `json:"messages"`
	Links    *Links    `json:"_links,omitempty"`
}
