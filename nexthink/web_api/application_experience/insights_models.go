package application_experience

// InsightCategory identifies the application experience narrative family.
type InsightCategory string

const (
	InsightAdoption      InsightCategory = "adoption"
	InsightHealth        InsightCategory = "health"
	InsightEfficiency    InsightCategory = "efficiency"
	InsightAffordability InsightCategory = "affordability"
	InsightDelight       InsightCategory = "delight"
)

type ApplicationInsightsRequest struct {
	ApplicationID     string                       `json:"applicationId"`
	CurrentTimeframe  string                       `json:"currentTimeframe"`
	PreviousTimeframe string                       `json:"previousTimeframe"`
	Timezone          string                       `json:"timezone"`
	Breakdowns        map[InsightCategory][]string `json:"breakdowns"`
	Type              InsightCategory              `json:"type,omitempty"`
}
type ApplicationInsight struct {
	ID         string          `json:"id"`
	Type       InsightCategory `json:"type"`
	Title      string          `json:"title"`
	Body       string          `json:"body"`
	Noteworthy bool            `json:"noteworthy"`
	PromptSlug string          `json:"promptSlug"`
}
type ApplicationInsightsResponse []ApplicationInsight
