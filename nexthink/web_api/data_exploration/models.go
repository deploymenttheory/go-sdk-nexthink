package data_exploration

import "encoding/json"

// TimeContext matches the UI's WAAS headers. Nil uses the current UTC time.
// UTCOffset is minutes west of UTC, matching JavaScript Date.getTimezoneOffset.
type TimeContext struct {
	TimeZone    string `json:"timeZone"`
	UTCOffset   int    `json:"utcOffset"`
	ISODateTime string `json:"isoDateTime"`
	AppName     string `json:"appName"`
}
type FieldInput struct {
	DMURI  string `json:"dmUri,omitempty"`
	Metric string `json:"metric,omitempty"`
}
type Duration struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}
type DateRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type TimeRange struct {
	Type                  string     `json:"type"`
	RelativeRange         *Duration  `json:"relativeRange,omitempty"`
	AbsoluteRange         *DateRange `json:"absoluteRange,omitempty"`
	AbsoluteDateRange     *DateRange `json:"absoluteDateRange,omitempty"`
	AbsoluteDateTimeRange *DateRange `json:"absoluteDateTimeRange,omitempty"`
	ByDuration            *Duration  `json:"byDuration,omitempty"`
}

// Polymorphic filters, metric expressions and timeseries options follow the UI's GraphQL input unions.
// Preserve these JSON documents instead of dropping subtype-specific fields.
type QueryInput struct {
	Query      string            `json:"query"`
	TimeRange  *TimeRange        `json:"timeRange,omitempty"`
	Filters    []json.RawMessage `json:"filters,omitempty"`
	Sort       []Sort            `json:"sort,omitempty"`
	Breakdowns []json.RawMessage `json:"breakdowns,omitempty"`
	Metrics    []json.RawMessage `json:"metrics,omitempty"`
	TimeSeries json.RawMessage   `json:"timeSeries,omitempty"`
}
type QueryRequest struct {
	QueryInput
	Limit             *int    `json:"limit,omitempty"`
	Offset            *int    `json:"offset,omitempty"`
	ShowDrilldown     bool    `json:"showDrilldown,omitempty"`
	FetchTotalCount   bool    `json:"fetchTotalCount,omitempty"`
	IncludeGlobalTime bool    `json:"includeGlobalTime,omitempty"`
	UniqueKey         *string `json:"uniqueKey,omitempty"`
}
type FieldsRequest struct {
	CollectionURI string       `json:"collectionUri,omitempty"`
	FieldURIs     []FieldInput `json:"fieldUris,omitempty"`
}
type SystemRatingsRequest struct {
	FieldURIs []FieldInput `json:"fieldUris"`
}
type FilterValuesInput struct {
	CollectionURI string            `json:"collectionUri"`
	FieldURI      string            `json:"fieldUri"`
	Filters       []json.RawMessage `json:"filters,omitempty"`
	Limit         *int              `json:"limit,omitempty"`
	Search        *string           `json:"search,omitempty"`
}
type FilterValuesRequest struct {
	Input           FilterValuesInput `json:"input"`
	FetchTotalCount bool              `json:"fetchTotalCount,omitempty"`
}
type BreakdownItem struct {
	DMURI string `json:"dmUri"`
}
type BreakdownOptions struct {
	BreakdownItems []BreakdownItem `json:"breakdownItems"`
}

// BreakdownFieldsVariables is deliberately limited to the query used by the UI's field discovery.
type BreakdownFieldsVariables struct {
	Query string `json:"query"`
}

