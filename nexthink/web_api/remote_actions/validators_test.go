package remote_actions

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
		_, _, err = s.GetForView(context.Background(), id)
		require.Error(t, err)
		require.Zero(t, mock.GetTotalCallCount())
	}
	require.NoError(t, ValidateUUID("opaque-id"))
}

func TestEmptyScriptDoesNotSend(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	_, _, err := s.InspectBashScript(context.Background(), nil)
	require.Error(t, err)
	_, _, err = s.InspectPowerShellScript(context.Background(), nil)
	require.Error(t, err)
	_, _, err = s.GetPowerShellSignature(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, mock.GetTotalCallCount())
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
	request := writeRequest(t, "Create")
	request.ID = "plain"
	require.Error(t, ValidateCreateRequest(request))
	request = writeRequest(t, "Create")
	request.Name = " "
	require.Error(t, ValidateCreateRequest(request))
	request = writeRequest(t, "Update")
	request.ID = ""
	require.Error(t, ValidateUpdateRequest(request))
}
