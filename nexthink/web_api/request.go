package web_api

// Request supplies path placeholders, query parameters and an optional JSON body.
// Tokens and cookies must be supplied through the client's provider, not Request.
type Request struct {
	PathParams  map[string]string
	Query       map[string]string
	Body        any
	ContentType string
}
