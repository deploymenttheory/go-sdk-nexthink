package investigations

import (
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"
)

// InvestigationInput saves a definition; it does not run the NQL query.
// Description is sent by the UI, but the tested server ignores it and returns an empty string.
type InvestigationInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	NQL         string `json:"nql"`
}
type NQLDefinition struct {
	UID              *string         `json:"uid"`
	Status           *string         `json:"status"`
	Dependencies     json.RawMessage `json:"dependencies"`
	CreationDate     *string         `json:"creationDate"`
	ModificationDate *string         `json:"modificationDate"`
	Query            string          `json:"query"`
	Owner            json.RawMessage `json:"owner"`
}
type Permissions struct {
	Share bool `json:"share"`
	Edit  bool `json:"edit"`
}

// UID is the identifier used by Get, Update, Delete and Export; List calls it contentId.
type Investigation struct {
	UID              string        `json:"uid"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	CreationDate     string        `json:"creationDate"`
	ModificationDate string        `json:"modificationDate"`
	NQL              NQLDefinition `json:"nql"`
	Permissions      Permissions   `json:"permissions"`
}

// ExportDocument is also the import payload. Import requires a name not already in use.
type ExportDocument struct {
	Name     string `json:"name"`
	NQLQuery string `json:"nqlQuery"`
}
type Summary struct {
	content_administration.Content
	Name string `json:"name"`
}
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []Summary                          `json:"rows"`
}

// QueryRequest executes or inspects NQL with the UI's time context. It creates no saved investigation.
type QueryRequest struct {
	Query     string `json:"query"`
	Limit     int    `json:"limit"`
	Editor    string `json:"editor"`
	Origin    string `json:"origin"`
	TimeZone  string `json:"timeZone"`
	UTCOffset int    `json:"utcOffset"`
	Now       string `json:"now"`
}
type LinkRequest struct {
	NQL    string `json:"nql"`
	Origin string `json:"origin"`
}
type LinkResponse struct {
	URL string `json:"url"`
}
type Timeframe struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type QueryMetadata struct {
	CollectionLabel string            `json:"collectionLabel"`
	Count           int               `json:"count"`
	GlobalTimeframe *Timeframe        `json:"globalTimeframe"`
	Now             string            `json:"now"`
	AdditionalInfo  map[string]string `json:"additionalInfo"`
}
type QueryTooltip struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	ShowTitle   bool   `json:"showTitle"`
	Collapsible bool   `json:"collapsible"`
}
type QueryColumn struct {
	Name              string         `json:"name"`
	Visible           bool           `json:"visible"`
	DataType          string         `json:"dataType"`
	DataPath          string         `json:"dataPath"`
	DMURI             string         `json:"dmUri"`
	Label             string         `json:"label"`
	Tooltip           []QueryTooltip `json:"tooltip"`
	Numeric           bool           `json:"numeric"`
	OriginalName      string         `json:"originalName"`
	WidestColumnValue string         `json:"widestColumnValue"`
}
type QueryMeta struct {
	Collection      string        `json:"collection"`
	CollectionLabel string        `json:"collectionLabel"`
	Columns         []QueryColumn `json:"columns"`
}
type QueryDateTime struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}
type QueryStatistics struct {
	NumberOfRows    int           `json:"number_of_rows"`
	QueryDuration   string        `json:"query_duration"`
	TimeZone        string        `json:"time_zone"`
	UTCOffset       string        `json:"utc_offset"`
	UserDateTimeNow QueryDateTime `json:"user_date_time_now"`
}
type QueryResult struct {
	Meta       QueryMeta           `json:"meta"`
	Data       [][]json.RawMessage `json:"data"`
	Statistics QueryStatistics     `json:"statistics"`
	Now        string              `json:"now"`
}
