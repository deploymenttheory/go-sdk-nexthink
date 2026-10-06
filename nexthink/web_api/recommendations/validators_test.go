package recommendations

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidationBeforeHTTP(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	for _, test := range []struct {
		id      string
		request *UpdateStatusRequest
	}{{"", &UpdateStatusRequest{Status: StatusNew}}, {"..", &UpdateStatusRequest{Status: StatusNew}}, {"fixture", nil}, {"fixture", &UpdateStatusRequest{Status: "unknown"}}} {
		_, _, err := s.UpdateStatus(context.Background(), test.id, test.request)
		require.Error(t, err)
	}
	require.Equal(t, 0, mock.GetTotalCallCount())
	for _, status := range []string{StatusNew, StatusInProgress, StatusDone, StatusDismissed} {
		require.NoError(t, validateUpdate("fixture", &UpdateStatusRequest{Status: status}))
	}
}
