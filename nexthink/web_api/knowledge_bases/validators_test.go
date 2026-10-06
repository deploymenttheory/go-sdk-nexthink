package knowledge_bases

import (
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestDownloadURLDecoding(t *testing.T) {
	url, err := (DownloadURLResponse{URL: base64.StdEncoding.EncodeToString([]byte("https://download.example.invalid/fixture"))}).DecodeURL()
	require.NoError(t, err)
	require.Equal(t, "https://download.example.invalid/fixture", url)
	_, err = (DownloadURLResponse{URL: "%%%"}).DecodeURL()
	require.Error(t, err)
}

func TestMultipartSendsSlicesOfWholeEncodedFile(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	// Splitting between base64 quanta must not add padding or encode either slice again.
	encoded := base64.StdEncoding.EncodeToString([]byte("synthetic CSV bytes"))
	chunks := []string{encoded[:5], encoded[5:]}
	for i, chunk := range chunks {
		mock.RegisterResponder("PUT", `=~^https://test\.eu\.nexthink\.cloud/apigateway/knowledge-manager/api/v1/file/multipart\?`, func(r *http.Request) (*http.Response, error) {
			b, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			require.Equal(t, chunk, string(b))
			require.Equal(t, "file + name.csv", r.URL.Query().Get("fileName"))
			require.Equal(t, "upload+/=", r.URL.Query().Get("uploadId"))
			return httpmock.NewJsonResponse(200, UploadedPart{PartNumber: i + 1, ETag: "fixture", Checksum: "fixture"})
		})
		_, _, err := NewService(transport).UploadPart(context.Background(), &UploadPartRequest{MultipartContext: MultipartContext{ContentID: "fixture-id", FileName: "file + name.csv", UploadID: "upload+/="}, PartNumber: i + 1, EncodedChunk: chunk})
		require.NoError(t, err)
	}
	require.Equal(t, 2, mock.GetTotalCallCount())
}

func TestRejectUnorderedMultipartCompletion(t *testing.T) {
	request := &CompleteMultipartRequest{MultipartContext: MultipartContext{ContentID: "fixture", FileName: "fixture.csv", UploadID: "fixture"}, Parts: []UploadedPart{{PartNumber: 2, ETag: "fixture", Checksum: "fixture"}, {PartNumber: 1, ETag: "fixture", Checksum: "fixture"}}}
	require.Error(t, ValidateComplete(request))
	request.Parts[1].PartNumber = 2
	require.Error(t, ValidateComplete(request))
	request.Parts[1].PartNumber = 3
	require.NoError(t, ValidateComplete(request))
}
