package collaboration_tools

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateCallInsightsRequest(t *testing.T) {
	require.Error(t, validateCallInsightsRequest(nil))
}

func TestTimeRangeValidation(t *testing.T) {
	require.NoError(t, validateDates("2026-01-01T00:00", "2026-01-02T00:00"))
	require.Error(t, validateDates("bad", "2026-01-02T00:00"))
	require.Error(t, validateDates("2026-01-01T00:00", "bad"))
	require.Error(t, validateDates("2026-01-02T00:00", "2026-01-01T00:00"))
	require.Error(t, validateDates("2026-01-01T00:00", "2026-01-01T00:00"))
}
