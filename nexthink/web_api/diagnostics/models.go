package diagnostics

import "encoding/json"

type Timeframe struct {
	Timezone      string `json:"timezone"`
	UTCOffsetMins int    `json:"utcOffsetMins"`
	Start         string `json:"start"`
	End           string `json:"end"`
}
type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Diagnostic filters vary with context and include hierarchical and technical dimensions.
type IssueDetails struct {
	ContextID string          `json:"contextId"`
	Timeframe Timeframe       `json:"timeframe"`
	Filters   json.RawMessage `json:"filters"`
	Params    []Parameter     `json:"params"`
}
type IssueReference struct {
	IssueDefinitionID   string          `json:"issueDefinitionId,omitempty"`
	AlertID             string          `json:"alertId,omitempty"`
	IssueQuery          string          `json:"issueQuery,omitempty"`
	IssueQueryParameter *Parameter      `json:"issueQueryParameter,omitempty"`
	Timeframe           Timeframe       `json:"timeframe"`
	Filters             json.RawMessage `json:"filters"`
}
type HierarchyLevelInput struct {
	HierarchyID string `json:"hierarchyId"`
	LevelDepth  int    `json:"levelDepth"`
}
type LocationBreakdownDimension string
type TechnicalBreakdownDimension string

type BinaryInfoRequest struct {
	Details *IssueDetails `json:"details" required:"true"`
}

type BinaryInfoResponseBinaryInfo struct {
	ProductName *string         `json:"productName"`
	Company     *string         `json:"company"`
	Platforms   json.RawMessage `json:"platforms"`
	HasUI       *bool           `json:"hasUI"`
}

type BinaryInfoResponse struct {
	BinaryInfo *BinaryInfoResponseBinaryInfo `json:"binaryInfo"`
}

type GetDiagnosticContextsRequest struct {
}

type GetDiagnosticContextsResponseDiagnosticContextDefinitionsParamsValidationRules struct {
	MaxItems      *int64          `json:"maxItems"`
	Min           *int64          `json:"min"`
	Max           *int64          `json:"max"`
	AllowedValues json.RawMessage `json:"allowedValues"`
}

type GetDiagnosticContextsResponseDiagnosticContextDefinitionsParams struct {
	Key             *string                                                                         `json:"key"`
	Type            *string                                                                         `json:"type"`
	Required        *bool                                                                           `json:"required"`
	DefaultValue    json.RawMessage                                                                 `json:"defaultValue"`
	ValidationRules *GetDiagnosticContextsResponseDiagnosticContextDefinitionsParamsValidationRules `json:"validationRules"`
}

type GetDiagnosticContextsResponseDiagnosticContextDefinitionsTimeFrameConfig struct {
	DefaultDuration     *string `json:"defaultDuration"`
	MaxDurationDays     *int64  `json:"maxDurationDays"`
	MaxDurationDays15m  *int64  `json:"maxDurationDays15m"`
	MaxRetentionDays    *int64  `json:"maxRetentionDays"`
	MaxRetentionDays15m *int64  `json:"maxRetentionDays15m"`
}

type GetDiagnosticContextsResponseDiagnosticContextDefinitions struct {
	ID              *string                                                                   `json:"id"`
	ScoreIDs        []string                                                                  `json:"scoreIds"`
	Params          []GetDiagnosticContextsResponseDiagnosticContextDefinitionsParams         `json:"params"`
	TimeFrameConfig *GetDiagnosticContextsResponseDiagnosticContextDefinitionsTimeFrameConfig `json:"timeFrameConfig"`
	Layout          *string                                                                   `json:"layout"`
}

type GetDiagnosticContextsResponseDiagnosticContextLayoutsWidgetsLabels struct {
	Key   *string         `json:"key"`
	Value json.RawMessage `json:"value"`
}

type GetDiagnosticContextsResponseDiagnosticContextLayoutsWidgetsParameters struct {
	Key   *string         `json:"key"`
	Value json.RawMessage `json:"value"`
}

type GetDiagnosticContextsResponseDiagnosticContextLayoutsWidgets struct {
	Name       *string                                                                  `json:"name"`
	Labels     []GetDiagnosticContextsResponseDiagnosticContextLayoutsWidgetsLabels     `json:"labels"`
	Parameters []GetDiagnosticContextsResponseDiagnosticContextLayoutsWidgetsParameters `json:"parameters"`
}

type GetDiagnosticContextsResponseDiagnosticContextLayouts struct {
	Name    *string                                                        `json:"name"`
	Widgets []GetDiagnosticContextsResponseDiagnosticContextLayoutsWidgets `json:"widgets"`
}

type GetDiagnosticContextsResponse struct {
	DiagnosticContextDefinitions []GetDiagnosticContextsResponseDiagnosticContextDefinitions `json:"diagnosticContextDefinitions"`
	DiagnosticContextLayouts     []GetDiagnosticContextsResponseDiagnosticContextLayouts     `json:"diagnosticContextLayouts"`
}

