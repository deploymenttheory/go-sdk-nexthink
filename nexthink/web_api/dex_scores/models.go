package dex_scores

import "encoding/json"

type DimensionFilter struct {
	DMURI         string   `json:"dmUri"`
	EntitiesNames []string `json:"entitiesNames"`
}
type Filter struct {
	TechnologyExperiences []string          `json:"technologyExperiences,omitempty"`
	SentimentExperiences  []string          `json:"sentimentExperiences,omitempty"`
	Dimensions            []DimensionFilter `json:"dimensions,omitempty"`
}
type Pagination struct {
	Offset int `json:"offset"`
	Size   int `json:"size"`
}
type OrderInput struct {
	Order   string `json:"order"`
	OrderBy string `json:"orderBy"`
}
type ScoreLevel string
type InvestigationKey string

const (
	ScoreLevelRootTechnologySentiment ScoreLevel       = "ROOT_TECHNOLOGY_SENTIMENT"
	ScoreLevelMetric                  ScoreLevel       = "METRIC"
	ScoreLevelApplicationAll          ScoreLevel       = "APPLICATION_ALL"
	ScoreLevelCollaborationAll        ScoreLevel       = "COLLABORATION_ALL"
	ScoreLevelEndpointAll             ScoreLevel       = "ENDPOINT_ALL"
	InvestigationEmployeesWithIssues  InvestigationKey = "EMPLOYEES_WITH_ISSUES"
	InvestigationTechnology           InvestigationKey = "TECHNOLOGY"
	InvestigationTimeLost             InvestigationKey = "TIME_LOST"
)

type GetCampaignRequest struct {
	Filter *Filter `json:"filter,omitempty"`
}

type GetCampaignResponseCampaignComplaints struct {
	Name  *string `json:"name"`
	Votes *int64  `json:"votes"`
}

type GetCampaignResponseCampaign struct {
	UUID                *string                                 `json:"uuid"`
	NQLID               *string                                 `json:"nqlId"`
	NumberOfRespondents *int64                                  `json:"numberOfRespondents"`
	Complaints          []GetCampaignResponseCampaignComplaints `json:"complaints"`
	ComplaintsNQLID     *string                                 `json:"complaintsNqlId"`
}

type GetCampaignResponse struct {
	Campaign *GetCampaignResponseCampaign `json:"campaign"`
}

type GetDeviceExperienceRequest struct {
	MetricID string  `json:"metricId" required:"true"`
	Filter   *Filter `json:"filter,omitempty"`
}

type GetDeviceExperienceResponseDeviceExperience struct {
	DevicesWithIssues *int64 `json:"devicesWithIssues"`
}

type GetDeviceExperienceResponse struct {
	DeviceExperience *GetDeviceExperienceResponseDeviceExperience `json:"deviceExperience"`
}

type GetDimensionBreakdownsRequest struct {
	Lang string `json:"lang" required:"true"`
}

type GetDimensionBreakdownsResponseDimensionBreakdowns struct {
	Key   *string         `json:"key"`
	NQLID *string         `json:"nqlId"`
	DMURI *string         `json:"dmUri"`
	Value json.RawMessage `json:"value"`
}

type GetDimensionBreakdownsResponse struct {
	DimensionBreakdowns []GetDimensionBreakdownsResponseDimensionBreakdowns `json:"dimensionBreakdowns"`
}

type GetDimensionsV2Request struct {
	DMURI      string      `json:"dmUri" required:"true"`
	Filter     *Filter     `json:"filter,omitempty"`
	Pagination *Pagination `json:"pagination" required:"true"`
	OrderInput *OrderInput `json:"orderInput" required:"true"`
}

type GetDimensionsV2ResponseDimensionsV2Dimensions struct {
	Name            *string  `json:"name"`
	TechnologyScore *float64 `json:"technologyScore"`
	SentimentScore  *float64 `json:"sentimentScore"`
}

type GetDimensionsV2ResponseDimensionsV2 struct {
	TotalSize  *int64                                          `json:"totalSize"`
	Dimensions []GetDimensionsV2ResponseDimensionsV2Dimensions `json:"dimensions"`
}

type GetDimensionsV2Response struct {
	DimensionsV2 *GetDimensionsV2ResponseDimensionsV2 `json:"dimensionsV2"`
}

type GetInvestigationUrlRequest struct {
	InvestigationKey InvestigationKey `json:"investigationKey" required:"true"`
	Filter           *Filter          `json:"filter,omitempty"`
	MetricID         *string          `json:"metricId,omitempty"`
}

type GetInvestigationUrlResponse struct {
	InvestigationURL *string `json:"investigationUrl"`
}

