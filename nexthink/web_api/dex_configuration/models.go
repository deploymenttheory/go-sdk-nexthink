package dex_configuration

import "encoding/json"

type Account struct {
	DEXApps int `json:"dexApps"`
}
type GetAccountResponse struct {
	Account *Account `json:"account"`
}
type ApplicationInput struct {
	UUID     string `json:"uuid"`
	Selected bool   `json:"selected"`
}
type UpdateApplicationsRequest struct {
	Applications []ApplicationInput `json:"apps"`
}
type ApplicationReference struct {
	UUID string `json:"uuid"`
}
type UpdateApplicationsResponse struct {
	Applications []ApplicationReference `json:"applications"`
}
type GetScoreMetricsRequest struct {
	Level *string `json:"level,omitempty"`
}
type ScoreMetricFields struct {
	ID                 float64  `json:"id"`
	Level              string   `json:"level"`
	Name               string   `json:"name"`
	MetricType         string   `json:"metricType"`
	Average            *float64 `json:"avg"`
	AverageDefault     *float64 `json:"avgDefault"`
	AverageEnabled     bool     `json:"avgEnabled"`
	Frustrating        *float64 `json:"frustrating"`
	FrustratingDefault *float64 `json:"frustratingDefault"`
	FrustratingEnabled bool     `json:"frustratingEnabled"`
	UUID               *string  `json:"uuid"`
	Selected           bool     `json:"selected"`
	AppType            *string  `json:"appType"`
	ReverseThresholds  bool     `json:"reverseThresholds"`
	WorkstationType    []string `json:"workstationType"`
	OSTypes            []string `json:"osTypes"`
}
type ScoreMetric struct {
	ScoreMetricFields
	Leaves []ScoreMetricFields `json:"leaves"`
}
type GetScoreMetricsResponse struct {
	Metrics []ScoreMetric `json:"ecScoreMetrics"`
}

// ScoreMetricInput is a sparse patch. Thresholds distinguish omitted (nil) from explicit JSON null (reset to default).
type ScoreMetricInput struct {
	ID                 float64         `json:"id"`
	Average            json.RawMessage `json:"avg,omitempty"`
	Frustrating        json.RawMessage `json:"frustrating,omitempty"`
	AverageEnabled     *bool           `json:"avgEnabled,omitempty"`
	FrustratingEnabled *bool           `json:"frustratingEnabled,omitempty"`
	Selected           *bool           `json:"selected,omitempty"`
}
type UpdateScoreMetricsRequest struct {
	Metrics []ScoreMetricInput `json:"ecScoreMetrics"`
}
type UpdateScoreMetricsResponse struct {
	Acknowledgment json.RawMessage `json:"ecScoreMetrics"`
}
type GetVDIOptInResponse struct {
	Enabled *bool `json:"optInVdi"`
}

// Opt-in mutation acknowledgments are preserved without assuming the scalar type.
type OptInVDIResponse struct {
	Acknowledgment json.RawMessage `json:"optInVdi"`
}
type GetMemoryMetricsOptInResponse struct {
	Enabled *bool `json:"optInMemoryMetrics"`
}
type OptInMemoryMetricsResponse struct {
	Acknowledgment json.RawMessage `json:"optInMemoryMetrics"`
}
type CampaignReference struct {
	NQLID string `json:"nqlId"`
}
type GetCampaignResponse struct {
	Campaign *CampaignReference `json:"configCampaign"`
}
type EnableCampaignRequest struct {
	NQLID string `json:"nqlId"`
}
type EnableCampaignResponse struct {
	Acknowledgment json.RawMessage `json:"enableCampaign"`
}

// Application describes whether an application participates in DEX scoring.
type Application struct {
	UUID        string   `json:"uuid"`
	Name        string   `json:"name"`
	Selected    bool     `json:"selected"`
	URLPatterns []string `json:"urlPatterns"`
	Binaries    []string `json:"binaries"`
}
type GetApplicationsResponse struct {
	Applications []Application `json:"applications"`
}
