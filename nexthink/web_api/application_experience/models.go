package application_experience

import "encoding/json"

// DesktopAppAutocompletionInput is the application GraphQL input observed in the lab schema.
type DesktopAppAutocompletionInput struct {
	FieldType string `json:"fieldType"`
	Input     string `json:"input"`
	Limit     *int   `json:"limit,omitempty"`
}

// DesktopAppBreakdownOrderingInput is the application GraphQL input observed in the lab schema.
type DesktopAppBreakdownOrderingInput struct {
	ByMetric    string `json:"byMetric"`
	BySortOrder string `json:"bySortOrder"`
}

// DesktopAppDimensionsFilterInput is the application GraphQL input observed in the lab schema.
type DesktopAppDimensionsFilterInput struct {
	OSAndVersion     []string `json:"osAndVersion,omitempty"`
	BinaryAndVersion []string `json:"binaryAndVersion,omitempty"`
	Country          []string `json:"country,omitempty"`
	State            []string `json:"state,omitempty"`
}

// DesktopAppFilterInput is the application GraphQL input observed in the lab schema.
type DesktopAppFilterInput struct {
	AppIDs []string `json:"appIds"`
}

// DesktopAppTimeframeFilterInput is the application GraphQL input observed in the lab schema.
type DesktopAppTimeframeFilterInput struct {
	LocalFrom string `json:"localFrom"`
	LocalTo   string `json:"localTo"`
	TimeZone  string `json:"timeZone"`
	LocalNow  string `json:"localNow"`
}

// NetworkAppTimeframeFilterInput is the application GraphQL input observed in the lab schema.
type NetworkAppTimeframeFilterInput struct {
	LocalFrom string `json:"localFrom"`
	LocalTo   string `json:"localTo"`
	LocalNow  string `json:"localNow"`
	TimeZone  string `json:"timeZone"`
}

// WebAppDimensionsFilterInput is the application GraphQL input observed in the lab schema.
type WebAppDimensionsFilterInput struct {
	URLs               []string `json:"urls,omitempty"`
	DeviceNames        []string `json:"deviceNames,omitempty"`
	Countries          []string `json:"countries,omitempty"`
	States             []string `json:"states,omitempty"`
	LocationTypes      []string `json:"locationTypes,omitempty"`
	Entities           []string `json:"entities,omitempty"`
	BrowserAndVersions []string `json:"browserAndVersions,omitempty"`
	OSAndVersions      []string `json:"osAndVersions,omitempty"`
	IspList            []string `json:"ispList,omitempty"`
	AdapterTypes       []string `json:"adapterTypes,omitempty"`
}

// WebAppTimeframeFilterInput is the application GraphQL input observed in the lab schema.
type WebAppTimeframeFilterInput struct {
	LocalFrom string `json:"localFrom"`
	LocalTo   string `json:"localTo"`
	LocalNow  string `json:"localNow"`
	TimeZone  string `json:"timeZone"`
}

// GetApplicationsOverviewDesktopInvestigationsRequest contains the variables for GetApplicationsOverviewDesktopInvestigations.
type GetApplicationsOverviewDesktopInvestigationsRequest struct {
	TimeFrameFilter DesktopAppTimeframeFilterInput `json:"timeFrameFilter"`
}

// GetBinarySuggestionRequest contains the variables for GetBinarySuggestion.
type GetBinarySuggestionRequest struct {
	Input DesktopAppAutocompletionInput `json:"input"`
}

// GetDeviceCentricMetricBreakdownRequest contains the variables for GetDeviceCentricMetricBreakdown.
type GetDeviceCentricMetricBreakdownRequest struct {
	DesktopAppExperienceID                    string                            `json:"desktopAppExperienceId"`
	DimensionToGroupBy                        string                            `json:"dimensionToGroupBy"`
	TimeFrameFilter                           DesktopAppTimeframeFilterInput    `json:"timeFrameFilter"`
	DimensionsFilterExcludingCurrentBreakdown DesktopAppDimensionsFilterInput   `json:"dimensionsFilterExcludingCurrentBreakdown"`
	OrderBy                                   *DesktopAppBreakdownOrderingInput `json:"orderBy,omitempty"`
	Limit                                     int                               `json:"limit"`
}

// GetFailedConnectionsRatioRequest contains the variables for GetFailedConnectionsRatio.
type GetFailedConnectionsRatioRequest struct {
	ID               string                           `json:"id"`
	TimeFrameFilter  DesktopAppTimeframeFilterInput   `json:"timeFrameFilter"`
	DimensionsFilter *DesktopAppDimensionsFilterInput `json:"dimensionsFilter,omitempty"`
}

// GetInsightsRequest contains the variables for GetInsights.
type GetInsightsRequest struct {
	ApplicationID    string                      `json:"applicationId"`
	ApplicationName  string                      `json:"applicationName"`
	DimensionsFilter WebAppDimensionsFilterInput `json:"dimensionsFilter"`
	InsightType      string                      `json:"insightType"`
	TimeFrameFilter  WebAppTimeframeFilterInput  `json:"timeFrameFilter"`
}