type GetLeavesRequest struct {
	MetricID string  `json:"metricId" required:"true"`
	Filter   *Filter `json:"filter,omitempty"`
}

type GetLeavesResponseLeavesThresholds struct {
	Experience      *string `json:"experience"`
	NumberOfDevices *int64  `json:"numberOfDevices"`
}

type GetLeavesResponseLeaves struct {
	ID                      *float64                            `json:"id"`
	Name                    *string                             `json:"name"`
	Score                   *float64                            `json:"score"`
	Improvement             *float64                            `json:"improvement"`
	TimeLost                *float64                            `json:"timeLost"`
	Thresholds              []GetLeavesResponseLeavesThresholds `json:"thresholds"`
	UUID                    *string                             `json:"uuid"`
	AppType                 *string                             `json:"appType"`
	TroubleshootURL         *string                             `json:"troubleshootUrl"`
	DevicesInvestigationURL *string                             `json:"devicesInvestigationUrl"`
	UsersInvestigationURL   *string                             `json:"usersInvestigationUrl"`
}

type GetLeavesResponse struct {
	Leaves []GetLeavesResponseLeaves `json:"leaves"`
}

type GetMetricThresholdRequest struct {
	MetricID string `json:"metricId" required:"true"`
}

type GetMetricThresholdResponseMetricThreshold struct {
	IsInverted  *bool    `json:"isInverted"`
	Unit        *string  `json:"unit"`
	Average     *float64 `json:"average"`
	Frustrating *float64 `json:"frustrating"`
}

type GetMetricThresholdResponse struct {
	MetricThreshold *GetMetricThresholdResponseMetricThreshold `json:"metricThreshold"`
}

type GetScoresRequest struct {
	ScoreLevel ScoreLevel `json:"scoreLevel" required:"true"`
	Filter     *Filter    `json:"filter,omitempty"`
}

type GetScoresResponseScoresThresholds struct {
	Experience      *string `json:"experience"`
	NumberOfDevices *int64  `json:"numberOfDevices"`
}

type GetScoresResponseScores struct {
	ID                      *float64                            `json:"id"`
	Name                    *string                             `json:"name"`
	Score                   *float64                            `json:"score"`
	NodeType                *string                             `json:"nodeType"`
	Improvement             *float64                            `json:"improvement"`
	TimeLost                *float64                            `json:"timeLost"`
	UUID                    *string                             `json:"uuid"`
	Thresholds              []GetScoresResponseScoresThresholds `json:"thresholds"`
	Level                   *string                             `json:"level"`
	AppType                 *string                             `json:"appType"`
	TroubleshootURL         *string                             `json:"troubleshootUrl"`
	DevicesInvestigationURL *string                             `json:"devicesInvestigationUrl"`
	UsersInvestigationURL   *string                             `json:"usersInvestigationUrl"`
}

type GetScoresResponse struct {
	Scores []GetScoresResponseScores `json:"scores"`
}

type GetTrendRequest struct {
	Filter *Filter `json:"filter,omitempty"`
}

type GetTrendResponseTrendDexTrend struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
	UserCount *int64          `json:"userCount"`
}

type GetTrendResponseTrendTechnologyTrend struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
	UserCount *int64          `json:"userCount"`
}

type GetTrendResponseTrendSentimentTrend struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
	UserCount *int64          `json:"userCount"`
}

type GetTrendResponseTrend struct {
	DexTrend        []GetTrendResponseTrendDexTrend        `json:"dexTrend"`
	TechnologyTrend []GetTrendResponseTrendTechnologyTrend `json:"technologyTrend"`
	SentimentTrend  []GetTrendResponseTrendSentimentTrend  `json:"sentimentTrend"`
}

type GetTrendResponse struct {
	Trend *GetTrendResponseTrend `json:"trend"`
}

type GetTrendDevicesRequest struct {
	MetricID string  `json:"metricId" required:"true"`
	Filter   *Filter `json:"filter,omitempty"`
}

type GetTrendDevicesResponseTrendDevices struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
}

type GetTrendDevicesResponse struct {
	TrendDevices []GetTrendDevicesResponseTrendDevices `json:"trendDevices"`
}

type GetTrendEmployeesWithIssuesRequest struct {
	MetricID string  `json:"metricId" required:"true"`
	Filter   *Filter `json:"filter,omitempty"`
}

type GetTrendEmployeesWithIssuesResponseTrendEmployeesWithIssues struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
	UserCount *int64          `json:"userCount"`
}

