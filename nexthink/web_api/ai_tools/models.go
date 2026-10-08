// Package ai_tools implements the AI Tools browser API recovered from the Nexthink UI.
package ai_tools

import "encoding/json"

// ToolRequest is the application configuration and governance document used by the UI.
// Null configuration sections disable that source of adoption monitoring. Preserve the
// existing configuration when changing governance; Update replaces the document.
type ToolRequest struct {
	Name                                  string                  `json:"name"`
	NQLID                                 string                  `json:"nqlId"`
	Description                           string                  `json:"description,omitempty"`
	WebConfiguration                      *WebConfiguration       `json:"webConfiguration"`
	BinaryConfiguration                   *BinaryConfiguration    `json:"binaryConfiguration"`
	CollectorConfiguration                *CollectorConfiguration `json:"collectorConfiguration"`
	BinaryWithNetworkTrafficConfiguration *NetworkConfiguration   `json:"binaryWithNetworkTrafficConfiguration"`
	Metadata                              *Metadata               `json:"metadata,omitempty"`
	Experience                            *Experience             `json:"experience,omitempty"`
	GovernanceStatus                      string                  `json:"governanceStatus,omitempty"`
	CompliancePolicy                      *CompliancePolicy       `json:"compliancePolicy,omitempty"`
	Domains                               []string                `json:"domains,omitempty"`
}

// CopilotRequest references existing connector credentials; it never contains their secret.
type CopilotRequest struct {
	Name             string            `json:"name"`
	NQLID            string            `json:"nqlId"`
	Description      string            `json:"description,omitempty"`
	APICredentials   []string          `json:"apiCredentials"`
	Metadata         *Metadata         `json:"metadata,omitempty"`
	Experience       *Experience       `json:"experience,omitempty"`
	GovernanceStatus string            `json:"governanceStatus,omitempty"`
	CompliancePolicy *CompliancePolicy `json:"compliancePolicy,omitempty"`
	Domains          []string          `json:"domains,omitempty"`
}
type Tool struct {
	ID                                    string                  `json:"id"`
	Name                                  string                  `json:"name"`
	NQLID                                 string                  `json:"nqlId"`
	Description                           string                  `json:"description,omitempty"`
	ToolOrigin                            string                  `json:"toolOrigin,omitempty"`
	WebConfiguration                      *WebConfiguration       `json:"webConfiguration,omitempty"`
	BinaryConfiguration                   *BinaryConfiguration    `json:"binaryConfiguration,omitempty"`
	CollectorConfiguration                *CollectorConfiguration `json:"collectorConfiguration,omitempty"`
	BinaryWithNetworkTrafficConfiguration *NetworkConfiguration   `json:"binaryWithNetworkTrafficConfiguration,omitempty"`
	APICredentials                        []string                `json:"apiCredentials,omitempty"`
	Metadata                              *Metadata               `json:"metadata,omitempty"`
	Experience                            *ExperienceResponse     `json:"experience,omitempty"`
	Domains                               []string                `json:"domains,omitempty"`
	DiscoveryInformation                  *DiscoveryInformation   `json:"discoveryInformation,omitempty"`
	GovernanceStatus                      string                  `json:"governanceStatus,omitempty"`
	CompliancePolicy                      *CompliancePolicy       `json:"compliancePolicy,omitempty"`
	Origin                                string                  `json:"origin,omitempty"`
	ContentType                           string                  `json:"contentType,omitempty"`
	CreatedAt                             string                  `json:"createdAt,omitempty"`
	Revision                              int                     `json:"_rev"`
}
type WebConfiguration struct {
	Rules []WebRule `json:"rules"`
}
type WebRule struct {
	ID                     string `json:"id"`
	URLName                string `json:"urlName"`
	BaseURLRegex           string `json:"baseUrlRegex"`
	MonitoredEndpointRegex string `json:"monitoredEndpointRegex"`
}
type BinaryConfiguration struct {
	BinaryNames []string `json:"binaryNames"`
}
type CollectorConfiguration struct {
	TeamsBotConfig TeamsBotConfiguration `json:"teamsBotConfig"`
}
type TeamsBotConfiguration struct {
	BotName string `json:"botName"`
}
type NetworkConfiguration struct {
	Rules []NetworkRule `json:"rules"`
}
type NetworkRule struct {
	ID         string `json:"id"`
	Domain     string `json:"domain"`
	DomainName string `json:"domainName"`
}
type Metadata struct {
	LicenseCount int `json:"licenseCount"`
}
type ExperienceResponse struct {
	Campaigns []ToolCampaign `json:"campaigns,omitempty"`
}
type Experience struct {
	Campaigns []ToolCampaign `json:"campaigns"`
}
type ToolCampaign struct {
	ToolCampaignNQLID string `json:"toolCampaignNqlId"`
	Enabled           bool   `json:"enabled"`
}
type CompliancePolicy struct {
	PolicyBehavior      string `json:"policyBehavior"`
	RedirectURL         string `json:"redirectUrl,omitempty"`
	RedirectLabel       string `json:"redirectLabel,omitempty"`
	CooldownPeriodHours int    `json:"cooldownPeriodHours,omitempty"`
}
type DiscoveryInformation struct {
	FirstDiscovered string  `json:"firstDiscovered,omitempty"`
	LastDiscovered  string  `json:"lastDiscovered,omitempty"`
	NumUsers        *int64  `json:"numUsers,omitempty"`
	Description     string  `json:"description,omitempty"`
	Country         string  `json:"country,omitempty"`
	Category        string  `json:"category,omitempty"`
	Evidence        string  `json:"evidence,omitempty"`
	RiskLevel       string  `json:"riskLevel,omitempty"`
	ConfidenceLevel float64 `json:"confidenceLevel,omitempty"`
}
type LegacyTool struct {
	ID       string `json:"id"`
	NQLID    string `json:"nqlId"`
	Name     string `json:"name"`
	ToolType string `json:"toolType"`
	AppType  string `json:"appType"`
	Origin   string `json:"origin"`
	Revision int    `json:"_rev"`
}
type License struct {
	Allowed   int  `json:"allowed"`
	Consumed  int  `json:"consumed"`
	CanCreate bool `json:"canCreate"`
}
type RedirectURL struct {
	Name        string `json:"name"`
	RedirectURL string `json:"redirectUrl"`
}
type CredentialsRequest struct {
	CredentialsReference string `json:"credsRef"`
}
type Insights struct {
	Message string `json:"message"`
}
type ToolInsightsRequest struct {
	ToolNQLID string `json:"toolNqlId"`
}
type GovernanceTrendPoint struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}
type GovernanceTrends struct {
	Pending     []GovernanceTrendPoint `json:"pending"`
	Recommended []GovernanceTrendPoint `json:"recommended"`
	Allowed     []GovernanceTrendPoint `json:"allowed"`
	Prohibited  []GovernanceTrendPoint `json:"prohibited"`
}
type GovernanceActiveUsers struct {
	ThisWeekCount     int64 `json:"thisWeekCount"`
	PreviousWeekCount int64 `json:"previousWeekCount"`
}
type GovernanceDashboard struct {
	Trends      GovernanceTrends      `json:"trends"`
	ActiveUsers GovernanceActiveUsers `json:"activeUsers"`
}

