package custom_field_values

// Field describes an editable manual field returned for an inventory object URI.
type Field struct {
	URI           string `json:"uri"`
	Label         string `json:"label"`
	FieldDataType string `json:"fieldDataType"`
}
type ListResponse struct {
	CustomFields []Field `json:"customFields"`
}

// Value is string-encoded. Boolean fields use "1"/"0"; an empty string clears the value.
type UpdateRequest struct {
	URI       string   `json:"uri"`
	Value     string   `json:"value"`
	ObjectIDs []string `json:"objectIds"`
}
