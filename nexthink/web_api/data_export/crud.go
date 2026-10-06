package data_export

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

func (s *Service) Start(ctx context.Context, request *StartRequest) (*StartResponse, *interfaces.Response, error) {
	if request == nil || strings.TrimSpace(request.Query) == "" || strings.TrimSpace(request.Config.FileName) == "" {
		return nil, nil, fmt.Errorf("query and export file name are required")
	}
	var result StartResponse
	resp, err := s.client.Post(ctx, Endpoint+"/export", request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

func (s *Service) GetStatus(ctx context.Context, id string) (*Status, *interfaces.Response, error) {
	if err := validation.PathSegment(id); err != nil {
		return nil, nil, err
	}
	var result Status
	resp, err := s.client.Get(ctx, Endpoint+"/status/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

type DataExportServiceInterface interface {
	Start(ctx context.Context, request *StartRequest) (*StartResponse, *interfaces.Response, error)
	GetStatus(ctx context.Context, id string) (*Status, *interfaces.Response, error)
}

var _ DataExportServiceInterface = (*Service)(nil)
