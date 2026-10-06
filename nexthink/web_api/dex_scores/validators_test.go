package dex_scores

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRequestVariables(t *testing.T) {
	type request struct {
		ID       string   `json:"id" required:"true"`
		Items    []string `json:"items" required:"true"`
		Optional *string  `json:"optional,omitempty"`
	}
	_, err := requestVariables(&request{})
	require.Error(t, err)
	_, err = requestVariables(&request{ID: "fixture"})
	require.Error(t, err)
	vars, err := requestVariables(&request{ID: "fixture", Items: []string{}})
	require.NoError(t, err)
	require.NotContains(t, vars, "optional")
	require.Equal(t, []any{}, vars["items"])
}
