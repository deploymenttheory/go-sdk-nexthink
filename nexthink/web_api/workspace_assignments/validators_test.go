package workspace_assignments

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidation(t *testing.T) {
	require.Error(t, validateID(""))
	require.Error(t, validateAssignment(nil))
	require.Error(t, validateAssignment(&AssignmentUpdate{State: "invalid", Revision: 1}))
	require.Error(t, validateAssignment(&AssignmentUpdate{AssigneeID: json.RawMessage(`42`), Revision: 1}))
	require.NoError(t, validateAssignment(&AssignmentUpdate{AssigneeID: json.RawMessage(`null`), Revision: 1}))
	q := assignmentQuery(&AssignmentOptions{Sources: []string{"a", "b"}, Sort: "priority"})
	assert.Equal(t, "?sort=priority&sources=a&sources=b", q)
}
