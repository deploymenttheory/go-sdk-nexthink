package client

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"resty.dev/v3"
)

func TestTypedResponsePreservesExactBody(t *testing.T) {
	for _, tc := range []struct {
		name       string
		size       int
		compressed bool
	}{{"small", 16, false}, {"large", 2 * 1024 * 1024, false}, {"gzip", 128 * 1024, true}} {
		t.Run(tc.name, func(t *testing.T) {
			document, err := json.Marshal(map[string]string{"id": "fixture", "message": strings.Repeat("x", tc.size), "unmodeled": "retained"})
			require.NoError(t, err)
			body := append([]byte(" \n"), document...)
			body = append(body, []byte("\n \t")...)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if tc.compressed {
					w.Header().Set("Content-Encoding", "gzip")
					writer := gzip.NewWriter(w)
					_, writeErr := writer.Write(body)
					require.NoError(t, writeErr)
					require.NoError(t, writer.Close())
					return
				}
				_, writeErr := w.Write(body)
				require.NoError(t, writeErr)
			}))
			defer server.Close()
			transport := setupTestClient(t, server.URL)
			var decoded testResponse
			response, err := transport.Get(context.Background(), "/fixture", nil, nil, &decoded)
			require.NoError(t, err)
			require.Equal(t, "fixture", decoded.ID)
			require.Len(t, decoded.Message, tc.size)
			require.True(t, bytes.Equal(body, response.Body), "raw bytes must survive typed decoding, including unmodeled fields and whitespace; expected %d bytes, got %d", len(body), len(response.Body))
			require.Equal(t, int64(len(body)), response.Size)
		})
	}
}

func TestMalformedTypedResponsePreservesBody(t *testing.T) {
	body := []byte(" \n{\"id\":\"fixture\", broken\n")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(body)
		require.NoError(t, err)
	}))
	defer server.Close()
	transport := setupTestClient(t, server.URL)
	var decoded testResponse
	response, err := transport.Get(context.Background(), "/fixture", nil, nil, &decoded)
	require.Error(t, err)
	require.NotNil(t, response)
	require.Equal(t, body, response.Body)
}

func TestGetBytesPreservesWhitespaceAndBinary(t *testing.T) {
	body := []byte{' ', '\n', 0, 0xff, 'x', '\r', '\t', ' '}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		_, err := w.Write(body)
		require.NoError(t, err)
	}))
	defer server.Close()
	transport := setupTestClient(t, server.URL)
	response, actual, err := transport.GetBytes(context.Background(), "/fixture", nil, nil)
	require.NoError(t, err)
	require.Equal(t, body, actual)
	require.Equal(t, body, response.Body)
}

func TestResponseBufferingRespectsBodyLimit(t *testing.T) {
	body := append([]byte(`{"message":"`), bytes.Repeat([]byte("x"), 4096)...)
	body = append(body, []byte(`"}`)...)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(body)
		require.NoError(t, err)
	}))
	defer server.Close()
	transport := setupTestClient(t, server.URL)
	transport.GetHTTPClient().SetResponseBodyLimit(128)
	var decoded testResponse
	_, err := transport.Get(context.Background(), "/fixture", nil, nil, &decoded)
	require.Error(t, err)
	require.ErrorIs(t, err, resty.ErrReadExceedsThresholdLimit)
}