// GetMetricBreakdownRequest contains the variables for GetMetricBreakdown.
type GetMetricBreakdownRequest struct {
	DesktopAppExperienceID                    string                            `json:"desktopAppExperienceId"`
	DimensionToGroupBy                        string                            `json:"dimensionToGroupBy"`
	TimeFrameFilter                           DesktopAppTimeframeFilterInput    `json:"timeFrameFilter"`
	DimensionsFilterExcludingCurrentBreakdown DesktopAppDimensionsFilterInput   `json:"dimensionsFilterExcludingCurrentBreakdown"`
	OrderBy                                   *DesktopAppBreakdownOrderingInput `json:"orderBy,omitempty"`
	Limit                                     int                               `json:"limit"`
}

// GetNumOfCrashesAndDevicesRequest contains the variables for GetNumOfCrashesAndDevices.
type GetNumOfCrashesAndDevicesRequest struct {
	ID               string                           `json:"id"`
	TimeFrameFilter  DesktopAppTimeframeFilterInput   `json:"timeFrameFilter"`
	DimensionsFilter *DesktopAppDimensionsFilterInput `json:"dimensionsFilter,omitempty"`
}

// GetNumOfCrashesAndEmployeesRequest contains the variables for GetNumOfCrashesAndEmployees.
type GetNumOfCrashesAndEmployeesRequest struct {
	ID               string                           `json:"id"`
	TimeFrameFilter  DesktopAppTimeframeFilterInput   `json:"timeFrameFilter"`
	DimensionsFilter *DesktopAppDimensionsFilterInput `json:"dimensionsFilter,omitempty"`
}

// GetNumOfDevicesRequest contains the variables for GetNumOfDevices.
type GetNumOfDevicesRequest struct {
	ID               string                           `json:"id"`
	TimeFrameFilter  DesktopAppTimeframeFilterInput   `json:"timeFrameFilter"`
	DimensionsFilter *DesktopAppDimensionsFilterInput `json:"dimensionsFilter,omitempty"`
}

// GetNumOfDevicesWithCrashesRequest contains the variables for GetNumOfDevicesWithCrashes.
type GetNumOfDevicesWithCrashesRequest struct {
	ID               string                           `json:"id"`
	TimeFrameFilter  DesktopAppTimeframeFilterInput   `json:"timeFrameFilter"`
	DimensionsFilter *DesktopAppDimensionsFilterInput `json:"dimensionsFilter,omitempty"`
}

// GetNumOfEmployeesRequest contains the variables for GetNumOfEmployees.
type GetNumOfEmployeesRequest struct {
	ID               string                           `json:"id"`
	TimeFrameFilter  DesktopAppTimeframeFilterInput   `json:"timeFrameFilter"`
	DimensionsFilter *DesktopAppDimensionsFilterInput `json:"dimensionsFilter,omitempty"`
}

// GetNumOfEmployeesWithCrashesRequest contains the variables for GetNumOfEmployeesWithCrashes.
type GetNumOfEmployeesWithCrashesRequest struct {
	ID               string                           `json:"id"`
	TimeFrameFilter  DesktopAppTimeframeFilterInput   `json:"timeFrameFilter"`
	DimensionsFilter *DesktopAppDimensionsFilterInput `json:"dimensionsFilter,omitempty"`
}

// OverviewDesktopTooltipsRequest contains the variables for OverviewDesktopTooltips.
type OverviewDesktopTooltipsRequest struct {
	DesktopTimeFrameFilter DesktopAppTimeframeFilterInput `json:"desktopTimeFrameFilter"`
}

// TilesDesktopCrashesPerEmployeeRequest contains the variables for TilesDesktopCrashesPerEmployee.
type TilesDesktopCrashesPerEmployeeRequest struct {
	TimeFrameFilter DesktopAppTimeframeFilterInput `json:"timeFrameFilter"`
}

// TilesDesktopNumberOfEmployeesRequest contains the variables for TilesDesktopNumberOfEmployees.
type TilesDesktopNumberOfEmployeesRequest struct {
	TimeFrameFilter DesktopAppTimeframeFilterInput `json:"timeFrameFilter"`
}

// TilesErrorCountRequest contains the variables for TilesErrorCount.
type TilesErrorCountRequest struct {
	TimeFrameFilter WebAppTimeframeFilterInput `json:"timeFrameFilter"`
}

// TilesFrustratingPageLoadsRequest contains the variables for TilesFrustratingPageLoads.
type TilesFrustratingPageLoadsRequest struct {
	TimeFrameFilter WebAppTimeframeFilterInput `json:"timeFrameFilter"`
}

// TilesNumberOfEmployeesRequest contains the variables for TilesNumberOfEmployees.
type TilesNumberOfEmployeesRequest struct {
	TimeFrameFilter WebAppTimeframeFilterInput `json:"timeFrameFilter"`
}

// TilesPageLoadTimeRequest contains the variables for TilesPageLoadTime.
type TilesPageLoadTimeRequest struct {
	TimeFrameFilter WebAppTimeframeFilterInput `json:"timeFrameFilter"`
}

