package monitors

import (
	"encoding/base64"
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"
)

type Configuration struct {
	Name                    string             `json:"name"`
	NQLID                   string             `json:"nqlId"`
	Description             string             `json:"description"`
	Notifications           []string           `json:"notifications"`
	MonitorType             string             `json:"monitorType"`
	Priority                string             `json:"priority"`
	Tags                    []Tag              `json:"tags"`
	MetricType              string             `json:"metricType"`
	PollingPeriodMillis     *int64             `json:"pollingPeriodMillis"`
	NQLQuery                string             `json:"nqlQuery"`
	NQLStoreUUID            string             `json:"nqlStoreUuid"`
	UpdateNQLQuery          bool               `json:"updateNqlQuery"`
	EmailConfiguration      EmailConfiguration `json:"emailConfiguration"`
	Origin                  string             `json:"origin"`
	TriggerConditionList    []TriggerCondition `json:"triggerConditionList"`
	FilterConditionList     []FilterCondition  `json:"filterConditionList"`
	GroupByList             []GroupBy          `json:"groupByList"`
	EvaluationWindowSeconds *int64             `json:"evaluationWindowSeconds"`
	AutoRecoveryMode        string             `json:"autoRecoveryMode"`
}
type Tag struct {
	Name  string  `json:"name"`
	Color *string `json:"color"`
}
type EmailAddress struct {
	Email string `json:"email"`
}
type EmailConfiguration struct {
	Recipients        []EmailAddress `json:"recipients"`
	CarbonCopies      []EmailAddress `json:"carbonCopies"`
	Subject           string         `json:"subject"`
	Content           string         `json:"content"`
	NotificationScope string         `json:"notificationScope"`
}
type StaticThreshold struct {
	Threshold float64 `json:"threshold"`
	Unit      *string `json:"unit"`
}
type ChangeThreshold struct {
	MetricAggregationTimeframe string  `json:"metricAggregationTimeframe"`
	MetricAggregationType      string  `json:"metricAggregationType"`
	ThresholdPercentage        float64 `json:"thresholdPercentage"`
}
type TimeOfDayThreshold struct {
	MetricAggregationTimeframe string  `json:"metricAggregationTimeframe"`
	MetricAggregationType      string  `json:"metricAggregationType"`
	ThresholdFactor            float64 `json:"thresholdFactor"`
}
type TriggerCondition struct {
	Name               string              `json:"name"`
	Label              string              `json:"label"`
	DataType           string              `json:"dataType"`
	MetricComparator   string              `json:"metricComparator"`
	StaticThreshold    *StaticThreshold    `json:"staticThreshold"`
	ChangeThreshold    *ChangeThreshold    `json:"changeThreshold"`
	TimeOfDayThreshold *TimeOfDayThreshold `json:"timeOfDayThreshold"`
}
type FilterCondition struct {
	CollectionDMURI     string          `json:"collectionDmUri"`
	FilterPropertyDMURI string          `json:"filterPropertyDmUri"`
	Comparator          string          `json:"comparator"`
	FilterValue         json.RawMessage `json:"filterValue"`
	FilterUnit          *string         `json:"filterUnit"`
	DataType            string          `json:"dataType"`
}
type GroupBy struct {
	PropertyDMURI string `json:"propertyDmUri"`
	Label         string `json:"label"`
	DataType      string `json:"dataType"`
	Type          string `json:"type"`
}

// UUID is omitted on create and required on update. It differs from the content/doc UUID.
type MonitorInput struct {
	Configuration
	UUID *string `json:"uuid,omitempty"`
}
type Monitor struct {
	Configuration
	UUID      string `json:"uuid"`
	ContentID string `json:"contentId"`
	DocUUID   string `json:"docUuid"`
	Revision  int    `json:"revision"`
}
type UpdateRequest struct {
	DocUUID  string       `json:"docUuid"`
	Revision int          `json:"revision"`
	Monitor  MonitorInput `json:"monitor"`
}
type DeleteRequest struct {
	DocUUID     string `json:"docUuid"`
	MonitorType string `json:"monitorType"`
}
type CreateResponse struct {
	ContentID *string `json:"addNqlMonitor"`
}
type GetResponse struct {
	Monitor *Monitor `json:"nqlMonitor"`
}

// Update and Delete returned null acknowledgments in the observed UI contract.
// Keep the scalar opaque instead of inventing a boolean success result.
type UpdateResponse struct {
	Acknowledgment json.RawMessage `json:"updateNqlMonitor"`
}
type DeleteResponse struct {
	Acknowledgment json.RawMessage `json:"deleteMonitor"`
}
type EmailConfig struct {
	Disabled bool `json:"disabled"`
}
type Summary struct {
	content_administration.Content
	MetricType  string      `json:"metricType"`
	EmailConfig EmailConfig `json:"emailConfig"`
	MonitorType string      `json:"monitorType"`
}
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []Summary                          `json:"rows"`
}

