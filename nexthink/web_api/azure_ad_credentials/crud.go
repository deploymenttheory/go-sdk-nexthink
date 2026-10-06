package azure_ad_credentials

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
	form := map[string]string{"tenant_id": request.TenantID, "client_id": request.ClientID, "ms_national_cloud": request.NationalCloud}
	if request.ClientSecret != "" {
		form["client_secret"] = request.ClientSecret
	}
	return s.client.PostForm(ctx, Endpoint+"/check-credentials", form, nil, nil)
}

type AzureADCredentialsServiceInterface interface {
	CheckCredentials(context.Context, *CheckCredentialsRequest) (*interfaces.Response, error)
}

var _ AzureADCredentialsServiceInterface = (*Service)(nil)
