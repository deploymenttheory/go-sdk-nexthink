package data_export

// ExportConfig follows the browser CSV export contract.
type ExportConfig struct {
	Delimiter        string `json:"delimiter"`
	StringDecoration string `json:"string_decoration"`
	IncludeColumns   bool   `json:"include_columns"`
	FormatType       string `json:"format_type"`
	FileType         string `json:"file_type"`
	FileName         string `json:"file_name"`
}
type StartRequest struct {
	Query  string       `json:"query"`
	Config ExportConfig `json:"config"`
}
type StartResponse struct {
	StatusID string `json:"statusId"`
}

// Result may contain a signed download URL. Do not forward the browser bearer token to that URL.
type Status struct {
	Status string  `json:"status"`
	Result *string `json:"result,omitempty"`
}
