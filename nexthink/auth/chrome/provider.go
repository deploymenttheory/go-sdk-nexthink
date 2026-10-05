// Package chrome provides opt-in access to a signed-in Nexthink tab on macOS.
// Chrome must allow JavaScript from Apple Events. Login and refresh remain owned
// by Chrome; this package never reads refresh tokens or browser cookies.
package chrome

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
)

type Provider struct{ origin string }

func New(origin string) (*Provider, error) {
	u, err := url.Parse(origin)
	if err != nil || u.Scheme != "https" || u.User != nil || u.Port() != "" || !strings.HasSuffix(u.Hostname(), ".nexthink.cloud") || (u.Path != "" && u.Path != "/") || u.RawQuery != "" || u.Fragment != "" {
		return nil, fmt.Errorf("Chrome provider requires an HTTPS Nexthink tenant origin")
	}
	return &Provider{origin: u.Scheme + "://" + u.Host}, nil
}

func (p *Provider) Token(ctx context.Context) (auth.Token, error) {
	raw, err := readSession(ctx, p.origin)
	if err != nil {
		return auth.Token{}, err
	}
	return parseSession(raw)
}

func parseSession(raw []byte) (auth.Token, error) {
	var session struct {
		Token     string `json:"token"`
		ExpiresAt int64  `json:"expiresAt"`
		Error     string `json:"error"`
	}
	if json.Unmarshal(raw, &session) != nil {
		return auth.Token{}, fmt.Errorf("Chrome returned an invalid session response")
	}
	if session.Error != "" || session.Token == "" {
		return auth.Token{}, fmt.Errorf("no signed-in Nexthink session in Chrome")
	}
	if session.ExpiresAt <= 0 {
		return auth.Token{}, fmt.Errorf("Chrome session is missing its token expiry")
	}
	t := auth.Token{Value: session.Token, ExpiresAt: time.Unix(session.ExpiresAt, 0)}
	return t, t.Validate()
}
