package snapshots

import "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"

// Definition is also the browser's portable JSON export. Retention approval must be explicit.
type Definition struct {
	Name                          string `json:"name"`
	Description                   string `json:"description"`
	NQLName                       string `json:"nqlName"`
	NQLQuery                      string `json:"nqlQuery"`
	Origin                        string `json:"origin"`
	ExtendedDataRetentionApproved bool   `json:"extendedDataRetentionApproved"`
	SnapshotVersion               int    `json:"snapshotVersion"`
}
type Snapshot struct {
	Definition
	ContentID         string `json:"contentId"`
	BCSRevisionNumber int    `json:"bcsRevisionNumber"`
}
type ListRow struct {
	content_administration.Content
	NQLID string `json:"nqlId"`
}
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []ListRow                          `json:"rows"`
}
