package data_exploration

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ graphql *graphql.Service }

func NewService(c interfaces.HTTPClient) *Service { return &Service{graphql: graphql.NewService(c)} }
func contextHeaders(options *TimeContext) (map[string]string, error) {
	value := TimeContext{TimeZone: "UTC", ISODateTime: time.Now().UTC().Format(time.RFC3339), AppName: "sdk"}
	if options != nil {
		if options.TimeZone != "" {
			value.TimeZone = options.TimeZone
		}
		value.UTCOffset = options.UTCOffset
		if options.ISODateTime != "" {
			value.ISODateTime = options.ISODateTime
		}
		if options.AppName != "" {
			value.AppName = options.AppName
		}
	}
	if _, err := time.Parse(time.RFC3339, value.ISODateTime); err != nil {
		return nil, fmt.Errorf("invalid ISO date time: %w", err)
	}
	return map[string]string{"x-nxt-waas-iso-date-time": value.ISODateTime, "x-nxt-waas-timezone": value.TimeZone, "x-nxt-waas-utc-offset": strconv.Itoa(value.UTCOffset), "x-nxt-waas-app-name": value.AppName}, nil
}
func variables(request any) (map[string]any, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	var result map[string]any
	err = json.Unmarshal(b, &result)
	return result, err
}
func (s *Service) Query(ctx context.Context, request *QueryRequest, options *TimeContext) (*QueryResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[QueryResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryQuery, OperationName: "NqlQuery", Variables: vars, Headers: headers})
}
func (s *Service) Inspect(ctx context.Context, request *QueryInput, options *TimeContext) (*InspectResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[InspectResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryInspect, OperationName: "inspect", Variables: vars, Headers: headers})
}
func (s *Service) GetFilterValues(ctx context.Context, request *FilterValuesRequest, options *TimeContext) (*GetFilterValuesResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetFilterValuesResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetFilterValues, OperationName: "FilterValues", Variables: vars, Headers: headers})
}
func (s *Service) ListFields(ctx context.Context, request *FieldsRequest, options *TimeContext) (*ListFieldsResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ListFieldsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryListFields, OperationName: "Fields", Variables: vars, Headers: headers})
}
func (s *Service) ListSystemRatings(ctx context.Context, request *SystemRatingsRequest, options *TimeContext) (*ListSystemRatingsResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ListSystemRatingsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryListSystemRatings, OperationName: "SystemRatings", Variables: vars, Headers: headers})
}
func (s *Service) ListOrganisationFields(ctx context.Context, options *TimeContext) (*ListOrganisationFieldsResponse, *interfaces.Response, error) {
	var vars map[string]any
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ListOrganisationFieldsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryListOrganisationFields, OperationName: "OrganisationFields", Variables: vars, Headers: headers})
}
func (s *Service) GetMenu(ctx context.Context, request *MenuRequest, options *TimeContext) (*GetMenuResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetMenuResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetMenu, OperationName: "Menu", Variables: vars, Headers: headers})
}
func (s *Service) ListBreakdownFields(ctx context.Context, request *BreakdownFieldsRequest, options *TimeContext) (*ListBreakdownFieldsResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ListBreakdownFieldsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryListBreakdownFields, OperationName: "BreakdownFields", Variables: vars, Headers: headers})
}
func (s *Service) GetBreakdownInsights(ctx context.Context, request *BreakdownInsightsRequest, options *TimeContext) (*GetBreakdownInsightsResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetBreakdownInsightsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetBreakdownInsights, OperationName: "BreakdownInsights", Variables: vars, Headers: headers})
}
func (s *Service) ListByDurations(ctx context.Context, request *ByDurationsRequest, options *TimeContext) (*ListByDurationsResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ListByDurationsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryListByDurations, OperationName: "ByDurations", Variables: vars, Headers: headers})
}
func (s *Service) GetOrganisation(ctx context.Context, request *OrganisationRequest, options *TimeContext) (*GetOrganisationResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetOrganisationResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetOrganisation, OperationName: "Organisation", Variables: vars, Headers: headers})
}
func (s *Service) GetItemMeta(ctx context.Context, request *ItemMetaRequest, options *TimeContext) (*GetItemMetaResponse, *interfaces.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, nil, err
	}
	vars, err := variables(request)
	if err != nil {
		return nil, nil, err
	}
	headers, err := contextHeaders(options)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetItemMetaResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetItemMeta, OperationName: "ItemMeta", Variables: vars, Headers: headers})
}

type DataExplorationServiceInterface interface {
	Query(ctx context.Context, request *QueryRequest, options *TimeContext) (*QueryResponse, *interfaces.Response, error)
	Inspect(ctx context.Context, request *QueryInput, options *TimeContext) (*InspectResponse, *interfaces.Response, error)
	GetFilterValues(ctx context.Context, request *FilterValuesRequest, options *TimeContext) (*GetFilterValuesResponse, *interfaces.Response, error)
	ListFields(ctx context.Context, request *FieldsRequest, options *TimeContext) (*ListFieldsResponse, *interfaces.Response, error)
	ListSystemRatings(ctx context.Context, request *SystemRatingsRequest, options *TimeContext) (*ListSystemRatingsResponse, *interfaces.Response, error)
	ListOrganisationFields(ctx context.Context, options *TimeContext) (*ListOrganisationFieldsResponse, *interfaces.Response, error)
	GetMenu(ctx context.Context, request *MenuRequest, options *TimeContext) (*GetMenuResponse, *interfaces.Response, error)
	ListBreakdownFields(ctx context.Context, request *BreakdownFieldsRequest, options *TimeContext) (*ListBreakdownFieldsResponse, *interfaces.Response, error)
	GetBreakdownInsights(ctx context.Context, request *BreakdownInsightsRequest, options *TimeContext) (*GetBreakdownInsightsResponse, *interfaces.Response, error)
	ListByDurations(ctx context.Context, request *ByDurationsRequest, options *TimeContext) (*ListByDurationsResponse, *interfaces.Response, error)
	GetOrganisation(ctx context.Context, request *OrganisationRequest, options *TimeContext) (*GetOrganisationResponse, *interfaces.Response, error)
	GetItemMeta(ctx context.Context, request *ItemMetaRequest, options *TimeContext) (*GetItemMetaResponse, *interfaces.Response, error)
}

var _ DataExplorationServiceInterface = (*Service)(nil)
