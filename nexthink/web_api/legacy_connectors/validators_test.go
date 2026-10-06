package legacy_connectors

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/legacy_connectors/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestCollaborationConfiguration(t *testing.T) {
	for _, enabled := range []*bool{nil, new(bool)} {
		transport, mock := testutil.NewTransport(t)
		request := contractLoad[ConfigurationInput](t, "CollaborationConfiguration_request")
		request.Enabled = enabled
		mock.RegisterResponder("POST", testutil.BaseURL+Endpoint+"/config/ms_teams-1", func(r *http.Request) (*http.Response, error) {
			body, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			var expected map[string]any
			require.NoError(t, json.Unmarshal(mocks.Fixture("CollaborationConfiguration_request"), &expected))
			if enabled != nil {
				expected["enabled"] = false
			}
			want, err := json.Marshal(expected)
			require.NoError(t, err)
			assert.JSONEq(t, string(want), string(body))
			return httpmock.NewStringResponse(200, "saved"), nil
		})
		_, _, err := NewService(transport).Update(context.Background(), "ms_teams-1", request)
		require.NoError(t, err)
	}
}
func TestSecretValidationBeforeTransport(t *testing.T) {
	ctx := context.Background()
	cases := []func(*Service) (*interfaces.Response, error){
		func(s *Service) (*interfaces.Response, error) { _, r, e := s.HasSecrets(ctx, "../bad"); return r, e },
		func(s *Service) (*interfaces.Response, error) {
			return s.UpdateSecrets(ctx, "../bad", &SecretRequest{})
		},
		func(s *Service) (*interfaces.Response, error) { return s.UpdateSecrets(ctx, "zoom-1", nil) },
		func(s *Service) (*interfaces.Response, error) {
			return s.UpdateSecrets(ctx, "zoom-1", &SecretRequest{})
		},
	}
	for _, call := range cases {
		transport, mock := testutil.NewTransport(t)
		response, err := call(NewService(transport))
		require.Error(t, err)
		assert.Nil(t, response)
		assert.Zero(t, mock.GetTotalCallCount())
	}
	request := contractLoad[ConfigurationInput](t, "CollaborationConfiguration_request")
	request.Mapping = nil
	require.Error(t, ValidateInput("ms_teams-1", request))
}