// TilesTransactionTimeRequest contains the variables for TilesTransactionTime.
type TilesTransactionTimeRequest struct {
	TimeFrameFilter WebAppTimeframeFilterInput `json:"timeFrameFilter"`
}

// TilesUsageTimeRequest contains the variables for TilesUsageTime.
type TilesUsageTimeRequest struct {
	TimeFrameFilter WebAppTimeframeFilterInput `json:"timeFrameFilter"`
}

// WebOverviewTooltipsRequest contains the variables for WebOverviewTooltips.
type WebOverviewTooltipsRequest struct {
	TimeFrameFilter WebAppTimeframeFilterInput `json:"timeFrameFilter"`
}

type GetApplicationsOverviewDesktopInvestigationsResponseDesktopAppOverviewInvestigationsNumberOfEmployeesAndDevices struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	URL  *string `json:"url"`
}

type GetApplicationsOverviewDesktopInvestigationsResponseDesktopAppOverviewInvestigationsCrashes struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	URL  *string `json:"url"`
}

type GetApplicationsOverviewDesktopInvestigationsResponseDesktopAppOverviewInvestigations struct {
	NumberOfEmployeesAndDevices *GetApplicationsOverviewDesktopInvestigationsResponseDesktopAppOverviewInvestigationsNumberOfEmployeesAndDevices `json:"numberOfEmployeesAndDevices"`
	Crashes                     *GetApplicationsOverviewDesktopInvestigationsResponseDesktopAppOverviewInvestigationsCrashes                     `json:"crashes"`
}

type GetApplicationsOverviewDesktopInvestigationsResponseDesktopAppOverview struct {
	Investigations *GetApplicationsOverviewDesktopInvestigationsResponseDesktopAppOverviewInvestigations `json:"investigations"`
}

type GetApplicationsOverviewDesktopInvestigationsResponse struct {
	DesktopAppOverview *GetApplicationsOverviewDesktopInvestigationsResponseDesktopAppOverview `json:"desktopAppOverview"`
}
type GetAvgNetworkResponseTimeResponseDesktopAppExperienceConnectivityNetworkResponseTimeValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetAvgNetworkResponseTimeResponseDesktopAppExperienceConnectivityNetworkResponseTime struct {
	AverageMS       *int64                                                                                                `json:"averageMs"`
	ValueTimeSeries []GetAvgNetworkResponseTimeResponseDesktopAppExperienceConnectivityNetworkResponseTimeValueTimeSeries `json:"valueTimeSeries"`
}

type GetAvgNetworkResponseTimeResponseDesktopAppExperienceConnectivity struct {
	NetworkResponseTime *GetAvgNetworkResponseTimeResponseDesktopAppExperienceConnectivityNetworkResponseTime `json:"networkResponseTime"`
}

type GetAvgNetworkResponseTimeResponseDesktopAppExperience struct {
	Connectivity *GetAvgNetworkResponseTimeResponseDesktopAppExperienceConnectivity `json:"connectivity"`
}

type GetAvgNetworkResponseTimeResponse struct {
	DesktopAppExperience *GetAvgNetworkResponseTimeResponseDesktopAppExperience `json:"desktopAppExperience"`
}
type GetBinarySuggestionResponseDesktopAppAutoComplete struct {
	Matches []string `json:"matches"`
}

type GetBinarySuggestionResponse struct {
	DesktopAppAutoComplete *GetBinarySuggestionResponseDesktopAppAutoComplete `json:"desktopAppAutoComplete"`
}
type GetDeviceCentricMetricBreakdownResponseDesktopAppExperienceMetricBreakdown struct {
	BreakdownKey           string   `json:"breakdownKey"`
	CountDevices           *int64   `json:"countDevices"`
	CrashesPerDevice       *float64 `json:"crashesPerDevice"`
	DevicesWithCrashes     *int64   `json:"devicesWithCrashes"`
	AvgNetResponseTimeMS   *int64   `json:"avgNetResponseTimeMs"`
	FailedConnectionsRatio *float64 `json:"failedConnectionsRatio"`
}

type GetDeviceCentricMetricBreakdownResponseDesktopAppExperience struct {
	MetricBreakdown []GetDeviceCentricMetricBreakdownResponseDesktopAppExperienceMetricBreakdown `json:"metricBreakdown"`
}

type GetDeviceCentricMetricBreakdownResponse struct {
	DesktopAppExperience *GetDeviceCentricMetricBreakdownResponseDesktopAppExperience `json:"desktopAppExperience"`
}
type GetFailedConnectionsRatioResponseDesktopAppExperienceConnectivityFailedConnectionsRatioRatioTimeSeries struct {
	NoHost             *float64 `json:"noHost"`
	NoService          *float64 `json:"noService"`
	RejectedConnection *float64 `json:"rejectedConnection"`
	Timestamp          string   `json:"timestamp"`
}

