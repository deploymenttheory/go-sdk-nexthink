package amplify

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
	"strconv"
	"strings"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(client interfaces.HTTPClient) *Service { return &Service{client: client} }

type AmplifyServiceInterface interface {
	CreateConfiguration(context.Context, *ConfigurationRequest) (Configuration, *interfaces.Response, error)
	UpdateConfiguration(context.Context, string, int, *ConfigurationRequest) (Configuration, *interfaces.Response, error)
	Search(ctx context.Context, request *SearchRequest) (*SearchResponse, *interfaces.Response, error)
	SearchDevices(ctx context.Context, request *SearchRequest) (*DeviceSearchResponse, *interfaces.Response, error)
	GetDeviceProperties(ctx context.Context, collectorUID string) (Properties, *interfaces.Response, error)
	GetDeviceUsers(ctx context.Context, collectorUID string) ([]Properties, *interfaces.Response, error)
	GetDevicePackages(ctx context.Context, collectorUID string) ([]Properties, *interfaces.Response, error)
	GetUserProperties(ctx context.Context, userUID string) (Properties, *interfaces.Response, error)
	GetUserDevices(ctx context.Context, userUID string) ([]Properties, *interfaces.Response, error)
	GetConfiguration(ctx context.Context) ([]Configuration, *interfaces.Response, error)
	PostInsights(context.Context, *InsightsRequest) (*interfaces.Response, error)
}

var _ AmplifyServiceInterface = (*Service)(nil)

// Search reads Amplify extension data.
func (s *Service) Search(ctx context.Context, request *SearchRequest) (*SearchResponse, *interfaces.Response, error) {
	if err := validateSearch(request); err != nil {
		return nil, nil, err
	}
	var result SearchResponse
	response, err := s.client.Post(ctx, EndpointSearch, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// SearchDevices reads Amplify extension data.
func (s *Service) SearchDevices(ctx context.Context, request *SearchRequest) (*DeviceSearchResponse, *interfaces.Response, error) {
	if err := validateSearch(request); err != nil {
		return nil, nil, err
	}
	var result DeviceSearchResponse
	response, err := s.client.Post(ctx, EndpointSearchDevices, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetDeviceProperties reads Amplify extension data.
// collectorUID must come from Search: deviceId is the Collector UID; userUid identifies a user.
func (s *Service) GetDeviceProperties(ctx context.Context, collectorUID string) (Properties, *interfaces.Response, error) {
	if err := validateID(collectorUID); err != nil {
		return nil, nil, err
	}
	var result Properties
	response, err := s.client.Get(ctx, fmt.Sprintf(EndpointGetDeviceProperties, url.PathEscape(collectorUID)), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// GetDeviceUsers reads Amplify extension data.
// collectorUID must come from Search: deviceId is the Collector UID; userUid identifies a user.
func (s *Service) GetDeviceUsers(ctx context.Context, collectorUID string) ([]Properties, *interfaces.Response, error) {
	if err := validateID(collectorUID); err != nil {
		return nil, nil, err
	}
	var result []Properties
	response, err := s.client.Get(ctx, fmt.Sprintf(EndpointGetDeviceUsers, url.PathEscape(collectorUID)), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// GetDevicePackages reads Amplify extension data.
// collectorUID must come from Search: deviceId is the Collector UID; userUid identifies a user.
func (s *Service) GetDevicePackages(ctx context.Context, collectorUID string) ([]Properties, *interfaces.Response, error) {
	if err := validateID(collectorUID); err != nil {
		return nil, nil, err
	}
	var result []Properties
	response, err := s.client.Get(ctx, fmt.Sprintf(EndpointGetDevicePackages, url.PathEscape(collectorUID)), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// GetUserProperties reads Amplify extension data.
// userUID must come from Search: deviceId is the Collector UID; userUid identifies a user.
func (s *Service) GetUserProperties(ctx context.Context, userUID string) (Properties, *interfaces.Response, error) {
	if err := validateID(userUID); err != nil {
		return nil, nil, err
	}
	var result Properties
	response, err := s.client.Get(ctx, fmt.Sprintf(EndpointGetUserProperties, url.PathEscape(userUID)), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// GetUserDevices reads Amplify extension data.
// userUID must come from Search: deviceId is the Collector UID; userUid identifies a user.
func (s *Service) GetUserDevices(ctx context.Context, userUID string) ([]Properties, *interfaces.Response, error) {
	if err := validateID(userUID); err != nil {
		return nil, nil, err
	}
	var result []Properties
	response, err := s.client.Get(ctx, fmt.Sprintf(EndpointGetUserDevices, url.PathEscape(userUID)), nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// GetConfiguration reads Amplify extension data.
func (s *Service) GetConfiguration(ctx context.Context) ([]Configuration, *interfaces.Response, error) {
	var result []Configuration
	response, err := s.client.Get(ctx, EndpointGetConfiguration, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// PostInsights records extension usage telemetry; it does not execute remote actions.
func (s *Service) PostInsights(ctx context.Context, request *InsightsRequest) (*interfaces.Response, error) {
	if request == nil || strings.TrimSpace(request.Action) == "" {
		return nil, fmt.Errorf("action is required")
	}
	return s.client.Post(ctx, EndpointPostInsights, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, nil)
}

// CreateConfiguration initializes the central Amplify configuration. Read
// GetConfiguration first; existing documents are edited with UpdateConfiguration.
func (s *Service) CreateConfiguration(ctx context.Context, request *ConfigurationRequest) (Configuration, *interfaces.Response, error) {
	if err := validateConfiguration(request); err != nil {
		return nil, nil, err
	}
	var result Configuration
	response, err := s.client.Post(ctx, EndpointCreateConfiguration, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// UpdateConfiguration replaces the entire ordered application list and reporting
// setting using the document id and revisionNumber returned by GetConfiguration.
// App creation, editing, deletion and reordering all use this operation. There is
// no separate per-application or configuration-document DELETE in the observed UI.
func (s *Service) UpdateConfiguration(ctx context.Context, id string, revisionNumber int, request *ConfigurationRequest) (Configuration, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if revisionNumber < 0 {
		return nil, nil, fmt.Errorf("revisionNumber must be non-negative")
	}
	if err := validateConfiguration(request); err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf(EndpointUpdateConfiguration, url.PathEscape(id)) + "?" + url.Values{"revisionNumber": {strconv.Itoa(revisionNumber)}}.Encode()
	var result Configuration
	response, err := s.client.Put(ctx, path, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}
