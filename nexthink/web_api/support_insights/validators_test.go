package support_insights

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateInsightsRequest(t *testing.T) { require.Error(t, validateInsightsRequest(nil)) }

func TestTimeRangeValidation(t *testing.T) {
	require.NoError(t, validateDates("2026-01-01T00:00", "2026-01-02T00:00"))
	require.Error(t, validateDates("bad", "2026-01-02T00:00"))
	require.Error(t, validateDates("2026-01-01T00:00", "bad"))
	require.Error(t, validateDates("2026-01-02T00:00", "2026-01-01T00:00"))
	require.Error(t, validateDates("2026-01-01T00:00", "2026-01-01T00:00"))
}
