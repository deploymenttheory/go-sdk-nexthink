package connector_credentials

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/connector_credentials/mocks"
	"github.com/stretchr/testify/require"
)

func TestInvalidRequestsDoNotReachTransport(t *testing.T) {
	ctx := context.Background()
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	for _, id := range []string{"", "..", "conn_cr-1/extra", "conn_cr-1?x", "other-1"} {
		_, _, err := s.Get(ctx, id)
		require.Error(t, err)
		_, _, err = s.Delete(ctx, id)
		require.Error(t, err)
		_, _, err = s.Create(ctx, id, load[CredentialInput](t, "Create_request"))
		require.Error(t, err)
		_, _, err = s.Update(ctx, id, load[CredentialInput](t, "Update_request"))
		require.Error(t, err)
	}
	_, _, err := s.Create(ctx, "conn_cr-42", nil)
	require.Error(t, err)
	for _, mutate := range []func(*CredentialInput){
		func(r *CredentialInput) { r.Config.ConnectorType = "conn_cr-99" },
		func(r *CredentialInput) { r.Config.RunTime = "" },
		func(r *CredentialInput) { r.Config.ConnectionDetails = nil },
		func(r *CredentialInput) { r.Config.ConnectionDetails[0].Key = "" },
	} {
		r := load[CredentialInput](t, "Create_request")
		mutate(r)
		_, _, err = s.Create(ctx, "conn_cr-42", r)
		require.Error(t, err)
		_, _, err = s.Update(ctx, "conn_cr-42", r)
		require.Error(t, err)
	}
	require.Zero(t, mock.GetTotalCallCount())
}
func TestSecretEnvelopeAndInputPreserved(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	request := load[CredentialInput](t, "Create_with_secret_request")
	before, err := json.Marshal(request)
	require.NoError(t, err)
	mock.RegisterResponder("POST", testutil.BaseURL+Endpoint+"/credential/conn_cr-42", func(r *http.Request) (*http.Response, error) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.JSONEq(t, string(mocks.Fixture("Create_with_secret_request")), string(body))
		return mocks.Responder(201, "Create_success")(r)
	})
	_, _, err = NewService(transport).Create(context.Background(), "conn_cr-42", request)
	require.NoError(t, err)
	after, err := json.Marshal(request)
	require.NoError(t, err)
	require.Equal(t, string(before), string(after))
}
