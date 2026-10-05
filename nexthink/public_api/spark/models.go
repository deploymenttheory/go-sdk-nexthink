package spark

// Part is a TEXT or FILE message part. FileContent is the API's encoded content string.
type Part struct {
	Type        string  `json:"type"`
	Text        string  `json:"text,omitempty"`
	MIMEType    string  `json:"mimeType,omitempty"`
	FileContent *string `json:"fileContent,omitempty"`
}
type Message struct {
	Parts []Part `json:"parts"`
}
type HandoffRequest struct {
	Message  Message           `json:"message"`
	Metadata map[string]string `json:"metadata,omitempty"`
}
