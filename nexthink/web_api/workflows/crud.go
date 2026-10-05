package workflows

import (
	"context"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ graphql *graphql.Service }

func NewService(
	c interfaces.HTTPClient,
) *Service {
	return &Service{graphql: graphql.NewService(c)}
}

// List uses the observed Nexthink management query.
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	return graphql.ExecuteData[ListResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{Query: queryList, OperationName: "GetWorkflowsListQuery"},
	)
}

// Get uses the observed Nexthink management query.
func (s *Service) Get(
	ctx context.Context,
	uuid string,
) (*GetResponse, *interfaces.Response, error) {
	if err := ValidateUUID(uuid); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryGet,
			OperationName: "GetWorkflowQuery",
			Variables:     map[string]any{"uuid": uuid},
		},
	)
}

// Export uses the observed Nexthink management query.
func (s *Service) Export(
	ctx context.Context,
	uuid string,
) (*ExportResponse, *interfaces.Response, error) {
	if err := ValidateUUID(uuid); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ExportResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryExport,
			OperationName: "ExportWorkflowQuery",
			Variables:     map[string]any{"uuid": uuid},
		},
	)
}

type WorkflowsServiceInterface interface {
	Delete(ctx context.Context, uuid string) (*DeleteResponse, *interfaces.Response, error)
	Update(
		ctx context.Context,
		request *WorkflowInput,
	) (*UpdateResponse, *interfaces.Response, error)
	Create(
		ctx context.Context,
		request *CreateRequest,
	) (*CreateResponse, *interfaces.Response, error)
	List(ctx context.Context) (*ListResponse, *interfaces.Response, error)
	Get(ctx context.Context, uuid string) (*GetResponse, *interfaces.Response, error)
	Export(ctx context.Context, uuid string) (*ExportResponse, *interfaces.Response, error)
}

var _ WorkflowsServiceInterface = (*Service)(nil)

// Create calls the UI management mutation. Inspect partial results even when
// err is non-nil; GraphQL can return both data and errors.
func (s *Service) Create(
	ctx context.Context,
	request *CreateRequest,
) (*CreateResponse, *interfaces.Response, error) {
	if err := ValidateCreateRequest(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"WorkflowInput": request.Workflow}
	if request.BuiltinLibraryUUID != nil {
		variables["builtinLibraryUuid"] = *request.BuiltinLibraryUUID
	}
	return graphql.ExecuteData[CreateResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryCreate,
			OperationName: "CreateWorkflowMutation",
			Variables:     variables,
		},
	)
}

// Update calls the UI management mutation. Inspect partial results even when
// err is non-nil; GraphQL can return both data and errors.
func (s *Service) Update(
	ctx context.Context,
	request *WorkflowInput,
) (*UpdateResponse, *interfaces.Response, error) {
	if err := ValidateUpdateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryUpdate,
			OperationName: "UpdateWorkflowMutation",
			Variables:     map[string]any{"WorkflowInput": request},
		},
	)
}

// Delete calls the UI management mutation. Inspect partial results even when
// err is non-nil; GraphQL can return both data and errors.
func (s *Service) Delete(
	ctx context.Context,
	uuid string,
) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateUUID(uuid); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DeleteResponse](
		ctx,
		s.graphql,
		operationID,
		graphql.GraphQLRequest{
			Query:         queryDelete,
			OperationName: "DeleteWorkflowMutation",
			Variables:     map[string]any{"uuid": uuid},
		},
	)
}
