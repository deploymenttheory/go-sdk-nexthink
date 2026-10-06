package investigations

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json"}
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Investigation, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json"}
	var result Investigation
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Create(ctx context.Context, request *InvestigationInput) (*Investigation, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json"}
	var result Investigation
	response, err := s.client.Post(ctx, Endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Update(ctx context.Context, id string, request *InvestigationInput) (*Investigation, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json"}
	var result Investigation
	response, err := s.client.Put(ctx, Endpoint+"/"+url.PathEscape(id), request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Export(ctx context.Context, id string) (*ExportDocument, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json"}
	var result ExportDocument
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id)+"/export", nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Import(ctx context.Context, request *ExportDocument) (*Investigation, *interfaces.Response, error) {
	if err := ValidateImport(request); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json"}
	var result Investigation
	response, err := s.client.Post(ctx, Endpoint+"/import", request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Delete(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	return s.client.Delete(ctx, Endpoint+"/"+url.PathEscape(id), nil, nil, nil)
}

type InvestigationsServiceInterface interface {
	CreateLink(context.Context, *LinkRequest) (*LinkResponse, *interfaces.Response, error)
	GetQueryMetadata(context.Context, *QueryRequest) (*QueryMetadata, *interfaces.Response, error)
	ExecuteQuery(context.Context, *QueryRequest) (*QueryResult, *interfaces.Response, error)
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Get(context.Context, string) (*Investigation, *interfaces.Response, error)
	Create(context.Context, *InvestigationInput) (*Investigation, *interfaces.Response, error)
	Update(context.Context, string, *InvestigationInput) (*Investigation, *interfaces.Response, error)
	Delete(context.Context, string) (*interfaces.Response, error)
	Export(context.Context, string) (*ExportDocument, *interfaces.Response, error)
	Import(context.Context, *ExportDocument) (*Investigation, *interfaces.Response, error)
}

var _ InvestigationsServiceInterface = (*Service)(nil)

func (s *Service) ExecuteQuery(ctx context.Context, request *QueryRequest) (*QueryResult, *interfaces.Response, error) {
	if request == nil || strings.TrimSpace(request.Query) == "" {
		return nil, nil, fmt.Errorf("query is required")
	}
	var result QueryResult
	response, err := s.client.Post(ctx, EndpointAnalysis+"/query/execute", request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

func (s *Service) GetQueryMetadata(ctx context.Context, request *QueryRequest) (*QueryMetadata, *interfaces.Response, error) {
	if request == nil || strings.TrimSpace(request.Query) == "" {
		return nil, nil, fmt.Errorf("query is required")
	}
	var result QueryMetadata
	response, err := s.client.Post(ctx, EndpointAnalysis+"/query/meta", request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

func (s *Service) CreateLink(ctx context.Context, request *LinkRequest) (*LinkResponse, *interfaces.Response, error) {
	if request == nil || strings.TrimSpace(request.NQL) == "" {
		return nil, nil, fmt.Errorf("query is required")
	}
	var result LinkResponse
	response, err := s.client.Post(ctx, EndpointAnalysis+"/link", request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
