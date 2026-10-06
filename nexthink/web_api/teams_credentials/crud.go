package teams_credentials

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
	return s.client.PostForm(ctx, Endpoint+"/subscription-manager/check-credentials", map[string]string{"tenant_id": request.TenantID, "client_id": request.ClientID, "client_secret": request.ClientSecret, "ms_national_cloud": request.NationalCloud}, nil, nil)
}

type TeamsCredentialsServiceInterface interface {
	CheckCredentials(context.Context, *CheckCredentialsRequest) (*interfaces.Response, error)
}

var _ TeamsCredentialsServiceInterface = (*Service)(nil)