// BreakdownInsightsVariables carries only server-side NQL inputs. UI display flags
// such as showDrilldown and fetchTotalCount are not fields of this GraphQL input type.
type BreakdownInsightsVariables struct {
	Query     string            `json:"query"`
	TimeRange *TimeRange        `json:"timeRange,omitempty"`
	Filters   []json.RawMessage `json:"filters,omitempty"`
	Sort      []Sort            `json:"sort,omitempty"`
	Limit     *int              `json:"limit,omitempty"`
	Metrics   []json.RawMessage `json:"metrics,omitempty"`
}
type BreakdownFieldsRequest struct {
	NQLVariables BreakdownFieldsVariables `json:"nqlVariables"`
	Input        *BreakdownOptions        `json:"input,omitempty"`
}
type BreakdownInsightsRequest struct {
	NQLVariables BreakdownInsightsVariables `json:"nqlVariables"`
	Input        *BreakdownOptions          `json:"input,omitempty"`
}
type ByDurationsRequest struct {
	TimeRange TimeRange `json:"timeRange"`
}
type OrganisationRequest struct {
	CollectionURI string `json:"collectionUri"`
}
type ItemMetaRequest struct {
	FieldURI FieldInput `json:"fieldUri"`
	Item     string     `json:"item"`
}
type MenuInput struct {
	DMURI  string  `json:"dmUri,omitempty"`
	Metric string  `json:"metric,omitempty"`
	Value  *string `json:"value,omitempty"`
}
type Identifier struct {
	ObjectURI  string          `json:"objectURI"`
	Identifier string          `json:"identifier"`
	Value      json.RawMessage `json:"value"`
}
type MenuContext struct {
	AppName           string       `json:"appName"`
	Identifiers       []Identifier `json:"identifiers"`
	Origin            string       `json:"origin"`
	Query             string       `json:"query"`
	ViewInvestigation *bool        `json:"viewInvestigation,omitempty"`
}
type MenuRequest struct {
	Inputs  []MenuInput  `json:"inputs"`
	Context *MenuContext `json:"context,omitempty"`
}
type Tooltip struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	ShowTitle   bool   `json:"showTitle"`
	Collapsible bool   `json:"collapsible"`
}
type Field struct {
	Tooltip           []Tooltip `json:"tooltip"`
	URI               string    `json:"uri"`
	FullLabel         string    `json:"fullLabel"`
	HasDataPermission bool      `json:"hasDataPermission"`
	IsStaticDataModel bool      `json:"isStaticDataModel"`
}
type Rating struct {
	RatingID             string          `json:"ratingId"`
	Label                string          `json:"label"`
	TargetFieldURI       string          `json:"targetFieldUri"`
	AverageThreshold     json.RawMessage `json:"averageThreshold"`
	FrustratingThreshold json.RawMessage `json:"frustratingThreshold"`
	ReverseThreshold     bool            `json:"reverseThreshold"`
}
type SystemRating struct {
	FieldURI string  `json:"fieldUri"`
	Rating   *Rating `json:"rating"`
}
type FilterValue struct {
	Value          json.RawMessage `json:"value"`
	FormattedValue string          `json:"formattedValue"`
}
type FilterValues struct {
	SupportsSearch bool          `json:"supportsSearch"`
	HasMoreData    bool          `json:"hasMoreData"`
	TotalCount     *int          `json:"totalCount,omitempty"`
	Values         []FilterValue `json:"values"`
}
type BreakdownField struct {
	DMURI             string `json:"dmUri"`
	Label             string `json:"label"`
	FullLabel         string `json:"fullLabel"`
	Category          string `json:"category"`
	IsStaticDataModel bool   `json:"isStaticDataModel"`
}
type BreakdownFields struct {
	Data []BreakdownField `json:"data"`
}
type BreakdownInsight struct {
	DMURI       string          `json:"dmUri"`
	Label       string          `json:"label"`
	FullLabel   string          `json:"fullLabel"`
	Category    string          `json:"category"`
	Description *string         `json:"description"`
	Metric      json.RawMessage `json:"metric"`
	Value       json.RawMessage `json:"value"`
}
type ExecutedQuery struct {
	ExecutedQuery string `json:"executedQuery"`
}
type BreakdownInsights struct {
	Data []BreakdownInsight `json:"data"`
	Meta ExecutedQuery      `json:"meta"`
}
type ByDuration struct {
	Duration
	IsDefault bool   `json:"isDefault"`
	Label     string `json:"label"`
}
type OrganisationField struct {
	Label string `json:"label"`
	DMURI string `json:"dmUri"`
}
type ItemMeta struct {
	Tooltip json.RawMessage `json:"tooltip"`
}
type MenuField struct {
	DMURI  *string `json:"dmUri"`
	Metric *string `json:"metric"`
}
type LinkConfig struct {
	URL    string `json:"url"`
	NewTab bool   `json:"newTab"`
}
type ActionConfig struct {
	ActionName string          `json:"actionName"`
	Params     json.RawMessage `json:"params"`
}
type MenuConfig struct {
	LinkConfig   *LinkConfig   `json:"linkConfig"`
	ActionConfig *ActionConfig `json:"actionConfig"`
}
type MenuItem struct {
	Section      *string     `json:"section"`
	Label        string      `json:"label"`
	Category     *string     `json:"category"`
	CategoryIcon *string     `json:"categoryIcon"`
	SubMenu      *string     `json:"subMenu"`
	SubMenuIcon  *string     `json:"subMenuIcon"`
	Type         string      `json:"type"`
	Icon         *string     `json:"icon"`
	Config       *MenuConfig `json:"config"`
}
type Menu struct {
	Field MenuField  `json:"field"`
	Items []MenuItem `json:"items"`
}
type Sort struct {
	NQLColumnName string `json:"nqlColumnName"`
	Direction     string `json:"direction"`
}
type GlobalTimeInfo struct {
	GlobalTimeframe DateRange `json:"globalTimeframe"`
}
type DMElementInfo struct {
	ParentCollectionName *string `json:"parentCollectionName"`
	CollectionName       *string `json:"collectionName"`
}
type ColumnMetric struct {
	Function string `json:"function"`
}
type Format struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
type Column struct {
	Name         string        `json:"name"`
	OriginalName string        `json:"originalName"`
	Label        string        `json:"label"`
	Visible      bool          `json:"visible"`
	Description  *string       `json:"description"`
	DataPath     string        `json:"dataPath"`
	DMURI        string        `json:"dmUri"`
	DataType     string        `json:"dataType"`
	IsBreakdown  bool          `json:"isBreakdown"`
	Metric       *ColumnMetric `json:"metric"`
	Tooltip      []Tooltip     `json:"tooltip"`
	Format       *Format       `json:"format"`
}
type TimeSeries struct {
	BucketSizeRaw     string `json:"bucketSizeRaw"`
	BucketSizeSeconds int64  `json:"bucketSizeSeconds"`
}
type QueryMeta struct {
	HasMoreData    bool            `json:"hasMoreData"`
	Collection     string          `json:"collection"`
	TotalCount     *int            `json:"totalCount,omitempty"`
	DMElementInfo  DMElementInfo   `json:"dmElementInfo"`
	Columns        []Column        `json:"columns"`
	ExecutedQuery  string          `json:"executedQuery"`
	GlobalTimeInfo *GlobalTimeInfo `json:"globalTimeInfo,omitempty"`
	TimeSeries     *TimeSeries     `json:"timeSeries"`
	Sort           []Sort          `json:"sort"`
	DrillDown      json.RawMessage `json:"drillDown,omitempty"`
}
type DateTimeParts struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}
type QueryStatistics struct {
	QueryDuration   string        `json:"queryDuration"`
	Rows            int           `json:"rows"`
	TimeZone        string        `json:"timeZone"`
	UTCOffset       string        `json:"utcOffset"`
	UserDateTimeNow DateTimeParts `json:"userDateTimeNow"`
}
type QueryResult struct {
	Meta       QueryMeta           `json:"meta"`
	Data       [][]json.RawMessage `json:"data"`
	Links      json.RawMessage     `json:"links"`
	Statistics QueryStatistics     `json:"statistics"`
}
type WhereOperand struct {
	ColumnName *string `json:"columnName,omitempty"`
	FieldURI   *string `json:"fieldUri,omitempty"`
	Operator   *string `json:"operator,omitempty"`
	NQLValue   *string `json:"nqlValue,omitempty"`
}
type WhereClause struct {
	Operator string       `json:"operator"`
	Left     WhereOperand `json:"left"`
	Right    WhereOperand `json:"right"`
}
type InspectResult struct {
	QueryToBeExecuted string            `json:"queryToBeExecuted"`
	Collection        string            `json:"collection"`
	HasSummarize      bool              `json:"hasSummarize"`
	Breakdowns        []BreakdownItem   `json:"breakdowns"`
	Metrics           []json.RawMessage `json:"metrics"`
	AllowedMetrics    []json.RawMessage `json:"allowedMetrics"`
	Wheres            []WhereClause     `json:"wheres"`
	Sort              []Sort            `json:"sort"`
	GlobalTimeInfo    *GlobalTimeInfo   `json:"globalTimeInfo"`
}
type QueryResponse struct {
	Result *QueryResult `json:"nql"`
}
type InspectResponse struct {
	Result *InspectResult `json:"inspect"`
}
type GetFilterValuesResponse struct {
	Result *FilterValues `json:"filterValues"`
}
type ListFieldsResponse struct {
	Result []Field `json:"fields"`
}
type ListSystemRatingsResponse struct {
	Result []SystemRating `json:"systemRatings"`
}
type ListOrganisationFieldsResponse struct {
	Result []Field `json:"organisationFields"`
}
type GetMenuResponse struct {
	Result []Menu `json:"menu"`
}
type ListBreakdownFieldsResponse struct {
	Result *BreakdownFields `json:"breakdownFields"`
}
type GetBreakdownInsightsResponse struct {
	Result *BreakdownInsights `json:"breakdownInsights"`
}
type ListByDurationsResponse struct {
	Result []ByDuration `json:"byDurations"`
}
type GetOrganisationResponse struct {
	Result []OrganisationField `json:"organisation"`
}
type GetItemMetaResponse struct {
	Result *ItemMeta `json:"itemMeta"`
}
