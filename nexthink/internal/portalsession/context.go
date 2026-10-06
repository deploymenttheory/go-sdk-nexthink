// Package portalsession scopes the legacy PortalServlet authentication dialect.
// It is internal: callers use typed legacy resource methods.
package portalsession

import "context"

type contextKey struct{}

const (
	Endpoint  = "/PortalServlet"
	GetAsset  = "/PortalApiServlet/nxportalbranding/getAsset"
	SaveAsset = "/PortalApiServlet/nxportalbranding/saveAsset"
)

func Allowed(path string) bool { return path == Endpoint || path == GetAsset || path == SaveAsset }

// Context marks only fixed-route PortalServlet requests; the transport validates
// the route before honoring this marker. Credentials remain request headers.
func Context(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKey{}, true)
}

func Is(ctx context.Context) bool {
	v, _ := ctx.Value(contextKey{}).(bool)
	return v
}
