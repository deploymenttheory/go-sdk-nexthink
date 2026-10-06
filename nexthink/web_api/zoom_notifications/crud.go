package zoom_notifications

import (
	"context"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// CheckCredentials performs the form-encoded UI check. Inspect response.StatusCode;
// unsuccessful HTTP statuses retain response metadata and return an error.
func (s *Service) CheckCredentials(ctx context.Context, request *CheckCredentialsRequest) (*interfaces.Response, error) {
	if err := ValidateCheckCredentials(request); err != nil {
		return nil, err
	}
	return s.client.PostForm(ctx, Endpoint+"/check-credentials", map[string]string{"jwt": request.JWT}, nil, nil)
}
func (s *Service) GetAppInfo(ctx context.Context) (*AppInfo, *interfaces.Response, error) {
	var result AppInfo
	response, err := s.client.Get(ctx, Endpoint+"/app-info", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type ZoomNotificationsServiceInterface interface {
	CheckCredentials(context.Context, *CheckCredentialsRequest) (*interfaces.Response, error)
	GetAppInfo(context.Context) (*AppInfo, *interfaces.Response, error)
}

var _ ZoomNotificationsServiceInterface = (*Service)(nil)
