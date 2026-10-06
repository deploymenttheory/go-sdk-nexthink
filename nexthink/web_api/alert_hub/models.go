package alert_hub

import "encoding/json"

type IssueTimeFrame struct {
	FromTimestamp int64  `json:"fromTimestamp"`
	ToTimestamp   int64  `json:"toTimestamp"`
	TimeZone      string `json:"timeZone"`
	DuringPast    string `json:"duringPast,omitempty"`
}
type IssueTimeFrameForWidget struct {
	Timezone      string `json:"timezone"`
	UTCOffsetMins int    `json:"utcOffsetMins"`
	Start         string `json:"start"`
	End           string `json:"end"`
}
type FilterInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type AlertOrderByInput struct {
	Field string `json:"field"`
	Order string `json:"order"`
}
type ComplexDataFieldInput struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	DataType string `json:"dataType,omitempty"`
	DataPath string `json:"dataPath,omitempty"`
	Visible  *bool  `json:"visible,omitempty"`
}
type MonitorConfigViewInput struct {
	AlertName string `json:"alertName"`
}

type AlertImpactAssessmentRequest struct {
	AlertUUID string          `json:"alertUuid" required:"true"`
	TimeFrame *IssueTimeFrame `json:"timeFrame" required:"true"`
}

type AlertImpactAssessmentResponseAlertImpactAssessment struct {
	ImpactLevel *string `json:"impactLevel"`
	Assessment  *string `json:"assessment"`
}

type AlertImpactAssessmentResponse struct {
	AlertImpactAssessment *AlertImpactAssessmentResponseAlertImpactAssessment `json:"alertImpactAssessment"`
}

type EventsDrillDownRequest struct {
	MonitorName string                  `json:"monitorName" required:"true"`
	StartedAt   int                     `json:"startedAt" required:"true"`
	Context     []ComplexDataFieldInput `json:"context" required:"true"`
}

type EventsDrillDownResponse struct {
	EventsDrillDown *string `json:"eventsDrillDown"`
}

type IssuesRequest struct {
	FilterInput *[]FilterInput     `json:"filterInput,omitempty"`
	OrderBy     *AlertOrderByInput `json:"orderBy,omitempty"`
	TimeFrame   *IssueTimeFrame    `json:"timeFrame,omitempty"`
}

type IssuesResponseIssuesListAssociations struct {
	Count *int64  `json:"count"`
	Type  *string `json:"type"`
}

type IssuesResponseIssuesListContext struct {
	Key      *string         `json:"key"`
	Value    json.RawMessage `json:"value"`
	DataType *string         `json:"dataType"`
	DataPath *string         `json:"dataPath"`
	Visible  *bool           `json:"visible"`
}

type IssuesResponseIssuesList struct {
	LastAlertUUID       *string                                `json:"lastAlertUuid"`
	ConfigUUID          *string                                `json:"configUuid"`
	NQLID               *string                                `json:"nqlId"`
	MonitorUUID         *string                                `json:"monitorUuid"`
	Priority            *string                                `json:"priority"`
	Status              *string                                `json:"status"`
	Name                *string                                `json:"name"`
	Occurrence          *int64                                 `json:"occurrence"`
	LastTriggerDateTime *int64                                 `json:"lastTriggerDateTime"`
	ResolvedAt          *int64                                 `json:"resolvedAt"`
	Tags                *string                                `json:"tags"`
	Origin              *string                                `json:"origin"`
	Associations        []IssuesResponseIssuesListAssociations `json:"associations"`
	ImpactType          *string                                `json:"impactType"`
	Context             []IssuesResponseIssuesListContext      `json:"context"`
	CustomActions       json.RawMessage                        `json:"customActions"`
	ContextHash         *string                                `json:"contextHash"`
	MetricType          *string                                `json:"metricType"`
	DetectionType       *string                                `json:"detectionType"`
	BenchmarkType       *string                                `json:"benchmarkType"`
	BenchmarkMetricUri  *string                                `json:"benchmarkMetricUri"`
	MainMetricUri       *string                                `json:"mainMetricUri"`
}

