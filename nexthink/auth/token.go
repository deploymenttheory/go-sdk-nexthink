// Package auth defines authentication providers shared by Nexthink clients.
package auth

import (
	"context"
	"errors"
	"time"
)

var ErrSessionExpired = errors.New("Nexthink session expired; sign in again or supply a fresh token")

// Token is an access token. A zero ExpiresAt means the provider does not know its expiry.
// Never log Value or serialize Token into application diagnostics.
type Token struct {
	Value     string
	ExpiresAt time.Time
}

func (t Token) Validate() error {
	if t.Value == "" {
		return errors.New("access token is empty")
	}
	if !t.ExpiresAt.IsZero() && !time.Now().Before(t.ExpiresAt) {
		return ErrSessionExpired
	}
	return nil
}

// TokenProvider must be safe for concurrent calls and observe context cancellation.
type TokenProvider interface {
	Token(context.Context) (Token, error)
}

type TokenProviderFunc func(context.Context) (Token, error)

func (f TokenProviderFunc) Token(ctx context.Context) (Token, error) { return f(ctx) }

// StaticToken uses a caller-managed access token without performing a login.
func StaticToken(value string, expiry time.Time) TokenProvider {
	return TokenProviderFunc(func(ctx context.Context) (Token, error) {
		if err := ctx.Err(); err != nil {
			return Token{}, err
		}
		t := Token{Value: value, ExpiresAt: expiry}
		return t, t.Validate()
	})
}
