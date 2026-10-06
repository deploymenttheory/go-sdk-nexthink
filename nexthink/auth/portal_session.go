package auth

import (
	"fmt"
	"strings"
)

// PortalSession supplies explicit per-call credentials to the fixed legacy portal
// routes. The SDK does not persist response cookies or acquire portal sessions.
// Root WebAPI construction still requires browser credentials or a token provider;
// that provider is not consulted for these legacy calls.
type PortalSession struct {
	XAuthToken string `json:"-"`
	Cookie     string `json:"-"`
}

func (PortalSession) String() string { return "PortalSession{credentials:redacted}" }
func (s *PortalSession) Validate() error {
	if s == nil || (strings.TrimSpace(s.XAuthToken) == "" && strings.TrimSpace(s.Cookie) == "") {
		return fmt.Errorf("explicit portal session cookie or x-auth-token is required")
	}
	for _, value := range []string{s.XAuthToken, s.Cookie} {
		if strings.ContainsAny(value, "\r\n\x00") {
			return fmt.Errorf("portal credentials contain invalid header characters")
		}
	}
	return nil
}
