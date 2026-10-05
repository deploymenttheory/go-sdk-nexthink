package graphql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestValidateRequest(t *testing.T) {
	for _, tc := range []struct{ id, query string }{{"graphql.workflows", " "}, {"graphql.missing", "{__typename}"}, {"shell.user", "{__typename}"}, {"https://other.test/graphql", "{__typename}"}} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(
			transport,
		).Execute(context.Background(), tc.id, GraphQLRequest{Query: tc.query})
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}
