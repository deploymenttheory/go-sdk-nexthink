package cci_insights

import "encoding/json"

type BinaryInsightAttributeInput struct {
	DMURI string `json:"dmUri"`
	Value string `json:"value"`
}
type BinaryInsightsInput struct {
	Attributes []BinaryInsightAttributeInput `json:"attributes"`
}

type GetBinaryInsightsRequest struct {
	Filters *BinaryInsightsInput `json:"filters" required:"true"`
}

type GetBinaryInsightsResponseBinaryInsightsBinaryInsightsAttributes struct {
	DMURI *string         `json:"dmUri"`
	Value json.RawMessage `json:"value"`
}

type GetBinaryInsightsResponseBinaryInsightsBinaryInsights struct {
	Attributes                 []GetBinaryInsightsResponseBinaryInsightsBinaryInsightsAttributes `json:"attributes"`
	TenantDeviceCount          *int64                                                            `json:"tenant_device_count"`
	IssueType                  *string                                                           `json:"issue_type"`
	InsightType                *string                                                           `json:"insight_type"`
	RootMetricValue            *float64                                                          `json:"root_metric_value"`
	GroupMetricValue           *float64                                                          `json:"group_metric_value"`
	TenantMetricValue          *float64                                                          `json:"tenant_metric_value"`
	RecBinaryVersion           *string                                                           `json:"rec_binary_version"`
	RecUpgrade                 *bool                                                             `json:"rec_upgrade"`
	RecTenantMetricImprovement *float64                                                          `json:"rec_tenant_metric_improvement"`
}

type GetBinaryInsightsResponseBinaryInsights struct {
	ResultCode     *string                                                 `json:"resultCode"`
	BinaryInsights []GetBinaryInsightsResponseBinaryInsightsBinaryInsights `json:"binaryInsights"`
}

type GetBinaryInsightsResponse struct {
	BinaryInsights *GetBinaryInsightsResponseBinaryInsights `json:"binaryInsights"`
}

type GetDiagnosticBinaryInsightsRequest struct {
	Filters *BinaryInsightsInput `json:"filters" required:"true"`
}

type GetDiagnosticBinaryInsightsResponseBinaryInsightsBinaryInsightsAttributes struct {
	DMURI *string         `json:"dmUri"`
	Value json.RawMessage `json:"value"`
}

type GetDiagnosticBinaryInsightsResponseBinaryInsightsBinaryInsights struct {
	Attributes                  []GetDiagnosticBinaryInsightsResponseBinaryInsightsBinaryInsightsAttributes `json:"attributes"`
	ParameterSetID              json.RawMessage                                                             `json:"parameter_set_id"`
	GroupID                     json.RawMessage                                                             `json:"group_id"`
	Region                      *string                                                                     `json:"region"`
	TenantDeviceCount           *int64                                                                      `json:"tenant_device_count"`
	GroupTimeframeStart         json.RawMessage                                                             `json:"group_timeframe_start"`
	GroupTimeframeDurationDays  *int64                                                                      `json:"group_timeframe_duration_days"`
	TenantTimeframeStart        json.RawMessage                                                             `json:"tenant_timeframe_start"`
	TenantTimeframeDurationDays *int64                                                                      `json:"tenant_timeframe_duration_days"`
	IssueType                   *string                                                                     `json:"issue_type"`
	InsightType                 *string                                                                     `json:"insight_type"`
	RootMetricValue             *float64                                                                    `json:"root_metric_value"`
	GroupMetricValue            *float64                                                                    `json:"group_metric_value"`
	TenantMetricValue           *float64                                                                    `json:"tenant_metric_value"`
	RecBinaryVersion            *string                                                                     `json:"rec_binary_version"`
	RecUpgrade                  *bool                                                                       `json:"rec_upgrade"`
	RecTenantMetricImprovement  *float64                                                                    `json:"rec_tenant_metric_improvement"`
}

type GetDiagnosticBinaryInsightsResponseBinaryInsights struct {
	ResultCode     *string                                                           `json:"resultCode"`
	BinaryInsights []GetDiagnosticBinaryInsightsResponseBinaryInsightsBinaryInsights `json:"binaryInsights"`
}

type GetDiagnosticBinaryInsightsResponse struct {
	BinaryInsights *GetDiagnosticBinaryInsightsResponseBinaryInsights `json:"binaryInsights"`
}

type GetDatasetRequest struct {
	Name string `json:"name" required:"true"`
}

type GetDatasetResponseDatasetSchema struct {
	Name *string `json:"name"`
	Type *string `json:"type"`
}

type GetDatasetResponseDataset struct {
	ResultCode *string                           `json:"resultCode"`
	Schema     []GetDatasetResponseDatasetSchema `json:"schema"`
	Rows       json.RawMessage                   `json:"rows"`
}

type GetDatasetResponse struct {
	Dataset *GetDatasetResponseDataset `json:"dataset"`
}
