package applications

import (
	"context"
	"net/url"
	"strconv"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

type ApplicationsServiceInterface interface {
	List(context.Context, *ListOptions) (*ListResponse, *interfaces.Response, error)
	Get(context.Context, string) (*Application, *interfaces.Response, error)
	Create(context.Context, *ApplicationInput) (*Application, *interfaces.Response, error)
	Update(context.Context, string, *ApplicationInput) (*Application, *interfaces.Response, error)
	Delete(context.Context, string, int) (*DeleteResponse, *interfaces.Response, error)
}

var _ ApplicationsServiceInterface = (*Service)(nil)

// List retrieves one page; use Links.Next or Total to request subsequent pages.
func (s *Service) List(ctx context.Context, options *ListOptions) (*ListResponse, *interfaces.Response, error) {
	if err := ValidateListOptions(options); err != nil {
		return nil, nil, err
	}
	page, size, sort := 0, 100, "asc"
	if options != nil {
		page = options.PageNumber
		if options.PageSize > 0 {
			size = options.PageSize
		}
		if options.SortName != "" {
			sort = options.SortName
		}
	}
	var result ListResponse
	response, err := s.client.Get(ctx, Endpoint+"/list", map[string]string{"page_number": strconv.Itoa(page), "page_size": strconv.Itoa(size), "sort_name": sort}, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Application, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result Application
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Create(ctx context.Context, request *ApplicationInput) (*Application, *interfaces.Response, error) {
	if err := ValidateInput(request, false); err != nil {
		return nil, nil, err
	}
	var result Application
	response, err := s.client.Post(ctx, Endpoint+"/", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Update(ctx context.Context, id string, request *ApplicationInput) (*Application, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	if err := ValidateInput(request, true); err != nil {
		return nil, nil, err
	}
	var result Application
	response, err := s.client.Put(ctx, Endpoint+"/"+url.PathEscape(id), request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Delete requires the latest revision to avoid deleting an application changed by another user.
func (s *Service) Delete(ctx context.Context, id string, revision int) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	if err := ValidateRevision(revision); err != nil {
		return nil, nil, err
	}
	var result DeleteResponse
	response, err := s.client.Delete(ctx, Endpoint+"/"+url.PathEscape(id)+"/"+strconv.Itoa(revision), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
