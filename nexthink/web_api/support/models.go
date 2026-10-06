package support

import "encoding/json"

// Polymorphic NQL values, event payloads and plugin-defined details retain their exact JSON.
type GetProfileResponseName struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseModel struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseManufacturer struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseOSName struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseOSBuild struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseHardwareType struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseMemory struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseLastSeen struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseLastIPAddress struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseLocationType struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseLocalIps struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponseLastConnectionType struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type GetProfileResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields   map[string]json.RawMessage            `json:"-"`
	Name               *GetProfileResponseName               `json:"name,omitempty"`
	Model              *GetProfileResponseModel              `json:"model,omitempty"`
	Manufacturer       *GetProfileResponseManufacturer       `json:"manufacturer,omitempty"`
	OSName             *GetProfileResponseOSName             `json:"osName,omitempty"`
	OSBuild            *GetProfileResponseOSBuild            `json:"osBuild,omitempty"`
	HardwareType       *GetProfileResponseHardwareType       `json:"hardwareType,omitempty"`
	Memory             *GetProfileResponseMemory             `json:"memory,omitempty"`
	LastSeen           *GetProfileResponseLastSeen           `json:"lastSeen,omitempty"`
	LastIPAddress      *GetProfileResponseLastIPAddress      `json:"lastIpAddress,omitempty"`
	LocationType       *GetProfileResponseLocationType       `json:"locationType,omitempty"`
	LocalIps           *GetProfileResponseLocalIps           `json:"localIps,omitempty"`
	LastConnectionType *GetProfileResponseLastConnectionType `json:"lastConnectionType,omitempty"`
}
type GetPlatformResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Platform         *string                    `json:"platform,omitempty"`
}
type ListUsersResponseUsersItemUsername struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type ListUsersResponseUsersItemTypeNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type ListUsersResponseUsersItemType struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage             `json:"-"`
	NQLData          *ListUsersResponseUsersItemTypeNQLData `json:"nqlData,omitempty"`
	ComputedValue    json.RawMessage                        `json:"computedValue,omitempty"`
}
type ListUsersResponseUsersItemFullName struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type ListUsersResponseUsersItemLastSeen struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
}
type ListUsersResponseUsersItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage          `json:"-"`
	ID               *string                             `json:"id,omitempty"`
	Username         *ListUsersResponseUsersItemUsername `json:"username,omitempty"`
	Type             *ListUsersResponseUsersItemType     `json:"type,omitempty"`
	FullName         *ListUsersResponseUsersItemFullName `json:"fullName,omitempty"`
	LastSeen         *ListUsersResponseUsersItemLastSeen `json:"lastSeen,omitempty"`
}
type ListUsersResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage    `json:"-"`
	URL              *string                       `json:"url,omitempty"`
	Users            *[]ListUsersResponseUsersItem `json:"users,omitempty"`
}
type SearchResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields  map[string]json.RawMessage `json:"-"`
	MatchingDeviceIds *[]json.RawMessage         `json:"matchingDeviceIds,omitempty"`
}

type SearchRequest struct {
	Name string `json:"name"`
}
