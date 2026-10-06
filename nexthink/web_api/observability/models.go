package observability

// Submission carries the browser telemetry protocol. ClientToken is the telemetry client token, not an OAuth client secret.
// Payload bytes are kept opaque because the UI sends newline-delimited events or encoded batches.
type Submission struct {
	Source        string `json:"source"`
	ClientToken   string `json:"clientToken"`
	Origin        string `json:"origin"`
	OriginVersion string `json:"originVersion"`
	RequestID     string `json:"requestId"`
	Encoding      string `json:"encoding,omitempty"`
	BatchTime     string `json:"batchTime,omitempty"`
	API           string `json:"api,omitempty"`
	RetryCount    *int   `json:"retryCount,omitempty"`
	RetryAfter    string `json:"retryAfter,omitempty"`
	ContentType   string `json:"contentType"`
	Payload       []byte `json:"payload"`
}
