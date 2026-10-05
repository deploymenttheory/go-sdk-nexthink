package workflows

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestInvalidUUIDDoesNotSend(t *testing.T) {
	for _, id := range []string{"", " ", "\t\n"} {
		transport, mock := testutil.NewTransport(t)
		s := NewService(transport)
		_, _, err := s.Get(context.Background(), id)
		require.Error(t, err)
		_, _, err = s.Export(context.Background(), id)
		require.Error(t, err)
		require.Zero(t, mock.GetTotalCallCount())
	}
	require.NoError(t, ValidateUUID("opaque-id"))
}

func TestInvalidWritesDoNotSend(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	_, _, err := s.Create(context.Background(), nil)
	require.Error(t, err)
	_, _, err = s.Update(context.Background(), nil)
	require.Error(t, err)
	_, _, err = s.Delete(context.Background(), " ")
	require.Error(t, err)
	require.Zero(t, mock.GetTotalCallCount())
}

func TestWriteValidation(t *testing.T) {
	for _, change := range []func(*CreateRequest){
		func(r *CreateRequest) { r.Workflow.ID = "plain" },
		func(r *CreateRequest) { r.Workflow.Name = " " },
		func(r *CreateRequest) { r.Workflow.Status = "" },
		func(r *CreateRequest) { r.Workflow.UUID = "existing" },
	} {
		request := &CreateRequest{Workflow: *writeRequest(t, "Create")}
		change(request)
		require.Error(t, ValidateCreateRequest(request))
	}
	request := writeRequest(t, "Update")
	request.UUID = ""
	require.Error(t, ValidateUpdateRequest(request))
}
