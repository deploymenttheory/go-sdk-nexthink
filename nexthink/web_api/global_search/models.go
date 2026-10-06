package global_search

import (
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
)

type SearchRequest struct {
	Search     string `json:"search"`
	MaxResults int    `json:"maxResults"`
}
type SearchResponse struct {
	Events []SearchEvent `json:"events"`
}

// SearchEvent is one object in the concatenated JSON response. The server sends
// provider information followed by independently completed category results.
type SearchEvent struct {
	LoadMoreURL string `json:"loadMoreUrl,omitempty"`
	// Raw preserves the original server object; marshaling a decoded event returns
	// this snapshot, including unknown provider fields.
	Raw               json.RawMessage `json:"-"`
	SearchInfoResults []CategoryInfo  `json:"searchInfoResults,omitempty"`
	Category          string          `json:"category,omitempty"`
	LabelKey          string          `json:"labelKey,omitempty"`
	CategoryOrder     int             `json:"categoryOrder,omitempty"`
	HasMoreData       bool            `json:"hasMoreData,omitempty"`
	Results           []SearchResult  `json:"results,omitempty"`
	ErrorResponse     *ProviderError  `json:"errorResponse,omitempty"`
}
type CategoryInfo struct {
	Category      string `json:"category"`
	CategoryOrder int    `json:"categoryOrder"`
}
type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type SearchResult struct {
	SuggestedInvestigations []Action   `json:"suggestedInvestigations,omitempty"`
	HiddenMetadata          []Metadata `json:"hiddenMetadata,omitempty"`
	Name                    string     `json:"name"`
	FoundByMetaKey          []Metadata `json:"foundByMetaKey,omitempty"`
	Actions                 []Action   `json:"actions,omitempty"`
	// Raw retains provider-specific fields without constraining their schema.
	Raw json.RawMessage `json:"-"`
}

func (r *SearchResult) UnmarshalJSON(data []byte) error {
	type plain SearchResult
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*r = SearchResult(decoded)
	r.Raw = append(r.Raw[:0], data...)
	return nil
}
func (r SearchResult) MarshalJSON() ([]byte, error) {
	if len(r.Raw) > 0 {
		return r.Raw, nil
	}
	type plain SearchResult
	return json.Marshal(plain(r))
}

type Action struct {
	IsDefault bool   `json:"isDefault"`
	LabelKey  string `json:"labelKey"`
	URL       string `json:"url"`
}
type ProviderError struct {
	Message string          `json:"message"`
	Code    json.RawMessage `json:"code,omitempty"`
	Source  string          `json:"source,omitempty"`
}

func (e *ProviderError) Error() string { return "global search: " + e.Message }

func (r *SearchEvent) UnmarshalJSON(data []byte) error {
	type plain SearchEvent
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*r = SearchEvent(decoded)
	r.Raw = append(r.Raw[:0], data...)
	return nil
}
func (r SearchEvent) MarshalJSON() ([]byte, error) {
	if len(r.Raw) > 0 {
		return r.Raw, nil
	}
	type plain SearchEvent
	return json.Marshal(plain(r))
}

// PortalSession is an alias of the shared explicit legacy credentials.
type PortalSession = auth.PortalSession

type PortalStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (e *PortalStatus) Error() string { return "legacy portal: " + e.Description }

type PortalResponse[T any] struct {
	ResultStatus *PortalStatus `json:"resultStatus"`
	Result       *T            `json:"result"`
}

// PortalAuthToken is credential material. Avoid logging or serializing it.
type PortalAuthToken struct {
	Token string `json:"token"`
}

func (PortalAuthToken) String() string { return "PortalAuthToken{redacted}" }

type LegacyDashboardSearchRequest struct {
	Search              string `json:"search"`
	MaxPersonalResults  int    `json:"maxPersonalResults"`
	MaxPublishedResults int    `json:"maxPublishedResults"`
	MaxRoleBasedResults int    `json:"maxRoleBasedResults"`
}
type LegacyDashboardSearchResults struct {
	Personal  []LegacyDashboard `json:"personal"`
	Published []LegacyDashboard `json:"published"`
	RoleBased []LegacyDashboard `json:"roleBased"`
}
type LegacyDashboard struct {
	DashboardName  string          `json:"dashboardName"`
	ModuleCategory string          `json:"moduleCategory"`
	ModuleUID      json.RawMessage `json:"moduleUid"`
	DashboardUID   json.RawMessage `json:"dashboardUid"`
}
