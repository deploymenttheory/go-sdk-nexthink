package monitors

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct {
	client  interfaces.HTTPClient
	graphql *graphql.Service
}

func NewService(c interfaces.HTTPClient) *Service {
	return &Service{client: c, graphql: graphql.NewService(c)}
}
func (s *Service) Create(ctx context.Context, request *MonitorInput) (*CreateResponse, *interfaces.Response, error) {
	if err := ValidateCreateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[CreateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryCreate, OperationName: "AddNqlMonitor", Variables: map[string]any{"monitor": request}})
}
func (s *Service) Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error) {
	if err := ValidateUpdateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryUpdate, OperationName: "UpdateNqlMonitor", Variables: map[string]any{"docUuid": request.DocUUID, "revision": request.Revision, "monitor": request.Monitor}})
}
func (s *Service) Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateDeleteRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DeleteResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryDelete, OperationName: "DeleteMonitor", Variables: map[string]any{"deleteMonitorInput": request}})
}
func (s *Service) Get(ctx context.Context, docUUID string) (*GetResponse, *interfaces.Response, error) {
	if err := ValidateID(docUUID); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGet, OperationName: "NqlMonitor", Variables: map[string]any{"docUuid": docUUID}})
}
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type MonitorsServiceInterface interface {
	SetActivity(ctx context.Context, request *SetActivityRequest) (*SetActivityResponse, *interfaces.Response, error)
	UpdateBuiltIn(ctx context.Context, request *UpdateRequest) (*UpdateBuiltInResponse, *interfaces.Response, error)
	GetLicense(ctx context.Context) (*GetLicenseResponse, *interfaces.Response, error)
	ListTags(ctx context.Context) (*ListTagsResponse, *interfaces.Response, error)
	GetMetadata(ctx context.Context, request *GetMetadataRequest) (*GetMetadataResponse, *interfaces.Response, error)
	AnalyzeQuery(ctx context.Context, request *AnalyzeQueryRequest) (*AnalyzeQueryResponse, *interfaces.Response, error)
	GetImpactQuery(ctx context.Context, request *GetImpactQueryRequest) (*GetImpactQueryResponse, *interfaces.Response, error)
	ListFilterFields(ctx context.Context, request *ListFilterFieldsRequest) (*ListFilterFieldsResponse, *interfaces.Response, error)
	Import(ctx context.Context, request *ImportRequest) (*ImportResponse, *interfaces.Response, error)
	ExportLibrary(ctx context.Context, docUUID string) (*ExportLibraryResponse, *interfaces.Response, error)
	Export(ctx context.Context, docUUID string) (*ExportResponse, *interfaces.Response, error)
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Create(ctx context.Context, request *MonitorInput) (*CreateResponse, *interfaces.Response, error)
	Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error)
	Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, *interfaces.Response, error)
	Get(ctx context.Context, docUUID string) (*GetResponse, *interfaces.Response, error)
}

var _ MonitorsServiceInterface = (*Service)(nil)

func (s *Service) Export(ctx context.Context, docUUID string) (*ExportResponse, *interfaces.Response, error) {
	if err := ValidateID(docUUID); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ExportResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryExport, OperationName: "Export", Variables: map[string]any{"docUuid": docUUID}})
}

func (s *Service) ExportLibrary(ctx context.Context, docUUID string) (*ExportLibraryResponse, *interfaces.Response, error) {
	if err := ValidateID(docUUID); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ExportLibraryResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryExportLibrary, OperationName: "ExportLibrary", Variables: map[string]any{"docUuid": docUUID}})
}

func (s *Service) Import(ctx context.Context, request *ImportRequest) (*ImportResponse, *interfaces.Response, error) {
	if request == nil || strings.TrimSpace(request.Content) == "" {
		return nil, nil, fmt.Errorf("export content is required")
	}
	return graphql.ExecuteData[ImportResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryImport, OperationName: "ImportMonitor", Variables: map[string]any{"content": request.Content}})
}

// SetActivity calls the observed ToggleActivity UI operation. It can change tenant configuration; inspect partial GraphQL results even when an error is returned.
func (s *Service) SetActivity(ctx context.Context, request *SetActivityRequest) (*SetActivityResponse, *interfaces.Response, error) {
	if err := validateManagementSetActivity(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[SetActivityResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: querySetActivity, OperationName: "ToggleActivity", Variables: variables})
}

// UpdateBuiltIn calls the observed UpdateBuiltInMonitor UI operation. It can change tenant configuration; inspect partial GraphQL results even when an error is returned.
func (s *Service) UpdateBuiltIn(ctx context.Context, request *UpdateRequest) (*UpdateBuiltInResponse, *interfaces.Response, error) {
	if err := validateManagementUpdateBuiltIn(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateBuiltInResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryUpdateBuiltIn, OperationName: "UpdateBuiltInMonitor", Variables: variables})
}

// GetLicense calls the observed License UI operation.
func (s *Service) GetLicense(ctx context.Context) (*GetLicenseResponse, *interfaces.Response, error) {

	return graphql.ExecuteData[GetLicenseResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetLicense, OperationName: "License"})
}

// ListTags calls the observed Tags UI operation.
func (s *Service) ListTags(ctx context.Context) (*ListTagsResponse, *interfaces.Response, error) {

	return graphql.ExecuteData[ListTagsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryListTags, OperationName: "Tags"})
}

// GetMetadata calls the observed MonitorMetaData UI operation.
func (s *Service) GetMetadata(ctx context.Context, request *GetMetadataRequest) (*GetMetadataResponse, *interfaces.Response, error) {
	if err := validateManagementGetMetadata(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetMetadataResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetMetadata, OperationName: "MonitorMetaData", Variables: variables})
}

// AnalyzeQuery calls the observed NqlQueryAnalysis UI operation.
func (s *Service) AnalyzeQuery(ctx context.Context, request *AnalyzeQueryRequest) (*AnalyzeQueryResponse, *interfaces.Response, error) {
	if err := validateManagementAnalyzeQuery(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[AnalyzeQueryResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryAnalyzeQuery, OperationName: "NqlQueryAnalysis", Variables: variables})
}

// GetImpactQuery calls the observed ImpactQuery UI operation.
func (s *Service) GetImpactQuery(ctx context.Context, request *GetImpactQueryRequest) (*GetImpactQueryResponse, *interfaces.Response, error) {
	if err := validateManagementGetImpactQuery(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetImpactQueryResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetImpactQuery, OperationName: "ImpactQuery", Variables: variables})
}

// ListFilterFields calls the observed FilterFields UI operation.
func (s *Service) ListFilterFields(ctx context.Context, request *ListFilterFieldsRequest) (*ListFilterFieldsResponse, *interfaces.Response, error) {
	if err := validateManagementListFilterFields(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ListFilterFieldsResponse](ctx, s.graphql, "graphql.visual_editor", graphql.GraphQLRequest{Query: queryListFilterFields, OperationName: "FilterFields", Variables: variables})
}
