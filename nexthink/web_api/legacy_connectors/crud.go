package legacy_connectors

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context, options *ListOptions) ([]Summary, *interfaces.Response, error) {
	params := map[string]string{}
	if options != nil {
		if options.Type != "" {
			params["type"] = options.Type
		}
		if options.Enabled != nil {
			params["enabled"] = strconv.FormatBool(*options.Enabled)
		}
	}
	var result []Summary
	response, err := s.client.Get(ctx, Endpoint+"/configs", params, nil, &result)
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
	response, err := s.client.Get(ctx, Endpoint+"/config/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Create and Update use the same legacy POST upsert; existence is not enforced.
func (s *Service) Create(ctx context.Context, id string, request *ConfigurationInput) (*WriteResult, *interfaces.Response, error) {
	return s.save(ctx, id, request)
}
func (s *Service) Update(ctx context.Context, id string, request *ConfigurationInput) (*WriteResult, *interfaces.Response, error) {
	return s.save(ctx, id, request)
}
func (s *Service) save(ctx context.Context, id string, request *ConfigurationInput) (*WriteResult, *interfaces.Response, error) {
	if err := ValidateInput(id, request); err != nil {
		return nil, nil, err
	}
	response, err := s.client.Post(ctx, Endpoint+"/config/"+url.PathEscape(id), request, map[string]string{"Content-Type": "application/json"}, nil)
	if err != nil {
		return nil, response, err
	}
	return &WriteResult{Message: string(response.Body)}, response, nil
}
func (s *Service) Delete(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	return s.client.Delete(ctx, Endpoint+"/config/"+url.PathEscape(id), nil, nil, nil)
}
func (s *Service) SaveSecrets(ctx context.Context, id string, request *SecretRequest) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	if request == nil || len(request.Entries) == 0 {
		return nil, fmt.Errorf("secret entries are required")
	}
	return s.client.Post(ctx, Endpoint+"/secret/"+url.PathEscape(id), request, map[string]string{"Content-Type": "application/json"}, nil)
}

// HasSecrets uses the status-only presence check shipped by the Teams and Zoom UIs.
// A 404 remains an error with response metadata; no secret values are decoded.
func (s *Service) HasSecrets(ctx context.Context, id string) (bool, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return false, nil, err
	}
	response, err := s.client.Get(ctx, Endpoint+"/secret/"+url.PathEscape(id), nil, nil, nil)
	if err != nil {
		return false, response, err
	}
	return response.StatusCode == 200, response, nil
}

// UpdateSecrets partially updates the supplied entries; SaveSecrets performs the POST save.
func (s *Service) UpdateSecrets(ctx context.Context, id string, request *SecretRequest) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	if request == nil || len(request.Entries) == 0 {
		return nil, fmt.Errorf("secret entries are required")
	}
	return s.client.Patch(ctx, Endpoint+"/secret/"+url.PathEscape(id), request, nil, nil)
}

type LegacyConnectorsServiceInterface interface {
	HasSecrets(context.Context, string) (bool, *interfaces.Response, error)
	UpdateSecrets(context.Context, string, *SecretRequest) (*interfaces.Response, error)
	List(context.Context, *ListOptions) ([]Summary, *interfaces.Response, error)
	Get(context.Context, string) (*Configuration, *interfaces.Response, error)
	Create(context.Context, string, *ConfigurationInput) (*WriteResult, *interfaces.Response, error)
	Update(context.Context, string, *ConfigurationInput) (*WriteResult, *interfaces.Response, error)
	Delete(context.Context, string) (*interfaces.Response, error)
	SaveSecrets(context.Context, string, *SecretRequest) (*interfaces.Response, error)
}

var _ LegacyConnectorsServiceInterface = (*Service)(nil)
