package network_insights

import "encoding/json"

type NodeIdentifier struct {
	ColumnName string `json:"columnName"`
	Name       string `json:"name"`
}
type GraphEdgeInput struct {
	TargetNode            NodeIdentifier `json:"targetNode"`
	MetricValue           string         `json:"metricValue"`
	SecondaryMetricValues []string       `json:"secondaryMetricValues"`
}
type GraphNodeInput struct {
	Identifier            NodeIdentifier   `json:"identifier"`
	MetricValue           string           `json:"metricValue"`
	SecondaryMetricValues []string         `json:"secondaryMetricValues"`
	Edges                 []GraphEdgeInput `json:"edges"`
}
type GraphData struct {
	Nodes []GraphNodeInput `json:"nodes"`
}
type GraphInsightsInput struct {
	MetricLabel           string    `json:"metricLabel"`
	ProtocolType          string    `json:"protocolType"`
	SecondaryMetricLabels []string  `json:"secondaryMetricLabels"`
	GraphData             GraphData `json:"graphData"`
}

type GetInsightsRequest struct {
	Input *GraphInsightsInput `json:"input" required:"true"`
}

type GetInsightsResponseGraphInsightsNodesIDentifier struct {
	ColumnName *string `json:"columnName"`
	Name       *string `json:"name"`
}

type GetInsightsResponseGraphInsightsNodesEdgesTargetNode struct {
	ColumnName *string `json:"columnName"`
	Name       *string `json:"name"`
}

type GetInsightsResponseGraphInsightsNodesEdges struct {
	TargetNode *GetInsightsResponseGraphInsightsNodesEdgesTargetNode `json:"targetNode"`
	Items      json.RawMessage                                       `json:"items"`
}

type GetInsightsResponseGraphInsightsNodes struct {
	IDentifier *GetInsightsResponseGraphInsightsNodesIDentifier `json:"identifier"`
	Items      json.RawMessage                                  `json:"items"`
	Edges      []GetInsightsResponseGraphInsightsNodesEdges     `json:"edges"`
}

type GetInsightsResponseGraphInsights struct {
	Nodes []GetInsightsResponseGraphInsightsNodes `json:"nodes"`
}

type GetInsightsResponse struct {
	GraphInsights *GetInsightsResponseGraphInsights `json:"graphInsights"`
}