// Module is the shared AI Tools experience-campaign configuration. Filters retain
// the UI's extensible organization-filter document shape.
type Module struct {
	ContentID string            `json:"contentId,omitempty"`
	Revision  *int              `json:"_rev,omitempty"`
	Campaigns []ModuleCampaign  `json:"campaigns"`
	Filters   []json.RawMessage `json:"filters"`
}
type ModuleCampaign struct {
	CampaignNQLID    string `json:"campaignNqlId"`
	CampaignModuleID string `json:"campaignModuleId"`
	OptIn            bool   `json:"optIn"`
	ExclusionNQL     string `json:"exclusionNql"`
}
type GoalRequest struct {
	Name        string         `json:"name"`
	NQLID       string         `json:"nqlId,omitempty"`
	Description string         `json:"description,omitempty"`
	Metrics     []GoalMetric   `json:"metrics"`
	Population  GoalPopulation `json:"population"`
	Timeline    GoalTimeline   `json:"timeline"`
	// Progress is omitted to preserve the current status; JSON null clears a manual status.
	Progress json.RawMessage `json:"progress,omitempty"`
}
type Goal struct {
	GoalID      string          `json:"goalId"`
	NQLID       string          `json:"nqlId"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Metrics     []GoalMetric    `json:"metrics"`
	Population  GoalPopulation  `json:"population"`
	Timeline    GoalTimeline    `json:"timeline"`
	Progress    json.RawMessage `json:"progress,omitempty"`
	Revision    int             `json:"_rev"`
	ContentType string          `json:"contentType,omitempty"`
	Origin      string          `json:"origin,omitempty"`
	CreatedAt   string          `json:"createdAt,omitempty"`
}
type GoalMetric struct {
	Type                 string        `json:"type"`
	Threshold            GoalThreshold `json:"threshold"`
	PopulationPercentage float64       `json:"populationPercentage"`
	ToolScope            GoalToolScope `json:"toolScope"`
}
type GoalThreshold struct {
	Target  float64 `json:"target"`
	Cadence string  `json:"cadence"`
}
type GoalToolScope struct {
	PolicyGroups    *[]string   `json:"policyGroups,omitempty"`
	IndividualTools *[]GoalTool `json:"individualTools,omitempty"`
}
type GoalTool struct {
	ToolID string `json:"toolId"`
}
type GoalPopulation struct {
	AllEmployees   *struct{}           `json:"allEmployees,omitempty"`
	SpecificGroups *GoalSpecificGroups `json:"specificGroups,omitempty"`
}
type GoalSpecificGroups struct {
	UserGroups []GoalUserGroup `json:"userGroups"`
}
type GoalUserGroup struct {
	DMURI  string   `json:"dmUri"`
	Values []string `json:"values"`
}
type GoalTimeline struct {
	StartDate string `json:"startDate"`
	// EndDate accepts an RFC3339 JSON string; JSON null removes the deadline.
	EndDate json.RawMessage `json:"endDate,omitempty"`
}
