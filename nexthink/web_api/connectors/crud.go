package connectors

import (
	"context"
	"net/url"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Connector, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result Connector
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Create saves a universal connector. The tested server forces Enabled=true, even
// when false is supplied. Choose the destination and schedule before calling.
func (s *Service) Create(ctx context.Context, request *ConnectorInput) (*Connector, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	var result Connector
	response, err := s.client.Post(ctx, Endpoint, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Update replaces the configuration and honors Enabled=false in the tested lab.
func (s *Service) Update(ctx context.Context, request *ConnectorInput) (*Connector, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	var result Connector
	response, err := s.client.Put(ctx, Endpoint+"/"+url.PathEscape(request.ContentID), request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) ListTemplates(ctx context.Context) ([]Template, *interfaces.Response, error) {
	var result []Template
	response, err := s.client.Get(ctx, Endpoint+"/templates", nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}
func (s *Service) GetTemplate(ctx context.Context, id string) (*Template, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result Template
	response, err := s.client.Get(ctx, Endpoint+"/templates/"+url.PathEscape(id), nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) ListManualCustomFields(ctx context.Context, dataModelObject string) ([]ManualCustomField, *interfaces.Response, error) {
	if err := ValidateDataModelObject(dataModelObject); err != nil {
		return nil, nil, err
	}
	var result []ManualCustomField
	response, err := s.client.Get(ctx, Endpoint+"/manualcustomfields", map[string]string{"dataModelObject": dataModelObject}, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}
func (s *Service) Delete(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	return s.client.Delete(ctx, Endpoint+"/"+url.PathEscape(id), nil, nil, nil)
}

type ConnectorsServiceInterface interface {
	StartTest(context.Context, *TestRequest) (*TestExecution, *interfaces.Response, error)
	GetTest(context.Context, string) (*TestResult, *interfaces.Response, error)
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Get(context.Context, string) (*Connector, *interfaces.Response, error)
	Create(context.Context, *ConnectorInput) (*Connector, *interfaces.Response, error)
	Update(context.Context, *ConnectorInput) (*Connector, *interfaces.Response, error)
	Delete(context.Context, string) (*interfaces.Response, error)
	ListTemplates(context.Context) ([]Template, *interfaces.Response, error)
	GetTemplate(context.Context, string) (*Template, *interfaces.Response, error)
	ListManualCustomFields(context.Context, string) ([]ManualCustomField, *interfaces.Response, error)
}

var _ ConnectorsServiceInterface = (*Service)(nil)

// StartTest starts asynchronous third-party test execution. Poll GetTest until
// COMPLETED or FAILED; the caller controls timeout, cancellation and interval.
func (s *Service) StartTest(ctx context.Context, request *TestRequest) (*TestExecution, *interfaces.Response, error) {
	if err := ValidateTest(request); err != nil {
		return nil, nil, err
	}
	var result TestExecution
	response, err := s.client.PostWithQuery(ctx, Endpoint+"/test", map[string]string{"async-test-execution": "true"}, request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) GetTest(ctx context.Context, id string) (*TestResult, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result TestResult
	response, err := s.client.Get(ctx, Endpoint+"/test/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
