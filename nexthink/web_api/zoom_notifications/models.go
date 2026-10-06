package zoom_notifications

// AppInfo retains both legacy and current presence flags returned by the controller.
type AppInfo struct {
	NotificationEndpointURL string  `json:"notificationEndpointUrl"`
	AccountID               *string `json:"accountId,omitempty"`
	ClientID                *string `json:"clientId,omitempty"`
	SecretTokenPresent      *bool   `json:"secretTokenPresent,omitempty"`
	ClientSecretPresent     *bool   `json:"clientSecretPresent,omitempty"`
	IsSecretTokenPresent    *bool   `json:"isSecretTokenPresent,omitempty"`
	IsClientSecretPresent   *bool   `json:"isClientSecretPresent,omitempty"`
}

// CheckCredentialsRequest is the JWT check still shipped by the Zoom UI.
// It is not an OAuth credentials validation endpoint.
type CheckCredentialsRequest struct {
	JWT string `json:"jwt"`
}