type GetFailedConnectionsRatioResponseDesktopAppExperienceConnectivityFailedConnectionsRatio struct {
	Ratio           *float64                                                                                                 `json:"ratio"`
	RatioTimeSeries []GetFailedConnectionsRatioResponseDesktopAppExperienceConnectivityFailedConnectionsRatioRatioTimeSeries `json:"ratioTimeSeries"`
}

type GetFailedConnectionsRatioResponseDesktopAppExperienceConnectivity struct {
	FailedConnectionsRatio *GetFailedConnectionsRatioResponseDesktopAppExperienceConnectivityFailedConnectionsRatio `json:"failedConnectionsRatio"`
}

type GetFailedConnectionsRatioResponseDesktopAppExperience struct {
	Connectivity *GetFailedConnectionsRatioResponseDesktopAppExperienceConnectivity `json:"connectivity"`
}

type GetFailedConnectionsRatioResponse struct {
	DesktopAppExperience *GetFailedConnectionsRatioResponseDesktopAppExperience `json:"desktopAppExperience"`
}
type GetInsightsResponseWebAppPerformanceInsightsInsightsBaselinePointsBaselineMetric struct {
	MetricName  string  `json:"metricName"`
	MetricValue float64 `json:"metricValue"`
}

type GetInsightsResponseWebAppPerformanceInsightsInsightsBaselinePoints struct {
	BaselineType   string                                                                            `json:"baselineType"`
	BaselineMetric *GetInsightsResponseWebAppPerformanceInsightsInsightsBaselinePointsBaselineMetric `json:"baselineMetric"`
}

type GetInsightsResponseWebAppPerformanceInsightsInsights struct {
	InsightType          string                                                               `json:"insightType"`
	InsightPoint         json.RawMessage                                                      `json:"insightPoint"`
	BaselinePoints       []GetInsightsResponseWebAppPerformanceInsightsInsightsBaselinePoints `json:"baselinePoints"`
	InsightAnalysisSteps []json.RawMessage                                                    `json:"insightAnalysisSteps"`
}

type GetInsightsResponseWebAppPerformanceInsights struct {
	Insights []GetInsightsResponseWebAppPerformanceInsightsInsights `json:"insights"`
}

type GetInsightsResponse struct {
	WebAppPerformanceInsights *GetInsightsResponseWebAppPerformanceInsights `json:"webAppPerformanceInsights"`
}
type GetMetricBreakdownResponseDesktopAppExperienceMetricBreakdown struct {
	BreakdownKey           string   `json:"breakdownKey"`
	CountEmployees         *int64   `json:"countEmployees"`
	CrashesPerEmployee     *float64 `json:"crashesPerEmployee"`
	EmployeesWithCrashes   *int64   `json:"employeesWithCrashes"`
	AvgNetResponseTimeMS   *int64   `json:"avgNetResponseTimeMs"`
	FailedConnectionsRatio *float64 `json:"failedConnectionsRatio"`
}

type GetMetricBreakdownResponseDesktopAppExperience struct {
	MetricBreakdown []GetMetricBreakdownResponseDesktopAppExperienceMetricBreakdown `json:"metricBreakdown"`
}

type GetMetricBreakdownResponse struct {
	DesktopAppExperience *GetMetricBreakdownResponseDesktopAppExperience `json:"desktopAppExperience"`
}
type GetNumOfCrashesAndDevicesResponseDesktopAppExperienceAdoptionNumDevicesValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetNumOfCrashesAndDevicesResponseDesktopAppExperienceAdoptionNumDevices struct {
	Total           *int64                                                                                   `json:"total"`
	ValueTimeSeries []GetNumOfCrashesAndDevicesResponseDesktopAppExperienceAdoptionNumDevicesValueTimeSeries `json:"valueTimeSeries"`
}

type GetNumOfCrashesAndDevicesResponseDesktopAppExperienceAdoption struct {
	NumDevices *GetNumOfCrashesAndDevicesResponseDesktopAppExperienceAdoptionNumDevices `json:"numDevices"`
}

type GetNumOfCrashesAndDevicesResponseDesktopAppExperienceReliabilityNumCrashesValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetNumOfCrashesAndDevicesResponseDesktopAppExperienceReliabilityNumCrashes struct {
	Total           *int64                                                                                      `json:"total"`
	ValueTimeSeries []GetNumOfCrashesAndDevicesResponseDesktopAppExperienceReliabilityNumCrashesValueTimeSeries `json:"valueTimeSeries"`
}

type GetNumOfCrashesAndDevicesResponseDesktopAppExperienceReliability struct {
	NumCrashes *GetNumOfCrashesAndDevicesResponseDesktopAppExperienceReliabilityNumCrashes `json:"numCrashes"`
}

type GetNumOfCrashesAndDevicesResponseDesktopAppExperience struct {
	Adoption    *GetNumOfCrashesAndDevicesResponseDesktopAppExperienceAdoption    `json:"adoption"`
	Reliability *GetNumOfCrashesAndDevicesResponseDesktopAppExperienceReliability `json:"reliability"`
}

