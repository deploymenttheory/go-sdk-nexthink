package global_search

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConcatenatedObjectsAndPartialErrors(t *testing.T) {
	raw := `{"category":"device","results":[{"name":"a}{b","foundByMetaKey":[{"key":"user","value":"fixture"}],"future":42}]}{"errorResponse":{"message":"provider unavailable","code":503,"source":"package"}}`
	r, err := decode([]byte(raw))
	require.Error(t, err)
	require.Len(t, r.Events, 2)
	assert.Equal(t, "a}{b", r.Events[0].Results[0].Name)
	assert.Equal(t, "fixture", r.Events[0].Results[0].FoundByMetaKey[0].Value)
	var provider *ProviderError
	require.ErrorAs(t, err, &provider)
	assert.Equal(t, "package", provider.Source)
	encoded, e := json.Marshal(r.Events[0])
	require.NoError(t, e)
	assert.JSONEq(t, `{"category":"device","results":[{"name":"a}{b","foundByMetaKey":[{"key":"user","value":"fixture"}],"future":42}]}`, string(encoded))
}
func TestMalformedTailRetainsEvents(t *testing.T) {
	for _, tail := range []string{`{`, `null`, `[]`} {
		r, err := decode([]byte(`{"category":"device"}` + tail))
		require.Error(t, err)
		require.Len(t, r.Events, 1)
	}
}
func TestValidation(t *testing.T) {
	for _, r := range []*SearchRequest{nil, {}, {Search: " ", MaxResults: 5}, {Search: "fixture", MaxResults: 0}, {Search: "fixture", MaxResults: -1}} {
		require.Error(t, validateRequest(r))
	}
	require.NoError(t, validateRequest(&SearchRequest{Search: "fixture", MaxResults: 5}))
}
