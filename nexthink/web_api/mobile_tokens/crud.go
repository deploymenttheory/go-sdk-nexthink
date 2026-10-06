package mobile_tokens

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

func (s *Service) List(ctx context.Context) ([]Token, *interfaces.Response, error) {
	var result []Token
	resp, err := s.client.Get(ctx, Endpoint+"/list", nil, nil, &result)
	return result, resp, err
}
func (s *Service) Get(ctx context.Context, jti string) (*Token, *interfaces.Response, error) {
	if err := validateID(jti); err != nil {
		return nil, nil, err
	}
	var result Token
	resp, err := s.client.Get(ctx, Endpoint+"/get", map[string]string{"jti": jti}, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) Create(ctx context.Context, r *CreateRequest) (*Token, *interfaces.Response, error) {
	if err := validateCreate(r); err != nil {
		return nil, nil, err
	}
	var result Token
	resp, err := s.client.Post(ctx, Endpoint+"/create", r, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) Update(ctx context.Context, r *UpdateRequest) (*Token, *interfaces.Response, error) {
	if err := validateUpdate(r); err != nil {
		return nil, nil, err
	}
	var result Token
	resp, err := s.client.Post(ctx, Endpoint+"/update", r, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
func (s *Service) Delete(ctx context.Context, r *DeleteRequest) (*interfaces.Response, error) {
	if r == nil {
		return nil, fmt.Errorf("request is required")
	}
	if err := validateID(r.JTI); err != nil {
		return nil, err
	}
	if r.Revision < 0 {
		return nil, fmt.Errorf("revision must not be negative")
	}
	return s.client.Post(ctx, Endpoint+"/delete", r, nil, nil)
}

type MobileTokensServiceInterface interface {
	List(context.Context) ([]Token, *interfaces.Response, error)
	Get(context.Context, string) (*Token, *interfaces.Response, error)
	Create(context.Context, *CreateRequest) (*Token, *interfaces.Response, error)
	Update(context.Context, *UpdateRequest) (*Token, *interfaces.Response, error)
	Delete(context.Context, *DeleteRequest) (*interfaces.Response, error)
}

var _ MobileTokensServiceInterface = (*Service)(nil)
