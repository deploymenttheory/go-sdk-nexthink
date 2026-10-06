package action_executions

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
func (v *ListRemoteActionsResponseItemTargeting) UnmarshalJSON(data []byte) error {
	type wire ListRemoteActionsResponseItemTargeting
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListRemoteActionsResponseItemTargeting(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListRemoteActionsResponseItemTargeting) MarshalJSON() ([]byte, error) {
	type wire ListRemoteActionsResponseItemTargeting
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListRemoteActionsResponseItemScriptInfo) UnmarshalJSON(data []byte) error {
	type wire ListRemoteActionsResponseItemScriptInfo
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListRemoteActionsResponseItemScriptInfo(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListRemoteActionsResponseItemScriptInfo) MarshalJSON() ([]byte, error) {
	type wire ListRemoteActionsResponseItemScriptInfo
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListRemoteActionsResponseItem) UnmarshalJSON(data []byte) error {
	type wire ListRemoteActionsResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListRemoteActionsResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListRemoteActionsResponseItem) MarshalJSON() ([]byte, error) {
	type wire ListRemoteActionsResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetRemoteActionResponseTargeting) UnmarshalJSON(data []byte) error {
	type wire GetRemoteActionResponseTargeting
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetRemoteActionResponseTargeting(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetRemoteActionResponseTargeting) MarshalJSON() ([]byte, error) {
	type wire GetRemoteActionResponseTargeting
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetRemoteActionResponseScriptInfo) UnmarshalJSON(data []byte) error {
	type wire GetRemoteActionResponseScriptInfo
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetRemoteActionResponseScriptInfo(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetRemoteActionResponseScriptInfo) MarshalJSON() ([]byte, error) {
	type wire GetRemoteActionResponseScriptInfo
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetRemoteActionResponse) UnmarshalJSON(data []byte) error {
	type wire GetRemoteActionResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetRemoteActionResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetRemoteActionResponse) MarshalJSON() ([]byte, error) {
	type wire GetRemoteActionResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListRemoteActionsForQueryResponseItemTargeting) UnmarshalJSON(data []byte) error {
	type wire ListRemoteActionsForQueryResponseItemTargeting
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListRemoteActionsForQueryResponseItemTargeting(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListRemoteActionsForQueryResponseItemTargeting) MarshalJSON() ([]byte, error) {
	type wire ListRemoteActionsForQueryResponseItemTargeting
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListRemoteActionsForQueryResponseItemScriptInfo) UnmarshalJSON(data []byte) error {
	type wire ListRemoteActionsForQueryResponseItemScriptInfo
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListRemoteActionsForQueryResponseItemScriptInfo(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListRemoteActionsForQueryResponseItemScriptInfo) MarshalJSON() ([]byte, error) {
	type wire ListRemoteActionsForQueryResponseItemScriptInfo
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListRemoteActionsForQueryResponseItem) UnmarshalJSON(data []byte) error {
	type wire ListRemoteActionsForQueryResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListRemoteActionsForQueryResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListRemoteActionsForQueryResponseItem) MarshalJSON() ([]byte, error) {
	type wire ListRemoteActionsForQueryResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ExecuteResponse) UnmarshalJSON(data []byte) error {
	type wire ExecuteResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ExecuteResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ExecuteResponse) MarshalJSON() ([]byte, error) {
	type wire ExecuteResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListActionsResponseItem) UnmarshalJSON(data []byte) error {
	type wire ListActionsResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListActionsResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListActionsResponseItem) MarshalJSON() ([]byte, error) {
	type wire ListActionsResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDeviceHistoryResponseItemExecutionDetailsItemStatusDetails) UnmarshalJSON(data []byte) error {
	type wire GetDeviceHistoryResponseItemExecutionDetailsItemStatusDetails
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDeviceHistoryResponseItemExecutionDetailsItemStatusDetails(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDeviceHistoryResponseItemExecutionDetailsItemStatusDetails) MarshalJSON() ([]byte, error) {
	type wire GetDeviceHistoryResponseItemExecutionDetailsItemStatusDetails
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDeviceHistoryResponseItemExecutionDetailsItemDeviceID) UnmarshalJSON(data []byte) error {
	type wire GetDeviceHistoryResponseItemExecutionDetailsItemDeviceID
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDeviceHistoryResponseItemExecutionDetailsItemDeviceID(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDeviceHistoryResponseItemExecutionDetailsItemDeviceID) MarshalJSON() ([]byte, error) {
	type wire GetDeviceHistoryResponseItemExecutionDetailsItemDeviceID
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDeviceHistoryResponseItemExecutionDetailsItem) UnmarshalJSON(data []byte) error {
	type wire GetDeviceHistoryResponseItemExecutionDetailsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDeviceHistoryResponseItemExecutionDetailsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDeviceHistoryResponseItemExecutionDetailsItem) MarshalJSON() ([]byte, error) {
	type wire GetDeviceHistoryResponseItemExecutionDetailsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDeviceHistoryResponseItem) UnmarshalJSON(data []byte) error {
	type wire GetDeviceHistoryResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDeviceHistoryResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDeviceHistoryResponseItem) MarshalJSON() ([]byte, error) {
	type wire GetDeviceHistoryResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *RemoteActionInput) UnmarshalJSON(data []byte) error {
	type wire RemoteActionInput
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = RemoteActionInput(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v RemoteActionInput) MarshalJSON() ([]byte, error) {
	type wire RemoteActionInput
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *RemoteActionOutput) UnmarshalJSON(data []byte) error {
	type wire RemoteActionOutput
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = RemoteActionOutput(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v RemoteActionOutput) MarshalJSON() ([]byte, error) {
	type wire RemoteActionOutput
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *RemoteActionTargetingEntity) UnmarshalJSON(data []byte) error {
	type wire RemoteActionTargetingEntity
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = RemoteActionTargetingEntity(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v RemoteActionTargetingEntity) MarshalJSON() ([]byte, error) {
	type wire RemoteActionTargetingEntity
	return encodeResponse(wire(v), v.AdditionalFields)
}