type GetDiagnosticOverviewRequest struct {
	ID string `json:"id" required:"true"`
}

type GetDiagnosticOverviewResponseGetDiagnosticOverview struct {
	Query                        *string `json:"query"`
	IncludeBinaryRecommendations *bool   `json:"includeBinaryRecommendations"`
	HasTotalDevices              *bool   `json:"hasTotalDevices"`
}

type GetDiagnosticOverviewResponse struct {
	GetDiagnosticOverview *GetDiagnosticOverviewResponseGetDiagnosticOverview `json:"getDiagnosticOverview"`
}

type GetImpactedAndTotalObjectsRequest struct {
	Details *IssueDetails `json:"details" required:"true"`
}

type GetImpactedAndTotalObjectsResponseImpactedAssociatedObjects struct {
	DrilldownQuery *string `json:"drilldownQuery"`
	Count          *int64  `json:"count"`
}

type GetImpactedAndTotalObjectsResponse struct {
	ImpactedAssociatedObjects *GetImpactedAndTotalObjectsResponseImpactedAssociatedObjects `json:"impactedAssociatedObjects"`
	TotalAssociatedObjects    *int64                                                       `json:"totalAssociatedObjects"`
}

type GetStandaloneDashboardRequest struct {
	IssueReference *IssueReference `json:"issueReference" required:"true"`
}

type GetStandaloneDashboardResponseGetStandaloneDashboard struct {
	Widgets json.RawMessage `json:"widgets"`
}

type GetStandaloneDashboardResponse struct {
	GetStandaloneDashboard *GetStandaloneDashboardResponseGetStandaloneDashboard `json:"getStandaloneDashboard"`
}

type HierarchyBreakdownDimensionsRequest struct {
}

type HierarchyBreakdownDimensionsResponseHierarchyBreakdownDimensions struct {
	HierarchyID   *string         `json:"hierarchyId"`
	HierarchyName *string         `json:"hierarchyName"`
	LevelDepth    *int64          `json:"levelDepth"`
	LevelName     *string         `json:"levelName"`
	NodeValues    json.RawMessage `json:"nodeValues"`
}

type HierarchyBreakdownDimensionsResponse struct {
	HierarchyBreakdownDimensions []HierarchyBreakdownDimensionsResponseHierarchyBreakdownDimensions `json:"hierarchyBreakdownDimensions"`
}

type IssueTimeseriesRequest struct {
	Details *IssueDetails `json:"details" required:"true"`
}

type IssueTimeseriesResponseIssueEventsTimeSeriesItems struct {
	Label *string         `json:"label"`
	Value json.RawMessage `json:"value"`
}

type IssueTimeseriesResponseIssueEventsTimeSeries struct {
	Scope json.RawMessage                                     `json:"scope"`
	Items []IssueTimeseriesResponseIssueEventsTimeSeriesItems `json:"items"`
}

type IssueTimeseriesResponse struct {
	IssueEventsTimeSeries *IssueTimeseriesResponseIssueEventsTimeSeries `json:"issueEventsTimeSeries"`
}

type TroubleshootingInsightsRequest struct {
	Details *IssueDetails `json:"details" required:"true"`
}

type TroubleshootingInsightsResponseTroubleshootingInsightsCausesBreakdownIndex struct {
	Name   *string         `json:"name"`
	Values json.RawMessage `json:"values"`
}

type TroubleshootingInsightsResponseTroubleshootingInsightsCausesBreakdownIntCols struct {
	Name   *string         `json:"name"`
	Values json.RawMessage `json:"values"`
}

type TroubleshootingInsightsResponseTroubleshootingInsightsCausesBreakdownFloatCols struct {
	Name   *string         `json:"name"`
	Values json.RawMessage `json:"values"`
}

type TroubleshootingInsightsResponseTroubleshootingInsightsCausesBreakdown struct {
	Index     *TroubleshootingInsightsResponseTroubleshootingInsightsCausesBreakdownIndex      `json:"index"`
	IntCols   []TroubleshootingInsightsResponseTroubleshootingInsightsCausesBreakdownIntCols   `json:"intCols"`
	FloatCols []TroubleshootingInsightsResponseTroubleshootingInsightsCausesBreakdownFloatCols `json:"floatCols"`
}

type TroubleshootingInsightsResponseTroubleshootingInsightsCauses struct {
	ID              *string                                                                `json:"id"`
	Name            *string                                                                `json:"name"`
	SubContextValue *string                                                                `json:"subContextValue"`
	Confidence      *float64                                                               `json:"confidence"`
	ConfidenceLevel *string                                                                `json:"confidenceLevel"`
	Breakdown       *TroubleshootingInsightsResponseTroubleshootingInsightsCausesBreakdown `json:"breakdown"`
}

type TroubleshootingInsightsResponseTroubleshootingInsights struct {
	Causes []TroubleshootingInsightsResponseTroubleshootingInsightsCauses `json:"causes"`
}

