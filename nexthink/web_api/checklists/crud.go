package checklists

import (
	"context"
	"net/url"
	"strconv"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Checklist, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result Checklist
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Create(ctx context.Context, request *ChecklistInput, options *CreateOptions) (*Checklist, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	params := map[string]string{}
	if options != nil && options.LibraryUUID != "" {
		if err := ValidateID(options.LibraryUUID); err != nil {
			return nil, nil, err
		}
		params["libraryUUID"] = options.LibraryUUID
	}
	var result Checklist
	response, err := s.client.PostWithQuery(ctx, Endpoint, params, request, map[string]string{"Content-Type": "application/json", "Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Update(ctx context.Context, id string, revision int, request *ChecklistInput) (*Checklist, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	if err := ValidateRevision(revision); err != nil {
		return nil, nil, err
	}
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	var result Checklist
	response, err := s.client.Put(ctx, Endpoint+"/"+url.PathEscape(id)+"?revisionNumber="+strconv.Itoa(revision), request, map[string]string{"Content-Type": "application/json", "Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Delete preserves the UI's nonempty DELETE body. The successful response is not JSON.
func (s *Service) Delete(ctx context.Context, id string, revision int) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	if err := ValidateRevision(revision); err != nil {
		return nil, err
	}
	return s.client.DeleteWithBody(ctx, Endpoint+"/"+url.PathEscape(id)+"?revisionNumber="+strconv.Itoa(revision), map[string]string{"a": "fix"}, map[string]string{"Content-Type": "application/json"}, nil)
}

type ChecklistsServiceInterface interface {
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Get(context.Context, string) (*Checklist, *interfaces.Response, error)
	Create(context.Context, *ChecklistInput, *CreateOptions) (*Checklist, *interfaces.Response, error)
	Update(context.Context, string, int, *ChecklistInput) (*Checklist, *interfaces.Response, error)
	Delete(context.Context, string, int) (*interfaces.Response, error)
}

var _ ChecklistsServiceInterface = (*Service)(nil)
