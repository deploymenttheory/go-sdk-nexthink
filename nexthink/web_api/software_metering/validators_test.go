package software_metering

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidationPreventsRequests(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	ctx := context.Background()
	for _, id := range []string{"", " "} {
		_, _, err := s.Get(ctx, id)
		require.Error(t, err)
		_, _, err = s.Delete(ctx, id)
		require.Error(t, err)
		_, _, err = s.Update(ctx, id, load[ConfigurationInput](t, "Create_input"))
		require.Error(t, err)
	}
	_, _, err := s.Create(ctx, nil)
	require.Error(t, err)
	for _, mutate := range []func(*ConfigurationInput){func(r *ConfigurationInput) { r.Name = "" }, func(r *ConfigurationInput) { r.NQLID = "missing_hash" }, func(r *ConfigurationInput) { r.LicenseType = "invalid" }, func(r *ConfigurationInput) { r.ApplicationUUIDs = nil }, func(r *ConfigurationInput) { r.ApplicationUUIDs = []string{" "} }, func(r *ConfigurationInput) { r.Thresholds = nil }, func(r *ConfigurationInput) { r.Thresholds[0].AppType = "invalid" }, func(r *ConfigurationInput) { r.Thresholds[0].MetricType = "invalid" }, func(r *ConfigurationInput) { r.Thresholds[0].TimeUnit = "invalid" }, func(r *ConfigurationInput) { r.Thresholds[0].Value = -1 }} {
		req := load[ConfigurationInput](t, "Create_input")
		mutate(req)
		_, _, err := s.Create(ctx, req)
		require.Error(t, err)
		_, _, err = s.Update(ctx, "id", req)
		require.Error(t, err)
	}
	assert.Zero(t, mock.GetTotalCallCount())
}
