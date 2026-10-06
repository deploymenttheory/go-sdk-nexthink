package library

import "encoding/json"

// Content describes a built-in Library item. Resource-specific dependency types are map keys.
type Content struct {
	LibraryUUID           string                   `json:"libraryUuid"`
	MinMTPVersion         *string                  `json:"minMtpVersion,omitempty"`
	LatestVersion         *Version                 `json:"latestVersion,omitempty"`
	VersionHistory        *[]Version               `json:"versionHistory,omitempty"`
	Name                  string                   `json:"name"`
	Description           *string                  `json:"description,omitempty"`
	Type                  string                   `json:"type"`
	FileName              *string                  `json:"fileName,omitempty"`
	IsInCustomPack        bool                     `json:"isInCustomPack"`
	IsInstalled           bool                     `json:"isInstalled"`
	LibraryPacks          *[]PackReference         `json:"libraryPacks,omitempty"`
	Images                *[]json.RawMessage       `json:"images,omitempty"`
	Licenses              *[]json.RawMessage       `json:"licenses,omitempty"`
	LicenseIDs            *[]LicenseReference      `json:"licenseIds,omitempty"`
	SystemContent         bool                     `json:"systemContent"`
	ProductArea           *string                  `json:"productArea,omitempty"`
	IsNewContent          bool                     `json:"isNewContent"`
	Dependencies          *map[string][]Dependency `json:"dependencies,omitempty"`
	CurrentVersion        *string                  `json:"currentVersion,omitempty"`
	ContentID             *string                  `json:"contentId,omitempty"`
	ResourceName          *string                  `json:"resourceName,omitempty"`
	IsUpdateAvailable     *bool                    `json:"isUpdateAvailable,omitempty"`
	UpdateAvailable       *bool                    `json:"updateAvailable,omitempty"`
	HasPermissionToImport *bool                    `json:"hasPermissionToImport,omitempty"`
}
type Version struct {
	Version *string `json:"version,omitempty"`
	Date    string  `json:"date"`
	Comment string  `json:"comment"`
}
type PackReference struct {
	Name        string `json:"name"`
	RedirectURL string `json:"redirectURL"`
}
type ContentsResponse struct {
	Data []Content `json:"data"`
}

// Pack includes both list and detail fields; optional detail fields are omitted from list responses.
type Pack struct {
	PackUUID               string                   `json:"packUuid"`
	Name                   string                   `json:"name"`
	ShortDescription       string                   `json:"shortDescription"`
	FileName               string                   `json:"fileName"`
	InstallationState      string                   `json:"installationState"`
	IsUpdateAvailable      bool                     `json:"isUpdateAvailable"`
	IsNewPack              bool                     `json:"isNewPack"`
	IsCustomPack           bool                     `json:"isCustomPack"`
	HasPermissionToInstall bool                     `json:"hasPermissionToInstall"`
	InstalledCount         int                      `json:"installedCount"`
	TotalCount             int                      `json:"totalCount"`
	LatestVersion          Version                  `json:"latestVersion"`
	BuiltinContent         map[string][]PackContent `json:"builtinContent"`
	VersionHistoryCount    *int                     `json:"versionHistoryCount,omitempty"`
	MinMTPVersion          *string                  `json:"minMtpVersion,omitempty"`
	VersionHistory         []Version                `json:"versionHistory,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	Images                 *[]json.RawMessage       `json:"images,omitempty"`
	CurrentVersion         *string                  `json:"currentVersion,omitempty"`
}
type PackContent struct {
	LibraryUUID       string  `json:"libraryUuid"`
	Name              string  `json:"name"`
	FileName          string  `json:"fileName"`
	Installed         bool    `json:"installed"`
	IsUpdateAvailable bool    `json:"isUpdateAvailable"`
	ContentID         *string `json:"contentId,omitempty"`
	ResourceName      *string `json:"resourceName,omitempty"`
	CurrentVersion    *string `json:"currentVersion,omitempty"`
}
type PacksResponse struct {
	Packs []Pack `json:"packs"`
}
type LocaleResponse struct {
	Locale string `json:"locale"`
}
type CreateCopyInfo struct {
	KubernetesServiceName string          `json:"k8s_service_name"`
	Port                  string          `json:"port"`
	BCSTypeName           string          `json:"bcs_type_name"`
	BCSSchemaName         string          `json:"bcs_schema_name"`
	BuiltinContent        json.RawMessage `json:"builtin_content"`
	ProductArea           string          `json:"product_area"`
	SupportSystemContent  bool            `json:"support_system_content"`
}

// ContentInstallationRequest installs a standard built-in item. Custom content uses InstallCustomContent.
type ContentInstallationRequest struct {
	FileName string `json:"fileName"`
}
type ContentUpdateRequest struct {
	ContentID string `json:"contentId"`
	FileName  string `json:"fileName"`
}
type Dependency struct {
	LibraryUUID      string `json:"libraryUuid"`
	Name             string `json:"name"`
	FileName         string `json:"fileName,omitempty"`
	InstalledFlag    *bool  `json:"installedFlag,omitempty"`
	TaggedForInstall *bool  `json:"taggedForInstall,omitempty"`
}
type DependenciesRequest struct {
	Dependencies map[string][]Dependency `json:"dependencies"`
}
type DependenciesResponse struct {
	Dependencies map[string][]Dependency `json:"dependencies,omitempty"`
}

// OperationResult preserves resource-dependent import/update response bodies. Empty responses remain nil.
type OperationResult = json.RawMessage

// InstallationStatus is the asynchronous custom pack/content installation state observed in the UI.
type InstallationStatus struct {
	IsInstalling bool                `json:"isInstalling"`
	Status       string              `json:"status"`
	Errors       []InstallationError `json:"errors"`
}
type InstallationError struct {
	ErrorCode      string          `json:"errorCode"`
	ErrorDetails   json.RawMessage `json:"errorDetails"`
	ErrorReference json.RawMessage `json:"errorReference"`
}

type LicenseReference struct {
	Name string `json:"name"`
}