type GetTrendEmployeesWithIssuesResponse struct {
	TrendEmployeesWithIssues []GetTrendEmployeesWithIssuesResponseTrendEmployeesWithIssues `json:"trendEmployeesWithIssues"`
}

type GetTrendImprovementRequest struct {
	MetricID string  `json:"metricId" required:"true"`
	Filter   *Filter `json:"filter,omitempty"`
}

type GetTrendImprovementResponseTrendImprovement struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
	UserCount *int64          `json:"userCount"`
}

type GetTrendImprovementResponse struct {
	TrendImprovement []GetTrendImprovementResponseTrendImprovement `json:"trendImprovement"`
}

type GetTrendScoreRequest struct {
	MetricID string  `json:"metricId" required:"true"`
	Filter   *Filter `json:"filter,omitempty"`
}

type GetTrendScoreResponseTrendScore struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
	UserCount *int64          `json:"userCount"`
}

type GetTrendScoreResponse struct {
	TrendScore []GetTrendScoreResponseTrendScore `json:"trendScore"`
}

type GetTrendTimeLostRequest struct {
	MetricID string  `json:"metricId" required:"true"`
	Filter   *Filter `json:"filter,omitempty"`
}

type GetTrendTimeLostResponseTrendTimeLost struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
	UserCount *int64          `json:"userCount"`
}

type GetTrendTimeLostResponse struct {
	TrendTimeLost []GetTrendTimeLostResponseTrendTimeLost `json:"trendTimeLost"`
}

type GetTrendWithRangeRequest struct {
	Filter     *Filter `json:"filter,omitempty"`
	StartTsSec *int64  `json:"startTsSec,omitempty"`
	EndTsSec   *int64  `json:"endTsSec,omitempty"`
}

type GetTrendWithRangeResponseTrendDexTrend struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
	UserCount *int64          `json:"userCount"`
}

type GetTrendWithRangeResponseTrendTechnologyTrend struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
	UserCount *int64          `json:"userCount"`
}

type GetTrendWithRangeResponseTrendSentimentTrend struct {
	StartDate json.RawMessage `json:"startDate"`
	EndDate   json.RawMessage `json:"endDate"`
	Value     json.RawMessage `json:"value"`
	UserCount *int64          `json:"userCount"`
}

type GetTrendWithRangeResponseTrend struct {
	DexTrend        []GetTrendWithRangeResponseTrendDexTrend        `json:"dexTrend"`
	TechnologyTrend []GetTrendWithRangeResponseTrendTechnologyTrend `json:"technologyTrend"`
	SentimentTrend  []GetTrendWithRangeResponseTrendSentimentTrend  `json:"sentimentTrend"`
}

type GetTrendWithRangeResponse struct {
	Trend *GetTrendWithRangeResponseTrend `json:"trend"`
}

type GetWhatsChangedRequest struct {
	StartTsSec int64   `json:"startTsSec" required:"true"`
	EndTsSec   int64   `json:"endTsSec" required:"true"`
	Filter     *Filter `json:"filter,omitempty"`
}

type GetWhatsChangedResponseWhatsChangedFromRootScore struct {
	Date  json.RawMessage `json:"date"`
	Score *float64        `json:"score"`
}

type GetWhatsChangedResponseWhatsChangedToRootScore struct {
	Date  json.RawMessage `json:"date"`
	Score *float64        `json:"score"`
}

type GetWhatsChangedResponseWhatsChangedKeyVariations struct {
	ParentID                *string `json:"parentId"`
	ParentName              *string `json:"parentName"`
	ParentUUID              *string `json:"parentUuid"`
	MetricLevel             *string `json:"metricLevel"`
	ID                      *string `json:"id"`
	Name                    *string `json:"name"`
	Impact                  *string `json:"impact"`
	ParentAppType           *string `json:"parentAppType"`
	DiagnoseURL             *string `json:"diagnoseUrl"`
	DevicesInvestigationURL *string `json:"devicesInvestigationUrl"`
	UsersInvestigationURL   *string `json:"usersInvestigationUrl"`
	ImpactInsights          *string `json:"impactInsights"`
}

type GetWhatsChangedResponseWhatsChanged struct {
	FromRootScore *GetWhatsChangedResponseWhatsChangedFromRootScore  `json:"fromRootScore"`
	ToRootScore   *GetWhatsChangedResponseWhatsChangedToRootScore    `json:"toRootScore"`
	KeyVariations []GetWhatsChangedResponseWhatsChangedKeyVariations `json:"keyVariations"`
}

type GetWhatsChangedResponse struct {
	WhatsChanged *GetWhatsChangedResponseWhatsChanged `json:"whatsChanged"`
}