type GetNumOfCrashesAndDevicesResponse struct {
	DesktopAppExperience *GetNumOfCrashesAndDevicesResponseDesktopAppExperience `json:"desktopAppExperience"`
}
type GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceAdoptionNumEmployeesValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceAdoptionNumEmployees struct {
	Total           *int64                                                                                       `json:"total"`
	ValueTimeSeries []GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceAdoptionNumEmployeesValueTimeSeries `json:"valueTimeSeries"`
}

type GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceAdoption struct {
	NumEmployees *GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceAdoptionNumEmployees `json:"numEmployees"`
}

type GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceReliabilityNumCrashesValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceReliabilityNumCrashes struct {
	Total           *int64                                                                                        `json:"total"`
	ValueTimeSeries []GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceReliabilityNumCrashesValueTimeSeries `json:"valueTimeSeries"`
}

type GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceReliability struct {
	NumCrashes *GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceReliabilityNumCrashes `json:"numCrashes"`
}

type GetNumOfCrashesAndEmployeesResponseDesktopAppExperience struct {
	Adoption    *GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceAdoption    `json:"adoption"`
	Reliability *GetNumOfCrashesAndEmployeesResponseDesktopAppExperienceReliability `json:"reliability"`
}

type GetNumOfCrashesAndEmployeesResponse struct {
	DesktopAppExperience *GetNumOfCrashesAndEmployeesResponseDesktopAppExperience `json:"desktopAppExperience"`
}
type GetNumOfDevicesResponseDesktopAppExperienceAdoptionNumDevicesValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetNumOfDevicesResponseDesktopAppExperienceAdoptionNumDevices struct {
	Total           *int64                                                                         `json:"total"`
	ValueTimeSeries []GetNumOfDevicesResponseDesktopAppExperienceAdoptionNumDevicesValueTimeSeries `json:"valueTimeSeries"`
}

type GetNumOfDevicesResponseDesktopAppExperienceAdoption struct {
	NumDevices *GetNumOfDevicesResponseDesktopAppExperienceAdoptionNumDevices `json:"numDevices"`
}

type GetNumOfDevicesResponseDesktopAppExperience struct {
	Adoption *GetNumOfDevicesResponseDesktopAppExperienceAdoption `json:"adoption"`
}

type GetNumOfDevicesResponse struct {
	DesktopAppExperience *GetNumOfDevicesResponseDesktopAppExperience `json:"desktopAppExperience"`
}
type GetNumOfDevicesWithCrashesResponseDesktopAppExperienceReliabilityDevicesWithCrashesValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetNumOfDevicesWithCrashesResponseDesktopAppExperienceReliabilityDevicesWithCrashes struct {
	Total           *int64                                                                                               `json:"total"`
	ValueTimeSeries []GetNumOfDevicesWithCrashesResponseDesktopAppExperienceReliabilityDevicesWithCrashesValueTimeSeries `json:"valueTimeSeries"`
}

type GetNumOfDevicesWithCrashesResponseDesktopAppExperienceReliability struct {
	DevicesWithCrashes *GetNumOfDevicesWithCrashesResponseDesktopAppExperienceReliabilityDevicesWithCrashes `json:"devicesWithCrashes"`
}

type GetNumOfDevicesWithCrashesResponseDesktopAppExperienceAdoptionNumDevicesValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetNumOfDevicesWithCrashesResponseDesktopAppExperienceAdoptionNumDevices struct {
	Total           *int64                                                                                    `json:"total"`
	ValueTimeSeries []GetNumOfDevicesWithCrashesResponseDesktopAppExperienceAdoptionNumDevicesValueTimeSeries `json:"valueTimeSeries"`
}

type GetNumOfDevicesWithCrashesResponseDesktopAppExperienceAdoption struct {
	NumDevices *GetNumOfDevicesWithCrashesResponseDesktopAppExperienceAdoptionNumDevices `json:"numDevices"`
}

type GetNumOfDevicesWithCrashesResponseDesktopAppExperience struct {
	Reliability *GetNumOfDevicesWithCrashesResponseDesktopAppExperienceReliability `json:"reliability"`
	Adoption    *GetNumOfDevicesWithCrashesResponseDesktopAppExperienceAdoption    `json:"adoption"`
}

type GetNumOfDevicesWithCrashesResponse struct {
	DesktopAppExperience *GetNumOfDevicesWithCrashesResponseDesktopAppExperience `json:"desktopAppExperience"`
}
type GetNumOfEmployeesResponseDesktopAppExperienceAdoptionNumEmployeesValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetNumOfEmployeesResponseDesktopAppExperienceAdoptionNumEmployees struct {
	Total           *int64                                                                             `json:"total"`
	ValueTimeSeries []GetNumOfEmployeesResponseDesktopAppExperienceAdoptionNumEmployeesValueTimeSeries `json:"valueTimeSeries"`
}

type GetNumOfEmployeesResponseDesktopAppExperienceAdoption struct {
	NumEmployees *GetNumOfEmployeesResponseDesktopAppExperienceAdoptionNumEmployees `json:"numEmployees"`
}

