package connectors

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/connectors/mocks"
	"github.com/stretchr/testify/require"
)

func TestInvalidRequestsDoNotReachTransport(t *testing.T) {
	ctx := context.Background()
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	for _, id := range []string{"", "..", "../fixture", "a/b", "a\\b"} {
		_, _, err := s.Get(ctx, id)
		require.Error(t, err)
		_, _, err = s.GetTemplate(ctx, id)
		require.Error(t, err)
		_, err = s.Delete(ctx, id)
		require.Error(t, err)
	}
	_, _, err := s.ListManualCustomFields(ctx, " ")
	require.Error(t, err)
	_, _, err = s.Create(ctx, nil)
	require.Error(t, err)
	_, _, err = s.Update(ctx, nil)
	require.Error(t, err)
	for _, mutate := range []func(*ConnectorInput){
		func(r *ConnectorInput) { r.ContentID = "not-a-uuid" },
		func(r *ConnectorInput) { r.Template.Name = "" },
		func(r *ConnectorInput) { r.General.Name = " " },
		func(r *ConnectorInput) { r.General.NQLID = "" },
		func(r *ConnectorInput) { r.Credentials.Reference = "" },
		func(r *ConnectorInput) { r.Credentials.Type = "" },
		func(r *ConnectorInput) { r.Scheduling.Cron = "" },
		func(r *ConnectorInput) { r.Scheduling.Timezone = "" },
		func(r *ConnectorInput) { r.Streams = nil },
		func(r *ConnectorInput) { r.Streams[0].Name = "" },
	} {
		request := load[ConnectorInput](t, "Create_request")
		mutate(request)
		_, _, err = s.Create(ctx, request)
		require.Error(t, err)
		_, _, err = s.Update(ctx, request)
		require.Error(t, err)
	}
	require.Zero(t, mock.GetTotalCallCount())
}

func TestEscapedPathAndQuery(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint+"/templates/template%3F%23", mocks.Responder(200, "GetTemplate_success"))
	_, _, err := NewService(transport).GetTemplate(context.Background(), "template?#")
	require.NoError(t, err)
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint+"/manualcustomfields?dataModelObject=device%2Fdevice%26other%3Dx", mocks.Responder(200, "ListManualCustomFields_success"))
	_, _, err = NewService(transport).ListManualCustomFields(context.Background(), "device/device&other=x")
	require.NoError(t, err)
	require.Equal(t, 2, mock.GetTotalCallCount())
}
