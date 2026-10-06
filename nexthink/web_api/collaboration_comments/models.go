package collaboration_comments

// IdentifierRequest resolves the comment document for the UI's resource identifiers.
type IdentifierRequest struct {
	Identifiers []string `json:"identifiers"`
	URL         string   `json:"url"`
}
type IdentifierResponse struct {
	Identifier *string `json:"identifier"`
}
type UserMention struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	Self     *bool  `json:"self,omitempty"`
}
type UserMentionsResponse struct {
	UserMentions []UserMention `json:"userMentions"`
}

// CreateMessageRequest requires the caller-generated UUID used for idempotent UI retries.
type CreateMessageRequest struct {
	ID           string        `json:"id"`
	Message      string        `json:"message"`
	UserMentions []UserMention `json:"userMentions"`
}
type EditMessageRequest struct {
	Message      string        `json:"message"`
	UserMentions []UserMention `json:"userMentions"`
}
type CommentOperation struct {
	Operation string              `json:"operation"`
	Data      *EditMessageRequest `json:"data,omitempty"`
}
type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Message struct {
	Text         string        `json:"text"`
	UserMentions []UserMention `json:"userMentions"`
}
type Permissions struct {
	CanBeArchived  bool `json:"canBeArchived"`
	CanBeDeleted   bool `json:"canBeDeleted"`
	CanBeEdited    bool `json:"canBeEdited"`
	CanBeRepliedTo bool `json:"canBeRepliedTo"`
}
type Reply struct {
	ID              string      `json:"id"`
	Author          Author      `json:"author"`
	Message         Message     `json:"message"`
	Status          string      `json:"status"`
	CreatedAt       string      `json:"createdAt"`
	MessageEditedAt *string     `json:"messageEditedAt,omitempty"`
	Permissions     Permissions `json:"permissions"`
}
type Comment struct {
	ID              string      `json:"id"`
	Author          Author      `json:"author"`
	Message         Message     `json:"message"`
	Status          string      `json:"status"`
	CreatedAt       string      `json:"createdAt"`
	MessageEditedAt *string     `json:"messageEditedAt,omitempty"`
	Permissions     Permissions `json:"permissions"`
	Replies         []Reply     `json:"replies"`
}
type CommentsResponse struct {
	Items []Comment `json:"items"`
}
