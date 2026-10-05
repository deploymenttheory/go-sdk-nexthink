package nexthink

import (
	"crypto/tls"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
)

type (
	clientOptions struct{ common, public, web []client.ClientOption }
	ClientOption  func(*clientOptions)
)

func appendOptions(common, specific []client.ClientOption) []client.ClientOption {
	return append(append([]client.ClientOption(nil), common...), specific...)
}

// WithPublicAPIOptions scopes transport overrides (such as a base URL) to the public API.
func WithPublicAPIOptions(options ...client.ClientOption) ClientOption {
	return func(c *clientOptions) { c.public = append(c.public, options...) }
}

// WithWebAPIOptions scopes transport overrides to the web API, keeping token destinations separate.
func WithWebAPIOptions(options ...client.ClientOption) ClientOption {
	return func(c *clientOptions) { c.web = append(c.web, options...) }
}

func withCommon(option client.ClientOption) ClientOption {
	return func(c *clientOptions) { c.common = append(c.common, option) }
}

func WithLogger(logger *zap.Logger) ClientOption { return withCommon(client.WithLogger(logger)) }

func WithTimeout(
	timeout time.Duration,
) ClientOption {
	return withCommon(client.WithTimeout(timeout))
}

func WithRetryCount(count int) ClientOption { return withCommon(client.WithRetryCount(count)) }

func WithTransport(transport http.RoundTripper) ClientOption {
	return withCommon(client.WithTransport(transport))
}
func WithProxy(proxy string) ClientOption { return withCommon(client.WithProxy(proxy)) }
func WithTLSClientConfig(config *tls.Config) ClientOption {
	return withCommon(client.WithTLSClientConfig(config))
}
func WithDebug() ClientOption { return withCommon(client.WithDebug()) }
