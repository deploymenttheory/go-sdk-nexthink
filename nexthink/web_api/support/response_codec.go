package support

import (
	"encoding/json"
	"reflect"
	"strings"
)

// decodeResponse retains unmodeled feature fields without replacing typed known properties.
func decodeResponse(data []byte, value any) (map[string]json.RawMessage, error) {
	if err := json.Unmarshal(data, value); err != nil {
		return nil, err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return nil, err
	}
	t := reflect.TypeOf(value).Elem()
	for i := 0; i < t.NumField(); i++ {
		key := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
		if key != "-" && string(extra[key]) != "null" {
			delete(extra, key)
		}
	}
	return extra, nil
}
func encodeResponse(value any, extra map[string]json.RawMessage) ([]byte, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	if len(extra) == 0 {
		return data, nil
	}
	var result map[string]json.RawMessage
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	t := reflect.TypeOf(value)
	known := map[string]bool{}
	for i := 0; i < t.NumField(); i++ {
		key := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
		known[key] = true
	}
	for key, val := range extra {
		_, present := result[key]
		if !known[key] || (!present && string(val) == "null") {
			result[key] = val
		}
	}
	return json.Marshal(result)
}
func (v *GetProfileResponseName) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseName
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseName(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseName) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseName
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseModel) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseModel
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseModel(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseModel) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseModel
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseManufacturer) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseManufacturer
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseManufacturer(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseManufacturer) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseManufacturer
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseOSName) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseOSName
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseOSName(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseOSName) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseOSName
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseOSBuild) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseOSBuild
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseOSBuild(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseOSBuild) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseOSBuild
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseHardwareType) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseHardwareType
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseHardwareType(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseHardwareType) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseHardwareType
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseMemory) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseMemory
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseMemory(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseMemory) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseMemory
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseLastSeen) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseLastSeen
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseLastSeen(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseLastSeen) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseLastSeen
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseLastIPAddress) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseLastIPAddress
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseLastIPAddress(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseLastIPAddress) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseLastIPAddress
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseLocationType) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseLocationType
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseLocationType(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseLocationType) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseLocationType
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseLocalIps) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseLocalIps
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseLocalIps(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseLocalIps) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseLocalIps
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponseLastConnectionType) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponseLastConnectionType
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponseLastConnectionType(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponseLastConnectionType) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponseLastConnectionType
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetProfileResponse) UnmarshalJSON(data []byte) error {
	type wire GetProfileResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetProfileResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetProfileResponse) MarshalJSON() ([]byte, error) {
	type wire GetProfileResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPlatformResponse) UnmarshalJSON(data []byte) error {
	type wire GetPlatformResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPlatformResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPlatformResponse) MarshalJSON() ([]byte, error) {
	type wire GetPlatformResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListUsersResponseUsersItemUsername) UnmarshalJSON(data []byte) error {
	type wire ListUsersResponseUsersItemUsername
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListUsersResponseUsersItemUsername(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListUsersResponseUsersItemUsername) MarshalJSON() ([]byte, error) {
	type wire ListUsersResponseUsersItemUsername
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListUsersResponseUsersItemTypeNQLData) UnmarshalJSON(data []byte) error {
	type wire ListUsersResponseUsersItemTypeNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListUsersResponseUsersItemTypeNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListUsersResponseUsersItemTypeNQLData) MarshalJSON() ([]byte, error) {
	type wire ListUsersResponseUsersItemTypeNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListUsersResponseUsersItemType) UnmarshalJSON(data []byte) error {
	type wire ListUsersResponseUsersItemType
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListUsersResponseUsersItemType(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListUsersResponseUsersItemType) MarshalJSON() ([]byte, error) {
	type wire ListUsersResponseUsersItemType
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListUsersResponseUsersItemFullName) UnmarshalJSON(data []byte) error {
	type wire ListUsersResponseUsersItemFullName
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListUsersResponseUsersItemFullName(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListUsersResponseUsersItemFullName) MarshalJSON() ([]byte, error) {
	type wire ListUsersResponseUsersItemFullName
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListUsersResponseUsersItemLastSeen) UnmarshalJSON(data []byte) error {
	type wire ListUsersResponseUsersItemLastSeen
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListUsersResponseUsersItemLastSeen(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListUsersResponseUsersItemLastSeen) MarshalJSON() ([]byte, error) {
	type wire ListUsersResponseUsersItemLastSeen
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListUsersResponseUsersItem) UnmarshalJSON(data []byte) error {
	type wire ListUsersResponseUsersItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListUsersResponseUsersItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListUsersResponseUsersItem) MarshalJSON() ([]byte, error) {
	type wire ListUsersResponseUsersItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListUsersResponse) UnmarshalJSON(data []byte) error {
	type wire ListUsersResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListUsersResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListUsersResponse) MarshalJSON() ([]byte, error) {
	type wire ListUsersResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *SearchResponse) UnmarshalJSON(data []byte) error {
	type wire SearchResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = SearchResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v SearchResponse) MarshalJSON() ([]byte, error) {
	type wire SearchResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
