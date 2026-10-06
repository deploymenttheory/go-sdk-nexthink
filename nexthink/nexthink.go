// Package nexthink is the single entry point for the Nexthink SDK.
package nexthink

import (
	"fmt"

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
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/applications"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/assets"
	web_campaigns "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/campaigns"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/checklists"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/collector_management"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/connector_credentials"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/connectors"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/custom_fields"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/dashboards"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/data_exporters"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/device_configuration"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/investigations"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/knowledge_bases"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/legacy_connectors"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/license"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/monitors"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/nql_editor"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/nql_queries"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/product_shell"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/ratings"
	web_remote_actions "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/remote_actions"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/rule_based_custom_fields"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/software_metering"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/webhooks"
	web_workflows "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workflows"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/writing_assistant"
	"go.uber.org/zap"
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
	DataExporters         *data_exporters.Service
	Webhooks              *webhooks.Service
	LegacyConnectors      *legacy_connectors.Service
	KnowledgeBases        *knowledge_bases.Service
	Connectors            *connectors.Service
	ConnectorCredentials  *connector_credentials.Service
	Investigations        *investigations.Service
	Checklists            *checklists.Service
	Ratings               *ratings.Service
	Dashboards            *dashboards.Service
	Assets                *assets.Service
	RuleBasedCustomFields *rule_based_custom_fields.Service
	Campaigns             *web_campaigns.Service
	Monitors              *monitors.Service
	CustomFields          *custom_fields.Service
	Applications          *applications.Service
	SoftwareMetering      *software_metering.Service
	WritingAssistant      *writing_assistant.Service
	transport             *client.Transport
	CollectorManagement   *collector_management.Service
	ProductShell          *product_shell.Service
	License               *license.Service
	NQLQueries            *nql_queries.Service
	ContentAdministration *content_administration.Service
	DeviceConfiguration   *device_configuration.Service
	NQLEditor             *nql_editor.Service
	GraphQL               *graphql.Service
	Workflows             *web_workflows.Service
	RemoteActions         *web_remote_actions.Service
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
		DataExporters:         data_exporters.NewService(transport),
		Webhooks:              webhooks.NewService(transport),
		LegacyConnectors:      legacy_connectors.NewService(transport),
		KnowledgeBases:        knowledge_bases.NewService(transport),
		Assets:                assets.NewService(transport),
		Dashboards:            dashboards.NewService(transport),
		Ratings:               ratings.NewService(transport),
		Checklists:            checklists.NewService(transport),
		Connectors:            connectors.NewService(transport),
		ConnectorCredentials:  connector_credentials.NewService(transport),
		Investigations:        investigations.NewService(transport),
		RuleBasedCustomFields: rule_based_custom_fields.NewService(transport),
		Campaigns:             web_campaigns.NewService(transport),
		Monitors:              monitors.NewService(transport),
		CustomFields:          custom_fields.NewService(transport),
		Applications:          applications.NewService(transport),
		SoftwareMetering:      software_metering.NewService(transport),
		WritingAssistant:      writing_assistant.NewService(transport),
		transport:             transport,
		CollectorManagement:   collector_management.NewService(transport),
		ProductShell:          product_shell.NewService(transport),
		License:               license.NewService(transport),
		NQLQueries:            nql_queries.NewService(transport),
		ContentAdministration: content_administration.NewService(transport),
		DeviceConfiguration:   device_configuration.NewService(transport),
		NQLEditor:             nql_editor.NewService(transport),
		GraphQL:               graphql.NewService(transport),
		Workflows:             web_workflows.NewService(transport),
		RemoteActions:         web_remote_actions.NewService(transport),
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
