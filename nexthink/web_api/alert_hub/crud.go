package alert_hub

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// AlertImpactAssessment executes the first-party UI operation AlertImpactAssessment.
func (s *Service) AlertImpactAssessment(ctx context.Context, request *AlertImpactAssessmentRequest) (*AlertImpactAssessmentResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[AlertImpactAssessmentResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryAlertImpactAssessment, OperationName: "AlertImpactAssessment", Variables: variables})
}

// EventsDrillDown executes the first-party UI operation EventsDrillDown.
func (s *Service) EventsDrillDown(ctx context.Context, request *EventsDrillDownRequest) (*EventsDrillDownResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[EventsDrillDownResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryEventsDrillDown, OperationName: "EventsDrillDown", Variables: variables})
}

// Issues executes the first-party UI operation Issues.
func (s *Service) Issues(ctx context.Context, request *IssuesRequest) (*IssuesResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[IssuesResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryIssues, OperationName: "Issues", Variables: variables})
}

// IssuesSelectedPeriod executes the first-party UI operation IssuesSelectedPeriod.
func (s *Service) IssuesSelectedPeriod(ctx context.Context, request *IssuesSelectedPeriodRequest) (*IssuesSelectedPeriodResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[IssuesSelectedPeriodResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryIssuesSelectedPeriod, OperationName: "IssuesSelectedPeriod", Variables: variables})
}

// IssuesTimeline executes the first-party UI operation IssuesTimeline.
func (s *Service) IssuesTimeline(ctx context.Context, request *IssuesTimelineRequest) (*IssuesTimelineResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[IssuesTimelineResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryIssuesTimeline, OperationName: "IssuesTimeline", Variables: variables})
}

// MonitorConfigView executes the first-party UI operation MonitorConfigView.
func (s *Service) MonitorConfigView(ctx context.Context, request *MonitorConfigViewRequest) (*MonitorConfigViewResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[MonitorConfigViewResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryMonitorConfigView, OperationName: "MonitorConfigView", Variables: variables})
}

// AlertTriggerInfo executes the first-party UI operation alertTriggerInfo.
func (s *Service) AlertTriggerInfo(ctx context.Context, request *AlertTriggerInfoRequest) (*AlertTriggerInfoResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[AlertTriggerInfoResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryAlertTriggerInfo, OperationName: "alertTriggerInfo", Variables: variables})
}

// AlertsImpactedAssociationOverTime executes the first-party UI operation alertsImpactedAssociationOverTime.
func (s *Service) AlertsImpactedAssociationOverTime(ctx context.Context, request *AlertsImpactedAssociationOverTimeRequest) (*AlertsImpactedAssociationOverTimeResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[AlertsImpactedAssociationOverTimeResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryAlertsImpactedAssociationOverTime, OperationName: "alertsImpactedAssociationOverTime", Variables: variables})
}

// GetAlertOnChangeTimeSeries executes the first-party UI operation getAlertOnChangeTimeSeries.
func (s *Service) GetAlertOnChangeTimeSeries(ctx context.Context, request *GetAlertOnChangeTimeSeriesRequest) (*GetAlertOnChangeTimeSeriesResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetAlertOnChangeTimeSeriesResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetAlertOnChangeTimeSeries, OperationName: "getAlertOnChangeTimeSeries", Variables: variables})
}

// GetIssuePeriods executes the first-party UI operation getIssuePeriods.
func (s *Service) GetIssuePeriods(ctx context.Context, request *GetIssuePeriodsRequest) (*GetIssuePeriodsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetIssuePeriodsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetIssuePeriods, OperationName: "getIssuePeriods", Variables: variables})
}

// GetStaticAlertTimeSeries executes the first-party UI operation getStaticAlertTimeSeries.
func (s *Service) GetStaticAlertTimeSeries(ctx context.Context, request *GetStaticAlertTimeSeriesRequest) (*GetStaticAlertTimeSeriesResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetStaticAlertTimeSeriesResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetStaticAlertTimeSeries, OperationName: "getStaticAlertTimeSeries", Variables: variables})
}

// Tags executes the first-party UI operation Tags.
func (s *Service) Tags(ctx context.Context, request *TagsRequest) (*TagsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[TagsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryTags, OperationName: "Tags", Variables: variables})
}

type AlertHubServiceInterface interface {
	AlertImpactAssessment(context.Context, *AlertImpactAssessmentRequest) (*AlertImpactAssessmentResponse, *interfaces.Response, error)
	EventsDrillDown(context.Context, *EventsDrillDownRequest) (*EventsDrillDownResponse, *interfaces.Response, error)
	Issues(context.Context, *IssuesRequest) (*IssuesResponse, *interfaces.Response, error)
	IssuesSelectedPeriod(context.Context, *IssuesSelectedPeriodRequest) (*IssuesSelectedPeriodResponse, *interfaces.Response, error)
	IssuesTimeline(context.Context, *IssuesTimelineRequest) (*IssuesTimelineResponse, *interfaces.Response, error)
	MonitorConfigView(context.Context, *MonitorConfigViewRequest) (*MonitorConfigViewResponse, *interfaces.Response, error)
	AlertTriggerInfo(context.Context, *AlertTriggerInfoRequest) (*AlertTriggerInfoResponse, *interfaces.Response, error)
	AlertsImpactedAssociationOverTime(context.Context, *AlertsImpactedAssociationOverTimeRequest) (*AlertsImpactedAssociationOverTimeResponse, *interfaces.Response, error)
	GetAlertOnChangeTimeSeries(context.Context, *GetAlertOnChangeTimeSeriesRequest) (*GetAlertOnChangeTimeSeriesResponse, *interfaces.Response, error)
	GetIssuePeriods(context.Context, *GetIssuePeriodsRequest) (*GetIssuePeriodsResponse, *interfaces.Response, error)
	GetStaticAlertTimeSeries(context.Context, *GetStaticAlertTimeSeriesRequest) (*GetStaticAlertTimeSeriesResponse, *interfaces.Response, error)
	Tags(context.Context, *TagsRequest) (*TagsResponse, *interfaces.Response, error)
}

var _ AlertHubServiceInterface = (*Service)(nil)
