// Package nexthink is the single entry point for the Nexthink SDK.
package nexthink

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/config"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/campaigns"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/data_management"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/enrichment"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/nql"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/remote_actions"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/spark"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/workflows"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/collector_management"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/device_configuration"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/license"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/nql_editor"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/nql_queries"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/product_shell"
)

type (
	AuthConfig         = config.AuthConfig
	ClientCredentials  = config.ClientCredentials
	BrowserCredentials = config.BrowserCredentials
)

// Client groups the public integration APIs and the browser-facing web APIs.
// A nil family means that no credentials were configured for that family.
type Client struct {
	PublicAPI *PublicAPIClient
	WebAPI    *WebAPIClient
	logger    *zap.Logger
}
type PublicAPIClient struct {
	transport      *client.Transport
	DataManagement *data_management.Service
	Spark          *spark.Service
	Campaigns      *campaigns.Service
	Enrichment     *enrichment.Service
	NQL            *nql.Service
	RemoteActions  *remote_actions.Service
	Workflows      *workflows.Service
}
type WebAPIClient struct {
	transport             *client.Transport
	CollectorManagement   *collector_management.Service
	ProductShell          *product_shell.Service
	License               *license.Service
	NQLQueries            *nql_queries.Service
	ContentAdministration *content_administration.Service
	DeviceConfiguration   *device_configuration.Service
	NQLEditor             *nql_editor.Service
	GraphQL               *graphql.Service
}

// NewClient validates each enabled API family before constructing any transport.
// Public credentials never substitute for a web session, or vice versa.
func NewClient(authConfig *AuthConfig, options ...ClientOption) (*Client, error) {
	if err := authConfig.Validate(); err != nil {
		return nil, fmt.Errorf("invalid authentication configuration: %w", err)
	}
	opts := clientOptions{}
	for _, option := range options {
		if option == nil {
			return nil, fmt.Errorf("client option cannot be nil")
		}
		option(&opts)
	}
	c := &Client{}
	if credentials := authConfig.PublicAPI; credentials != nil {
		transport, err := client.NewTransport(
			credentials.ClientID,
			credentials.ClientSecret,
			authConfig.Instance,
			authConfig.Region,
			appendOptions(opts.common, opts.public)...)
		if err != nil {
			return nil, fmt.Errorf("create public API transport: %w", err)
		}
		c.PublicAPI = newPublicAPIClient(transport)
		c.logger = transport.GetLogger()
	}
	if credentials := authConfig.WebAPI; credentials != nil {
		provider := credentials.TokenProvider
		if provider == nil {
			provider = auth.StaticToken(credentials.AccessToken, credentials.ExpiresAt)
		}
		origin := fmt.Sprintf(
			"https://%s.%s.nexthink.cloud",
			authConfig.Instance,
			authConfig.Region,
		)
		transport, err := client.NewTransportWithTokenProvider(
			origin,
			provider,
			appendOptions(opts.common, opts.web)...)
		if err != nil {
			return nil, fmt.Errorf("create web API transport: %w", err)
		}
		c.WebAPI = newWebAPIClient(transport)
		if c.logger == nil {
			c.logger = transport.GetLogger()
		}
	}
	return c, nil
}

func newPublicAPIClient(transport *client.Transport) *PublicAPIClient {
	return &PublicAPIClient{
		transport:      transport,
		DataManagement: data_management.NewService(transport),
		Spark:          spark.NewService(transport),
		Campaigns:      campaigns.NewService(transport),
		Enrichment:     enrichment.NewService(transport),
		NQL:            nql.NewService(transport),
		RemoteActions:  remote_actions.NewService(transport),
		Workflows:      workflows.NewService(transport),
	}
}

func newWebAPIClient(transport *client.Transport) *WebAPIClient {
	return &WebAPIClient{
		transport:             transport,
		CollectorManagement:   collector_management.NewService(transport),
		ProductShell:          product_shell.NewService(transport),
		License:               license.NewService(transport),
		NQLQueries:            nql_queries.NewService(transport),
		ContentAdministration: content_administration.NewService(transport),
		DeviceConfiguration:   device_configuration.NewService(transport),
		NQLEditor:             nql_editor.NewService(transport),
		GraphQL:               graphql.NewService(transport),
	}
}

// NewClientFromEnv selects public, web or both via NEXTHINK_API (default public).
func NewClientFromEnv(options ...ClientOption) (*Client, error) {
	cfg, err := AuthConfigFromEnv()
	if err != nil {
		return nil, err
	}
	return NewClient(cfg, options...)
}
func AuthConfigFromEnv() (*AuthConfig, error) { return config.AuthConfigFromEnv() }
func (c *Client) GetLogger() *zap.Logger      { return c.logger }
func (c *PublicAPIClient) GetTokenManager() *client.TokenManager {
	return c.transport.GetTokenManager()
}
func (c *PublicAPIClient) RefreshToken() error { return c.transport.RefreshToken() }
func (c *PublicAPIClient) InvalidateToken()    { c.transport.InvalidateToken() }
