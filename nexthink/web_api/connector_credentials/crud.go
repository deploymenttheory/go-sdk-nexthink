package connector_credentials

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context) ([]Summary, *interfaces.Response, error) {
	var result []Summary
	response, err := s.client.Get(ctx, EndpointList, map[string]string{"enabled": "true", "type": "conn_cr"}, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}
func (s *Service) ListIDs(ctx context.Context) ([]string, *interfaces.Response, error) {
	var result []string
	response, err := s.client.Get(ctx, Endpoint+"/config/conn_cr/all", nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Credential, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result Credential
	response, err := s.client.Get(ctx, Endpoint+"/config/"+url.PathEscape(id), nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Create uses the UI save upsert. It can both create and overwrite a credential.
func (s *Service) Create(ctx context.Context, id string, request *CredentialInput) (*SavedCredential, *interfaces.Response, error) {
	if err := ValidateInput(id, request); err != nil {
		return nil, nil, err
	}
	var result SavedCredential
	response, err := s.client.Post(ctx, Endpoint+"/credential/"+url.PathEscape(id), request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Update uses the UI save upsert. It can both create and overwrite a credential.
func (s *Service) Update(ctx context.Context, id string, request *CredentialInput) (*SavedCredential, *interfaces.Response, error) {
	if err := ValidateInput(id, request); err != nil {
		return nil, nil, err
	}
	var result SavedCredential
	response, err := s.client.Post(ctx, Endpoint+"/credential/"+url.PathEscape(id), request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Delete follows the UI: clear configuration/secrets and disable via POST. The ID remains allocated.
func (s *Service) Delete(ctx context.Context, id string) (*SavedCredential, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	request := CredentialInput{Config: ConfigurationInput{ConnectionDetails: []ConnectionDetail{}, Mapping: []json.RawMessage{}, RunTime: "23:30", Enabled: false, ConnectorType: id}, Secret: &Secret{Entries: []SecretEntry{}}}
	var result SavedCredential
	response, err := s.client.Post(ctx, Endpoint+"/credential/"+url.PathEscape(id), &request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type ConnectorCredentialsServiceInterface interface {
	List(context.Context) ([]Summary, *interfaces.Response, error)
	ListIDs(context.Context) ([]string, *interfaces.Response, error)
	Get(context.Context, string) (*Credential, *interfaces.Response, error)
	Create(context.Context, string, *CredentialInput) (*SavedCredential, *interfaces.Response, error)
	Update(context.Context, string, *CredentialInput) (*SavedCredential, *interfaces.Response, error)
	Delete(context.Context, string) (*SavedCredential, *interfaces.Response, error)
}

var _ ConnectorCredentialsServiceInterface = (*Service)(nil)
