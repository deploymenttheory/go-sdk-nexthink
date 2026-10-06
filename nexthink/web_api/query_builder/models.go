package query_builder

import "encoding/json"

// Transform selects the destination object, time series, or metric for an NQL rewrite.
type Transform struct {
	ObjectURI     string `json:"objectURI,omitempty"`
	TimeSeriesURI string `json:"timeSeriesURI,omitempty"`
	Metric        string `json:"metric,omitempty"`
	// AdditionalFields preserves UI-provided transform attributes not yet modeled.
	AdditionalFields map[string]json.RawMessage `json:"-"`
}
type TransformRequest struct {
	Query     string    `json:"query"`
	Transform Transform `json:"transform"`
}
type QueryResponse struct {
	Query string `json:"query"`
}
type DestinationsRequest struct {
	Query string `json:"query"`
	Kind  string `json:"kind,omitempty"`
}
type ContextIdentifier struct {
	ObjectURI  string `json:"objectURI"`
	Identifier string `json:"identifier"`
}

// Conditions contains selected result rows; values retain their JSON scalar types.
type Conditions struct {
	Identifiers []ContextIdentifier `json:"identifiers"`
	Values      [][]json.RawMessage `json:"values"`
}
type DrilldownRequest struct {
	Query      string      `json:"query"`
	Conditions *Conditions `json:"conditions"`
	Transform  Transform   `json:"transform"`
}
type Destination struct {
	CategoryLabel string    `json:"categoryLabel"`
	Label         string    `json:"label"`
	Kind          string    `json:"kind"`
	Transform     Transform `json:"transform"`
}
type DestinationsResponse struct {
	Destinations       []Destination       `json:"destinations"`
	ContextIdentifiers []ContextIdentifier `json:"contextIdentifiers"`
	AdditionalInfo     map[string]string   `json:"additionalInfo"`
}
