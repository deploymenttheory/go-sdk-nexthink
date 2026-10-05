package data_management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestValidateDeleteDevices(t *testing.T) {
	for _, req := range []*DeleteDevicesRequest{nil, {}, {Devices: make([]Device, 101)}, {Devices: []Device{{UID: " ", Name: "fixture"}}}, {Devices: []Device{{UID: "invalid", Name: " "}}}} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(transport).DeleteDevices(context.Background(), req, "")
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
	for _, size := range []int{1, 100} {
		devices := make([]Device, size)
		for i := range devices {
			devices[i] = Device{UID: "invalid", Name: "fixture"}
		}
		require.NoError(t, ValidateDeleteDevicesRequest(&DeleteDevicesRequest{Devices: devices}))
	}
}
