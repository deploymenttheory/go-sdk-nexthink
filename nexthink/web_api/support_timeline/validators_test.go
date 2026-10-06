package support_timeline

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateTimeRange(t *testing.T)          { require.Error(t, validateTimeRange(nil)) }
func TestValidateAlertsRequest(t *testing.T)      { require.Error(t, validateAlertsRequest(nil)) }
func TestValidateApplicationRequest(t *testing.T) { require.Error(t, validateApplicationRequest(nil)) }

func TestTimeRangeValidation(t *testing.T) {
	require.NoError(t, validateDates("2026-01-01T00:00", "2026-01-02T00:00"))
	require.Error(t, validateDates("bad", "2026-01-02T00:00"))
	require.Error(t, validateDates("2026-01-01T00:00", "bad"))
	require.Error(t, validateDates("2026-01-02T00:00", "2026-01-01T00:00"))
	require.Error(t, validateDates("2026-01-01T00:00", "2026-01-01T00:00"))
}