type IssuesResponseIssues struct {
	Count         *int64                     `json:"count"`
	OpenCount     *int64                     `json:"openCount"`
	CriticalCount *int64                     `json:"criticalCount"`
	List          []IssuesResponseIssuesList `json:"list"`
}

type IssuesResponse struct {
	Issues *IssuesResponseIssues `json:"issues"`
}

type IssuesSelectedPeriodRequest struct {
	FilterInput *[]FilterInput     `json:"filterInput,omitempty"`
	OrderBy     *AlertOrderByInput `json:"orderBy,omitempty"`
	TimeFrame   *IssueTimeFrame    `json:"timeFrame,omitempty"`
}

type IssuesSelectedPeriodResponseIssuesSelectedPeriodListAssociations struct {
	Count *int64  `json:"count"`
	Type  *string `json:"type"`
}

type IssuesSelectedPeriodResponseIssuesSelectedPeriodListContext struct {
	Key      *string         `json:"key"`
	Value    json.RawMessage `json:"value"`
	DataType *string         `json:"dataType"`
	DataPath *string         `json:"dataPath"`
	Visible  *bool           `json:"visible"`
}

type IssuesSelectedPeriodResponseIssuesSelectedPeriodList struct {
	LastAlertUUID       *string                                                            `json:"lastAlertUuid"`
	ConfigUUID          *string                                                            `json:"configUuid"`
	NQLID               *string                                                            `json:"nqlId"`
	MonitorUUID         *string                                                            `json:"monitorUuid"`
	Priority            *string                                                            `json:"priority"`
	Status              *string                                                            `json:"status"`
	Name                *string                                                            `json:"name"`
	Occurrence          *int64                                                             `json:"occurrence"`
	LastTriggerDateTime *int64                                                             `json:"lastTriggerDateTime"`
	ResolvedAt          *int64                                                             `json:"resolvedAt"`
	Tags                *string                                                            `json:"tags"`
	Origin              *string                                                            `json:"origin"`
	Associations        []IssuesSelectedPeriodResponseIssuesSelectedPeriodListAssociations `json:"associations"`
	ImpactType          *string                                                            `json:"impactType"`
	Context             []IssuesSelectedPeriodResponseIssuesSelectedPeriodListContext      `json:"context"`
	CustomActions       json.RawMessage                                                    `json:"customActions"`
	ContextHash         *string                                                            `json:"contextHash"`
	MetricType          *string                                                            `json:"metricType"`
	DetectionType       *string                                                            `json:"detectionType"`
	BenchmarkType       *string                                                            `json:"benchmarkType"`
	BenchmarkMetricUri  *string                                                            `json:"benchmarkMetricUri"`
	MainMetricUri       *string                                                            `json:"mainMetricUri"`
}

type IssuesSelectedPeriodResponseIssuesSelectedPeriod struct {
	Count         *int64                                                 `json:"count"`
	OpenCount     *int64                                                 `json:"openCount"`
	CriticalCount *int64                                                 `json:"criticalCount"`
	List          []IssuesSelectedPeriodResponseIssuesSelectedPeriodList `json:"list"`
}

type IssuesSelectedPeriodResponse struct {
	IssuesSelectedPeriod *IssuesSelectedPeriodResponseIssuesSelectedPeriod `json:"issuesSelectedPeriod"`
}

type IssuesTimelineRequest struct {
	FilterInput *[]FilterInput  `json:"filterInput,omitempty"`
	TimeFrame   *IssueTimeFrame `json:"timeFrame,omitempty"`
}

type IssuesTimelineResponseIssuesTimelineIssuesDataPointsValues struct {
	OpenLow      *int64 `json:"openLow"`
	OpenMedium   *int64 `json:"openMedium"`
	OpenHigh     *int64 `json:"openHigh"`
	OpenCritical *int64 `json:"openCritical"`
	OpenTotal    *int64 `json:"openTotal"`
}

type IssuesTimelineResponseIssuesTimelineIssuesDataPoints struct {
	FromDate json.RawMessage                                             `json:"fromDate"`
	ToDate   json.RawMessage                                             `json:"toDate"`
	Values   *IssuesTimelineResponseIssuesTimelineIssuesDataPointsValues `json:"values"`
}

