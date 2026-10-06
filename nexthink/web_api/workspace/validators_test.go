package workspace

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidation(t *testing.T) {
	require.Error(t, validateID(""))
	require.Error(t, validateID(".."))
	require.NoError(t, validateID("a/b ?"))
	_, err := pageQuery(&PageOptions{Limit: -1})
	require.Error(t, err)
	q, err := pageQuery(&PageOptions{Before: "opaque+/=", Limit: 25})
	require.NoError(t, err)
	assert.Equal(t, "?before=opaque%2B%2F%3D&limit=25", q)
	require.Error(t, validateChat(nil))
	r := load[ChatRequest](t, "Chat_request")
	r.Source = "invalid"
	require.Error(t, validateChat(r))
	r.Source = "workspace"
	r.Messages = nil
	require.Error(t, validateChat(r))
}
