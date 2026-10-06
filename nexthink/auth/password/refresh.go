package password

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (p *Provider) refresh(ctx context.Context, old session) (session, error) {
	// Endpoint and client ID originate only from the observed browser exchange.
	if old.TokenURL != tokenURL(p.origin) || old.ClientID == "" {
		return session{}, ErrRenewalFailed
	}
	form := url.Values{"grant_type": {"refresh_token"}, "client_id": {old.ClientID}, "refresh_token": {old.RefreshToken}}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, old.TokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return session{}, ErrRenewalFailed
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	resp, err := p.httpClient.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return session{}, ctx.Err()
		}
		return session{}, ErrRenewalFailed
	}
	defer resp.Body.Close()
	const maxResponse = 1 << 20
	raw, err := io.ReadAll(io.LimitReader(resp.Body, maxResponse+1))
	if err != nil || len(raw) > maxResponse {
		return session{}, ErrRenewalFailed
	}
	var result struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int64  `json:"expires_in"`
		TokenType    string `json:"token_type"`
		Error        string `json:"error"`
	}
	if json.Unmarshal(raw, &result) != nil {
		return session{}, ErrRenewalFailed
	}
	if resp.StatusCode == http.StatusBadRequest && (result.Error == "invalid_grant" || result.Error == "invalid_token") {
		return session{}, errRefreshRejected
	}
	if resp.StatusCode != http.StatusOK {
		return session{}, fmt.Errorf("%w: HTTP %d", ErrRenewalFailed, resp.StatusCode)
	}
	if !strings.EqualFold(result.TokenType, "Bearer") || result.ExpiresIn <= 0 || result.ExpiresIn > int64((365*24*time.Hour)/time.Second) {
		return session{}, ErrRenewalFailed
	}
	next := old
	next.AccessToken = result.AccessToken
	next.ExpiresAt = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)
	if result.RefreshToken != "" {
		next.RefreshToken = result.RefreshToken
	}
	if validateSession(next) != nil {
		return session{}, ErrRenewalFailed
	}
	return next, nil
}