type TroubleshootingInsightsResponse struct {
	TroubleshootingInsights *TroubleshootingInsightsResponseTroubleshootingInsights `json:"troubleshootingInsights"`
}

type GetConfigurationRequest struct {
}

type GetConfigurationResponseConfiguration struct {
	StaticNow                       json.RawMessage `json:"staticNow"`
	TenantTimezone                  *string         `json:"tenantTimezone"`
	IsMtpOnly                       *bool           `json:"isMtpOnly"`
	BinaryRecommendationMaxRequests *int64          `json:"binaryRecommendationMaxRequests"`
}

type GetConfigurationResponse struct {
	Configuration *GetConfigurationResponseConfiguration `json:"configuration"`
}

type IssueEventsAndAssociatedObjectsByHierarchyBreakdownRequest struct {
	Details        *IssueDetails        `json:"details" required:"true"`
	HierarchyLevel *HierarchyLevelInput `json:"hierarchyLevel" required:"true"`
}

type IssueEventsAndAssociatedObjectsByHierarchyBreakdownResponseIssueEventsAndAssociatedObjectsByHierarchyBreakdownItems struct {
	Label *string         `json:"label"`
	Value json.RawMessage `json:"value"`
}

type IssueEventsAndAssociatedObjectsByHierarchyBreakdownResponseIssueEventsAndAssociatedObjectsByHierarchyBreakdown struct {
	Label *string                                                                                                               `json:"label"`
	Items []IssueEventsAndAssociatedObjectsByHierarchyBreakdownResponseIssueEventsAndAssociatedObjectsByHierarchyBreakdownItems `json:"items"`
}

type IssueEventsAndAssociatedObjectsByHierarchyBreakdownResponse struct {
	IssueEventsAndAssociatedObjectsByHierarchyBreakdown *IssueEventsAndAssociatedObjectsByHierarchyBreakdownResponseIssueEventsAndAssociatedObjectsByHierarchyBreakdown `json:"issueEventsAndAssociatedObjectsByHierarchyBreakdown"`
}

type IssueEventsAndAssociatedObjectsByLocationBreakdownRequest struct {
	Details   *IssueDetails              `json:"details" required:"true"`
	Dimension LocationBreakdownDimension `json:"dimension" required:"true"`
}

type IssueEventsAndAssociatedObjectsByLocationBreakdownResponseIssueEventsAndAssociatedObjectsByLocationBreakdownItems struct {
	Label *string         `json:"label"`
	Value json.RawMessage `json:"value"`
}

type IssueEventsAndAssociatedObjectsByLocationBreakdownResponseIssueEventsAndAssociatedObjectsByLocationBreakdown struct {
	Label *string                                                                                                             `json:"label"`
	Items []IssueEventsAndAssociatedObjectsByLocationBreakdownResponseIssueEventsAndAssociatedObjectsByLocationBreakdownItems `json:"items"`
}

type IssueEventsAndAssociatedObjectsByLocationBreakdownResponse struct {
	IssueEventsAndAssociatedObjectsByLocationBreakdown []IssueEventsAndAssociatedObjectsByLocationBreakdownResponseIssueEventsAndAssociatedObjectsByLocationBreakdown `json:"issueEventsAndAssociatedObjectsByLocationBreakdown"`
}

type IssueEventsAndAssociatedObjectsByTechnicalBreakdownRequest struct {
	Details   *IssueDetails               `json:"details" required:"true"`
	Dimension TechnicalBreakdownDimension `json:"dimension" required:"true"`
}

type IssueEventsAndAssociatedObjectsByTechnicalBreakdownResponseIssueEventsAndAssociatedObjectsByTechnicalBreakdownItems struct {
	Label *string         `json:"label"`
	Value json.RawMessage `json:"value"`
}

type IssueEventsAndAssociatedObjectsByTechnicalBreakdownResponseIssueEventsAndAssociatedObjectsByTechnicalBreakdown struct {
	Label *string                                                                                                               `json:"label"`
	Items []IssueEventsAndAssociatedObjectsByTechnicalBreakdownResponseIssueEventsAndAssociatedObjectsByTechnicalBreakdownItems `json:"items"`
}

type IssueEventsAndAssociatedObjectsByTechnicalBreakdownResponse struct {
	IssueEventsAndAssociatedObjectsByTechnicalBreakdown []IssueEventsAndAssociatedObjectsByTechnicalBreakdownResponseIssueEventsAndAssociatedObjectsByTechnicalBreakdown `json:"issueEventsAndAssociatedObjectsByTechnicalBreakdown"`
}

type GetDashboardRequest struct {
	IssueReference *IssueReference `json:"issueReference" required:"true"`
}

type GetDashboardResponseGetDashboard struct {
	Widgets json.RawMessage `json:"widgets"`
}

type GetDashboardResponse struct {
	GetDashboard *GetDashboardResponseGetDashboard `json:"getDashboard"`
}
