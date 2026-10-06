package cci_benchmarks

import "encoding/json"

type Source struct {
	Name string `json:"name"`
}
type BenchmarkQuery struct {
	Source  Source              `json:"source"`
	Metrics []string            `json:"metrics"`
	Filters map[string][]string `json:"filters,omitempty"`
}
type QueryRequest struct {
	Queries  []BenchmarkQuery `json:"queries"`
	TimeZone string           `json:"-"`
}
type Metadata struct {
	Unit          string `json:"unit"`
	MetricName    string `json:"metricName"`
	Type          string `json:"type"`
	BenchmarkName string `json:"benchmarkName"`
	URI           string `json:"uri"`
}
type Result struct {
	Metadata Metadata          `json:"metadata"`
	Keys     map[string]string `json:"keys"`
	Value    *float64          `json:"value"`
}

// QueryResponse retains partial benchmark results alongside service-defined errors.
type QueryResponse struct {
	Results []Result          `json:"results"`
	Errors  []json.RawMessage `json:"errors,omitempty"`
}
