package user_communication_integrations

// Content selects an Entra tenant directly or through a configured Azure connector.
type Content struct {
	AzureTenantID    string `json:"azureTenantId,omitempty"`
	AzureConnectorID string `json:"azureConnectorId,omitempty"`
	WelcomeMessage   string `json:"welcomeMessage,omitempty"`
}

// IntegrationInput mirrors both UI modes. Name may be empty when using a connector.
type IntegrationInput struct {
	Name    string  `json:"name"`
	Content Content `json:"content"`
}
type Integration struct {
	ID string `json:"id"`
	IntegrationInput
}
type AzureConnector struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TenantID string `json:"tenantId"`
}
