package custom_fields

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
	"net/url"

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
func (s *Service) Create(ctx context.Context, request *CreateRequest) (*CreateResponse, *interfaces.Response, error) {
	if err := ValidateCreateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[CreateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryCreate, OperationName: "CreateCustomField", Variables: map[string]any{"input": request}})
}
func (s *Service) Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error) {
	if err := ValidateUpdateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryUpdate, OperationName: "UpdateCustomField", Variables: map[string]any{"input": request}})
}
func (s *Service) Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateDeleteRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DeleteResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryDelete, OperationName: "DeleteCustomField", Variables: map[string]any{"input": request}})
}
func (s *Service) Get(ctx context.Context, docUID, fieldType string) (*GetResponse, *interfaces.Response, error) {
	if err := ValidateID(docUID); err != nil {
		return nil, nil, err
	}
	if err := ValidateType(fieldType); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGet, OperationName: "GetCustomField", Variables: map[string]any{"docUid": docUID, "type": fieldType}})
}
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type CustomFieldsServiceInterface interface {
	Import(ctx context.Context, request *ImportRequest) (*ImportedField, *interfaces.Response, error)
	Export(ctx context.Context, fieldType, id string) (*ExportDocument, *interfaces.Response, error)
	GetValidationPatterns(ctx context.Context) (*ValidationPatterns, *interfaces.Response, error)
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Create(ctx context.Context, request *CreateRequest) (*CreateResponse, *interfaces.Response, error)
	Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error)
	Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, *interfaces.Response, error)
	Get(ctx context.Context, docUID, fieldType string) (*GetResponse, *interfaces.Response, error)
}

var _ CustomFieldsServiceInterface = (*Service)(nil)

func (s *Service) GetValidationPatterns(ctx context.Context) (*ValidationPatterns, *interfaces.Response, error) {
	var result ValidationPatterns
	resp, err := s.client.Get(ctx, EndpointREST+"/regex", nil, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

func (s *Service) Export(ctx context.Context, fieldType, id string) (*ExportDocument, *interfaces.Response, error) {
	if fieldType != "MANUAL" && fieldType != "COMPUTED" {
		return nil, nil, fmt.Errorf("field type must be MANUAL or COMPUTED")
	}
	if err := validation.PathSegment(id); err != nil {
		return nil, nil, err
	}
	var result ExportDocument
	resp, err := s.client.Get(ctx, EndpointREST+"/export/"+url.PathEscape(fieldType)+"/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

func (s *Service) Import(ctx context.Context, request *ImportRequest) (*ImportedField, *interfaces.Response, error) {
	if request == nil || !json.Valid([]byte(request.ContentFile)) {
		return nil, nil, fmt.Errorf("contentFile must contain a valid JSON definition")
	}
	var result ImportedField
	resp, err := s.client.Post(ctx, EndpointREST+"/import", request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
