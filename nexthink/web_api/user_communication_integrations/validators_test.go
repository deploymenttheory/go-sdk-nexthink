package user_communication_integrations

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidationBeforeTransport(t *testing.T) {
	ctx := context.Background()
	valid := &IntegrationInput{Content: Content{AzureConnectorID: "fixture"}}
	require.NoError(t, ValidateInput(valid))
	require.NoError(t, ValidateInput(&IntegrationInput{Content: Content{AzureTenantID: "tenant"}}))
	cases := []func(*Service) (*interfaces.Response, error){
		func(s *Service) (*interfaces.Response, error) { _, r, e := s.Get(ctx, "../bad"); return r, e },
		func(s *Service) (*interfaces.Response, error) { _, r, e := s.Update(ctx, "../bad", valid); return r, e },
		func(s *Service) (*interfaces.Response, error) { _, r, e := s.Create(ctx, nil); return r, e },
		func(s *Service) (*interfaces.Response, error) {
			_, r, e := s.Create(ctx, &IntegrationInput{})
			return r, e
		},
		func(s *Service) (*interfaces.Response, error) { _, r, e := s.Update(ctx, "fixture", nil); return r, e },
		func(s *Service) (*interfaces.Response, error) { return s.Delete(ctx, "../bad") },
	}
	for _, call := range cases {
		transport, mock := testutil.NewTransport(t)
		response, err := call(NewService(transport))
		require.Error(t, err)
		assert.Nil(t, response)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}