type IssuesTimelineResponseIssuesTimeline struct {
	IssuesDataPoints []IssuesTimelineResponseIssuesTimelineIssuesDataPoints `json:"issuesDataPoints"`
	Scope            json.RawMessage                                        `json:"scope"`
}

type IssuesTimelineResponse struct {
	IssuesTimeline *IssuesTimelineResponseIssuesTimeline `json:"issuesTimeline"`
}

type MonitorConfigViewRequest struct {
	MonitorConfigViewInput *MonitorConfigViewInput `json:"monitorConfigViewInput" required:"true"`
}

type MonitorConfigViewResponse struct {
	MonitorConfigView *string `json:"monitorConfigView"`
}

type AlertTriggerInfoRequest struct {
	AlertUUID string `json:"alertUuid" required:"true"`
}

type AlertTriggerInfoResponseAlertTriggerInfoMetrics struct {
	Label             *string         `json:"label"`
	Value             json.RawMessage `json:"value"`
	ValueDataType     *string         `json:"valueDataType"`
	Threshold         *string         `json:"threshold"`
	ThresholdDataType *string         `json:"thresholdDataType"`
}

type AlertTriggerInfoResponseAlertTriggerInfo struct {
	MonitorMetricType *string                                           `json:"monitorMetricType"`
	IssueDescription  *string                                           `json:"issueDescription"`
	Metrics           []AlertTriggerInfoResponseAlertTriggerInfoMetrics `json:"metrics"`
}

type AlertTriggerInfoResponse struct {
	AlertTriggerInfo *AlertTriggerInfoResponseAlertTriggerInfo `json:"alertTriggerInfo"`
}

type AlertsImpactedAssociationOverTimeRequest struct {
	NQLID       string          `json:"nqlId" required:"true"`
	ContextHash string          `json:"contextHash" required:"true"`
	TimeFrame   *IssueTimeFrame `json:"timeFrame" required:"true"`
}

type AlertsImpactedAssociationOverTimeResponseAlertsImpactedAssociationOverTimeDataPoints struct {
	DateTime json.RawMessage `json:"dateTime"`
	Open     *int64          `json:"open"`
	Devices  *int64          `json:"devices"`
	Users    *int64          `json:"users"`
}

type AlertsImpactedAssociationOverTimeResponseAlertsImpactedAssociationOverTime struct {
	DataPoints []AlertsImpactedAssociationOverTimeResponseAlertsImpactedAssociationOverTimeDataPoints `json:"dataPoints"`
}

type AlertsImpactedAssociationOverTimeResponse struct {
	AlertsImpactedAssociationOverTime *AlertsImpactedAssociationOverTimeResponseAlertsImpactedAssociationOverTime `json:"alertsImpactedAssociationOverTime"`
}

type GetAlertOnChangeTimeSeriesRequest struct {
	MonitorUUID  string                   `json:"monitorUuid" required:"true"`
	ContextHash  string                   `json:"contextHash" required:"true"`
	AlertContext *string                  `json:"alertContext,omitempty"`
	TimeFrame    *IssueTimeFrameForWidget `json:"timeFrame" required:"true"`
	MonitorType  *string                  `json:"monitorType,omitempty"`
}

type GetAlertOnChangeTimeSeriesResponseTimeSeriesDataMetricDataPoints struct {
	Time  json.RawMessage `json:"time"`
	Value json.RawMessage `json:"value"`
}

type GetAlertOnChangeTimeSeriesResponseTimeSeriesDataMetric struct {
	DataPoints []GetAlertOnChangeTimeSeriesResponseTimeSeriesDataMetricDataPoints `json:"dataPoints"`
}

type GetAlertOnChangeTimeSeriesResponseTimeSeriesDataBaselineDataPoints struct {
	Time  json.RawMessage `json:"time"`
	Value json.RawMessage `json:"value"`
}

