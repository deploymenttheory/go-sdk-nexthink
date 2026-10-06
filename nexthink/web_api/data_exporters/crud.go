package data_exporters

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context, options *ListOptions) ([]Configuration, *interfaces.Response, error) {
	query := map[string]string{"$deleted": "false"}
	if options != nil && options.Deleted != nil {
		query["$deleted"] = strconv.FormatBool(*options.Deleted)
	}
	var result []Configuration
	response, err := s.client.Get(ctx, Endpoint+"/all", query, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Configuration, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result Configuration
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Create(ctx context.Context, request *ConfigurationInput) (*WriteResult, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	response, err := s.client.Post(ctx, Endpoint, request, map[string]string{"Content-Type": "application/json"}, nil)
	if err != nil {
		return nil, response, err
	}
	return &WriteResult{Message: string(response.Body)}, response, nil
}
func (s *Service) Update(ctx context.Context, request *ConfigurationInput) (*WriteResult, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	response, err := s.client.Post(ctx, Endpoint, request, map[string]string{"Content-Type": "application/json"}, nil)
	if err != nil {
		return nil, response, err
	}
	return &WriteResult{Message: string(response.Body)}, response, nil
}
func (s *Service) Delete(ctx context.Context, id string) (*WriteResult, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	response, err := s.client.DeleteWithBody(ctx, Endpoint+"/"+url.PathEscape(id), struct{}{}, map[string]string{"Content-Type": "application/json"}, nil)
	if err != nil {
		return nil, response, err
	}
	return &WriteResult{Message: string(response.Body)}, response, nil
}
func (s *Service) GetCustomerInfo(ctx context.Context) (*CustomerInfo, *interfaces.Response, error) {
	var result CustomerInfo
	response, err := s.client.Get(ctx, Endpoint+"/customer-info", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) ListStatuses(ctx context.Context) (*StatusList, *interfaces.Response, error) {
	var result StatusList
	response, err := s.client.Get(ctx, Endpoint+"/status/all", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) GetPlaceholders(ctx context.Context, nql string) (*Placeholders, *interfaces.Response, error) {
	if strings.TrimSpace(nql) == "" {
		return nil, nil, fmt.Errorf("NQL is required")
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(nql))
	var result Placeholders
	response, err := s.client.Get(ctx, Endpoint+"/queries/placeholders/"+url.PathEscape(encoded), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) StartTest(ctx context.Context, request *TestRequest) (*TestExecution, *interfaces.Response, error) {
	if err := ValidateTest(request); err != nil {
		return nil, nil, err
	}
	var result TestExecution
	response, err := s.client.Post(ctx, Endpoint+"/send-test", request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) GetTest(ctx context.Context, id, executionID string) (*ExecutionStatus, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	if err := ValidateID(executionID); err != nil {
		return nil, nil, err
	}
	var result ExecutionStatus
	response, err := s.client.Get(ctx, EndpointExecutionStatus+"/"+url.PathEscape(id)+"/"+url.PathEscape(executionID), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type DataExportersServiceInterface interface {
	List(ctx context.Context, options *ListOptions) ([]Configuration, *interfaces.Response, error)
	Get(ctx context.Context, id string) (*Configuration, *interfaces.Response, error)
	Create(ctx context.Context, request *ConfigurationInput) (*WriteResult, *interfaces.Response, error)
	Update(ctx context.Context, request *ConfigurationInput) (*WriteResult, *interfaces.Response, error)
	Delete(ctx context.Context, id string) (*WriteResult, *interfaces.Response, error)
	GetCustomerInfo(ctx context.Context) (*CustomerInfo, *interfaces.Response, error)
	ListStatuses(ctx context.Context) (*StatusList, *interfaces.Response, error)
	GetPlaceholders(ctx context.Context, nql string) (*Placeholders, *interfaces.Response, error)
	StartTest(ctx context.Context, request *TestRequest) (*TestExecution, *interfaces.Response, error)
	GetTest(ctx context.Context, id, executionID string) (*ExecutionStatus, *interfaces.Response, error)
}

var _ DataExportersServiceInterface = (*Service)(nil)
