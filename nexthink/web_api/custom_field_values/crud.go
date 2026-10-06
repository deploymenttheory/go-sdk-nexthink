package custom_field_values

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context, uri string) (*ListResponse, *interfaces.Response, error) {
	if strings.TrimSpace(uri) == "" {
		return nil, nil, fmt.Errorf("inventory object URI is required")
	}
	var result ListResponse
	resp, err := s.client.Get(ctx, Endpoint, map[string]string{"uri": uri}, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) Update(ctx context.Context, r *UpdateRequest) (*interfaces.Response, error) {
	if err := validateUpdate(r); err != nil {
		return nil, err
	}
	return s.client.Put(ctx, Endpoint, r, map[string]string{"Content-Type": "application/json"}, nil)
}

// ValidateCSV uses dryRun=true and does not update values.
func (s *Service) ValidateCSV(ctx context.Context, name string, reader io.Reader, size int64) ([]string, *interfaces.Response, error) {
	return s.upload(ctx, name, reader, size, true)
}

// ImportCSV is asynchronous; HTTP 202 acknowledges acceptance, not completed enrichment.
func (s *Service) ImportCSV(ctx context.Context, name string, reader io.Reader, size int64) ([]string, *interfaces.Response, error) {
	return s.upload(ctx, name, reader, size, false)
}
func (s *Service) upload(ctx context.Context, name string, reader io.Reader, size int64, dryRun bool) ([]string, *interfaces.Response, error) {
	if err := validateUpload(name, reader, size); err != nil {
		return nil, nil, err
	}
	path := Endpoint
	if dryRun {
		path += "?dryRun=true"
	}
	var result []string
	resp, err := s.client.PostMultipart(ctx, path, "file", name, reader, size, nil, nil, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type CustomFieldValuesServiceInterface interface {
	List(context.Context, string) (*ListResponse, *interfaces.Response, error)
	Update(context.Context, *UpdateRequest) (*interfaces.Response, error)
	ValidateCSV(context.Context, string, io.Reader, int64) ([]string, *interfaces.Response, error)
	ImportCSV(context.Context, string, io.Reader, int64) ([]string, *interfaces.Response, error)
}

var _ CustomFieldValuesServiceInterface = (*Service)(nil)
