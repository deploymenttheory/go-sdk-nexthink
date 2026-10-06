package workspace_tasks

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidation(t *testing.T) {
	require.Error(t, validateID(""))
	require.Error(t, validateTask(nil))
	require.Error(t, validateTask(&TaskRequest{Title: "test"}))
	require.NoError(t, validateTask(&TaskRequest{Title: "test", Prompt: "Define SDK", Enabled: false, Schedule: TaskSchedule{Type: "once", Datetime: "2099-01-01T00:00:00Z"}}))
}
