package campaigns

import (
	"context"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct {
	graphql *graphql.Service
	client  interfaces.HTTPClient
}

func NewService(c interfaces.HTTPClient) *Service {
	return &Service{graphql: graphql.NewService(c), client: c}
}
func (s *Service) Create(ctx context.Context, request *CreateRequest) (*CreateResponse, *interfaces.Response, error) {
	if err := ValidateCreateRequest(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"campaign": request.Campaign}
	if len(request.Metadata) > 0 {
		variables["metadata"] = request.Metadata
	}
	return graphql.ExecuteData[CreateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryCreate, OperationName: "CreateCampaign", Variables: variables})
}
func (s *Service) Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error) {
	if err := ValidateUpdateRequest(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"campaign": request.Campaign}
	if request.RemoveQuestions != nil {
		variables["removeQuestions"] = request.RemoveQuestions
	}
	if request.AddQuestions != nil {
		variables["addQuestions"] = request.AddQuestions
	}
	if request.Owner != nil {
		variables["owner"] = *request.Owner
	}
	return graphql.ExecuteData[UpdateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryUpdate, OperationName: "EditCampaign", Variables: variables})
}
func (s *Service) Delete(ctx context.Context, contentID string) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateID(contentID); err != nil {
		return nil, nil, err
	}

	return graphql.ExecuteData[DeleteResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryDelete, OperationName: "DeleteCampaign", Variables: map[string]any{"contentId": contentID}})
}
func (s *Service) Get(ctx context.Context, contentID string) (*GetResponse, *interfaces.Response, error) {
	if err := ValidateID(contentID); err != nil {
		return nil, nil, err
	}

	return graphql.ExecuteData[GetResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGet, OperationName: "FetchCampaignDoc", Variables: map[string]any{"contentId": contentID}})
}
func (s *Service) List(ctx context.Context, options *ListOptions) (*ListResponse, *interfaces.Response, error) {
	if err := ValidateListOptions(options); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{}
	if options != nil {
		if options.PageNumber != nil {
			variables["pageNumber"] = *options.PageNumber
		}
		if options.Offset != nil {
			variables["offset"] = *options.Offset
		}
	}
	return graphql.ExecuteData[ListResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryList, OperationName: "FetchCampaignsDocuments", Variables: variables})
}

type CampaignsServiceInterface interface {
	GetBranding(ctx context.Context) (*GetBrandingResponse, *interfaces.Response, error)
	UpdateBranding(ctx context.Context, request *UpdateBrandingRequest) (*UpdateBrandingResponse, *interfaces.Response, error)
	SetStatus(ctx context.Context, request *SetStatusRequest) (*SetStatusResponse, *interfaces.Response, error)
	GetByNQLID(ctx context.Context, request *GetByNQLIDRequest) (*GetByNQLIDResponse, *interfaces.Response, error)
	GetFromLibrary(ctx context.Context, request *GetFromLibraryRequest) (*GetFromLibraryResponse, *interfaces.Response, error)
	GetWithV6(ctx context.Context, request *GetWithV6Request) (*GetWithV6Response, *interfaces.Response, error)
	Create(ctx context.Context, request *CreateRequest) (*CreateResponse, *interfaces.Response, error)
	Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error)
	Delete(ctx context.Context, contentID string) (*DeleteResponse, *interfaces.Response, error)
	Get(ctx context.Context, contentID string) (*GetResponse, *interfaces.Response, error)
	List(ctx context.Context, options *ListOptions) (*ListResponse, *interfaces.Response, error)
}

var _ CampaignsServiceInterface = (*Service)(nil)

// GetBranding calls the observed GetBranding UI operation.
func (s *Service) GetBranding(ctx context.Context) (*GetBrandingResponse, *interfaces.Response, error) {

	return graphql.ExecuteData[GetBrandingResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetBranding, OperationName: "GetBranding"})
}

// UpdateBranding calls the observed updateBranding UI operation. It can change tenant configuration; inspect partial GraphQL results even when an error is returned.
func (s *Service) UpdateBranding(ctx context.Context, request *UpdateBrandingRequest) (*UpdateBrandingResponse, *interfaces.Response, error) {
	if err := validateManagementUpdateBranding(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateBrandingResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryUpdateBranding, OperationName: "updateBranding", Variables: variables})
}

// SetStatus calls the observed ChangeCampaignStatus UI operation. It can change tenant configuration; inspect partial GraphQL results even when an error is returned.
func (s *Service) SetStatus(ctx context.Context, request *SetStatusRequest) (*SetStatusResponse, *interfaces.Response, error) {
	if err := validateManagementSetStatus(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[SetStatusResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: querySetStatus, OperationName: "ChangeCampaignStatus", Variables: variables})
}

// GetByNQLID calls the observed FetchCampaignDocByNqlId UI operation.
func (s *Service) GetByNQLID(ctx context.Context, request *GetByNQLIDRequest) (*GetByNQLIDResponse, *interfaces.Response, error) {
	if err := validateManagementGetByNQLID(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetByNQLIDResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetByNQLID, OperationName: "FetchCampaignDocByNqlId", Variables: variables})
}

// GetFromLibrary calls the observed LibraryContentByUuid UI operation.
func (s *Service) GetFromLibrary(ctx context.Context, request *GetFromLibraryRequest) (*GetFromLibraryResponse, *interfaces.Response, error) {
	if err := validateManagementGetFromLibrary(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetFromLibraryResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetFromLibrary, OperationName: "LibraryContentByUuid", Variables: variables})
}

// GetWithV6 calls the observed campaignDocWithV6 UI operation.
func (s *Service) GetWithV6(ctx context.Context, request *GetWithV6Request) (*GetWithV6Response, *interfaces.Response, error) {
	if err := validateManagementGetWithV6(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetWithV6Response](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetWithV6, OperationName: "campaignDocWithV6", Variables: variables})
}

// GetFeatures returns campaign permissions, license volume and UI feature flags.
func (s *Service) GetFeatures(ctx context.Context) (*Features, *interfaces.Response, error) {
	var result Features
	response, err := s.client.Get(ctx, EndpointFeatures, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
