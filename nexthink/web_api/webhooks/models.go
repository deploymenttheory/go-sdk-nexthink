package webhooks

import "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/connector_credentials"

// Webhook is the saved configuration and UI write schema. UUID is caller-generated.
// Create/Update are the same POST upsert. Save does not itself send a webhook.
type Webhook struct {
	UUID           string          `json:"uuid"`
	Name           string          `json:"name"`
	Type           string          `json:"type"`
	Description    string          `json:"description"`
	Enabled        bool            `json:"enabled"`
	NQLCondition   string          `json:"nqlCondition"`
	CredentialID   string          `json:"credentialId"`
	Communications []Communication `json:"communications"`
}
type Communication struct {
	TriggeredAlert string `json:"triggeredAlert"`
	PayloadFormat  string `json:"payloadFormat"`
	// Payload contains original bytes; encoding/json base64-encodes exactly once.
	Payload      []byte `json:"payload"`
	ResourcePath string `json:"resourcePath"`
	HTTPMethod   string `json:"httpMethod"`
}
type WriteResult struct {
	Message string `json:"message"`
}
type Availability struct {
	Available int `json:"available"`
}

// TestRequest contacts the configured third party immediately. Payload is the
// plain text supplied by the UI test form, unlike Communication.Payload.
type TestRequest struct {
	ConnectorConfig connector_credentials.Credential `json:"connectorConfig"`
	ResourceURL     string                           `json:"resourceUrl"`
	WebhookType     string                           `json:"webhookType"`
	Payload         string                           `json:"payload"`
	HTTPMethod      string                           `json:"httpMethod"`
}
