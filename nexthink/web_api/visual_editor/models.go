package visual_editor

import "encoding/json"

type Collection struct {
	URI   string `json:"uri"`
	Label string `json:"label"`
	Type  string `json:"type"`
}
type Column struct {
	URI      string `json:"uri"`
	Label    string `json:"label"`
	Fullname string `json:"fullname"`
	Type     string `json:"type"`
	IsMetric bool   `json:"isMetric"`
	IsListed bool   `json:"isListed"`
}
type AggregateDescriptor struct {
	AggregateLabel string `json:"aggregateLabel"`
	AggregateAlias string `json:"aggregateAlias"`
	URI            string `json:"uri"`
}
type FilterCollection struct {
	Type                                string                `json:"type"`
	URI                                 string                `json:"uri"`
	Label                               string                `json:"label"`
	Category                            string                `json:"category"`
	HasTimeframe                        bool                  `json:"hasTimeframe"`
	MaxQueryableTimeframeInDays         *int                  `json:"maxQueryableTimeframeInDays"`
	MinTimeframePrecision               *string               `json:"minTimeframePrecision"`
	DefaultAggregateDescriptors         []AggregateDescriptor `json:"defaultAggregateDescriptors"`
	DefaultAggregateDurationDescriptors json.RawMessage       `json:"defaultAggregateDurationDescriptors"`
	DefaultNQLQuery                     *string               `json:"defaultNqlQuery"`
	IsPunctualEvent                     bool                  `json:"isPunctualEvent"`
}