type GetNumOfEmployeesResponseDesktopAppExperience struct {
	Adoption *GetNumOfEmployeesResponseDesktopAppExperienceAdoption `json:"adoption"`
}

type GetNumOfEmployeesResponse struct {
	DesktopAppExperience *GetNumOfEmployeesResponseDesktopAppExperience `json:"desktopAppExperience"`
}
type GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceReliabilityEmployeesWithCrashesValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceReliabilityEmployeesWithCrashes struct {
	Total           *int64                                                                                                   `json:"total"`
	ValueTimeSeries []GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceReliabilityEmployeesWithCrashesValueTimeSeries `json:"valueTimeSeries"`
}

type GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceReliability struct {
	EmployeesWithCrashes *GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceReliabilityEmployeesWithCrashes `json:"employeesWithCrashes"`
}

type GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceAdoptionNumEmployeesValueTimeSeries struct {
	MetricValue *int64 `json:"metricValue"`
	Timestamp   string `json:"timestamp"`
}

type GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceAdoptionNumEmployees struct {
	Total           *int64                                                                                        `json:"total"`
	ValueTimeSeries []GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceAdoptionNumEmployeesValueTimeSeries `json:"valueTimeSeries"`
}

type GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceAdoption struct {
	NumEmployees *GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceAdoptionNumEmployees `json:"numEmployees"`
}

type GetNumOfEmployeesWithCrashesResponseDesktopAppExperience struct {
	Reliability *GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceReliability `json:"reliability"`
	Adoption    *GetNumOfEmployeesWithCrashesResponseDesktopAppExperienceAdoption    `json:"adoption"`
}

type GetNumOfEmployeesWithCrashesResponse struct {
	DesktopAppExperience *GetNumOfEmployeesWithCrashesResponseDesktopAppExperience `json:"desktopAppExperience"`
}
type OverviewDesktopTooltipsResponseDesktopAppOverviewNumberOfEmployeesPerAppMetrics struct {
	AppID string   `json:"appId"`
	Value *float64 `json:"value"`
}

type OverviewDesktopTooltipsResponseDesktopAppOverviewNumberOfEmployees struct {
	PerAppMetrics []OverviewDesktopTooltipsResponseDesktopAppOverviewNumberOfEmployeesPerAppMetrics `json:"perAppMetrics"`
}

type OverviewDesktopTooltipsResponseDesktopAppOverviewCrashesPerEmployeePerAppMetrics struct {
	AppID string   `json:"appId"`
	Value *float64 `json:"value"`
}

type OverviewDesktopTooltipsResponseDesktopAppOverviewCrashesPerEmployee struct {
	PerAppMetrics []OverviewDesktopTooltipsResponseDesktopAppOverviewCrashesPerEmployeePerAppMetrics `json:"perAppMetrics"`
}

type OverviewDesktopTooltipsResponseDesktopAppOverview struct {
	NumberOfEmployees  *OverviewDesktopTooltipsResponseDesktopAppOverviewNumberOfEmployees  `json:"numberOfEmployees"`
	CrashesPerEmployee *OverviewDesktopTooltipsResponseDesktopAppOverviewCrashesPerEmployee `json:"crashesPerEmployee"`
}

type OverviewDesktopTooltipsResponse struct {
	DesktopAppOverview *OverviewDesktopTooltipsResponseDesktopAppOverview `json:"desktopAppOverview"`
}
type TilesDesktopCrashesPerEmployeeResponseOverviewTilesPerAppMetrics struct {
	AppID         string   `json:"appId"`
	AppName       string   `json:"appName"`
	CurrentValue  *float64 `json:"currentValue"`
	PreviousValue *float64 `json:"previousValue"`
}

type TilesDesktopCrashesPerEmployeeResponseOverviewTiles struct {
	PerAppMetrics []TilesDesktopCrashesPerEmployeeResponseOverviewTilesPerAppMetrics `json:"perAppMetrics"`
}

type TilesDesktopCrashesPerEmployeeResponseOverview struct {
	Tiles *TilesDesktopCrashesPerEmployeeResponseOverviewTiles `json:"tiles"`
}

type TilesDesktopCrashesPerEmployeeResponse struct {
	Overview *TilesDesktopCrashesPerEmployeeResponseOverview `json:"overview"`
}
type TilesDesktopNumberOfEmployeesResponseOverviewTilesPerAppMetrics struct {
	AppID         string   `json:"appId"`
	CurrentValue  *float64 `json:"currentValue"`
	PreviousValue *float64 `json:"previousValue"`
}

type TilesDesktopNumberOfEmployeesResponseOverviewTiles struct {
	PerAppMetrics []TilesDesktopNumberOfEmployeesResponseOverviewTilesPerAppMetrics `json:"perAppMetrics"`
}

type TilesDesktopNumberOfEmployeesResponseOverview struct {
	Tiles *TilesDesktopNumberOfEmployeesResponseOverviewTiles `json:"tiles"`
}

type TilesDesktopNumberOfEmployeesResponse struct {
	Overview *TilesDesktopNumberOfEmployeesResponseOverview `json:"overview"`
}
type TilesErrorCountResponseOverviewTilesPerAppMetricsThreshold struct {
	Level string `json:"level"`
}

