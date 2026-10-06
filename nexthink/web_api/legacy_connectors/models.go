package legacy_connectors

import "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/connector_credentials"

type ConnectionDetail = connector_credentials.ConnectionDetail
type Mapping struct {
	DestinationField string `json:"destinationField"`
	NexthinkField    string `json:"nexthinkField"`
	Enabled          *bool  `json:"enabled,omitempty"`
}

// ConfigurationInput is the legacy config body. ConnectorType belongs in the URL.
type ConfigurationInput struct {
	ConnectionDetails    []ConnectionDetail `json:"connectionDetails"`
	ConnectorName        string             `json:"connectorName,omitempty"`
	ConnectorDescription *string            `json:"connectorDescription,omitempty"`
	Enabled              *bool              `json:"enabled,omitempty"`
	Mapping              []Mapping          `json:"mapping"`
	RunTime              string             `json:"runTime"`
	TimeZone             string             `json:"timeZone,omitempty"`
}
type Configuration struct {
	ConfigurationInput
	ConnectorType string `json:"connectorType"`
}
type ListOptions struct {
	Type    string
	Enabled *bool
}
type WriteResult struct {
	Message string `json:"message"`
}

// SecretRequest writes secrets. HasSecrets checks presence without decoding secret values.
type SecretRequest struct {
	Entries []connector_credentials.SecretEntry `json:"entries"`
	Tags    map[string]string                   `json:"tags,omitempty"`
}

// Summary includes nullable names present only in the collection response.
type Summary struct {
	Configuration
	ConnectorName        *string `json:"connectorName"`
	ConnectorDescription *string `json:"connectorDescription"`
}
