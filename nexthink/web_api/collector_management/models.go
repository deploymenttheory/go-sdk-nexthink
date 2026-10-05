package collector_management

import "encoding/json"

type DownloadLinks struct {
	Windows      *InstallerDownloadLinkData `json:"windowsInstallerDownloadLinkData"`
	MacOS        *InstallerDownloadLinkData `json:"macOsInstallerDownloadLinkData"`
	VDIExtension *InstallerDownloadLinkData `json:"vdiExtInstallerDownloadLinkData"`
}
type UpdateConfiguration struct {
	ConfigRevision     json.RawMessage `json:"configRevision"`
	BCSConfig          json.RawMessage `json:"bcsConfig"`
	ProductConfigFlags json.RawMessage `json:"productConfigFlags"`
}

// QueryValue is an inventory facet returned by the updater query endpoints.
type QueryValue struct {
	Value string `json:"value"`
	Count int    `json:"count"`
}

// InstallerDownloadLinkData contains short-lived download links. Do not log URLs.
type InstallerDownloadLinkData struct {
	PlatformName  string `json:"platformName"`
	Version       string `json:"version"`
	FileURL       string `json:"fileUrl"`
	SignatureName string `json:"signatureName"`
	SignatureURL  string `json:"signatureUrl"`
}