type TilesErrorCountResponseOverviewTilesPerAppMetrics struct {
	AppID         string                                                      `json:"appId"`
	CurrentValue  *float64                                                    `json:"currentValue"`
	PreviousValue *float64                                                    `json:"previousValue"`
	Threshold     *TilesErrorCountResponseOverviewTilesPerAppMetricsThreshold `json:"threshold"`
}

type TilesErrorCountResponseOverviewTiles struct {
	PerAppMetrics []TilesErrorCountResponseOverviewTilesPerAppMetrics `json:"perAppMetrics"`
}

type TilesErrorCountResponseOverview struct {
	Tiles *TilesErrorCountResponseOverviewTiles `json:"tiles"`
}

type TilesErrorCountResponse struct {
	Overview *TilesErrorCountResponseOverview `json:"overview"`
}
type TilesFrustratingPageLoadsResponseOverviewTilesPerAppMetricsThreshold struct {
	Level string `json:"level"`
}

type TilesFrustratingPageLoadsResponseOverviewTilesPerAppMetrics struct {
	AppID         string                                                                `json:"appId"`
	CurrentValue  *float64                                                              `json:"currentValue"`
	PreviousValue *float64                                                              `json:"previousValue"`
	Threshold     *TilesFrustratingPageLoadsResponseOverviewTilesPerAppMetricsThreshold `json:"threshold"`
}

type TilesFrustratingPageLoadsResponseOverviewTiles struct {
	PerAppMetrics []TilesFrustratingPageLoadsResponseOverviewTilesPerAppMetrics `json:"perAppMetrics"`
}

type TilesFrustratingPageLoadsResponseOverview struct {
	Tiles *TilesFrustratingPageLoadsResponseOverviewTiles `json:"tiles"`
}

type TilesFrustratingPageLoadsResponse struct {
	Overview *TilesFrustratingPageLoadsResponseOverview `json:"overview"`
}
type TilesNumberOfEmployeesResponseOverviewTilesPerAppMetricsThreshold struct {
	Level string `json:"level"`
}

type TilesNumberOfEmployeesResponseOverviewTilesPerAppMetrics struct {
	AppID         string                                                             `json:"appId"`
	CurrentValue  *float64                                                           `json:"currentValue"`
	PreviousValue *float64                                                           `json:"previousValue"`
	Threshold     *TilesNumberOfEmployeesResponseOverviewTilesPerAppMetricsThreshold `json:"threshold"`
}

type TilesNumberOfEmployeesResponseOverviewTiles struct {
	PerAppMetrics []TilesNumberOfEmployeesResponseOverviewTilesPerAppMetrics `json:"perAppMetrics"`
}

type TilesNumberOfEmployeesResponseOverview struct {
	Tiles *TilesNumberOfEmployeesResponseOverviewTiles `json:"tiles"`
}

type TilesNumberOfEmployeesResponse struct {
	Overview *TilesNumberOfEmployeesResponseOverview `json:"overview"`
}
type TilesPageLoadTimeResponseOverviewTilesPerAppMetricsThreshold struct {
	Level string `json:"level"`
}

type TilesPageLoadTimeResponseOverviewTilesPerAppMetrics struct {
	AppID         string                                                        `json:"appId"`
	CurrentValue  *float64                                                      `json:"currentValue"`
	PreviousValue *float64                                                      `json:"previousValue"`
	Threshold     *TilesPageLoadTimeResponseOverviewTilesPerAppMetricsThreshold `json:"threshold"`
}

type TilesPageLoadTimeResponseOverviewTiles struct {
	PerAppMetrics []TilesPageLoadTimeResponseOverviewTilesPerAppMetrics `json:"perAppMetrics"`
}

type TilesPageLoadTimeResponseOverview struct {
	Tiles *TilesPageLoadTimeResponseOverviewTiles `json:"tiles"`
}

type TilesPageLoadTimeResponse struct {
	Overview *TilesPageLoadTimeResponseOverview `json:"overview"`
}
type TilesTransactionTimeResponseOverviewTilesPerAppMetricsThreshold struct {
	Level string `json:"level"`
}

type TilesTransactionTimeResponseOverviewTilesPerAppMetrics struct {
	AppID         string                                                           `json:"appId"`
	CurrentValue  *float64                                                         `json:"currentValue"`
	PreviousValue *float64                                                         `json:"previousValue"`
	Threshold     *TilesTransactionTimeResponseOverviewTilesPerAppMetricsThreshold `json:"threshold"`
}

type TilesTransactionTimeResponseOverviewTiles struct {
	PerAppMetrics []TilesTransactionTimeResponseOverviewTilesPerAppMetrics `json:"perAppMetrics"`
}

type TilesTransactionTimeResponseOverview struct {
	Tiles *TilesTransactionTimeResponseOverviewTiles `json:"tiles"`
}

