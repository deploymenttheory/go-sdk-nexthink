package graphql

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestExecuteData(t *testing.T) {
	for _, tt := range []struct {
		name, body                            string
		wantError, wantGraphQLError, wantData bool
	}{
		{"success", `{"data":{"value":4}}`, false, false, true},
		{"partial", `{"data":{"value":4},"errors":[{"message":"partial failure"}]}`, true, true, true},
		{"errors only", `{"errors":[{"message":"failed"}]}`, true, true, false},
		{"null data", `{"data":null}`, true, false, false},
		{"missing data", `{}`, true, false, false},
		{"wrong type", `{"data":{"value":"bad"}}`, true, false, false},
		{"wrong type with GraphQL error", `{"data":{"value":"bad"},"errors":[{"message":"failed"}]}`, true, true, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(
				"POST",
				testutil.BaseURL+EndpointWorkflows,
				func(_ *http.Request) (*http.Response, error) {
					response := httpmock.NewStringResponse(200, tt.body)
					response.Header.Set("Content-Type", "application/json")
					return response, nil
				},
			)
			type data struct {
				Value int `json:"value"`
			}
			result, response, err := ExecuteData[data](
				context.Background(),
				NewService(transport),
				"graphql.workflows",
				GraphQLRequest{Query: "query { value }"},
			)
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
			if tt.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			if tt.wantGraphQLError {
				var gqlErrors GraphQLErrors
				require.ErrorAs(t, err, &gqlErrors)
			}
			if tt.wantData {
				require.NotNil(t, result)
				assert.Equal(t, 4, result.Value)
			} else {
				assert.Nil(t, result)
			}
		})
	}
}
