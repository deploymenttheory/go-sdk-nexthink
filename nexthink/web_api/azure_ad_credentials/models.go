package azure_ad_credentials

// CheckCredentialsRequest is encoded as a form, never as JSON on the wire.
// NationalCloud is GLOBAL or US_L4 in the UI. An empty ClientSecret uses saved credentials.
type CheckCredentialsRequest struct {
	TenantID      string `json:"tenant_id"`
	ClientID      string `json:"client_id"`
	ClientSecret  string `json:"client_secret,omitempty"`
	NationalCloud string `json:"ms_national_cloud"`
}
