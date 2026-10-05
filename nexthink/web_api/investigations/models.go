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
