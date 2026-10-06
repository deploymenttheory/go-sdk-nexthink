package support_checklists

import "encoding/json"

// Polymorphic NQL values, event payloads and plugin-defined details retain their exact JSON.
type ListPropertiesResponseItemValuesItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Status           *string                    `json:"status,omitempty"`
}
type ListPropertiesResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields    map[string]json.RawMessage              `json:"-"`
	Label               *string                                 `json:"label,omitempty"`
	LabelTranslationKey *string                                 `json:"labelTranslationKey,omitempty"`
	Values              *[]ListPropertiesResponseItemValuesItem `json:"values,omitempty"`
}
type ListPropertiesResponse []ListPropertiesResponseItem
type ListResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ID               *string                    `json:"id,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type ListResponse []ListResponseItem

// GetResponse contains evaluated checklist categories and values, matching the UI property table.
type GetResponse []ListPropertiesResponseItem