type ExportFile struct {
	FileName string `json:"fileName"`
	Content  string `json:"content"`
}
type LibraryExport struct {
	ContentExportData  ExportFile `json:"contentExportData"`
	MetadataExportData ExportFile `json:"metadataExportData"`
}
type ExportResponse struct {
	Export *ExportFile `json:"export"`
}
type ExportLibraryResponse struct {
	Export *LibraryExport `json:"exportLibrary"`
}

// ImportRequest.Content is decoded JSON text, not the export's base64 content.
type ImportRequest struct {
	Content string `json:"content"`
}
type ImportResponse struct {
	Acknowledgment json.RawMessage `json:"importMonitor"`
}

// DecodeContent decodes the base64 export into the JSON text accepted by Import.
func (f ExportFile) DecodeContent() (string, error) {
	b, err := base64.StdEncoding.DecodeString(f.Content)
	return string(b), err
}

type SetActivityRequest struct {
	DocUUID  string `json:"docUuid"`
	Activity string `json:"activity"`
}
type SetActivityResponse struct {
	Acknowledgment json.RawMessage `json:"toggleActivity"`
}
type UpdateBuiltInResponse struct {
	Acknowledgment json.RawMessage `json:"updateBuiltInMonitor"`
}
type License struct {
	FreeSlot         bool `json:"freeSlot"`
	FreeOnChangeSlot bool `json:"freeOnChangeSlot"`
	Volume           int  `json:"volume"`
	OnChangeVolume   int  `json:"onChangeVolume"`
	TotalMonitors    int  `json:"totalMonitors"`
}
type GetLicenseResponse struct {
	License *License `json:"license"`
}
type ListTagsResponse struct {
	Tags []Tag `json:"tags"`
}
type GetMetadataRequest struct {
	NQLQuery    string  `json:"nqlQuery"`
	DocUUID     *string `json:"docUuid,omitempty"`
	MonitorType string  `json:"monitorType,omitempty"`
}
type MetricMetadata struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	DataType string `json:"dataType"`
}
type CollectionMetadata struct {
	Type     string `json:"type"`
	URI      string `json:"uri"`
	Label    string `json:"label"`
	Category string `json:"category"`
}
type Metadata struct {
	MetricsMetadata     []MetricMetadata     `json:"metricsMetaData"`
	Collections         []CollectionMetadata `json:"collections"`
	GroupByOptions      []MetricMetadata     `json:"groupByOptions"`
	HasAssociation      bool                 `json:"hasAssociation"`
	SourceCollectionURI string               `json:"sourceCollectionUri"`
	AutoRecoveryOptions []string             `json:"autoRecoveryOptions"`
}
type GetMetadataResponse struct {
	Metadata *Metadata `json:"monitorMetaData"`
}
type AnalyzeQueryRequest struct {
	NQLQuery    string `json:"nqlQuery"`
	MonitorType string `json:"monitorType,omitempty"`
	MetricType  string `json:"metricType,omitempty"`
}
type QueryAnalysis struct {
	HasEventSource          bool     `json:"hasEventSource"`
	AvailableFrequencies    []int64  `json:"availableFrequencies"`
	SummarizeByFields       []string `json:"summarizeByFields"`
	HasSupportedContextSize bool     `json:"hasSupportedContextSize"`
	DuringPastMillis        *int64   `json:"duringPastMillis"`
	MaxDuringPastMillis     *int64   `json:"maxDuringPastMillis"`
}
type AnalyzeQueryResponse struct {
	Analysis *QueryAnalysis `json:"nqlQueryAnalysis"`
}
type ImpactQueryInput struct {
	Query             string             `json:"query"`
	TriggerConditions []TriggerCondition `json:"triggerConditions"`
}
type GetImpactQueryRequest struct {
	Input ImpactQueryInput `json:"impactQueryInput"`
}
type GetImpactQueryResponse struct {
	Query *string `json:"impactQuery"`
}
type CollectionDataInput struct {
	CollectionURI         string  `json:"collectionUri"`
	Type                  string  `json:"type"`
	VNQLPath              *string `json:"vnqlPath,omitempty"`
	CollectionOrMetricURI *string `json:"collectionOrMetricUri,omitempty"`
}
type ListFilterFieldsRequest struct {
	Collection CollectionDataInput `json:"collectionDataInput"`
}
type AutoCompletable struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}
type FilterField struct {
	URI             *string          `json:"uri"`
	Label           string           `json:"label"`
	Type            *string          `json:"type"`
	Comparators     []string         `json:"comparators"`
	Units           []string         `json:"units"`
	Category        *string          `json:"category"`
	SubCategory     *string          `json:"subCategory"`
	AutoCompletable *AutoCompletable `json:"autoCompletable"`
	VNQLPath        *string          `json:"vnqlPath"`
}
type ListFilterFieldsResponse struct {
	Fields []FilterField `json:"filterFields"`
}
