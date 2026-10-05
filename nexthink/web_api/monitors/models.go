package monitors

import (
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
