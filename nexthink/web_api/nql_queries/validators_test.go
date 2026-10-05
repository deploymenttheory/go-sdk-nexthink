package nql_queries

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestInvalidQueryWrites(t *testing.T) {
	for _, update := range []bool{false, true} {
		for _, req := range []*SaveQueryRequest{nil, {}, {NQLAPIID: "#q", Name: "q"}, {NQLAPIID: "#q", NQL: "devices"}, {Name: "q", NQL: "devices"}, {NQLAPIID: "#q", Name: "q", NQL: " "}} {
			transport, mock := testutil.NewTransport(t)
			s := NewService(transport)
			if update {
				result, resp, err := s.Update(context.Background(), req)
				require.Error(t, err)
				assert.Nil(t, result)
				assert.Nil(t, resp)
			} else {
				result, resp, err := s.Create(context.Background(), req)
				require.Error(t, err)
				assert.Nil(t, result)
				assert.Nil(t, resp)
			}
			assert.Zero(t, mock.GetTotalCallCount())
		}
	}
}

func TestInvalidContentIDs(t *testing.T) {
	for _, id := range []string{"", " ", ".", "..", "a/b", `a\b`} {
		transport, mock := testutil.NewTransport(t)
		s := NewService(transport)
		result, resp, err := s.Get(context.Background(), id)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		resp, err = s.Delete(context.Background(), id)
		require.Error(t, err)
		assert.Nil(t, resp)
		result, resp, err = s.Update(
			context.Background(),
			&SaveQueryRequest{ContentID: id, NQLAPIID: "#q", Name: "q", NQL: "devices"},
		)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}
