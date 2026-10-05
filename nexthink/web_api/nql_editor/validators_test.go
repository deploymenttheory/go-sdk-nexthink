package nql_editor

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestInvalidEditorRequests(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	for _, req := range []*ValidationRequest{nil, {}, {Document: Document{URI: "inmemory://fixture", LanguageID: "nql"}, Rules: json.RawMessage("broken")}} {
		result, resp, err := s.Validate(context.Background(), req, "")
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
	}
	for _, req := range []*PositionRequest{nil, {}, {Document: Document{URI: "fixture", LanguageID: "nql"}, Position: Position{Line: -1}}, {Document: Document{URI: "fixture", LanguageID: "nql"}, Position: Position{Character: -1}}} {
		result, resp, err := s.Complete(context.Background(), req, "")
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		hover, resp, err := s.Hover(context.Background(), req)
		require.Error(t, err)
		assert.Nil(t, hover)
		assert.Nil(t, resp)
	}
	for _, item := range []CompletionItem{nil, {}, {"label": json.RawMessage(`42`)}, {"label": json.RawMessage(`""`)}, {"label": json.RawMessage(`"name"`), "data": json.RawMessage("broken")}} {
		result, resp, err := s.Resolve(context.Background(), item)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
	}
	assert.Zero(t, mock.GetTotalCallCount())
	require.NoError(
		t,
		ValidateDocument(Document{URI: "inmemory://empty", LanguageID: "nql", Text: ""}),
	)
}
