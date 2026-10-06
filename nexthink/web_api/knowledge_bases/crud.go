package knowledge_bases

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetContents lists indexed content; newly uploaded items may not appear immediately.
func (s *Service) GetContents(ctx context.Context) ([]Content, *interfaces.Response, error) {
	var result []Content
	response, err := s.client.Get(ctx, Endpoint+"/contents", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// Create registers an uploaded CSV for asynchronous processing (202 observed).
func (s *Service) Create(ctx context.Context, request *CreateRequest) (*interfaces.Response, error) {
	if err := ValidateCreate(request); err != nil {
		return nil, err
	}
	return s.client.Post(ctx, Endpoint+"/knowledgebase", request, map[string]string{"Content-Type": "application/json"}, nil)
}
func (s *Service) Delete(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	return s.client.Delete(ctx, Endpoint+"/knowledgebase/"+url.PathEscape(id), nil, nil, nil)
}
func (s *Service) GetDownloadURL(ctx context.Context, id string) (*DownloadURLResponse, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result DownloadURLResponse
	response, err := s.client.Get(ctx, Endpoint+"/knowledgebase/"+url.PathEscape(id)+"/download-url", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) UploadFile(ctx context.Context, request *UploadRequest) (*UploadResponse, *interfaces.Response, error) {
	if err := ValidateUpload(request); err != nil {
		return nil, nil, err
	}
	var result UploadResponse
	response, err := s.client.PostWithQuery(ctx, Endpoint+"/file/upload", map[string]string{"originalFileName": request.OriginalFileName, "type": "KB"}, base64.StdEncoding.EncodeToString(request.Data), map[string]string{"Content-Type": "application/octet-stream"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) StartMultipartUpload(ctx context.Context, request *StartMultipartRequest) (*MultipartUpload, *interfaces.Response, error) {
	if request == nil || strings.TrimSpace(request.OriginalFileName) == "" {
		return nil, nil, fmt.Errorf("original filename is required")
	}
	var result MultipartUpload
	response, err := s.client.PostWithQuery(ctx, Endpoint+"/file/multipart", map[string]string{"type": "KB"}, request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func multipartQuery(r MultipartContext) url.Values {
	return url.Values{"contentId": {r.ContentID}, "fileName": {r.FileName}, "uploadId": {r.UploadID}, "type": {"KB"}}
}
func (s *Service) UploadPart(ctx context.Context, request *UploadPartRequest) (*UploadedPart, *interfaces.Response, error) {
	if err := ValidatePart(request); err != nil {
		return nil, nil, err
	}
	query := multipartQuery(request.MultipartContext)
	query.Set("partNumber", strconv.Itoa(request.PartNumber))
	var result UploadedPart
	response, err := s.client.Put(ctx, Endpoint+"/file/multipart?"+query.Encode(), request.EncodedChunk, map[string]string{"Content-Type": "application/octet-stream"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) CompleteMultipartUpload(ctx context.Context, request *CompleteMultipartRequest) (*CompleteMultipartResponse, *interfaces.Response, error) {
	if err := ValidateComplete(request); err != nil {
		return nil, nil, err
	}
	query := map[string]string{}
	for key, values := range multipartQuery(request.MultipartContext) {
		query[key] = values[0]
	}
	var result CompleteMultipartResponse
	response, err := s.client.PostWithQuery(ctx, Endpoint+"/file/multipart/complete", query, request.Parts, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type KnowledgeBasesServiceInterface interface {
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	GetContents(context.Context) ([]Content, *interfaces.Response, error)
	Create(context.Context, *CreateRequest) (*interfaces.Response, error)
	Delete(context.Context, string) (*interfaces.Response, error)
	GetDownloadURL(context.Context, string) (*DownloadURLResponse, *interfaces.Response, error)
	UploadFile(context.Context, *UploadRequest) (*UploadResponse, *interfaces.Response, error)
	StartMultipartUpload(context.Context, *StartMultipartRequest) (*MultipartUpload, *interfaces.Response, error)
	UploadPart(context.Context, *UploadPartRequest) (*UploadedPart, *interfaces.Response, error)
	CompleteMultipartUpload(context.Context, *CompleteMultipartRequest) (*CompleteMultipartResponse, *interfaces.Response, error)
}

var _ KnowledgeBasesServiceInterface = (*Service)(nil)
