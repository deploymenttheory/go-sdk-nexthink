package nql_queries

import (
	"context"
	"net/url"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type NQLQueriesServiceInterface interface {
	Get(context.Context, string) (*SavedQuery, *interfaces.Response, error)
	Create(context.Context, *SaveQueryRequest) (*SavedQuery, *interfaces.Response, error)
	Update(context.Context, *SaveQueryRequest) (*SavedQuery, *interfaces.Response, error)
	Delete(context.Context, string) (*interfaces.Response, error)
}

var _ NQLQueriesServiceInterface = (*Service)(nil)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

func (s *Service) Get(
	ctx context.Context,
	contentID string,
) (*SavedQuery, *interfaces.Response, error) {
	if err := ValidateContentID(contentID); err != nil {
		return nil, nil, err
	}
	var result SavedQuery
	resp, err := s.client.Get(
		ctx,
		EndpointQueries+"/"+url.PathEscape(contentID),
		nil,
		map[string]string{"Accept": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

func (s *Service) Create(
	ctx context.Context,
	req *SaveQueryRequest,
) (*SavedQuery, *interfaces.Response, error) {
	if err := ValidateSaveQueryRequest(req, false); err != nil {
		return nil, nil, err
	}
	var result SavedQuery
	resp, err := s.client.Post(
		ctx,
		EndpointQueries,
		req,
		map[string]string{"Accept": "application/json", "Content-Type": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

func (s *Service) Update(
	ctx context.Context,
	req *SaveQueryRequest,
) (*SavedQuery, *interfaces.Response, error) {
	if err := ValidateSaveQueryRequest(req, true); err != nil {
		return nil, nil, err
	}
	var result SavedQuery
	resp, err := s.client.Put(
		ctx,
		EndpointQueries+"/"+url.PathEscape(req.ContentID),
		req,
		map[string]string{"Accept": "application/json", "Content-Type": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

func (s *Service) Delete(ctx context.Context, contentID string) (*interfaces.Response, error) {
	if err := ValidateContentID(contentID); err != nil {
		return nil, err
	}
	return s.client.Delete(
		ctx,
		EndpointQueries+"/"+url.PathEscape(contentID),
		nil,
		map[string]string{"Accept": "application/json"},
		nil,
	)
}
