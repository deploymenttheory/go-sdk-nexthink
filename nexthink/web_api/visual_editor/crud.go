package visual_editor

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

func (s *Service) ListBCOs(ctx context.Context, uri string) ([]Collection, *interfaces.Response, error) {
	if strings.TrimSpace(uri) == "" {
		return nil, nil, fmt.Errorf("collection URI is required")
	}
	var result []Collection
	resp, err := s.client.Get(ctx, Endpoint+"/collections/bco", map[string]string{"uri": uri}, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (s *Service) GetDefaultColumns(ctx context.Context, uri string) ([]Column, *interfaces.Response, error) {
	if strings.TrimSpace(uri) == "" {
		return nil, nil, fmt.Errorf("collection URI is required")
	}
	var result []Column
	resp, err := s.client.Get(ctx, Endpoint+"/columns", map[string]string{"uri": uri}, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (s *Service) ListFilterCollections(ctx context.Context, uri string) ([]FilterCollection, *interfaces.Response, error) {
	if strings.TrimSpace(uri) == "" {
		return nil, nil, fmt.Errorf("collection URI is required")
	}
	var result []FilterCollection
	resp, err := s.client.Get(ctx, Endpoint+"/filters/collections", map[string]string{"uri": uri}, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type VisualEditorServiceInterface interface {
	ListBCOs(ctx context.Context, uri string) ([]Collection, *interfaces.Response, error)
	GetDefaultColumns(ctx context.Context, uri string) ([]Column, *interfaces.Response, error)
	ListFilterCollections(ctx context.Context, uri string) ([]FilterCollection, *interfaces.Response, error)
}

var _ VisualEditorServiceInterface = (*Service)(nil)
