package data_exporters

import "encoding/json"

// Write requests use numeric enums; responses use their symbolic names.
type ConfigurationInput struct {
	UUID                         string                  `json:"uuid"`
	Name                         string                  `json:"name"`
	NQLID                        string                  `json:"nqlId"`
	Description                  string                  `json:"description"`
	ConnectorID                  string                  `json:"connectorId"`
	Format                       int                     `json:"format"`
	ThirdParty                   string                  `json:"thirdParty"`
	Enabled                      bool                    `json:"enabled"`
	Scheduling                   string                  `json:"scheduling"`
	QueryInfos                   []QueryInput            `json:"queryInfos"`
	MaxSendTime                  int64                   `json:"maxSendTime"`
	SkipNQLQueryClauseValidation bool                    `json:"skipNqlQueryClauseValidation"`
	FileInfo                     FileInput               `json:"fileInfo"`
	AdditionalConfigs            []AdditionalConfigInput `json:"additionalConfigs"`
}
type AdditionalConfigInput struct {
	Key   string `json:"configKey"`
	Value string `json:"configValue"`
}
type AdditionalConfig struct {
	Key   string  `json:"configKey"`
	Value *string `json:"configValue,omitempty"`
}
type QueryInput struct {
	ID                string                  `json:"queryId"`
	Text              string                  `json:"queryText"`
	Name              string                  `json:"queryName"`
	Type              int                     `json:"queryType"`
	AdditionalConfigs []AdditionalConfigInput `json:"additionalConfigs"`
}
type Query struct {
	ID                string             `json:"queryId"`
	Text              string             `json:"queryText"`
	Name              string             `json:"queryName"`
	Type              string             `json:"queryType"`
	AdditionalConfigs []AdditionalConfig `json:"additionalConfigs,omitempty"`
}
type FileInput struct {
	Name   string `json:"fileName"`
	Format int    `json:"fileFormat"`
	Size   int64  `json:"fileSize"`
}
type FileInfo struct {
	Name   string `json:"fileName"`
	Format string `json:"fileFormat"`
	Size   int64  `json:"fileSize"`
}
type Configuration struct {
	UUID                         string             `json:"uuid"`
	Name                         string             `json:"name"`
	NQLID                        string             `json:"nqlId"`
	Description                  string             `json:"description"`
	ConnectorID                  string             `json:"connectorId"`
	Format                       string             `json:"format"`
	ThirdParty                   string             `json:"thirdParty"`
	Enabled                      bool               `json:"enabled"`
	Scheduling                   string             `json:"scheduling"`
	QueryInfos                   []Query            `json:"queryInfos"`
	MaxSendTime                  int64              `json:"maxSendTime"`
	SkipNQLQueryClauseValidation bool               `json:"skipNqlQueryClauseValidation"`
	FileInfo                     FileInfo           `json:"fileInfo"`
	AdditionalConfigs            []AdditionalConfig `json:"additionalConfigs"`
	LastUpdate                   *int64             `json:"lastUpdate,omitempty"`
	Frequency                    *string            `json:"frequency,omitempty"`
}
type WriteResult struct {
	Message string `json:"message"`
}
type ListOptions struct{ Deleted *bool }
type CustomerInfo struct {
	TimeZone string `json:"timeZone"`
}
type Placeholders struct {
	Placeholders []string `json:"placeholders"`
}

// Status records vary by export protocol and execution phase. Preserve their full schema.
type StatusList struct {
	Statuses []json.RawMessage `json:"dataExporterExecutionStatusList"`
}
type ExecutionStatus map[string]json.RawMessage
type TestRequest struct {
	UUID              string                  `json:"uuid"`
	Name              string                  `json:"name"`
	ConnectorID       string                  `json:"connectorId"`
	Format            int                     `json:"format"`
	ThirdParty        string                  `json:"thirdParty"`
	Scheduling        string                  `json:"scheduling"`
	QueryInfo         QueryInput              `json:"queryInfo"`
	FileInfo          FileInput               `json:"fileInfo"`
	AdditionalConfigs []AdditionalConfigInput `json:"additionalConfigs"`
}
type TestExecution struct {
	ExecutionUUID string `json:"executionUuid"`
}
