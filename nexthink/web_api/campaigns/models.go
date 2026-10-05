package campaigns

import "encoding/json"

// CampaignInput is the management create body. Create saves a draft; publication is separate.
type CampaignInput struct {
	Name                   string             `json:"name"`
	Description            string             `json:"description"`
	NQLID                  string             `json:"nqlId"`
	Triggers               []string           `json:"triggers"`
	SenderTitle            *string            `json:"senderTitle,omitempty"`
	SenderName             *string            `json:"senderName,omitempty"`
	SenderImage            *string            `json:"senderImage,omitempty"`
	InvestigationQuery     *string            `json:"investigationQuery,omitempty"`
	IsImageGenerated       *bool              `json:"isImageGenerated,omitempty"`
	Priority               string             `json:"priority"`
	IsMandatory            bool               `json:"isMandatory"`
	Questions              []QuestionInput    `json:"questions"`
	IsMultiLanguageEnabled bool               `json:"isMultiLanguageEnabled"`
	MultiLanguageInfo      *MultiLanguageInfo `json:"multiLanguageInfo,omitempty"`
	Translations           []Translation      `json:"translations,omitempty"`
	QuietPeriodDuration    *string            `json:"quietPeriodDuration,omitempty"`
	IsParametric           bool               `json:"isParametric"`
	Parameters             []Parameter        `json:"parameters,omitempty"`
}
type Campaign struct {
	BCSUID                 string             `json:"bcsUid"`
	ContentID              string             `json:"contentId"`
	Status                 string             `json:"status"`
	PublishedDate          json.RawMessage    `json:"publishedDate"`
	Name                   string             `json:"name"`
	Description            string             `json:"description"`
	NQLID                  string             `json:"nqlId"`
	Triggers               []string           `json:"triggers"`
	SenderTitle            *string            `json:"senderTitle"`
	SenderName             *string            `json:"senderName"`
	SenderImage            *string            `json:"senderImage"`
	InvestigationQuery     *string            `json:"investigationQuery"`
	IsImageGenerated       *bool              `json:"isImageGenerated"`
	Priority               string             `json:"priority"`
	IsMandatory            bool               `json:"isMandatory"`
	Questions              []Question         `json:"questions"`
	IsMultiLanguageEnabled bool               `json:"isMultiLanguageEnabled"`
	MultiLanguageInfo      *MultiLanguageInfo `json:"multiLanguageInfo"`
	Translations           []Translation      `json:"translations"`
	QuietPeriodDuration    *string            `json:"quietPeriodDuration"`
	IsParametric           bool               `json:"isParametric"`
	Parameters             []Parameter        `json:"parameters"`
}
type QuestionInput struct {
	ID          string        `json:"id"`
	NQLID       string        `json:"nqlId"`
	Type        string        `json:"type"`
	Number      int           `json:"number"`
	Text        string        `json:"text"`
	Choices     []ChoiceInput `json:"choices,omitempty"`
	Comment     string        `json:"comment"`
	DefaultNext *string       `json:"defaultNext,omitempty"`
}
type Question struct {
	ID          string   `json:"id"`
	NQLID       string   `json:"nqlId"`
	Type        string   `json:"type"`
	Number      int      `json:"number"`
	Text        string   `json:"text"`
	Choices     []Choice `json:"choices"`
	Comment     string   `json:"comment"`
	DefaultNext *string  `json:"defaultNext"`
}
type ChoiceInput struct {
	Label string          `json:"label"`
	Text  string          `json:"text"`
	Value json.RawMessage `json:"value,omitempty"`
	Next  *string         `json:"next,omitempty"`
}
type MultiLanguageInfo struct {
	DefaultLanguage     string   `json:"defaultLanguage"`
	AdditionalLanguages []string `json:"additionalLanguages"`
	BaseLanguage        string   `json:"baseLanguage"`
}
type Translation struct {
	Language  string                `json:"language"`
	Status    string                `json:"status"`
	Questions []QuestionTranslation `json:"questions"`
}
type QuestionTranslation struct {
	ID      string              `json:"id"`
	Comment string              `json:"comment"`
	Text    string              `json:"text"`
	Choices []ChoiceTranslation `json:"choices"`
}
type ChoiceTranslation struct {
	Text string `json:"text"`
}
type Parameter struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type CreateRequest struct {
	Campaign CampaignInput   `json:"campaign"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
}
type EditCampaignInput struct {
	CampaignInput
	ContentID string `json:"contentId"`
	BCSUID    string `json:"bcsUid"`
	Status    string `json:"status"`
}
type UpdateRequest struct {
	Campaign        EditCampaignInput `json:"campaign"`
	RemoveQuestions []string          `json:"removeQuestions,omitempty"`
	AddQuestions    []QuestionInput   `json:"addQuestions,omitempty"`
	Owner           *string           `json:"owner,omitempty"`
}
type ListOptions struct {
	PageNumber *int
	Offset     *int
}
type Summary struct {
	Name      string `json:"name"`
	NQLID     string `json:"nqlId"`
	BCSUID    string `json:"bcsUid"`
	ContentID string `json:"contentId"`
}
type ListResponse struct {
	Campaigns []Summary `json:"listCampaignDocs"`
}
type GetResponse struct {
	Campaign *Campaign `json:"campaignDoc"`
}
type CreateResponse struct {
	Campaign *Campaign `json:"createCampaign"`
}
type UpdateResponse struct {
	Campaign *Campaign `json:"editCampaign"`
}
type DeleteResponse struct {
	ContentID *string `json:"deleteCampaignDoc"`
}

type Choice struct {
	Label string          `json:"label"`
	Text  string          `json:"text"`
	Value json.RawMessage `json:"value"`
	Next  *string         `json:"next"`
}
