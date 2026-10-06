package workflow_executions

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
func (v *GetTimelineResponseItem) UnmarshalJSON(data []byte) error {
	type wire GetTimelineResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetTimelineResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetTimelineResponseItem) MarshalJSON() ([]byte, error) {
	type wire GetTimelineResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetTimelineV2ResponseItem) UnmarshalJSON(data []byte) error {
	type wire GetTimelineV2ResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetTimelineV2ResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetTimelineV2ResponseItem) MarshalJSON() ([]byte, error) {
	type wire GetTimelineV2ResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListActivitiesResponseItem) UnmarshalJSON(data []byte) error {
	type wire ListActivitiesResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListActivitiesResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListActivitiesResponseItem) MarshalJSON() ([]byte, error) {
	type wire ListActivitiesResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetRemoteActionDetailsResponse) UnmarshalJSON(data []byte) error {
	type wire GetRemoteActionDetailsResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetRemoteActionDetailsResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetRemoteActionDetailsResponse) MarshalJSON() ([]byte, error) {
	type wire GetRemoteActionDetailsResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCustomFieldsDetailsResponse) UnmarshalJSON(data []byte) error {
	type wire GetCustomFieldsDetailsResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCustomFieldsDetailsResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCustomFieldsDetailsResponse) MarshalJSON() ([]byte, error) {
	type wire GetCustomFieldsDetailsResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCampaignDetailsResponse) UnmarshalJSON(data []byte) error {
	type wire GetCampaignDetailsResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCampaignDetailsResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCampaignDetailsResponse) MarshalJSON() ([]byte, error) {
	type wire GetCampaignDetailsResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetFunctionDetailsResponse) UnmarshalJSON(data []byte) error {
	type wire GetFunctionDetailsResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetFunctionDetailsResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetFunctionDetailsResponse) MarshalJSON() ([]byte, error) {
	type wire GetFunctionDetailsResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetMessageDetailsResponse) UnmarshalJSON(data []byte) error {
	type wire GetMessageDetailsResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetMessageDetailsResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetMessageDetailsResponse) MarshalJSON() ([]byte, error) {
	type wire GetMessageDetailsResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSAPIDetailsResponse) UnmarshalJSON(data []byte) error {
	type wire GetSAPIDetailsResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSAPIDetailsResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSAPIDetailsResponse) MarshalJSON() ([]byte, error) {
	type wire GetSAPIDetailsResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListWorkflowsResponseItemTriggerMethods) UnmarshalJSON(data []byte) error {
	type wire ListWorkflowsResponseItemTriggerMethods
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListWorkflowsResponseItemTriggerMethods(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListWorkflowsResponseItemTriggerMethods) MarshalJSON() ([]byte, error) {
	type wire ListWorkflowsResponseItemTriggerMethods
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListWorkflowsResponseItemVersionsItemParametersItem) UnmarshalJSON(data []byte) error {
	type wire ListWorkflowsResponseItemVersionsItemParametersItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListWorkflowsResponseItemVersionsItemParametersItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListWorkflowsResponseItemVersionsItemParametersItem) MarshalJSON() ([]byte, error) {
	type wire ListWorkflowsResponseItemVersionsItemParametersItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListWorkflowsResponseItemVersionsItem) UnmarshalJSON(data []byte) error {
	type wire ListWorkflowsResponseItemVersionsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListWorkflowsResponseItemVersionsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListWorkflowsResponseItemVersionsItem) MarshalJSON() ([]byte, error) {
	type wire ListWorkflowsResponseItemVersionsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListWorkflowsResponseItem) UnmarshalJSON(data []byte) error {
	type wire ListWorkflowsResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListWorkflowsResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListWorkflowsResponseItem) MarshalJSON() ([]byte, error) {
	type wire ListWorkflowsResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetWorkflowResponseTriggerMethods) UnmarshalJSON(data []byte) error {
	type wire GetWorkflowResponseTriggerMethods
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetWorkflowResponseTriggerMethods(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetWorkflowResponseTriggerMethods) MarshalJSON() ([]byte, error) {
	type wire GetWorkflowResponseTriggerMethods
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetWorkflowResponseVersionsItemParametersItem) UnmarshalJSON(data []byte) error {
	type wire GetWorkflowResponseVersionsItemParametersItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetWorkflowResponseVersionsItemParametersItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetWorkflowResponseVersionsItemParametersItem) MarshalJSON() ([]byte, error) {
	type wire GetWorkflowResponseVersionsItemParametersItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetWorkflowResponseVersionsItem) UnmarshalJSON(data []byte) error {
	type wire GetWorkflowResponseVersionsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetWorkflowResponseVersionsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetWorkflowResponseVersionsItem) MarshalJSON() ([]byte, error) {
	type wire GetWorkflowResponseVersionsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetWorkflowResponse) UnmarshalJSON(data []byte) error {
	type wire GetWorkflowResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetWorkflowResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetWorkflowResponse) MarshalJSON() ([]byte, error) {
	type wire GetWorkflowResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetResponse) UnmarshalJSON(data []byte) error {
	type wire GetResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetResponse) MarshalJSON() ([]byte, error) {
	type wire GetResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetHistoryResponseItem) UnmarshalJSON(data []byte) error {
	type wire GetHistoryResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetHistoryResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetHistoryResponseItem) MarshalJSON() ([]byte, error) {
	type wire GetHistoryResponseItem
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
func (v *ExecuteNQLResponse) UnmarshalJSON(data []byte) error {
	type wire ExecuteNQLResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ExecuteNQLResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ExecuteNQLResponse) MarshalJSON() ([]byte, error) {
	type wire ExecuteNQLResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
