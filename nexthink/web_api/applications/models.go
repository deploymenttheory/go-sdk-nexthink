package applications

import "encoding/json"

// ApplicationInput is the UI create/update body. Updates require the current revision.
// Omitted web/desktop/network sections keep the request appropriate to the app type.
type ApplicationInput struct {
	Name                     string                `json:"name"`
	Category                 string                `json:"category"`
	Revision                 int                   `json:"revision"`
	URLs                     []URLPattern          `json:"urls,omitempty"`
	KeyPages                 []json.RawMessage     `json:"keyPages,omitempty"`
	Transactions             []json.RawMessage     `json:"transactions,omitempty"`
	URLTable                 *EnabledSetting       `json:"urlTable,omitempty"`
	SoftNavs                 *EnabledSetting       `json:"softNavs,omitempty"`
	AdoptEnabled             *EnabledSetting       `json:"adoptEnabled,omitempty"`
	Appex360Enabled          *EnabledSetting       `json:"appex360Enabled,omitempty"`
	EngageCampaign           json.RawMessage       `json:"engageCampaign,omitempty"`
	NumAvailableLicenses     *int                  `json:"numAvailableLicenses,omitempty"`
	Thresholds               *Thresholds           `json:"thresholds,omitempty"`
	Desktop                  *DesktopConfiguration `json:"desktop,omitempty"`
	Network                  *NetworkConfiguration `json:"network,omitempty"`
	KeyPageIdentifiers       []json.RawMessage     `json:"keyPageIdentifiers,omitempty"`
	ApplicationTemplateRef   *TemplateReference    `json:"applicationTemplateRef,omitempty"`
	HardNavigation           *HardNavigation       `json:"hardNavigation,omitempty"`
	FunctionalErrorsEnabled  bool                  `json:"functionalErrorsEnabled,omitempty"`
	URLSanitizationRuleSetID *string               `json:"urlSanitizationRuleSetId,omitempty"`
}

// Application retains nulls in the response and opaque nested configuration documents.
type Application struct {
	ID                       string                `json:"id"`
	Name                     string                `json:"name"`
	Category                 string                `json:"category"`
	Revision                 int                   `json:"revision"`
	URLs                     []URLPattern          `json:"urls"`
	KeyPages                 []json.RawMessage     `json:"keyPages"`
	Transactions             []json.RawMessage     `json:"transactions"`
	URLTable                 *EnabledSetting       `json:"urlTable"`
	SoftNavs                 *EnabledSetting       `json:"softNavs"`
	AdoptEnabled             *EnabledSetting       `json:"adoptEnabled"`
	Appex360Enabled          *EnabledSetting       `json:"appex360Enabled"`
	EngageCampaign           json.RawMessage       `json:"engageCampaign"`
	NumAvailableLicenses     *int                  `json:"numAvailableLicenses"`
	Thresholds               *Thresholds           `json:"thresholds"`
	Desktop                  *DesktopConfiguration `json:"desktop"`
	Network                  *NetworkConfiguration `json:"network"`
	KeyPageIdentifiers       []json.RawMessage     `json:"keyPageIdentifiers"`
	ApplicationTemplateRef   *TemplateReference    `json:"applicationTemplateRef"`
	HardNavigation           *HardNavigation       `json:"hardNavigation"`
	FunctionalErrorsEnabled  bool                  `json:"functionalErrorsEnabled"`
	URLSanitizationRuleSetID *string               `json:"urlSanitizationRuleSetId"`
	Options                  json.RawMessage       `json:"options"`
	IsCampaignConfigured     bool                  `json:"isCampaignConfigured"`
	Campaigns                json.RawMessage       `json:"campaigns"`
	CustomMetrics            []json.RawMessage     `json:"customMetrics"`
}
type URLPattern struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Pattern string `json:"pattern"`
}
type EnabledSetting struct {
	Enabled bool `json:"enabled"`
}
type DesktopConfiguration struct {
	ExactMatchBinaryNamesIgnoreCase []string `json:"exactMatchBinaryNamesIgnoreCase"`
	IsDeviceCentric                 bool     `json:"isDeviceCentric"`
}
type NetworkConfiguration struct {
	Rules []json.RawMessage `json:"rules"`
}
type TemplateReference struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type HardNavigation struct {
	Algorithm string `json:"algorithm"`
}
type Threshold struct {
	Average     float64 `json:"average"`
	Frustrating float64 `json:"frustrating"`
}
type Thresholds struct {
	PageLoadTimeSec        Threshold `json:"pageLoadTimeSec"`
	TransactionDurationSec Threshold `json:"transactionDurationSec"`
}
type ListOptions struct {
	PageNumber int
	PageSize   int
	SortName   string
}
type Link struct {
	Href string `json:"href"`
}
type Links struct {
	Self Link  `json:"self"`
	Next *Link `json:"next,omitempty"`
}
type ListResponse struct {
	Total int           `json:"_total"`
	Items []Application `json:"items"`
	Links Links         `json:"_links"`
}

// DeleteResponse is the revision returned by the application deletion endpoint.
type DeleteResponse int

type TemplateSelector struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Reference string `json:"reference,omitempty"`
}
type TemplateKeyPage struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Selectors []TemplateSelector `json:"selectors"`
}
type TemplateKeyPageSelector struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Template struct {
	ID                         string                    `json:"id"`
	Name                       string                    `json:"name"`
	Type                       string                    `json:"type"`
	GlobalAppID                string                    `json:"globalAppId"`
	KeyPages                   []TemplateKeyPage         `json:"keyPages"`
	KeyPageSelectors           []TemplateKeyPageSelector `json:"keyPageSelectors"`
	FunctionalErrors           []json.RawMessage         `json:"functionalErrors"`
	FunctionalErrorsHeuristics json.RawMessage           `json:"functionalErrorsHeuristics"`
	Revision                   int                       `json:"_rev"`
	SelectorSystem             json.RawMessage           `json:"selectorSystem,omitempty"`
}
type TemplateLink struct {
	Href string `json:"href"`
}
type TemplateList struct {
	Total int                     `json:"_total"`
	Items []Template              `json:"items"`
	Links map[string]TemplateLink `json:"_links"`
}
