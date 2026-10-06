package recommendations

// Recommendation follows the first-party Forge frontend schema. Category and
// lifecycle status remain strings so newer server values are not discarded.
type Recommendation struct {
	RecommendationID string     `json:"recommendationId"`
	Title            string     `json:"title"`
	Subtitle         string     `json:"subtitle"`
	Category         string     `json:"category"`
	CreatedAt        float64    `json:"createdAt"`
	Statistics       Statistics `json:"statistics"`
	Content          Content    `json:"content"`
	Lifecycle        Lifecycle  `json:"lifecycle"`
}
type Statistics struct {
	NumAbandonedConversations float64 `json:"numAbandonedConversations"`
	NumEscalatedConversations float64 `json:"numEscalatedConversations"`
}
type Content struct {
	WhatEmployeesAreAsking   string `json:"whatEmployeesAreAsking"`
	WhySparkIsNotResolvingIt string `json:"whySparkIsNotResolvingIt"`
	WhatIsRecommended        string `json:"whatIsRecommended"`
}
type Lifecycle struct {
	Status   string  `json:"status"`
	Assignee *string `json:"assignee"`
	Note     *string `json:"note,omitempty"`
}
type UpdateStatusRequest struct {
	Status string  `json:"status"`
	Note   *string `json:"note,omitempty"`
}