type TilesTransactionTimeResponse struct {
	Overview *TilesTransactionTimeResponseOverview `json:"overview"`
}
type TilesUsageTimeResponseOverviewTilesPerAppMetricsThreshold struct {
	Level string `json:"level"`
}

type TilesUsageTimeResponseOverviewTilesPerAppMetrics struct {
	AppID         string                                                     `json:"appId"`
	CurrentValue  *float64                                                   `json:"currentValue"`
	PreviousValue *float64                                                   `json:"previousValue"`
	Threshold     *TilesUsageTimeResponseOverviewTilesPerAppMetricsThreshold `json:"threshold"`
}

type TilesUsageTimeResponseOverviewTiles struct {
	PerAppMetrics []TilesUsageTimeResponseOverviewTilesPerAppMetrics `json:"perAppMetrics"`
}

type TilesUsageTimeResponseOverview struct {
	Tiles *TilesUsageTimeResponseOverviewTiles `json:"tiles"`
}

type TilesUsageTimeResponse struct {
	Overview *TilesUsageTimeResponseOverview `json:"overview"`
}
type WebOverviewTooltipsResponseOverviewPageLoadTimeMSPerAppMetrics struct {
	AppID string   `json:"appId"`
	Value *float64 `json:"value"`
}

type WebOverviewTooltipsResponseOverviewPageLoadTimeMS struct {
	PerAppMetrics []WebOverviewTooltipsResponseOverviewPageLoadTimeMSPerAppMetrics `json:"perAppMetrics"`
}

type WebOverviewTooltipsResponseOverviewTransactionDurationMSPerAppMetrics struct {
	AppID string   `json:"appId"`
	Value *float64 `json:"value"`
}

type WebOverviewTooltipsResponseOverviewTransactionDurationMS struct {
	PerAppMetrics []WebOverviewTooltipsResponseOverviewTransactionDurationMSPerAppMetrics `json:"perAppMetrics"`
}

type WebOverviewTooltipsResponseOverviewNumberOfErrorsPerAppMetrics struct {
	AppID string   `json:"appId"`
	Value *float64 `json:"value"`
}

type WebOverviewTooltipsResponseOverviewNumberOfErrors struct {
	PerAppMetrics []WebOverviewTooltipsResponseOverviewNumberOfErrorsPerAppMetrics `json:"perAppMetrics"`
}

type WebOverviewTooltipsResponseOverviewUsageTimePerEmployeeMSPerAppMetrics struct {
	AppID string   `json:"appId"`
	Value *float64 `json:"value"`
}

type WebOverviewTooltipsResponseOverviewUsageTimePerEmployeeMS struct {
	PerAppMetrics []WebOverviewTooltipsResponseOverviewUsageTimePerEmployeeMSPerAppMetrics `json:"perAppMetrics"`
}

type WebOverviewTooltipsResponseOverviewNumberOfEmployeesPerAppMetrics struct {
	AppID string   `json:"appId"`
	Value *float64 `json:"value"`
}

type WebOverviewTooltipsResponseOverviewNumberOfEmployees struct {
	PerAppMetrics []WebOverviewTooltipsResponseOverviewNumberOfEmployeesPerAppMetrics `json:"perAppMetrics"`
}

type WebOverviewTooltipsResponseOverviewNumberOfFrustratingPageloadsPerAppMetrics struct {
	AppID string   `json:"appId"`
	Value *float64 `json:"value"`
}

type WebOverviewTooltipsResponseOverviewNumberOfFrustratingPageloads struct {
	PerAppMetrics []WebOverviewTooltipsResponseOverviewNumberOfFrustratingPageloadsPerAppMetrics `json:"perAppMetrics"`
}

type WebOverviewTooltipsResponseOverview struct {
	PageLoadTimeMS               *WebOverviewTooltipsResponseOverviewPageLoadTimeMS               `json:"pageLoadTimeMs"`
	TransactionDurationMS        *WebOverviewTooltipsResponseOverviewTransactionDurationMS        `json:"transactionDurationMs"`
	NumberOfErrors               *WebOverviewTooltipsResponseOverviewNumberOfErrors               `json:"numberOfErrors"`
	UsageTimePerEmployeeMS       *WebOverviewTooltipsResponseOverviewUsageTimePerEmployeeMS       `json:"usageTimePerEmployeeMs"`
	NumberOfEmployees            *WebOverviewTooltipsResponseOverviewNumberOfEmployees            `json:"numberOfEmployees"`
	NumberOfFrustratingPageloads *WebOverviewTooltipsResponseOverviewNumberOfFrustratingPageloads `json:"numberOfFrustratingPageloads"`
}

type WebOverviewTooltipsResponse struct {
	Overview *WebOverviewTooltipsResponseOverview `json:"overview"`
}

// GetAvgNetworkResponseTimeRequest selects a desktop application and time range.
type GetAvgNetworkResponseTimeRequest struct {
	ID               string                           `json:"id"`
	TimeFrameFilter  DesktopAppTimeframeFilterInput   `json:"timeFrameFilter"`
	DimensionsFilter *DesktopAppDimensionsFilterInput `json:"dimensionsFilter,omitempty"`
}
