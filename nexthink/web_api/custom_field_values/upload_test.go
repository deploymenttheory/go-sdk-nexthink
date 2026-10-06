package custom_field_values

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/custom_field_values/mocks"
)

func TestCSVMultipartContracts(t *testing.T) {
	var input struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
	require.NoError(t, json.Unmarshal(mocks.Fixture("CSV_input"), &input))
	for _, dryRun := range []bool{true, false} {
		name := "ImportCSV"
		path := Endpoint
		if dryRun {
			name = "ValidateCSV"
			path += "?dryRun=true"
		}
		t.Run(name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				reader, err := r.MultipartReader()
				require.NoError(t, err)
				part, err := reader.NextPart()
				require.NoError(t, err)
				assert.Equal(t, "file", part.FormName())
				assert.Equal(t, input.Name, part.FileName())
				data, err := io.ReadAll(part)
				require.NoError(t, err)
				assert.Equal(t, input.Content, string(data))
				_, err = reader.NextPart()
				assert.ErrorIs(t, err, io.EOF)
				return mocks.Responder(202, "CSV_success")(r)
			})
			service := NewService(transport)
			call := service.ImportCSV
			if dryRun {
				call = service.ValidateCSV
			}
			result, response, err := call(context.Background(), input.Name, strings.NewReader(input.Content), int64(len(input.Content)))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 202, response.StatusCode)
			assert.Empty(t, result)
			assert.NotNil(t, result)
			for _, code := range []int{400, 401, 403, 413} {
				mock.RegisterResponder("POST", testutil.BaseURL+path, mocks.Responder(code, "CSV_error"))
				_, res, err := call(context.Background(), input.Name, strings.NewReader(input.Content), int64(len(input.Content)))
				require.Error(t, err)
				require.NotNil(t, res)
				assert.Equal(t, code, res.StatusCode)
			}
			mock.RegisterResponder("POST", testutil.BaseURL+path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
			_, _, err = call(context.Background(), input.Name, strings.NewReader(input.Content), int64(len(input.Content)))
			require.Error(t, err)
		})
	}
}
func TestCSVValidationPreventsRequest(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	for _, call := range []func(context.Context, string, io.Reader, int64) ([]string, *interfaces.Response, error){s.ValidateCSV, s.ImportCSV} {
		_, res, err := call(context.Background(), "fixture.csv", nil, 10)
		require.Error(t, err)
		assert.Nil(t, res)
		_, res, err = call(context.Background(), "", strings.NewReader("x"), 1)
		require.Error(t, err)
		assert.Nil(t, res)
		_, res, err = call(context.Background(), "fixture.csv", strings.NewReader(""), 0)
		require.Error(t, err)
		assert.Nil(t, res)
	}
	assert.Zero(t, mock.GetTotalCallCount())
}
