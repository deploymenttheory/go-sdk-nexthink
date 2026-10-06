package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// AuthConfig holds authentication configuration for the Nexthink API
//
// Nexthink API docs: https://docs.nexthink.com/api/getting-authentication-token
type AuthConfig struct {
	// ClientID is the OAuth2 client ID
	ClientID string

	// ClientSecret is the OAuth2 client secret
	ClientSecret string

	// Instance is the Nexthink instance name
	Instance string

	// Region is the Nexthink region (us, eu, pac, meta)
	Region string

	// TokenURL is the optional custom token endpoint URL
	// If not provided, it will be constructed from Instance and Region
	TokenURL string

	// Scope is the OAuth2 scope (defaults to service:integration)
	Scope string
}

// TokenResponse represents the OAuth2 token response
//
// Nexthink API docs: https://docs.nexthink.com/api/getting-authentication-token
type TokenResponse struct {
	TokenType   string `json:"token_type"`   // "Bearer"
	ExpiresIn   int    `json:"expires_in"`   // Token lifetime in seconds (900 = 15 minutes)
	AccessToken string `json:"access_token"` // The access token
	Scope       string `json:"scope"`        // "service:integration"
}

// TokenManager handles OAuth2 token lifecycle
//
// Nexthink API docs: https://docs.nexthink.com/api/getting-authentication-token
type TokenManager struct {
	authConfig    *AuthConfig
	logger        *zap.Logger
	client        *resty.Client
	currentToken  *TokenResponse
	tokenExpiry   time.Time
	mu            sync.RWMutex
	refreshBuffer time.Duration
	refreshing    chan struct{}
	refreshErr    error
}

// NewTokenManager creates a new token manager
func NewTokenManager(authConfig *AuthConfig, client *resty.Client, logger *zap.Logger) *TokenManager {
	httpClient := *client.Client()
	httpClient.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	return &TokenManager{
		authConfig:    authConfig,
		logger:        logger,
		client:        resty.NewWithClient(&httpClient).SetTimeout(client.Timeout()).SetDebug(false),
		refreshBuffer: TokenRefreshBuffer * time.Second,
	}
}

// Validate checks if the auth configuration is valid
func (a *AuthConfig) Validate() error {
	if a.ClientID == "" {
		return fmt.Errorf("client ID is required")
	}
	if a.ClientSecret == "" {
		return fmt.Errorf("client secret is required")
	}
	if a.Instance == "" {
		return fmt.Errorf("instance name is required")
	}
	if a.Region == "" {
		return fmt.Errorf("region is required")
	}

	// Validate region
	validRegions := map[string]bool{
		RegionUS:   true,
		RegionEU:   true,
		RegionPAC:  true,
		RegionMETA: true,
	}
	if !validRegions[a.Region] {
		return fmt.Errorf("invalid region: %s (must be one of: us, eu, pac, meta)", a.Region)
	}

	return nil
}

// GetTokenURL returns the token endpoint URL
func (a *AuthConfig) GetTokenURL() string {
	if a.TokenURL != "" {
		return a.TokenURL
	}
	return fmt.Sprintf(DefaultTokenURLTemplate, a.Instance, a.Region)
}

// GetScope returns the OAuth2 scope
func (a *AuthConfig) GetScope() string {
	if a.Scope != "" {
		return a.Scope
	}
	return ScopeServiceIntegration
}

// GenerateBasicAuth generates the Base64 encoded Basic auth string from clientId:clientSecret
func (a *AuthConfig) GenerateBasicAuth() string {
	credentials := fmt.Sprintf("%s:%s", a.ClientID, a.ClientSecret)
	return base64.StdEncoding.EncodeToString([]byte(credentials))
}

// GetToken returns a cached token, refreshing it when needed.
func (tm *TokenManager) GetToken() (string, error) {
	return tm.GetTokenContext(context.Background())
}

// GetTokenContext observes cancellation while obtaining or waiting for a token.
func (tm *TokenManager) GetTokenContext(ctx context.Context) (string, error) {
	return tm.token(ctx, false)
}

// RefreshToken explicitly requests a new token, even if the cached token is valid.
func (tm *TokenManager) RefreshToken() (string, error) {
	return tm.token(context.Background(), true)
}

