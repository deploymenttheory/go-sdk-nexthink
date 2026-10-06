package teams_credentials

// CheckCredentialsRequest is encoded as a form, never as JSON on the wire.
// NationalCloud is the UI's ms_national_cloud value (GLOBAL by default).
type CheckCredentialsRequest struct {
	TenantID      string `json:"tenant_id"`
	ClientID      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
	NationalCloud string `json:"ms_national_cloud"`
}
