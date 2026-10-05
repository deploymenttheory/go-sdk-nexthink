// Package experimental exposes observed, undocumented Nexthink browser APIs.
// Contracts may change independently of the public API. See Operations for
// the discovery evidence and authentication status of each implemented route.
package experimental

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Client struct {
	transport           *client.Transport
	CollectorManagement *CollectorManagementService
	ProductShell        *ProductShellService
	License             *LicenseService
	NQLQueries          *NQLQueriesService
}

// NewClient uses an explicitly supplied identity on the tenant's browser origin.
// Public client credentials can be supplied via the public client's TokenManager;
// whether that identity is authorized is determined by each endpoint.
func NewClient(instance, region string, provider auth.TokenProvider, options ...client.ClientOption) (*Client, error) {
	if err := client.ValidateTransportConfig("provider", "provider", instance, region); err != nil {
		return nil, err
	}
	if strings.ContainsAny(instance, "/\\:@?#") {
		return nil, fmt.Errorf("invalid instance name")
	}
	t, err := client.NewTransportWithTokenProvider(fmt.Sprintf("https://%s.%s.nexthink.cloud", instance, region), provider, options...)
	if err != nil {
		return nil, err
	}
	c := &Client{transport: t}
	c.CollectorManagement = &CollectorManagementService{c}
	c.ProductShell = &ProductShellService{c}
	c.License = &LicenseService{c}
	c.NQLQueries = &NQLQueriesService{c}
	return c, nil
}

// Request supplies path placeholders, query parameters and an optional JSON body.
// Tokens and cookies must be supplied through the client's provider, not Request.
type Request struct {
	PathParams  map[string]string
	Query       map[string]string
	Body        any
	ContentType string
}

// Do executes an observed operation from the catalog and retains the original JSON.
// Operations whose schemas remain undocumented expose raw JSON rather than
// inventing types or silently discarding fields.
func (c *Client) Do(ctx context.Context, operationID string, request Request) (json.RawMessage, *interfaces.Response, error) {
	op, ok := operationByID(operationID)
	if !ok {
		return nil, nil, fmt.Errorf("unknown experimental operation %q", operationID)
	}
	path := op.Path
	for key, value := range request.PathParams {
		if value == "" || value == "." || value == ".." || strings.ContainsAny(value, "/\\") || !strings.Contains(path, "{"+key+"}") {
			return nil, nil, fmt.Errorf("invalid path parameter %q", key)
		}
		path = strings.ReplaceAll(path, "{"+key+"}", url.PathEscape(value))
	}
	if strings.ContainsAny(path, "{}") {
		return nil, nil, fmt.Errorf("missing path parameters for %s", operationID)
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	if request.ContentType != "" {
		headers["Content-Type"] = request.ContentType
	}
	var result json.RawMessage
	var resp *interfaces.Response
	var err error
	switch op.Method {
	case "GET":
		if request.Body != nil {
			return nil, nil, fmt.Errorf("GET operation does not accept a body")
		}
		resp, err = c.transport.Get(ctx, path, request.Query, headers, &result)
	case "POST":
		resp, err = c.transport.PostWithQuery(ctx, path, request.Query, request.Body, headers, &result)
	case "PUT", "PATCH", "DELETE":
		if len(request.Query) > 0 {
			v := url.Values{}
			for k, x := range request.Query {
				v.Set(k, x)
			}
			path += "?" + v.Encode()
		}
		switch op.Method {
		case "PUT":
			resp, err = c.transport.Put(ctx, path, request.Body, headers, &result)
		case "PATCH":
			resp, err = c.transport.Patch(ctx, path, request.Body, headers, &result)
		case "DELETE":
			resp, err = c.transport.DeleteWithBody(ctx, path, request.Body, headers, &result)
		}
	default:
		return nil, nil, fmt.Errorf("unsupported operation method %q", op.Method)
	}
	return result, resp, err
}

func decode[T any](raw json.RawMessage, resp *interfaces.Response, err error) (*T, *interfaces.Response, error) {
	if err != nil {
		return nil, resp, err
	}
	var result T
	if err = json.Unmarshal(raw, &result); err != nil {
		return nil, resp, fmt.Errorf("decode experimental response: %w", err)
	}
	return &result, resp, nil
}

type CollectorManagementService struct{ client *Client }
type DownloadLinks struct {
	Windows      json.RawMessage `json:"windowsInstallerDownloadLinkData"`
	MacOS        json.RawMessage `json:"macOsInstallerDownloadLinkData"`
	VDIExtension json.RawMessage `json:"vdiExtInstallerDownloadLinkData"`
}
type UpdateConfiguration struct {
	ConfigRevision     json.RawMessage `json:"configRevision"`
	BCSConfig          json.RawMessage `json:"bcsConfig"`
	ProductConfigFlags json.RawMessage `json:"productConfigFlags"`
}

func (s *CollectorManagementService) GetDownloadLinks(ctx context.Context) (*DownloadLinks, *interfaces.Response, error) {
	return decode[DownloadLinks](s.client.Do(ctx, "collector.download_links", Request{}))
}
func (s *CollectorManagementService) GetUpdateConfiguration(ctx context.Context) (*UpdateConfiguration, *interfaces.Response, error) {
	return decode[UpdateConfiguration](s.client.Do(ctx, "collector.update_configuration", Request{}))
}

type ProductShellService struct{ client *Client }

func (s *ProductShellService) GetMenu(ctx context.Context) (json.RawMessage, *interfaces.Response, error) {
	return s.client.Do(ctx, "shell.menu", Request{Body: map[string]any{}})
}
func (s *ProductShellService) GetUser(ctx context.Context) (json.RawMessage, *interfaces.Response, error) {
	return s.client.Do(ctx, "shell.user", Request{})
}
func (s *ProductShellService) GetModules(ctx context.Context) (json.RawMessage, *interfaces.Response, error) {
	return s.client.Do(ctx, "shell.modules", Request{})
}
func (s *ProductShellService) GetConfiguration(ctx context.Context) (json.RawMessage, *interfaces.Response, error) {
	return s.client.Do(ctx, "shell.configuration", Request{})
}
func (s *ProductShellService) GetFlag(ctx context.Context, flag string) (json.RawMessage, *interfaces.Response, error) {
	return s.client.Do(ctx, "shell.flag", Request{PathParams: map[string]string{"flag": flag}})
}
func (s *ProductShellService) GetDynamicMenu(ctx context.Context, menu string) (json.RawMessage, *interfaces.Response, error) {
	return s.client.Do(ctx, "shell.dynamic_menu", Request{PathParams: map[string]string{"menu": menu}})
}

type LicenseService struct{ client *Client }

func (s *LicenseService) GetFeatureStatus(ctx context.Context, feature string) (json.RawMessage, *interfaces.Response, error) {
	return s.client.Do(ctx, "license.feature_status", Request{PathParams: map[string]string{"feature": feature}})
}
