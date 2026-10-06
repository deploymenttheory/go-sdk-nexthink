package nexthink

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type ownedAuthFixture struct{ calls atomic.Int32 }

func (f *ownedAuthFixture) Close() error { f.calls.Add(1); return errCloseFixture }
func (*ownedAuthFixture) Token(context.Context) (auth.Token, error) {
	return auth.Token{Value: "fixture-token"}, nil
}

var errCloseFixture = errors.New("fixture close failure")

func TestCloseOnlyOwnedProvider(t *testing.T) {
	caller := &ownedAuthFixture{}
	c, err := NewClient(&AuthConfig{Instance: "fixture", Region: "eu", WebAPI: &BrowserCredentials{TokenProvider: caller}}, WithLogger(zap.NewNop()))
	require.NoError(t, err)
	require.NoError(t, c.Close())
	assert.Zero(t, caller.calls.Load())
	owned := &ownedAuthFixture{}
	c = &Client{ownedAuth: owned}
	var group sync.WaitGroup
	for range 20 {
		group.Go(func() { assert.ErrorIs(t, c.Close(), errCloseFixture) })
	}
	group.Wait()
	assert.EqualValues(t, 1, owned.calls.Load())
	var nilClient *Client
	require.NoError(t, nilClient.Close())
}

func TestPasswordClientConstructionIsLazyAndOwned(t *testing.T) {
	cfg := &AuthConfig{Instance: "fixture", Region: "eu", WebAPI: &BrowserCredentials{UsernamePassword: &UsernamePasswordCredentials{Username: "local@example.invalid", Password: "fixture-password", BrowserExecutablePath: "/fixture/missing-browser"}}}
	c, err := NewClient(cfg, WithLogger(zap.NewNop()))
	require.NoError(t, err, "construction must not launch the browser")
	require.NotNil(t, c.WebAPI)
	require.NotNil(t, c.ownedAuth)
	require.NoError(t, c.Close())
	require.NoError(t, c.Close())
	provider, ok := c.ownedAuth.(auth.TokenProvider)
	require.True(t, ok)
	_, err = provider.Token(context.Background())
	require.Error(t, err, "closed provider must not launch a browser")
}
