package product_configuration

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
	"strings"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(client interfaces.HTTPClient) *Service { return &Service{client: client} }

func validateKey(key string) error {
	if strings.TrimSpace(key) == "" || key == "." || key == ".." {
		return fmt.Errorf("configuration key is required")
	}
	return nil
}
func validateConfiguration(request InstanceConfiguration) error {
	if len(request) != 1 {
		return fmt.Errorf("exactly one configuration key/value is required")
	}
	for key, value := range request {
		if err := validateKey(key); err != nil {
			return err
		}
		if !json.Valid(value) {
			return fmt.Errorf("configuration value must be valid JSON")
		}
	}
	return nil
}

// GetInstance returns the configuration value for the supplied key.
func (s *Service) GetInstance(ctx context.Context, key string) (*InstanceResponse, *interfaces.Response, error) {
	if err := validateKey(key); err != nil {
		return nil, nil, err
	}
	var result InstanceResponse
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(key), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CreateInstance initializes a previously absent tenant configuration key.
func (s *Service) CreateInstance(ctx context.Context, request InstanceConfiguration) (*interfaces.Response, error) {
	if err := validateConfiguration(request); err != nil {
		return nil, err
	}
	return s.client.Post(ctx, Endpoint, request, map[string]string{"Content-Type": "application/json"}, nil)
}

// UpdateInstance changes an existing tenant configuration key. Sending the existing
// value may be rejected with "Update requests must actually update something".
func (s *Service) UpdateInstance(ctx context.Context, key string, request InstanceConfiguration) (*interfaces.Response, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}
	if err := validateConfiguration(request); err != nil {
		return nil, err
	}
	if _, ok := request[key]; !ok {
		return nil, fmt.Errorf("request key must match URL key")
	}
	return s.client.Put(ctx, Endpoint+"/"+url.PathEscape(key), request, map[string]string{"Content-Type": "application/json"}, nil)
}
