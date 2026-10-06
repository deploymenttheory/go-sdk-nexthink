package dashboards

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

const EndpointAsyncProxy = "/apigateway/proxy/request/dash-graphql-gateway"

// WithAsyncProxy returns an independent service using the UI's long-running GraphQL proxy.
// Existing service configuration and REST routes are unchanged. Redirects follow the client's policy.
func (s *Service) WithAsyncProxy() *Service {
	if _, ok := s.client.(*asyncProxyClient); ok {
		return &Service{client: s.client}
	}
	return &Service{client: &asyncProxyClient{HTTPClient: s.client}}
}

type asyncProxyClient struct{ interfaces.HTTPClient }

func (c *asyncProxyClient) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	if path == Endpoint {
		path = EndpointAsyncProxy
		h := make(map[string]string, len(headers)+1)
		for k, v := range headers {
			h[k] = v
		}
		h["x-nxt-waas-allow-long-running"] = "true"
		headers = h
	}
	return c.HTTPClient.Post(ctx, path, body, headers, result)
}
