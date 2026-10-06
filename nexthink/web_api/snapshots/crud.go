package snapshots

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
	"net/url"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	resp, err := s.client.Get(ctx, EndpointList, nil, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Snapshot, *interfaces.Response, error) {
	if err := validation.PathSegment(id); err != nil {
		return nil, nil, err
	}
	var result Snapshot
	resp, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) Create(ctx context.Context, r *Definition) (*Snapshot, *interfaces.Response, error) {
	if err := validateDefinition(r); err != nil {
		return nil, nil, err
	}
	var result Snapshot
	resp, err := s.client.Post(ctx, Endpoint, r, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) Update(ctx context.Context, id string, r *Definition) (*Snapshot, *interfaces.Response, error) {
	if err := validation.PathSegment(id); err != nil {
		return nil, nil, err
	}
	if err := validateDefinition(r); err != nil {
		return nil, nil, err
	}
	var result Snapshot
	resp, err := s.client.Put(ctx, Endpoint+"/"+url.PathEscape(id), r, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) Delete(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := validation.PathSegment(id); err != nil {
		return nil, err
	}
	return s.client.Delete(ctx, Endpoint+"/"+url.PathEscape(id), nil, nil, nil)
}

// Export projects the same GET response used by the UI into a portable definition.
func (s *Service) Export(ctx context.Context, id string) (*Definition, *interfaces.Response, error) {
	r, resp, err := s.Get(ctx, id)
	if err != nil {
		return nil, resp, err
	}
	return &r.Definition, resp, nil
}

// Import uses the create endpoint. The caller explicitly supplies retention approval.
func (s *Service) Import(ctx context.Context, r *Definition) (*Snapshot, *interfaces.Response, error) {
	return s.Create(ctx, r)
}

type SnapshotsServiceInterface interface {
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Get(context.Context, string) (*Snapshot, *interfaces.Response, error)
	Create(context.Context, *Definition) (*Snapshot, *interfaces.Response, error)
	Update(context.Context, string, *Definition) (*Snapshot, *interfaces.Response, error)
	Delete(context.Context, string) (*interfaces.Response, error)
	Export(context.Context, string) (*Definition, *interfaces.Response, error)
	Import(context.Context, *Definition) (*Snapshot, *interfaces.Response, error)
}

var _ SnapshotsServiceInterface = (*Service)(nil)
