package client

import (
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/portalsession"
	"net/url"
	"resty.dev/v3"
	"strings"
)

// preparePortalSession validates the fixed legacy route on every request/retry.
// No bearer provider or OAuth refresh is consulted for this explicit dialect.
func preparePortalSession(req *resty.Request, method, path string) (bool, error) {
	if !portalsession.Is(req.Context()) {
		return false, nil
	}
	u, err := url.Parse(path)
	if err != nil || method != "POST" || !portalsession.Allowed(u.Path) || u.RawPath != "" || u.RawQuery != "" || u.Fragment != "" {
		return false, fmt.Errorf("portal session credentials are restricted to explicitly supported legacy POST routes")
	}
	req.SetAuthToken("")
	req.Header.Del("Authorization")
	if req.HeaderAuthorizationKey != "" {
		req.Header.Del(req.HeaderAuthorizationKey)
	}
	return true, nil
}

func portalCredentialHeader(name string) bool {
	return strings.EqualFold(name, "Cookie") || strings.EqualFold(name, "Authorization") || strings.EqualFold(name, "x-auth-token")
}
