package content_sharing

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_sharing/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestInvalidGrantsDoNotReachTransport(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	service := NewService(transport)
	for _, grant := range []ProfileGrant{{RoleUUID: "fixture-role"}, {RoleUUID: "", Actions: []string{}}, {RoleUUID: "  ", Actions: []string{"view"}}} {
		_, err := service.SetProfiles(context.Background(), "fixture", []ShareContent{{ContentID: "fixture-content", ResourceName: "fixture-resource", Profiles: []ProfileGrant{grant}}})
		require.Error(t, err)
	}
	_, _, err := service.SetLegacyProfiles(context.Background(), &LegacyOptions{Service: "fixture", ContentID: "fixture-content"}, &LegacyUpdateRequest{Profiles: []LegacyProfileGrant{{ProfileID: 7}}})
	require.Error(t, err)
	assert.Zero(t, mock.GetTotalCallCount())
}

func TestExplicitEmptyActionsAreTransmittedForRevocation(t *testing.T) {
	t.Run("modern", func(t *testing.T) {
		transport, mock := testutil.NewTransport(t)
		mock.RegisterResponder("POST", testutil.BaseURL+Endpoint+"/v3/share/profiles?contentKey=fixture", func(r *http.Request) (*http.Response, error) {
			body, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			assert.JSONEq(t, `[{"contentId":"fixture-content","contentName":"Fixture","resourceName":"fixture-resource","profiles":[{"roleUuid":"fixture-role","actions":[]}]}]`, string(body))
			return httpmock.NewStringResponse(204, ""), nil
		})
		_, err := NewService(transport).SetProfiles(context.Background(), "fixture", []ShareContent{{ContentID: "fixture-content", ContentName: "Fixture", ResourceName: "fixture-resource", Profiles: []ProfileGrant{{RoleUUID: "fixture-role", Actions: []string{}}}}})
		require.NoError(t, err)
		assert.Equal(t, 1, mock.GetTotalCallCount())
	})
	t.Run("legacy", func(t *testing.T) {
		transport, mock := testutil.NewTransport(t)
		mock.RegisterResponder("POST", testutil.BaseURL+EndpointLegacy+"/profiles?bcsName=Fixture&contentId=fixture-content&service=fixture-service&tag=fixture", func(r *http.Request) (*http.Response, error) {
			body, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			assert.JSONEq(t, `{"version":3,"profiles":[{"profileId":7,"actions":[]}]}`, string(body))
			return mocks.Responder(200, "SetLegacyProfiles_success")(r)
		})
		_, _, err := NewService(transport).SetLegacyProfiles(context.Background(), &LegacyOptions{Service: "fixture-service", ContentID: "fixture-content", BCSName: "Fixture", Tag: "fixture"}, &LegacyUpdateRequest{Version: 3, Profiles: []LegacyProfileGrant{{ProfileID: 7, Actions: []string{}}}})
		require.NoError(t, err)
		assert.Equal(t, 1, mock.GetTotalCallCount())
	})
}
