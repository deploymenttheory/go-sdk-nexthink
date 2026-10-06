package connector_credentials

import "encoding/json"

// Credential is configuration metadata. Secret values are not returned by Get/List.
type Credential struct {
	Enabled           bool               `json:"enabled"`
	Mapping           []json.RawMessage  `json:"mapping"`
	RunTime           string             `json:"runTime"`
	TimeZone          string             `json:"timeZone"`
	ConnectorType     string             `json:"connectorType"`
	ConnectionDetails []ConnectionDetail `json:"connectionDetails"`
}
type ConnectionDetail struct {
	Key   string `json:"connectionKey"`
	Value string `json:"connectionValue"`
}

// CredentialInput matches the UI save envelope. Omit Secret when keeping existing
// secrets. Create and Update use the same server upsert; neither enforces existence.
type CredentialInput struct {
	Config ConfigurationInput `json:"config"`
	Secret *Secret            `json:"secret,omitempty"`
	Tags   map[string]string  `json:"tags,omitempty"`
}
type ConfigurationInput struct {
	ConnectionDetails []ConnectionDetail `json:"connectionDetails"`
	Mapping           []json.RawMessage  `json:"mapping"`
	RunTime           string             `json:"runTime"`
	Enabled           bool               `json:"enabled"`
	ConnectorType     string             `json:"connectorType,omitempty"`
}

// Secret contains write-only values. Never log or publish a populated input.
type Secret struct {
	Entries []SecretEntry `json:"entries"`
}
type SecretEntry struct {
	Key   string `json:"secretKey"`
	Value string `json:"secretValue"`
}

// SavedCredential is the save response, which is smaller than the Get response.
type SavedCredential struct {
	ConnectionDetails []ConnectionDetail `json:"connectionDetails,omitempty"`
	RunTime           string             `json:"runTime"`
	ConnectorType     string             `json:"connectorType,omitempty"`
	Enabled           bool               `json:"enabled"`
}

// Summary includes nullable names exposed only by the collection endpoint.
type Summary struct {
	Credential
	ConnectorName        *string `json:"connectorName"`
	ConnectorDescription *string `json:"connectorDescription"`
}