func (tm *TokenManager) token(ctx context.Context, force bool) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	tm.mu.Lock()
	if done := tm.refreshing; done != nil {
		tm.mu.Unlock()
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-done:
		}
		tm.mu.RLock()
		defer tm.mu.RUnlock()
		if tm.refreshErr != nil {
			return "", tm.refreshErr
		}
		if tm.currentToken == nil {
			return "", fmt.Errorf("token invalidated during refresh")
		}
		return tm.currentToken.AccessToken, nil
	}
	if !force && tm.currentToken != nil && time.Now().Add(tm.refreshBuffer).Before(tm.tokenExpiry) {
		token := tm.currentToken.AccessToken
		tm.mu.Unlock()
		return token, nil
	}
	done := make(chan struct{})
	tm.refreshing = done
	tm.mu.Unlock()
	result, err := tm.requestToken(ctx)
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.refreshErr = err
	if err == nil {
		tm.currentToken = result
		lifetime := time.Duration(result.ExpiresIn) * time.Second
		tm.tokenExpiry = time.Now().Add(lifetime)
		tm.refreshBuffer = min(time.Duration(TokenRefreshBuffer)*time.Second, lifetime/10)
	}
	tm.refreshing = nil
	close(done)
	if err != nil {
		return "", err
	}
	return result.AccessToken, nil
}

func (tm *TokenManager) requestToken(ctx context.Context) (*TokenResponse, error) {
	resp, err := tm.client.R().SetContext(ctx).SetDebug(false).
		SetHeader("Content-Type", ContentTypeFormURLEncoded).
		SetHeader("Authorization", "Basic "+tm.authConfig.GenerateBasicAuth()).
		SetFormData(map[string]string{"grant_type": GrantTypeClientCredentials, "scope": tm.authConfig.GetScope()}).
		Post(tm.authConfig.GetTokenURL())
	if err != nil {
		return nil, fmt.Errorf("request access token: %w", err)
	}
	if !IsResponseSuccess(toInterfaceResponse(resp)) {
		return nil, fmt.Errorf("token request failed with status %d", resp.StatusCode())
	}
	var result TokenResponse
	if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
		return nil, fmt.Errorf("invalid token response JSON: %w", err)
	}
	if result.AccessToken == "" || result.ExpiresIn <= 0 || !strings.EqualFold(result.TokenType, "Bearer") {
		return nil, fmt.Errorf("invalid token response: nonempty bearer token and positive expires_in required")
	}
	return &result, nil
}

// InvalidateToken clears the current token, forcing a refresh on next use
func (tm *TokenManager) InvalidateToken() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.currentToken = nil
	tm.tokenExpiry = time.Time{}
	tm.logger.Info("Access token invalidated")
}

// SetupAuthentication configures the resty client with OAuth2 bearer token authentication
//
// Nexthink API docs: https://docs.nexthink.com/api/getting-authentication-token
func SetupAuthentication(client *resty.Client, authConfig *AuthConfig, logger *zap.Logger) (*TokenManager, error) {
	if err := authConfig.Validate(); err != nil {
		logger.Error("Authentication validation failed", zap.Error(err))
		return nil, fmt.Errorf("authentication validation failed: %w", err)
	}

	tokenManager := NewTokenManager(authConfig, client, logger)

	// Fetch initial token
	token, err := tokenManager.RefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to obtain initial access token: %w", err)
	}

	client.SetAuthToken(token)

	// Add request middleware to ensure token is valid before each request
	client.AddRequestMiddleware(func(c *resty.Client, req *resty.Request) error {
		if scoped, err := preparePortalSession(req, req.Method, req.URL); err != nil || scoped {
			return err
		}
		token, err := tokenManager.GetTokenContext(req.Context())
		if err != nil {
			logger.Error("Failed to get valid token for request", zap.Error(err))
			return fmt.Errorf("failed to get valid token: %w", err)
		}
		req.SetAuthToken(token)
		return nil
	})

	logger.Info("OAuth2 authentication configured successfully",
		zap.String("instance", authConfig.Instance),
		zap.String("region", authConfig.Region),
		zap.String("scope", authConfig.GetScope()))

	return tokenManager, nil
}

// Token implements auth.TokenProvider for reuse with another explicit API host.
func (tm *TokenManager) Token(ctx context.Context) (auth.Token, error) {
	value, err := tm.GetTokenContext(ctx)
	if err != nil {
		return auth.Token{}, err
	}
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return auth.Token{Value: value, ExpiresAt: tm.tokenExpiry}, nil
}
