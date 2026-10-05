package campaigns

import (
	"context"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ graphql *graphql.Service }

func NewService(c interfaces.HTTPClient) *Service { return &Service{graphql: graphql.NewService(c)} }
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
	Create(ctx context.Context, request *CreateRequest) (*CreateResponse, *interfaces.Response, error)
	Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error)
	Delete(ctx context.Context, contentID string) (*DeleteResponse, *interfaces.Response, error)
	Get(ctx context.Context, contentID string) (*GetResponse, *interfaces.Response, error)
	List(ctx context.Context, options *ListOptions) (*ListResponse, *interfaces.Response, error)
}

var _ CampaignsServiceInterface = (*Service)(nil)
