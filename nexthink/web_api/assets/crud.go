package assets

import (
	"context"
	"encoding/base64"
	"net/url"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

type AssetsServiceInterface interface {
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Create(context.Context, *UploadRequest) (*CreateResponse, *interfaces.Response, error)
	Update(context.Context, string, *UploadRequest) (*interfaces.Response, error)
	Delete(context.Context, string) (*interfaces.Response, error)
	GetSignedURL(context.Context, string) (*SignedURLResponse, *interfaces.Response, error)
}

var _ AssetsServiceInterface = (*Service)(nil)

func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Create sends text/plain, not JSON or multipart. Data must contain unencoded file bytes.
func (s *Service) Create(ctx context.Context, request *UploadRequest) (*CreateResponse, *interfaces.Response, error) {
	if err := ValidateUpload(request); err != nil {
		return nil, nil, err
	}
	var result CreateResponse
	response, err := s.client.Post(ctx, Endpoint, dataURL(request), uploadHeaders(request), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Update replaces file bytes; the filename header does not rename an existing asset. The successful response has no JSON body.
func (s *Service) Update(ctx context.Context, id string, request *UploadRequest) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	if err := ValidateUpload(request); err != nil {
		return nil, err
	}
	return s.client.Put(ctx, Endpoint+"/"+url.PathEscape(id), dataURL(request), uploadHeaders(request), nil)
}
func (s *Service) Delete(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	return s.client.Delete(ctx, Endpoint+"/"+url.PathEscape(id), nil, nil, nil)
}

// GetSignedURL is the observed asset read/download operation. It does not download the file.
func (s *Service) GetSignedURL(ctx context.Context, id string) (*SignedURLResponse, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result SignedURLResponse
	response, err := s.client.Post(ctx, EndpointSignedURL, map[string]string{"assetId": id}, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func dataURL(r *UploadRequest) string {
	return "data:" + r.MediaType + ";base64," + base64.StdEncoding.EncodeToString(r.Data)
}
func uploadHeaders(r *UploadRequest) map[string]string {
	return map[string]string{"Content-Type": "text/plain", "Accept": "application/json", "x-nxt-asset-name": r.Name}
}
