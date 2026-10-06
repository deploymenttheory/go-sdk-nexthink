package vdi

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateTimelineRequest(t *testing.T) { require.Error(t, validateTimelineRequest(nil)) }
func TestValidateHealthRequest(t *testing.T)   { require.Error(t, validateHealthRequest(nil)) }
func TestValidateHostnameRequest(t *testing.T) { require.Error(t, validateHostnameRequest(nil)) }

func TestTimeRangeValidation(t *testing.T) {
	require.NoError(t, validateDates("2026-01-01T00:00:00", "2026-01-02T00:00:00"))
	require.Error(t, validateDates("bad", "2026-01-02T00:00:00"))
	require.Error(t, validateDates("2026-01-01T00:00:00", "bad"))
	require.Error(t, validateDates("2026-01-02T00:00:00", "2026-01-01T00:00:00"))
	require.Error(t, validateDates("2026-01-01T00:00:00", "2026-01-01T00:00:00"))
}
