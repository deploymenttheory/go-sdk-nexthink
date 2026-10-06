package support_timeline

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
func (v *GetAlertsAndErrorsResponse) UnmarshalJSON(data []byte) error {
	type wire GetAlertsAndErrorsResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetAlertsAndErrorsResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetAlertsAndErrorsResponse) MarshalJSON() ([]byte, error) {
	type wire GetAlertsAndErrorsResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseCPUUsageResultTimeSeriesItemDataUsage) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseCPUUsageResultTimeSeriesItemDataUsage
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseCPUUsageResultTimeSeriesItemDataUsage(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseCPUUsageResultTimeSeriesItemDataUsage) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseCPUUsageResultTimeSeriesItemDataUsage
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseCPUUsageResultTimeSeriesItemDataHighUsageBinaries) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseCPUUsageResultTimeSeriesItemDataHighUsageBinaries
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseCPUUsageResultTimeSeriesItemDataHighUsageBinaries(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseCPUUsageResultTimeSeriesItemDataHighUsageBinaries) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseCPUUsageResultTimeSeriesItemDataHighUsageBinaries
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseCPUUsageResultTimeSeriesItemData) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseCPUUsageResultTimeSeriesItemData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseCPUUsageResultTimeSeriesItemData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseCPUUsageResultTimeSeriesItemData) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseCPUUsageResultTimeSeriesItemData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseCPUUsageResultTimeSeriesItem) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseCPUUsageResultTimeSeriesItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseCPUUsageResultTimeSeriesItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseCPUUsageResultTimeSeriesItem) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseCPUUsageResultTimeSeriesItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseCPUUsageResult) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseCPUUsageResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseCPUUsageResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseCPUUsageResult) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseCPUUsageResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseCPUUsage) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseCPUUsage
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseCPUUsage(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseCPUUsage) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseCPUUsage
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataUsage) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataUsage
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataUsage(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataUsage) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataUsage
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataInstalled) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataInstalled
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataInstalled(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataInstalled) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataInstalled
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemorySwapRate) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemorySwapRate
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemorySwapRate(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemorySwapRate) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemorySwapRate
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDiskQueueLength) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDiskQueueLength
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDiskQueueLength(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDiskQueueLength) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDiskQueueLength
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemoryPressure) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemoryPressure
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemoryPressure(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemoryPressure) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemoryPressure
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationHighMemoryPressure) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationHighMemoryPressure
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationHighMemoryPressure(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationHighMemoryPressure) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationHighMemoryPressure
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationMediumMemoryPressure) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationMediumMemoryPressure
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationMediumMemoryPressure(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationMediumMemoryPressure) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationMediumMemoryPressure
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemBinaryName) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemBinaryName
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemBinaryName(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemBinaryName) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemBinaryName
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemUsage) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemUsage
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemUsage(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemUsage) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemUsage
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItem) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItem) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinaries) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinaries
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinaries(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinaries) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinaries
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItemData) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItemData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItemData) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItemData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResultTimeSeriesItem) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResultTimeSeriesItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResultTimeSeriesItem) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResultTimeSeriesItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsageResult) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsageResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsageResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsageResult) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsageResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseMemoryUsage) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseMemoryUsage
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseMemoryUsage(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseMemoryUsage) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseMemoryUsage
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseGPUSlotOne) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseGPUSlotOne
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseGPUSlotOne(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseGPUSlotOne) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseGPUSlotOne
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseGPUSlotTwo) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseGPUSlotTwo
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseGPUSlotTwo(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseGPUSlotTwo) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseGPUSlotTwo
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseDiskResult) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseDiskResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseDiskResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseDiskResult) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseDiskResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseDisk) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseDisk
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseDisk(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseDisk) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseDisk
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseDriveSpaceResult) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseDriveSpaceResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseDriveSpaceResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseDriveSpaceResult) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseDriveSpaceResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseDriveSpace) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseDriveSpace
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseDriveSpace(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseDriveSpace) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseDriveSpace
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseNPUResult) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseNPUResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseNPUResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseNPUResult) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseNPUResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponseNPU) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponseNPU
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponseNPU(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponseNPU) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponseNPU
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetPerformanceResponse) UnmarshalJSON(data []byte) error {
	type wire GetPerformanceResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetPerformanceResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetPerformanceResponse) MarshalJSON() ([]byte, error) {
	type wire GetPerformanceResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseWifiResult) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseWifiResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseWifiResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseWifiResult) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseWifiResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseWifi) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseWifi
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseWifi(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseWifi) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseWifi
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseEthernetResultTimeSeriesItemDataLocalIPAddressesItem) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseEthernetResultTimeSeriesItemDataLocalIPAddressesItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseEthernetResultTimeSeriesItemDataLocalIPAddressesItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseEthernetResultTimeSeriesItemDataLocalIPAddressesItem) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseEthernetResultTimeSeriesItemDataLocalIPAddressesItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseEthernetResultTimeSeriesItemDataMacAddressItem) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseEthernetResultTimeSeriesItemDataMacAddressItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseEthernetResultTimeSeriesItemDataMacAddressItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseEthernetResultTimeSeriesItemDataMacAddressItem) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseEthernetResultTimeSeriesItemDataMacAddressItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseEthernetResultTimeSeriesItemData) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseEthernetResultTimeSeriesItemData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseEthernetResultTimeSeriesItemData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseEthernetResultTimeSeriesItemData) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseEthernetResultTimeSeriesItemData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseEthernetResultTimeSeriesItem) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseEthernetResultTimeSeriesItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseEthernetResultTimeSeriesItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseEthernetResultTimeSeriesItem) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseEthernetResultTimeSeriesItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseEthernetResult) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseEthernetResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseEthernetResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseEthernetResult) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseEthernetResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseEthernet) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseEthernet
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseEthernet(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseEthernet) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseEthernet
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseBluetoothResult) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseBluetoothResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseBluetoothResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseBluetoothResult) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseBluetoothResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseBluetooth) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseBluetooth
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseBluetooth(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseBluetooth) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseBluetooth
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsFailedConnectionRatio) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsFailedConnectionRatio
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsFailedConnectionRatio(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsFailedConnectionRatio) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsFailedConnectionRatio
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsConnectionEstablishmentTime) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsConnectionEstablishmentTime
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsConnectionEstablishmentTime(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsConnectionEstablishmentTime) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsConnectionEstablishmentTime
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnections) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnections
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnections(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnections) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnections
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsNumberOfConnections) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsNumberOfConnections
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsNumberOfConnections(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsNumberOfConnections) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsNumberOfConnections
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsIncomingTraffic) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsIncomingTraffic
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsIncomingTraffic(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsIncomingTraffic) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsIncomingTraffic
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsOutgoingTraffic) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsOutgoingTraffic
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsOutgoingTraffic(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsOutgoingTraffic) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsOutgoingTraffic
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnections) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnections
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnections(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnections) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnections
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnectionsResultTimeSeriesItemData) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnectionsResultTimeSeriesItemData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnectionsResultTimeSeriesItemData) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItemData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnectionsResultTimeSeriesItem) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnectionsResultTimeSeriesItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnectionsResultTimeSeriesItem) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnectionsResultTimeSeriesItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnectionsResult) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnectionsResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnectionsResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnectionsResult) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnectionsResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseConnections) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseConnections
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseConnections(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseConnections) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseConnections
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseVpnResult) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseVpnResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseVpnResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseVpnResult) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseVpnResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseVpn) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseVpn
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseVpn(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseVpn) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseVpn
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseDesktopApplications) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseDesktopApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseDesktopApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseDesktopApplications) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseDesktopApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseWebApplications) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseWebApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseWebApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseWebApplications) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseWebApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponseNetworkApplications) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponseNetworkApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponseNetworkApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponseNetworkApplications) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponseNetworkApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectivityResponse) UnmarshalJSON(data []byte) error {
	type wire GetConnectivityResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectivityResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectivityResponse) MarshalJSON() ([]byte, error) {
	type wire GetConnectivityResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetApplicationConnectivityResponseDesktopApplications) UnmarshalJSON(data []byte) error {
	type wire GetApplicationConnectivityResponseDesktopApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetApplicationConnectivityResponseDesktopApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetApplicationConnectivityResponseDesktopApplications) MarshalJSON() ([]byte, error) {
	type wire GetApplicationConnectivityResponseDesktopApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetApplicationConnectivityResponseWebApplications) UnmarshalJSON(data []byte) error {
	type wire GetApplicationConnectivityResponseWebApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetApplicationConnectivityResponseWebApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetApplicationConnectivityResponseWebApplications) MarshalJSON() ([]byte, error) {
	type wire GetApplicationConnectivityResponseWebApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetApplicationConnectivityResponseNetworkApplications) UnmarshalJSON(data []byte) error {
	type wire GetApplicationConnectivityResponseNetworkApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetApplicationConnectivityResponseNetworkApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetApplicationConnectivityResponseNetworkApplications) MarshalJSON() ([]byte, error) {
	type wire GetApplicationConnectivityResponseNetworkApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetApplicationConnectivityResponse) UnmarshalJSON(data []byte) error {
	type wire GetApplicationConnectivityResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetApplicationConnectivityResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetApplicationConnectivityResponse) MarshalJSON() ([]byte, error) {
	type wire GetApplicationConnectivityResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponseInstallationEventsResult) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponseInstallationEventsResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponseInstallationEventsResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponseInstallationEventsResult) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponseInstallationEventsResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponseInstallationEvents) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponseInstallationEvents
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponseInstallationEvents(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponseInstallationEvents) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponseInstallationEvents
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBootsEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBootsEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBootsEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBootsEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBootsEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBoots) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBoots
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBoots(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBoots) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBoots
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponseSystemEventsResultTimeSeriesItemData) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponseSystemEventsResultTimeSeriesItemData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponseSystemEventsResultTimeSeriesItemData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponseSystemEventsResultTimeSeriesItemData) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponseSystemEventsResultTimeSeriesItemData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponseSystemEventsResultTimeSeriesItem) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponseSystemEventsResultTimeSeriesItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponseSystemEventsResultTimeSeriesItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponseSystemEventsResultTimeSeriesItem) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponseSystemEventsResultTimeSeriesItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponseSystemEventsResult) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponseSystemEventsResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponseSystemEventsResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponseSystemEventsResult) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponseSystemEventsResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponseSystemEvents) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponseSystemEvents
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponseSystemEvents(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponseSystemEvents) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponseSystemEvents
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponseActionsResult) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponseActionsResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponseActionsResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponseActionsResult) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponseActionsResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponseActions) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponseActions
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponseActions(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponseActions) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponseActions
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActivityResponse) UnmarshalJSON(data []byte) error {
	type wire GetActivityResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActivityResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActivityResponse) MarshalJSON() ([]byte, error) {
	type wire GetActivityResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetApplicationsResponseDesktopApplications) UnmarshalJSON(data []byte) error {
	type wire GetApplicationsResponseDesktopApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetApplicationsResponseDesktopApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetApplicationsResponseDesktopApplications) MarshalJSON() ([]byte, error) {
	type wire GetApplicationsResponseDesktopApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetApplicationsResponseWebApplications) UnmarshalJSON(data []byte) error {
	type wire GetApplicationsResponseWebApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetApplicationsResponseWebApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetApplicationsResponseWebApplications) MarshalJSON() ([]byte, error) {
	type wire GetApplicationsResponseWebApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetApplicationsResponseNetworkApplications) UnmarshalJSON(data []byte) error {
	type wire GetApplicationsResponseNetworkApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetApplicationsResponseNetworkApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetApplicationsResponseNetworkApplications) MarshalJSON() ([]byte, error) {
	type wire GetApplicationsResponseNetworkApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetApplicationsResponse) UnmarshalJSON(data []byte) error {
	type wire GetApplicationsResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetApplicationsResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetApplicationsResponse) MarshalJSON() ([]byte, error) {
	type wire GetApplicationsResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsResponseItemUsername) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsResponseItemUsername
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsResponseItemUsername(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsResponseItemUsername) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsResponseItemUsername
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsResponseItemBucketsItemDataInteractionDuration) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsResponseItemBucketsItemDataInteractionDuration
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsResponseItemBucketsItemDataInteractionDuration(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsResponseItemBucketsItemDataInteractionDuration) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsResponseItemBucketsItemDataInteractionDuration
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsResponseItemBucketsItemDataInteraction) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsResponseItemBucketsItemDataInteraction
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsResponseItemBucketsItemDataInteraction(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsResponseItemBucketsItemDataInteraction) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsResponseItemBucketsItemDataInteraction
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemType) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemType
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemType(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemType) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemType
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemDuration) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemDuration
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemDuration(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemDuration) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemDuration
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsResponseItemBucketsItemDataLifecycle) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsResponseItemBucketsItemDataLifecycle
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsResponseItemBucketsItemDataLifecycle(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsResponseItemBucketsItemDataLifecycle) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsResponseItemBucketsItemDataLifecycle
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsResponseItemBucketsItemData) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsResponseItemBucketsItemData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsResponseItemBucketsItemData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsResponseItemBucketsItemData) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsResponseItemBucketsItemData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsResponseItemBucketsItem) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsResponseItemBucketsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsResponseItemBucketsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsResponseItemBucketsItem) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsResponseItemBucketsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsResponseItem) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsResponseItem) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCollaborationResponseDesktopApplications) UnmarshalJSON(data []byte) error {
	type wire GetCollaborationResponseDesktopApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCollaborationResponseDesktopApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCollaborationResponseDesktopApplications) MarshalJSON() ([]byte, error) {
	type wire GetCollaborationResponseDesktopApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCollaborationResponseWebApplications) UnmarshalJSON(data []byte) error {
	type wire GetCollaborationResponseWebApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCollaborationResponseWebApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCollaborationResponseWebApplications) MarshalJSON() ([]byte, error) {
	type wire GetCollaborationResponseWebApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCollaborationResponseNetworkApplications) UnmarshalJSON(data []byte) error {
	type wire GetCollaborationResponseNetworkApplications
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCollaborationResponseNetworkApplications(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCollaborationResponseNetworkApplications) MarshalJSON() ([]byte, error) {
	type wire GetCollaborationResponseNetworkApplications
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCollaborationResponseTeamsCallsResult) UnmarshalJSON(data []byte) error {
	type wire GetCollaborationResponseTeamsCallsResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCollaborationResponseTeamsCallsResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCollaborationResponseTeamsCallsResult) MarshalJSON() ([]byte, error) {
	type wire GetCollaborationResponseTeamsCallsResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCollaborationResponseTeamsCalls) UnmarshalJSON(data []byte) error {
	type wire GetCollaborationResponseTeamsCalls
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCollaborationResponseTeamsCalls(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCollaborationResponseTeamsCalls) MarshalJSON() ([]byte, error) {
	type wire GetCollaborationResponseTeamsCalls
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCollaborationResponseZoomCallsResult) UnmarshalJSON(data []byte) error {
	type wire GetCollaborationResponseZoomCallsResult
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCollaborationResponseZoomCallsResult(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCollaborationResponseZoomCallsResult) MarshalJSON() ([]byte, error) {
	type wire GetCollaborationResponseZoomCallsResult
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCollaborationResponseZoomCalls) UnmarshalJSON(data []byte) error {
	type wire GetCollaborationResponseZoomCalls
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCollaborationResponseZoomCalls(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCollaborationResponseZoomCalls) MarshalJSON() ([]byte, error) {
	type wire GetCollaborationResponseZoomCalls
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCollaborationResponse) UnmarshalJSON(data []byte) error {
	type wire GetCollaborationResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCollaborationResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCollaborationResponse) MarshalJSON() ([]byte, error) {
	type wire GetCollaborationResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetErrorsDrilldownResponseApplicationCrashes) UnmarshalJSON(data []byte) error {
	type wire GetErrorsDrilldownResponseApplicationCrashes
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetErrorsDrilldownResponseApplicationCrashes(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetErrorsDrilldownResponseApplicationCrashes) MarshalJSON() ([]byte, error) {
	type wire GetErrorsDrilldownResponseApplicationCrashes
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetErrorsDrilldownResponseSystemCrashes) UnmarshalJSON(data []byte) error {
	type wire GetErrorsDrilldownResponseSystemCrashes
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetErrorsDrilldownResponseSystemCrashes(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetErrorsDrilldownResponseSystemCrashes) MarshalJSON() ([]byte, error) {
	type wire GetErrorsDrilldownResponseSystemCrashes
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetErrorsDrilldownResponseHardResets) UnmarshalJSON(data []byte) error {
	type wire GetErrorsDrilldownResponseHardResets
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetErrorsDrilldownResponseHardResets(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetErrorsDrilldownResponseHardResets) MarshalJSON() ([]byte, error) {
	type wire GetErrorsDrilldownResponseHardResets
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetErrorsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetErrorsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetErrorsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetErrorsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetErrorsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetFreezesDrilldownResponseBinaryFreezes) UnmarshalJSON(data []byte) error {
	type wire GetFreezesDrilldownResponseBinaryFreezes
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetFreezesDrilldownResponseBinaryFreezes(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetFreezesDrilldownResponseBinaryFreezes) MarshalJSON() ([]byte, error) {
	type wire GetFreezesDrilldownResponseBinaryFreezes
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetFreezesDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetFreezesDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetFreezesDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetFreezesDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetFreezesDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetAlertsDrilldownResponseEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetAlertsDrilldownResponseEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetAlertsDrilldownResponseEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetAlertsDrilldownResponseEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetAlertsDrilldownResponseEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetAlertsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetAlertsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetAlertsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetAlertsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetAlertsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActionsDrilldownResponseWorkflows) UnmarshalJSON(data []byte) error {
	type wire GetActionsDrilldownResponseWorkflows
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActionsDrilldownResponseWorkflows(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActionsDrilldownResponseWorkflows) MarshalJSON() ([]byte, error) {
	type wire GetActionsDrilldownResponseWorkflows
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActionsDrilldownResponseRemoteActions) UnmarshalJSON(data []byte) error {
	type wire GetActionsDrilldownResponseRemoteActions
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActionsDrilldownResponseRemoteActions(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActionsDrilldownResponseRemoteActions) MarshalJSON() ([]byte, error) {
	type wire GetActionsDrilldownResponseRemoteActions
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetActionsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetActionsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetActionsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetActionsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetActionsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootDuration) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootDuration
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootDuration(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootDuration) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootDuration
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootType) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootType
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootType(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootType) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootType
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsDrilldownResponseSystemBootsDataEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsDrilldownResponseSystemBootsDataEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsDrilldownResponseSystemBootsDataEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsDrilldownResponseSystemBootsDataEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsDrilldownResponseSystemBootsDataEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsDrilldownResponseSystemBootsData) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsDrilldownResponseSystemBootsData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsDrilldownResponseSystemBootsData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsDrilldownResponseSystemBootsData) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsDrilldownResponseSystemBootsData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsDrilldownResponseSystemBoots) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsDrilldownResponseSystemBoots
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsDrilldownResponseSystemBoots(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsDrilldownResponseSystemBoots) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsDrilldownResponseSystemBoots
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootTime) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootTime
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootTime(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootTime) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootTime
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootDuration) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootDuration
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootDuration(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootDuration) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootDuration
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootType) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootType
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootType(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootType) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootType
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsAndSuspendsDrilldownResponseSystemBootsData) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBootsData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsAndSuspendsDrilldownResponseSystemBootsData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsAndSuspendsDrilldownResponseSystemBootsData) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBootsData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsAndSuspendsDrilldownResponseSystemBoots) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBoots
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsAndSuspendsDrilldownResponseSystemBoots(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsAndSuspendsDrilldownResponseSystemBoots) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSystemBoots
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsAndSuspendsDrilldownResponseSuspends) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSuspends
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsAndSuspendsDrilldownResponseSuspends(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsAndSuspendsDrilldownResponseSuspends) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsAndSuspendsDrilldownResponseSuspends
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetSystemBootsAndSuspendsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetSystemBootsAndSuspendsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetSystemBootsAndSuspendsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetSystemBootsAndSuspendsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetSystemBootsAndSuspendsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetTeamsCallsDrilldownResponseCalls) UnmarshalJSON(data []byte) error {
	type wire GetTeamsCallsDrilldownResponseCalls
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetTeamsCallsDrilldownResponseCalls(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetTeamsCallsDrilldownResponseCalls) MarshalJSON() ([]byte, error) {
	type wire GetTeamsCallsDrilldownResponseCalls
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetTeamsCallsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetTeamsCallsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetTeamsCallsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetTeamsCallsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetTeamsCallsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetZoomCallsDrilldownResponseCalls) UnmarshalJSON(data []byte) error {
	type wire GetZoomCallsDrilldownResponseCalls
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetZoomCallsDrilldownResponseCalls(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetZoomCallsDrilldownResponseCalls) MarshalJSON() ([]byte, error) {
	type wire GetZoomCallsDrilldownResponseCalls
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetZoomCallsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetZoomCallsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetZoomCallsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetZoomCallsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetZoomCallsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDesktopConnectivityDrilldownResponseEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetDesktopConnectivityDrilldownResponseEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDesktopConnectivityDrilldownResponseEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDesktopConnectivityDrilldownResponseEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetDesktopConnectivityDrilldownResponseEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDesktopConnectivityDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetDesktopConnectivityDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDesktopConnectivityDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDesktopConnectivityDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetDesktopConnectivityDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetWebConnectivityDrilldownResponseEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetWebConnectivityDrilldownResponseEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetWebConnectivityDrilldownResponseEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetWebConnectivityDrilldownResponseEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetWebConnectivityDrilldownResponseEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetWebConnectivityDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetWebConnectivityDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetWebConnectivityDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetWebConnectivityDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetWebConnectivityDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDesktopApplicationsDrilldownResponseEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetDesktopApplicationsDrilldownResponseEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDesktopApplicationsDrilldownResponseEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDesktopApplicationsDrilldownResponseEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetDesktopApplicationsDrilldownResponseEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDesktopApplicationsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetDesktopApplicationsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDesktopApplicationsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDesktopApplicationsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetDesktopApplicationsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetWebApplicationsDrilldownResponseEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetWebApplicationsDrilldownResponseEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetWebApplicationsDrilldownResponseEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetWebApplicationsDrilldownResponseEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetWebApplicationsDrilldownResponseEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetWebApplicationsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetWebApplicationsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetWebApplicationsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetWebApplicationsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetWebApplicationsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetInstallationsDrilldownResponseServiceEventsConfigurationChanges) UnmarshalJSON(data []byte) error {
	type wire GetInstallationsDrilldownResponseServiceEventsConfigurationChanges
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetInstallationsDrilldownResponseServiceEventsConfigurationChanges(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetInstallationsDrilldownResponseServiceEventsConfigurationChanges) MarshalJSON() ([]byte, error) {
	type wire GetInstallationsDrilldownResponseServiceEventsConfigurationChanges
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetInstallationsDrilldownResponseServiceEventsInstallations) UnmarshalJSON(data []byte) error {
	type wire GetInstallationsDrilldownResponseServiceEventsInstallations
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetInstallationsDrilldownResponseServiceEventsInstallations(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetInstallationsDrilldownResponseServiceEventsInstallations) MarshalJSON() ([]byte, error) {
	type wire GetInstallationsDrilldownResponseServiceEventsInstallations
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetInstallationsDrilldownResponseServiceEvents) UnmarshalJSON(data []byte) error {
	type wire GetInstallationsDrilldownResponseServiceEvents
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetInstallationsDrilldownResponseServiceEvents(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetInstallationsDrilldownResponseServiceEvents) MarshalJSON() ([]byte, error) {
	type wire GetInstallationsDrilldownResponseServiceEvents
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetInstallationsDrilldownResponsePackageEventsInstalls) UnmarshalJSON(data []byte) error {
	type wire GetInstallationsDrilldownResponsePackageEventsInstalls
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetInstallationsDrilldownResponsePackageEventsInstalls(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetInstallationsDrilldownResponsePackageEventsInstalls) MarshalJSON() ([]byte, error) {
	type wire GetInstallationsDrilldownResponsePackageEventsInstalls
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetInstallationsDrilldownResponsePackageEventsUninstalls) UnmarshalJSON(data []byte) error {
	type wire GetInstallationsDrilldownResponsePackageEventsUninstalls
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetInstallationsDrilldownResponsePackageEventsUninstalls(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetInstallationsDrilldownResponsePackageEventsUninstalls) MarshalJSON() ([]byte, error) {
	type wire GetInstallationsDrilldownResponsePackageEventsUninstalls
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetInstallationsDrilldownResponsePackageEvents) UnmarshalJSON(data []byte) error {
	type wire GetInstallationsDrilldownResponsePackageEvents
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetInstallationsDrilldownResponsePackageEvents(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetInstallationsDrilldownResponsePackageEvents) MarshalJSON() ([]byte, error) {
	type wire GetInstallationsDrilldownResponsePackageEvents
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetInstallationsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetInstallationsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetInstallationsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetInstallationsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetInstallationsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetEthernetDrilldownResponseEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetEthernetDrilldownResponseEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetEthernetDrilldownResponseEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetEthernetDrilldownResponseEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetEthernetDrilldownResponseEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetEthernetDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetEthernetDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetEthernetDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetEthernetDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetEthernetDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetWiFiDrilldownResponseWifi) UnmarshalJSON(data []byte) error {
	type wire GetWiFiDrilldownResponseWifi
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetWiFiDrilldownResponseWifi(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetWiFiDrilldownResponseWifi) MarshalJSON() ([]byte, error) {
	type wire GetWiFiDrilldownResponseWifi
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetWiFiDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetWiFiDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetWiFiDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetWiFiDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetWiFiDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryLocationTypeNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryLocationTypeNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryLocationTypeNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryLocationTypeNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryLocationTypeNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryLocationType) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryLocationType
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryLocationType(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryLocationType) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryLocationType
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCountNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCountNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCountNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCountNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCountNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCount) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCount
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCount(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCount) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCount
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCountNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCountNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCountNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCountNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCountNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCount) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCount
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCount(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCount) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCount
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCountNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCountNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCountNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCountNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCountNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCount) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCount
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCount(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCount) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCount
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTrafficNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTrafficNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTrafficNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTrafficNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTrafficNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTraffic) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTraffic
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTraffic(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTraffic) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTraffic
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTrafficNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTrafficNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTrafficNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTrafficNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTrafficNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTraffic) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTraffic
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTraffic(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTraffic) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTraffic
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCountNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCountNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCountNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCountNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCountNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCount) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCount
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCount(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCount) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCount
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsSummary) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummary
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsSummary(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsSummary) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsSummary
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryNameNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryNameNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryNameNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryNameNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryNameNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryName) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryName
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryName(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryName) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryName
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTrafficNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTrafficNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTrafficNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTrafficNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTrafficNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTraffic) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTraffic
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTraffic(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTraffic) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTraffic
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItem) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItem) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseConnectionEvents) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseConnectionEvents
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseConnectionEvents(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseConnectionEvents) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseConnectionEvents
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverageNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverageNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverageNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverageNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverageNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverage) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverage
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverage(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverage) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverage
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatioNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatioNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatioNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatioNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatioNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatio) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatio
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatio(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatio) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatio
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatioNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatioNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatioNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatioNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatioNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatio) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatio
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatio(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatio) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatio
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatioNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatioNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatioNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatioNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatioNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatio) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatio
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatio(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatio) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatio
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatioNQLData) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatioNQLData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatioNQLData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatioNQLData) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatioNQLData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatio) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatio
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatio(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatio) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatio
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEventsSummary) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEventsSummary
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEventsSummary(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEventsSummary) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEventsSummary
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponseTcpEvents) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponseTcpEvents
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponseTcpEvents(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponseTcpEvents) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponseTcpEvents
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetConnectionsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetConnectionsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetConnectionsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetConnectionsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetConnectionsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetNetworkApplicationDrilldownResponseEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetNetworkApplicationDrilldownResponseEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetNetworkApplicationDrilldownResponseEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetNetworkApplicationDrilldownResponseEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetNetworkApplicationDrilldownResponseEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetNetworkApplicationDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetNetworkApplicationDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetNetworkApplicationDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetNetworkApplicationDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetNetworkApplicationDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCPUDrilldownResponseCPUUsage) UnmarshalJSON(data []byte) error {
	type wire GetCPUDrilldownResponseCPUUsage
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCPUDrilldownResponseCPUUsage(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCPUDrilldownResponseCPUUsage) MarshalJSON() ([]byte, error) {
	type wire GetCPUDrilldownResponseCPUUsage
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCPUDrilldownResponseDevices) UnmarshalJSON(data []byte) error {
	type wire GetCPUDrilldownResponseDevices
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCPUDrilldownResponseDevices(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCPUDrilldownResponseDevices) MarshalJSON() ([]byte, error) {
	type wire GetCPUDrilldownResponseDevices
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCPUDrilldownResponseInterrupts) UnmarshalJSON(data []byte) error {
	type wire GetCPUDrilldownResponseInterrupts
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCPUDrilldownResponseInterrupts(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCPUDrilldownResponseInterrupts) MarshalJSON() ([]byte, error) {
	type wire GetCPUDrilldownResponseInterrupts
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetCPUDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetCPUDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetCPUDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetCPUDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetCPUDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetMemoryDrilldownResponseMemoryUsage) UnmarshalJSON(data []byte) error {
	type wire GetMemoryDrilldownResponseMemoryUsage
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetMemoryDrilldownResponseMemoryUsage(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetMemoryDrilldownResponseMemoryUsage) MarshalJSON() ([]byte, error) {
	type wire GetMemoryDrilldownResponseMemoryUsage
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetMemoryDrilldownResponseDevices) UnmarshalJSON(data []byte) error {
	type wire GetMemoryDrilldownResponseDevices
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetMemoryDrilldownResponseDevices(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetMemoryDrilldownResponseDevices) MarshalJSON() ([]byte, error) {
	type wire GetMemoryDrilldownResponseDevices
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetMemoryDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetMemoryDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetMemoryDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetMemoryDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetMemoryDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDiskPerformanceDrilldownResponseEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetDiskPerformanceDrilldownResponseEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDiskPerformanceDrilldownResponseEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDiskPerformanceDrilldownResponseEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetDiskPerformanceDrilldownResponseEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDiskPerformanceDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetDiskPerformanceDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDiskPerformanceDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDiskPerformanceDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetDiskPerformanceDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDriveSpaceDrilldownResponseEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetDriveSpaceDrilldownResponseEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDriveSpaceDrilldownResponseEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDriveSpaceDrilldownResponseEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetDriveSpaceDrilldownResponseEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetDriveSpaceDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetDriveSpaceDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetDriveSpaceDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetDriveSpaceDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetDriveSpaceDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetGPUDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetGPUDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetGPUDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetGPUDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetGPUDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetNPUDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetNPUDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetNPUDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetNPUDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetNPUDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseUserDetailsDataFullName) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseUserDetailsDataFullName
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseUserDetailsDataFullName(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseUserDetailsDataFullName) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseUserDetailsDataFullName
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseUserDetailsDataDepartment) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseUserDetailsDataDepartment
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseUserDetailsDataDepartment(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseUserDetailsDataDepartment) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseUserDetailsDataDepartment
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseUserDetailsDataType) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseUserDetailsDataType
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseUserDetailsDataType(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseUserDetailsDataType) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseUserDetailsDataType
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseUserDetailsData) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseUserDetailsData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseUserDetailsData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseUserDetailsData) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseUserDetailsData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseUserDetails) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseUserDetails
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseUserDetails(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseUserDetails) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseUserDetails
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseInteractionDataDuration) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseInteractionDataDuration
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseInteractionDataDuration(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseInteractionDataDuration) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseInteractionDataDuration
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseInteractionData) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseInteractionData
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseInteractionData(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseInteractionData) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseInteractionData
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseInteraction) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseInteraction
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseInteraction(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseInteraction) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseInteraction
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTime) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTime
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTime(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTime) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTime
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemType) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemType
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseLifeCyclesEventsItemType(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseLifeCyclesEventsItemType) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemType
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopVisible) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopVisible
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopVisible(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopVisible) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopVisible
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopReady) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopReady
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopReady(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopReady) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopReady
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemRemoteSessionStartupDuration) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemRemoteSessionStartupDuration
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseLifeCyclesEventsItemRemoteSessionStartupDuration(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseLifeCyclesEventsItemRemoteSessionStartupDuration) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemRemoteSessionStartupDuration
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemProfileLoadDuration) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemProfileLoadDuration
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseLifeCyclesEventsItemProfileLoadDuration(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseLifeCyclesEventsItemProfileLoadDuration) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemProfileLoadDuration
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemGpoLoadDuration) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemGpoLoadDuration
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseLifeCyclesEventsItemGpoLoadDuration(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseLifeCyclesEventsItemGpoLoadDuration) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemGpoLoadDuration
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemLogonScriptDuration) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemLogonScriptDuration
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseLifeCyclesEventsItemLogonScriptDuration(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseLifeCyclesEventsItemLogonScriptDuration) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItemLogonScriptDuration
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseLifeCyclesEventsItem) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseLifeCyclesEventsItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseLifeCyclesEventsItem) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseLifeCyclesEventsItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponseLifeCycles) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponseLifeCycles
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponseLifeCycles(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponseLifeCycles) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponseLifeCycles
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetUserInteractionsDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetUserInteractionsDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetUserInteractionsDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetUserInteractionsDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetUserInteractionsDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetRoundTripTimeDrilldownResponseRtt) UnmarshalJSON(data []byte) error {
	type wire GetRoundTripTimeDrilldownResponseRtt
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetRoundTripTimeDrilldownResponseRtt(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetRoundTripTimeDrilldownResponseRtt) MarshalJSON() ([]byte, error) {
	type wire GetRoundTripTimeDrilldownResponseRtt
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetRoundTripTimeDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetRoundTripTimeDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetRoundTripTimeDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetRoundTripTimeDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetRoundTripTimeDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetNetworkLatencyDrilldownResponseNetworkLatency) UnmarshalJSON(data []byte) error {
	type wire GetNetworkLatencyDrilldownResponseNetworkLatency
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetNetworkLatencyDrilldownResponseNetworkLatency(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetNetworkLatencyDrilldownResponseNetworkLatency) MarshalJSON() ([]byte, error) {
	type wire GetNetworkLatencyDrilldownResponseNetworkLatency
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *GetNetworkLatencyDrilldownResponse) UnmarshalJSON(data []byte) error {
	type wire GetNetworkLatencyDrilldownResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = GetNetworkLatencyDrilldownResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v GetNetworkLatencyDrilldownResponse) MarshalJSON() ([]byte, error) {
	type wire GetNetworkLatencyDrilldownResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