type GetAlertOnChangeTimeSeriesResponseTimeSeriesDataBaseline struct {
	DataPoints []GetAlertOnChangeTimeSeriesResponseTimeSeriesDataBaselineDataPoints `json:"dataPoints"`
}

type GetAlertOnChangeTimeSeriesResponseTimeSeriesDataThresholdDataPoints struct {
	Time  json.RawMessage `json:"time"`
	Value json.RawMessage `json:"value"`
}

type GetAlertOnChangeTimeSeriesResponseTimeSeriesDataThreshold struct {
	DataPoints []GetAlertOnChangeTimeSeriesResponseTimeSeriesDataThresholdDataPoints `json:"dataPoints"`
}

type GetAlertOnChangeTimeSeriesResponseTimeSeriesData struct {
	Metric    *GetAlertOnChangeTimeSeriesResponseTimeSeriesDataMetric    `json:"metric"`
	Baseline  *GetAlertOnChangeTimeSeriesResponseTimeSeriesDataBaseline  `json:"baseline"`
	Threshold *GetAlertOnChangeTimeSeriesResponseTimeSeriesDataThreshold `json:"threshold"`
}

type GetAlertOnChangeTimeSeriesResponseTimeSeriesIssuePeriodsPeriods struct {
	FromDate json.RawMessage `json:"fromDate"`
	ToDate   json.RawMessage `json:"toDate"`
}

type GetAlertOnChangeTimeSeriesResponseTimeSeriesIssuePeriods struct {
	Periods []GetAlertOnChangeTimeSeriesResponseTimeSeriesIssuePeriodsPeriods `json:"periods"`
	Count   *int64                                                            `json:"count"`
}

type GetAlertOnChangeTimeSeriesResponse struct {
	TimeSeriesData         *GetAlertOnChangeTimeSeriesResponseTimeSeriesData         `json:"timeSeriesData"`
	TimeSeriesIssuePeriods *GetAlertOnChangeTimeSeriesResponseTimeSeriesIssuePeriods `json:"timeSeriesIssuePeriods"`
}

type GetIssuePeriodsRequest struct {
	MonitorUUID string          `json:"monitorUuid" required:"true"`
	ContextHash string          `json:"contextHash" required:"true"`
	TimeFrame   *IssueTimeFrame `json:"timeFrame" required:"true"`
}

type GetIssuePeriodsResponseIssuePeriodsPeriods struct {
	FromDate json.RawMessage `json:"fromDate"`
	ToDate   json.RawMessage `json:"toDate"`
}

type GetIssuePeriodsResponseIssuePeriods struct {
	Periods []GetIssuePeriodsResponseIssuePeriodsPeriods `json:"periods"`
	Count   *int64                                       `json:"count"`
}

type GetIssuePeriodsResponse struct {
	IssuePeriods *GetIssuePeriodsResponseIssuePeriods `json:"issuePeriods"`
}

type GetStaticAlertTimeSeriesRequest struct {
	MonitorUUID string                   `json:"monitorUuid" required:"true"`
	ContextHash string                   `json:"contextHash" required:"true"`
	TimeFrame   *IssueTimeFrameForWidget `json:"timeFrame" required:"true"`
}

type GetStaticAlertTimeSeriesResponseTimeSeriesIssuePeriodsPeriods struct {
	FromDate json.RawMessage `json:"fromDate"`
	ToDate   json.RawMessage `json:"toDate"`
}

type GetStaticAlertTimeSeriesResponseTimeSeriesIssuePeriods struct {
	Periods []GetStaticAlertTimeSeriesResponseTimeSeriesIssuePeriodsPeriods `json:"periods"`
	Count   *int64                                                          `json:"count"`
}

type GetStaticAlertTimeSeriesResponse struct {
	TimeSeriesIssuePeriods *GetStaticAlertTimeSeriesResponseTimeSeriesIssuePeriods `json:"timeSeriesIssuePeriods"`
}

type TagsRequest struct {
}

type TagsResponseTags struct {
	Name  *string `json:"name"`
	Color *string `json:"color"`
}

type TagsResponse struct {
	Tags []TagsResponseTags `json:"tags"`
}
