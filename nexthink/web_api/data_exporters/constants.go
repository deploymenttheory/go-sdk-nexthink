package data_exporters

const Endpoint = "/apigateway/api/v1/data-exporter-config/data-exporter"
const EndpointExecutionStatus = "/apigateway/data-exporter-execution-status/api/v1"

// Numeric enums required by configuration and test writes. Reads use symbolic names.
const (
	FormatFile     = 0
	FormatPayload  = 1
	FileFormatCSV  = 0
	FileFormatJSON = 1
	QueryScheduled = 0
	QueryStreaming = 1
)
