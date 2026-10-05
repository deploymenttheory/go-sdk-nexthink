package chrome

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
)

func TestSessionParsing(t *testing.T) {
	value, err := parseSession([]byte(fmt.Sprintf(`{"token":"private-token","expiresAt":%d}`, time.Now().Add(time.Minute).Unix())))
	if err != nil || value.Value != "private-token" {
		t.Fatal("valid session rejected")
	}
	_, err = parseSession([]byte(`{"token":"private-token","expiresAt":1}`))
	if !errors.Is(err, auth.ErrSessionExpired) {
		t.Fatal("expired token accepted")
	}
	_, err = parseSession([]byte(`private-token`))
	if err == nil || strings.Contains(err.Error(), "private-token") {
		t.Fatal("invalid session leaked or accepted")
	}
}

func TestOriginValidation(t *testing.T) {
	for _, origin := range []string{"http://tenant.eu.nexthink.cloud", "https://tenant.eu.nexthink.cloud.evil.test", "https://user:password@tenant.eu.nexthink.cloud", "https://tenant.eu.nexthink.cloud/path"} {
		if _, err := New(origin); err == nil {
			t.Errorf("accepted %s", origin)
		}
	}
	if _, err := New("https://tenant.eu.nexthink.cloud/"); err != nil {
		t.Fatal(err)
	}
}
