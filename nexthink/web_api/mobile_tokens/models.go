package mobile_tokens

// Token contains an enrollment credential. Do not log JWTToken.
type Token struct {
	Name           string  `json:"name"`
	JTI            string  `json:"jti"`
	JWTToken       *string `json:"jwtToken,omitempty"`
	ExpirationDate string  `json:"expirationDate"`
	Revision       int     `json:"revision"`
}
type CreateRequest struct {
	Name           string `json:"name"`
	ExpirationDate string `json:"expirationDate"`
}
type UpdateRequest struct {
	JTI      string `json:"jti"`
	Name     string `json:"name"`
	Revision int    `json:"revision"`
}
type DeleteRequest struct {
	JTI      string `json:"jti"`
	Revision int    `json:"revision"`
}
