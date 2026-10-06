package workspace_agents

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidation(t *testing.T) {
	require.Error(t, validateID(""))
	require.Error(t, validateSkill(nil))
	require.Error(t, validateSkill(&SkillRequest{Name: "test"}))
	require.NoError(t, validateSkill(&SkillRequest{Name: "test", Instructions: "define SDK"}))
}
