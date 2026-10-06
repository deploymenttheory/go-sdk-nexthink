package content_sharing

import (
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_sharing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// The legacy UI checks this string discriminator on an HTTP200 response.
// Keep the package's documented inspect-Status contract, including raw metadata.
func TestLegacyBusinessFailurePreservesCode(t *testing.T) {
	for _, tt := range contractCases(t) {
		if tt.name != "GetLegacyProfiles" && tt.name != "SetLegacyProfiles" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, mocks.Responder(200, "LegacyBusinessFailure"))
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			require.IsType(t, &LegacyProfilesResponse{}, result)
			value := result.(*LegacyProfilesResponse)
			assert.False(t, value.Status.Success)
			assert.Equal(t, "profile_view_domain_restriction_error", value.Status.Code)
			assert.Len(t, value.Status.RestrictedProfiles, 1)
			assert.Nil(t, value.Result)
			encoded, err := json.Marshal(value)
			require.NoError(t, err)
			assert.JSONEq(t, string(mocks.Fixture("LegacyBusinessFailure")), string(encoded))
			assert.JSONEq(t, string(mocks.Fixture("LegacyBusinessFailure")), string(response.Body))
			assert.Equal(t, 200, response.StatusCode)
		})
	}
}
