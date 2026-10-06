package support_timeline

import "encoding/json"

// Polymorphic NQL values, event payloads and plugin-defined details retain their exact JSON.
type GetAlertsAndErrorsResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Alerts           *[]json.RawMessage         `json:"alerts,omitempty"`
	Errors           *[]json.RawMessage         `json:"errors,omitempty"`
	Freezes          *[]json.RawMessage         `json:"freezes,omitempty"`
}
type GetPerformanceResponseCPUUsageResultTimeSeriesItemDataUsage struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
	ComputedValue    json.RawMessage            `json:"computedValue,omitempty"`
	ComputedStatus   *string                    `json:"computedStatus,omitempty"`
}
type GetPerformanceResponseCPUUsageResultTimeSeriesItemDataHighUsageBinaries struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TotalCount       *int64                     `json:"totalCount,omitempty"`
	TopBinaries      *[]json.RawMessage         `json:"topBinaries,omitempty"`
}
type GetPerformanceResponseCPUUsageResultTimeSeriesItemData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields  map[string]json.RawMessage                                               `json:"-"`
	Usage             *GetPerformanceResponseCPUUsageResultTimeSeriesItemDataUsage             `json:"usage,omitempty"`
	HighUsageBinaries *GetPerformanceResponseCPUUsageResultTimeSeriesItemDataHighUsageBinaries `json:"highUsageBinaries,omitempty"`
}
type GetPerformanceResponseCPUUsageResultTimeSeriesItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                              `json:"-"`
	StartTime        *string                                                 `json:"startTime,omitempty"`
	EndTime          *string                                                 `json:"endTime,omitempty"`
	Data             *GetPerformanceResponseCPUUsageResultTimeSeriesItemData `json:"data,omitempty"`
}
type GetPerformanceResponseCPUUsageResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                            `json:"-"`
	TimeSeries       *[]GetPerformanceResponseCPUUsageResultTimeSeriesItem `json:"timeSeries,omitempty"`
}
type GetPerformanceResponseCPUUsage struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage            `json:"-"`
	Result           *GetPerformanceResponseCPUUsageResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage           `json:"meta,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataUsage struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
	ComputedValue    json.RawMessage            `json:"computedValue,omitempty"`
	ComputedStatus   *string                    `json:"computedStatus,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataInstalled struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemorySwapRate struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
	ComputedValue    json.RawMessage            `json:"computedValue,omitempty"`
	ComputedStatus   *string                    `json:"computedStatus,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDiskQueueLength struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ComputedStatus   *string                    `json:"computedStatus,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemoryPressure struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
	ComputedValue    json.RawMessage            `json:"computedValue,omitempty"`
	ComputedStatus   *string                    `json:"computedStatus,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationHighMemoryPressure struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
	ComputedStatus   *string                    `json:"computedStatus,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationMediumMemoryPressure struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemBinaryName struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemUsage struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                                           `json:"-"`
	BinaryName       *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemBinaryName `json:"binaryName,omitempty"`
	Usage            *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItemUsage      `json:"usage,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinaries struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                                   `json:"-"`
	TotalCount       *int64                                                                                       `json:"totalCount,omitempty"`
	TopBinaries      *[]GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinariesTopBinariesItem `json:"topBinaries,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItemData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields             map[string]json.RawMessage                                                             `json:"-"`
	Usage                        *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataUsage                        `json:"usage,omitempty"`
	Installed                    *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataInstalled                    `json:"installed,omitempty"`
	MemorySwapRate               *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemorySwapRate               `json:"memorySwapRate,omitempty"`
	DiskQueueLength              *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDiskQueueLength              `json:"diskQueueLength,omitempty"`
	MemoryPressure               *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataMemoryPressure               `json:"memoryPressure,omitempty"`
	DurationHighMemoryPressure   *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationHighMemoryPressure   `json:"durationHighMemoryPressure,omitempty"`
	DurationMediumMemoryPressure *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataDurationMediumMemoryPressure `json:"durationMediumMemoryPressure,omitempty"`
	HighUsageBinaries            *GetPerformanceResponseMemoryUsageResultTimeSeriesItemDataHighUsageBinaries            `json:"highUsageBinaries,omitempty"`
}
type GetPerformanceResponseMemoryUsageResultTimeSeriesItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                 `json:"-"`
	StartTime        *string                                                    `json:"startTime,omitempty"`
	EndTime          *string                                                    `json:"endTime,omitempty"`
	Data             *GetPerformanceResponseMemoryUsageResultTimeSeriesItemData `json:"data,omitempty"`
}
type GetPerformanceResponseMemoryUsageResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                               `json:"-"`
	TimeSeries       *[]GetPerformanceResponseMemoryUsageResultTimeSeriesItem `json:"timeSeries,omitempty"`
}
type GetPerformanceResponseMemoryUsage struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage               `json:"-"`
	Result           *GetPerformanceResponseMemoryUsageResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage              `json:"meta,omitempty"`
}
type GetPerformanceResponseGPUSlotOne struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetPerformanceResponseGPUSlotTwo struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetPerformanceResponseDiskResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TimeSeries       *[]json.RawMessage         `json:"timeSeries,omitempty"`
}
type GetPerformanceResponseDisk struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage        `json:"-"`
	Result           *GetPerformanceResponseDiskResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage       `json:"meta,omitempty"`
}
type GetPerformanceResponseDriveSpaceResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TimeSeries       *[]json.RawMessage         `json:"timeSeries,omitempty"`
}
type GetPerformanceResponseDriveSpace struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage              `json:"-"`
	Result           *GetPerformanceResponseDriveSpaceResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage             `json:"meta,omitempty"`
}
type GetPerformanceResponseNPUResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TimeSeries       *[]json.RawMessage         `json:"timeSeries,omitempty"`
}
type GetPerformanceResponseNPU struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage       `json:"-"`
	Result           *GetPerformanceResponseNPUResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage      `json:"meta,omitempty"`
}
type GetPerformanceResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage         `json:"-"`
	CPUUsage         *GetPerformanceResponseCPUUsage    `json:"cpuUsage,omitempty"`
	MemoryUsage      *GetPerformanceResponseMemoryUsage `json:"memoryUsage,omitempty"`
	GPUSlotOne       *GetPerformanceResponseGPUSlotOne  `json:"gpuSlotOne,omitempty"`
	GPUSlotTwo       *GetPerformanceResponseGPUSlotTwo  `json:"gpuSlotTwo,omitempty"`
	Disk             *GetPerformanceResponseDisk        `json:"disk,omitempty"`
	DriveSpace       *GetPerformanceResponseDriveSpace  `json:"driveSpace,omitempty"`
	NPU              *GetPerformanceResponseNPU         `json:"npu,omitempty"`
}
type GetConnectivityResponseWifiResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TimeSeries       *[]json.RawMessage         `json:"timeSeries,omitempty"`
}
type GetConnectivityResponseWifi struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage         `json:"-"`
	Result           *GetConnectivityResponseWifiResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage        `json:"meta,omitempty"`
}
type GetConnectivityResponseEthernetResultTimeSeriesItemDataLocalIPAddressesItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
}
type GetConnectivityResponseEthernetResultTimeSeriesItemDataMacAddressItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
}
type GetConnectivityResponseEthernetResultTimeSeriesItemData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                     `json:"-"`
	LocalIPAddresses *[]GetConnectivityResponseEthernetResultTimeSeriesItemDataLocalIPAddressesItem `json:"localIpAddresses,omitempty"`
	MacAddress       *[]GetConnectivityResponseEthernetResultTimeSeriesItemDataMacAddressItem       `json:"macAddress,omitempty"`
}
type GetConnectivityResponseEthernetResultTimeSeriesItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                               `json:"-"`
	StartTime        *string                                                  `json:"startTime,omitempty"`
	EndTime          *string                                                  `json:"endTime,omitempty"`
	Data             *GetConnectivityResponseEthernetResultTimeSeriesItemData `json:"data,omitempty"`
}
type GetConnectivityResponseEthernetResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                             `json:"-"`
	TimeSeries       *[]GetConnectivityResponseEthernetResultTimeSeriesItem `json:"timeSeries,omitempty"`
}
type GetConnectivityResponseEthernet struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage             `json:"-"`
	Result           *GetConnectivityResponseEthernetResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage            `json:"meta,omitempty"`
}
type GetConnectivityResponseBluetoothResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TimeSeries       *[]json.RawMessage         `json:"timeSeries,omitempty"`
}
type GetConnectivityResponseBluetooth struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage              `json:"-"`
	Result           *GetConnectivityResponseBluetoothResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage             `json:"meta,omitempty"`
}
type GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsFailedConnectionRatio struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
	ComputedStatus   *string                    `json:"computedStatus,omitempty"`
}
type GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsConnectionEstablishmentTime struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
	ComputedStatus   *string                    `json:"computedStatus,omitempty"`
}
type GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnections struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields            map[string]json.RawMessage                                                                           `json:"-"`
	FailedConnectionRatio       *GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsFailedConnectionRatio       `json:"failedConnectionRatio,omitempty"`
	ConnectionEstablishmentTime *GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnectionsConnectionEstablishmentTime `json:"connectionEstablishmentTime,omitempty"`
}
type GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsNumberOfConnections struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
}
type GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsIncomingTraffic struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
	ComputedStatus   *string                    `json:"computedStatus,omitempty"`
}
type GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsOutgoingTraffic struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	NQLValue         json.RawMessage            `json:"nqlValue,omitempty"`
	ComputedStatus   *string                    `json:"computedStatus,omitempty"`
}
type GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnections struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields    map[string]json.RawMessage                                                                `json:"-"`
	NumberOfConnections *GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsNumberOfConnections `json:"numberOfConnections,omitempty"`
	IncomingTraffic     *GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsIncomingTraffic     `json:"incomingTraffic,omitempty"`
	OutgoingTraffic     *GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnectionsOutgoingTraffic     `json:"outgoingTraffic,omitempty"`
	LocationType        *map[string]json.RawMessage                                                               `json:"locationType,omitempty"`
}
type GetConnectivityResponseConnectionsResultTimeSeriesItemData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                `json:"-"`
	TcpConnections   *GetConnectivityResponseConnectionsResultTimeSeriesItemDataTcpConnections `json:"tcpConnections,omitempty"`
	Connections      *GetConnectivityResponseConnectionsResultTimeSeriesItemDataConnections    `json:"connections,omitempty"`
}
type GetConnectivityResponseConnectionsResultTimeSeriesItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                  `json:"-"`
	StartTime        *string                                                     `json:"startTime,omitempty"`
	EndTime          *string                                                     `json:"endTime,omitempty"`
	Data             *GetConnectivityResponseConnectionsResultTimeSeriesItemData `json:"data,omitempty"`
}
type GetConnectivityResponseConnectionsResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                `json:"-"`
	TimeSeries       *[]GetConnectivityResponseConnectionsResultTimeSeriesItem `json:"timeSeries,omitempty"`
}
type GetConnectivityResponseConnections struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                `json:"-"`
	Result           *GetConnectivityResponseConnectionsResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage               `json:"meta,omitempty"`
}
type GetConnectivityResponseVpnResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TimeSeries       *[]json.RawMessage         `json:"timeSeries,omitempty"`
}
type GetConnectivityResponseVpn struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage        `json:"-"`
	Result           *GetConnectivityResponseVpnResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage       `json:"meta,omitempty"`
}
type GetConnectivityResponseDesktopApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetConnectivityResponseWebApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetConnectivityResponseNetworkApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetConnectivityResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields    map[string]json.RawMessage                  `json:"-"`
	Wifi                *GetConnectivityResponseWifi                `json:"wifi,omitempty"`
	Ethernet            *GetConnectivityResponseEthernet            `json:"ethernet,omitempty"`
	Bluetooth           *GetConnectivityResponseBluetooth           `json:"bluetooth,omitempty"`
	Connections         *GetConnectivityResponseConnections         `json:"connections,omitempty"`
	Vpn                 *GetConnectivityResponseVpn                 `json:"vpn,omitempty"`
	DesktopApplications *GetConnectivityResponseDesktopApplications `json:"desktopApplications,omitempty"`
	WebApplications     *GetConnectivityResponseWebApplications     `json:"webApplications,omitempty"`
	NetworkApplications *GetConnectivityResponseNetworkApplications `json:"networkApplications,omitempty"`
}
type GetApplicationConnectivityResponseDesktopApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetApplicationConnectivityResponseWebApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetApplicationConnectivityResponseNetworkApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetApplicationConnectivityResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields    map[string]json.RawMessage                             `json:"-"`
	DesktopApplications *GetApplicationConnectivityResponseDesktopApplications `json:"desktopApplications,omitempty"`
	WebApplications     *GetApplicationConnectivityResponseWebApplications     `json:"webApplications,omitempty"`
	NetworkApplications *GetApplicationConnectivityResponseNetworkApplications `json:"networkApplications,omitempty"`
}
type GetActivityResponseInstallationEventsResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TimeSeries       *[]json.RawMessage         `json:"timeSeries,omitempty"`
}
type GetActivityResponseInstallationEvents struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                   `json:"-"`
	Result           *GetActivityResponseInstallationEventsResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage                  `json:"meta,omitempty"`
}
type GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBootsEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Datetime         *string                    `json:"datetime,omitempty"`
	Type             *string                    `json:"type,omitempty"`
	Cardinality      *int64                     `json:"cardinality,omitempty"`
}
type GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBoots struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                      `json:"-"`
	Events           *[]GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBootsEventsItem `json:"events,omitempty"`
	Status           *string                                                                         `json:"status,omitempty"`
	Count            *int64                                                                          `json:"count,omitempty"`
}
type GetActivityResponseSystemEventsResultTimeSeriesItemData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                          `json:"-"`
	SystemBoots      *GetActivityResponseSystemEventsResultTimeSeriesItemDataSystemBoots `json:"systemBoots,omitempty"`
}
type GetActivityResponseSystemEventsResultTimeSeriesItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                               `json:"-"`
	StartTime        *string                                                  `json:"startTime,omitempty"`
	EndTime          *string                                                  `json:"endTime,omitempty"`
	Data             *GetActivityResponseSystemEventsResultTimeSeriesItemData `json:"data,omitempty"`
}
type GetActivityResponseSystemEventsResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                             `json:"-"`
	TimeSeries       *[]GetActivityResponseSystemEventsResultTimeSeriesItem `json:"timeSeries,omitempty"`
}
type GetActivityResponseSystemEvents struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage             `json:"-"`
	Result           *GetActivityResponseSystemEventsResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage            `json:"meta,omitempty"`
}
type GetActivityResponseActionsResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TimeSeries       *[]json.RawMessage         `json:"timeSeries,omitempty"`
}
type GetActivityResponseActions struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage        `json:"-"`
	Result           *GetActivityResponseActionsResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage       `json:"meta,omitempty"`
}
type GetActivityResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields   map[string]json.RawMessage             `json:"-"`
	InstallationEvents *GetActivityResponseInstallationEvents `json:"installationEvents,omitempty"`
	SystemEvents       *GetActivityResponseSystemEvents       `json:"systemEvents,omitempty"`
	Actions            *GetActivityResponseActions            `json:"actions,omitempty"`
}
type GetApplicationsResponseDesktopApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetApplicationsResponseWebApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetApplicationsResponseNetworkApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetApplicationsResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields    map[string]json.RawMessage                  `json:"-"`
	DesktopApplications *GetApplicationsResponseDesktopApplications `json:"desktopApplications,omitempty"`
	WebApplications     *GetApplicationsResponseWebApplications     `json:"webApplications,omitempty"`
	NetworkApplications *GetApplicationsResponseNetworkApplications `json:"networkApplications,omitempty"`
}
type GetUserInteractionsResponseItemUsername struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetUserInteractionsResponseItemBucketsItemDataInteractionDuration struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetUserInteractionsResponseItemBucketsItemDataInteraction struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                         `json:"-"`
	Duration         *GetUserInteractionsResponseItemBucketsItemDataInteractionDuration `json:"duration,omitempty"`
	IsUserActive     *bool                                                              `json:"isUserActive,omitempty"`
}
type GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemType struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemDuration struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
}
type GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                 `json:"-"`
	Datetime         *string                                                                    `json:"datetime,omitempty"`
	Type             *GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemType     `json:"type,omitempty"`
	Duration         *GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItemDuration `json:"duration,omitempty"`
	Cardinality      *int64                                                                     `json:"cardinality,omitempty"`
}
type GetUserInteractionsResponseItemBucketsItemDataLifecycle struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                           `json:"-"`
	Events           *[]GetUserInteractionsResponseItemBucketsItemDataLifecycleEventsItem `json:"events,omitempty"`
	Count            *int64                                                               `json:"count,omitempty"`
}
type GetUserInteractionsResponseItemBucketsItemData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                 `json:"-"`
	Interaction      *GetUserInteractionsResponseItemBucketsItemDataInteraction `json:"interaction,omitempty"`
	Lifecycle        *GetUserInteractionsResponseItemBucketsItemDataLifecycle   `json:"lifecycle,omitempty"`
}
type GetUserInteractionsResponseItemBucketsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                      `json:"-"`
	Datetime         *string                                         `json:"datetime,omitempty"`
	Data             *GetUserInteractionsResponseItemBucketsItemData `json:"data,omitempty"`
}
type GetUserInteractionsResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                    `json:"-"`
	Username         *GetUserInteractionsResponseItemUsername      `json:"username,omitempty"`
	UserID           *string                                       `json:"userId,omitempty"`
	Buckets          *[]GetUserInteractionsResponseItemBucketsItem `json:"buckets,omitempty"`
}
type GetUserInteractionsResponse []GetUserInteractionsResponseItem
type GetCollaborationResponseDesktopApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetCollaborationResponseWebApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetCollaborationResponseNetworkApplications struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage  `json:"-"`
	Result           *[]json.RawMessage          `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage `json:"meta,omitempty"`
}
type GetCollaborationResponseTeamsCallsResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TimeSeries       *[]json.RawMessage         `json:"timeSeries,omitempty"`
}
type GetCollaborationResponseTeamsCalls struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                `json:"-"`
	Result           *GetCollaborationResponseTeamsCallsResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage               `json:"meta,omitempty"`
}
type GetCollaborationResponseZoomCallsResult struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	TimeSeries       *[]json.RawMessage         `json:"timeSeries,omitempty"`
}
type GetCollaborationResponseZoomCalls struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage               `json:"-"`
	Result           *GetCollaborationResponseZoomCallsResult `json:"result,omitempty"`
	Meta             *map[string]json.RawMessage              `json:"meta,omitempty"`
}
type GetCollaborationResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields    map[string]json.RawMessage                   `json:"-"`
	DesktopApplications *GetCollaborationResponseDesktopApplications `json:"desktopApplications,omitempty"`
	WebApplications     *GetCollaborationResponseWebApplications     `json:"webApplications,omitempty"`
	NetworkApplications *GetCollaborationResponseNetworkApplications `json:"networkApplications,omitempty"`
	TeamsCalls          *GetCollaborationResponseTeamsCalls          `json:"teamsCalls,omitempty"`
	ZoomCalls           *GetCollaborationResponseZoomCalls           `json:"zoomCalls,omitempty"`
}
type GetErrorsDrilldownResponseApplicationCrashes struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	Events               *[]json.RawMessage         `json:"events,omitempty"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetErrorsDrilldownResponseSystemCrashes struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	Events               *[]json.RawMessage         `json:"events,omitempty"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetErrorsDrilldownResponseHardResets struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	Events               *[]json.RawMessage         `json:"events,omitempty"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetErrorsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields   map[string]json.RawMessage                    `json:"-"`
	ApplicationCrashes *GetErrorsDrilldownResponseApplicationCrashes `json:"applicationCrashes,omitempty"`
	SystemCrashes      *GetErrorsDrilldownResponseSystemCrashes      `json:"systemCrashes,omitempty"`
	HardResets         *GetErrorsDrilldownResponseHardResets         `json:"hardResets,omitempty"`
}
type GetFreezesDrilldownResponseBinaryFreezes struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Events           *[]json.RawMessage         `json:"events,omitempty"`
}
type GetFreezesDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                `json:"-"`
	BinaryFreezes    *GetFreezesDrilldownResponseBinaryFreezes `json:"binaryFreezes,omitempty"`
}
type GetAlertsDrilldownResponseEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Datetime         *string                    `json:"datetime,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetAlertsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage              `json:"-"`
	Events               *[]GetAlertsDrilldownResponseEventsItem `json:"events,omitempty"`
	InvestigationQueries *map[string]string                      `json:"investigationQueries,omitempty"`
}
type GetActionsDrilldownResponseWorkflows struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Events           *[]json.RawMessage         `json:"events,omitempty"`
}
type GetActionsDrilldownResponseRemoteActions struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Events           *[]json.RawMessage         `json:"events,omitempty"`
}
type GetActionsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                `json:"-"`
	Workflows        *GetActionsDrilldownResponseWorkflows     `json:"workflows,omitempty"`
	RemoteActions    *GetActionsDrilldownResponseRemoteActions `json:"remoteActions,omitempty"`
}
type GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootDuration struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootType struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetSystemBootsDrilldownResponseSystemBootsDataEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                            `json:"-"`
	BootTime         *string                                                               `json:"bootTime,omitempty"`
	BootDuration     *GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootDuration `json:"bootDuration,omitempty"`
	BootType         *GetSystemBootsDrilldownResponseSystemBootsDataEventsItemBootType     `json:"bootType,omitempty"`
}
type GetSystemBootsDrilldownResponseSystemBootsData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                  `json:"-"`
	Events           *[]GetSystemBootsDrilldownResponseSystemBootsDataEventsItem `json:"events,omitempty"`
}
type GetSystemBootsDrilldownResponseSystemBoots struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                      `json:"-"`
	Data                 *GetSystemBootsDrilldownResponseSystemBootsData `json:"data,omitempty"`
	InvestigationQueries *map[string]string                              `json:"investigationQueries,omitempty"`
}
type GetSystemBootsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                  `json:"-"`
	SystemBoots      *GetSystemBootsDrilldownResponseSystemBoots `json:"systemBoots,omitempty"`
}
type GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootTime struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootDuration struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootType struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                       `json:"-"`
	BootTime         *GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootTime     `json:"bootTime,omitempty"`
	BootDuration     *GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootDuration `json:"bootDuration,omitempty"`
	BootType         *GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItemBootType     `json:"bootType,omitempty"`
}
type GetSystemBootsAndSuspendsDrilldownResponseSystemBootsData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                             `json:"-"`
	Events           *[]GetSystemBootsAndSuspendsDrilldownResponseSystemBootsDataEventsItem `json:"events,omitempty"`
}
type GetSystemBootsAndSuspendsDrilldownResponseSystemBoots struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                                 `json:"-"`
	Data                 *GetSystemBootsAndSuspendsDrilldownResponseSystemBootsData `json:"data,omitempty"`
	InvestigationQueries *map[string]string                                         `json:"investigationQueries,omitempty"`
}
type GetSystemBootsAndSuspendsDrilldownResponseSuspends struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Events           *[]json.RawMessage         `json:"events,omitempty"`
}
type GetSystemBootsAndSuspendsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                             `json:"-"`
	SystemBoots      *GetSystemBootsAndSuspendsDrilldownResponseSystemBoots `json:"systemBoots,omitempty"`
	Suspends         *GetSystemBootsAndSuspendsDrilldownResponseSuspends    `json:"suspends,omitempty"`
}
type GetTeamsCallsDrilldownResponseCalls struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	Events               *[]json.RawMessage         `json:"events,omitempty"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetTeamsCallsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage           `json:"-"`
	Calls            *GetTeamsCallsDrilldownResponseCalls `json:"calls,omitempty"`
}
type GetZoomCallsDrilldownResponseCalls struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	Events               *[]json.RawMessage         `json:"events,omitempty"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetZoomCallsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage          `json:"-"`
	Calls            *GetZoomCallsDrilldownResponseCalls `json:"calls,omitempty"`
}
type GetDesktopConnectivityDrilldownResponseEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Datetime         *string                    `json:"datetime,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetDesktopConnectivityDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                           `json:"-"`
	Events               *[]GetDesktopConnectivityDrilldownResponseEventsItem `json:"events,omitempty"`
	InvestigationQueries *map[string]string                                   `json:"investigationQueries,omitempty"`
}
type GetWebConnectivityDrilldownResponseEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Datetime         *string                    `json:"datetime,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetWebConnectivityDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                       `json:"-"`
	Events               *[]GetWebConnectivityDrilldownResponseEventsItem `json:"events,omitempty"`
	InvestigationQueries *map[string]string                               `json:"investigationQueries,omitempty"`
}
type GetDesktopApplicationsDrilldownResponseEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Datetime         *string                    `json:"datetime,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetDesktopApplicationsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                           `json:"-"`
	Events               *[]GetDesktopApplicationsDrilldownResponseEventsItem `json:"events,omitempty"`
	InvestigationQueries *map[string]string                                   `json:"investigationQueries,omitempty"`
}
type GetWebApplicationsDrilldownResponseEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Datetime         *string                    `json:"datetime,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetWebApplicationsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                       `json:"-"`
	Events               *[]GetWebApplicationsDrilldownResponseEventsItem `json:"events,omitempty"`
	InvestigationQueries *map[string]string                               `json:"investigationQueries,omitempty"`
}
type GetInstallationsDrilldownResponseServiceEventsConfigurationChanges struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Events           *[]json.RawMessage         `json:"events,omitempty"`
}
type GetInstallationsDrilldownResponseServiceEventsInstallations struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Events           *[]json.RawMessage         `json:"events,omitempty"`
}
type GetInstallationsDrilldownResponseServiceEvents struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                                          `json:"-"`
	ConfigurationChanges *GetInstallationsDrilldownResponseServiceEventsConfigurationChanges `json:"configurationChanges,omitempty"`
	Installations        *GetInstallationsDrilldownResponseServiceEventsInstallations        `json:"installations,omitempty"`
}
type GetInstallationsDrilldownResponsePackageEventsInstalls struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Events           *[]json.RawMessage         `json:"events,omitempty"`
}
type GetInstallationsDrilldownResponsePackageEventsUninstalls struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Events           *[]json.RawMessage         `json:"events,omitempty"`
}
type GetInstallationsDrilldownResponsePackageEvents struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                `json:"-"`
	Installs         *GetInstallationsDrilldownResponsePackageEventsInstalls   `json:"installs,omitempty"`
	Uninstalls       *GetInstallationsDrilldownResponsePackageEventsUninstalls `json:"uninstalls,omitempty"`
}
type GetInstallationsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                      `json:"-"`
	ServiceEvents    *GetInstallationsDrilldownResponseServiceEvents `json:"serviceEvents,omitempty"`
	PackageEvents    *GetInstallationsDrilldownResponsePackageEvents `json:"packageEvents,omitempty"`
}
type GetEthernetDrilldownResponseEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Datetime         *string                    `json:"datetime,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetEthernetDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                `json:"-"`
	Events               *[]GetEthernetDrilldownResponseEventsItem `json:"events,omitempty"`
	InvestigationQueries *map[string]string                        `json:"investigationQueries,omitempty"`
}
type GetWiFiDrilldownResponseWifi struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	Events               *[]json.RawMessage         `json:"events,omitempty"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetWiFiDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage    `json:"-"`
	Wifi             *GetWiFiDrilldownResponseWifi `json:"wifi,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryLocationTypeNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryLocationType struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                 `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseConnectionEventsSummaryLocationTypeNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                    `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCountNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCount struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                     `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCountNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                        `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCountNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCount struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCountNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                   `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCountNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCount struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                    `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCountNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                       `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTrafficNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTraffic struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                    `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTrafficNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                       `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTrafficNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTraffic struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                    `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTrafficNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                       `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCountNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCount struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCountNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                   `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsSummary struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                              `json:"-"`
	LocationType     *GetConnectionsDrilldownResponseConnectionEventsSummaryLocationType     `json:"locationType,omitempty"`
	DestinationCount *GetConnectionsDrilldownResponseConnectionEventsSummaryDestinationCount `json:"destinationCount,omitempty"`
	BinaryCount      *GetConnectionsDrilldownResponseConnectionEventsSummaryBinaryCount      `json:"binaryCount,omitempty"`
	ConnectionCount  *GetConnectionsDrilldownResponseConnectionEventsSummaryConnectionCount  `json:"connectionCount,omitempty"`
	IncomingTraffic  *GetConnectionsDrilldownResponseConnectionEventsSummaryIncomingTraffic  `json:"incomingTraffic,omitempty"`
	OutgoingTraffic  *GetConnectionsDrilldownResponseConnectionEventsSummaryOutgoingTraffic  `json:"outgoingTraffic,omitempty"`
	DomainCount      *GetConnectionsDrilldownResponseConnectionEventsSummaryDomainCount      `json:"domainCount,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryNameNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryName struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                             `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryNameNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                                `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTrafficNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTraffic struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                               `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTrafficNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                                  `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                        `json:"-"`
	BinaryName       *GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemBinaryName   `json:"binaryName,omitempty"`
	TotalTraffic     *GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItemTotalTraffic `json:"totalTraffic,omitempty"`
}
type GetConnectionsDrilldownResponseConnectionEvents struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields  map[string]json.RawMessage                                              `json:"-"`
	Summary           *GetConnectionsDrilldownResponseConnectionEventsSummary                 `json:"summary,omitempty"`
	BinariesByTraffic *[]GetConnectionsDrilldownResponseConnectionEventsBinariesByTrafficItem `json:"binariesByTraffic,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverageNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverage struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                                `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverageNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                                   `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatioNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatio struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                    `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatioNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                       `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatioNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatio struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                    `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatioNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                       `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatioNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatio struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                       `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatioNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                          `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatioNQLData struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatio struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                                      `json:"-"`
	NQLData          *GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatioNQLData `json:"nqlData,omitempty"`
	Status           *string                                                                         `json:"status,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEventsSummary struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields                   map[string]json.RawMessage                                                         `json:"-"`
	ConnectionEstablishmentTimeAverage *GetConnectionsDrilldownResponseTcpEventsSummaryConnectionEstablishmentTimeAverage `json:"connectionEstablishmentTimeAverage,omitempty"`
	FailedConnectionsRatio             *GetConnectionsDrilldownResponseTcpEventsSummaryFailedConnectionsRatio             `json:"failedConnectionsRatio,omitempty"`
	NoHostConnectionsRatio             *GetConnectionsDrilldownResponseTcpEventsSummaryNoHostConnectionsRatio             `json:"noHostConnectionsRatio,omitempty"`
	NoServiceConnectionsRatio          *GetConnectionsDrilldownResponseTcpEventsSummaryNoServiceConnectionsRatio          `json:"noServiceConnectionsRatio,omitempty"`
	RejectedConnectionsRatio           *GetConnectionsDrilldownResponseTcpEventsSummaryRejectedConnectionsRatio           `json:"rejectedConnectionsRatio,omitempty"`
}
type GetConnectionsDrilldownResponseTcpEvents struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields              map[string]json.RawMessage                       `json:"-"`
	Summary                       *GetConnectionsDrilldownResponseTcpEventsSummary `json:"summary,omitempty"`
	TopBinariesByFailedConEstTime *[]json.RawMessage                               `json:"topBinariesByFailedConEstTime,omitempty"`
	TopBinariesByFailedCon        *[]json.RawMessage                               `json:"topBinariesByFailedCon,omitempty"`
	TopDomainsByFailedConEstTime  *[]json.RawMessage                               `json:"topDomainsByFailedConEstTime,omitempty"`
	TopDomainByFailedConnection   *[]json.RawMessage                               `json:"topDomainByFailedConnection,omitempty"`
}
type GetConnectionsDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                       `json:"-"`
	ConnectionEvents *GetConnectionsDrilldownResponseConnectionEvents `json:"connectionEvents,omitempty"`
	TcpEvents        *GetConnectionsDrilldownResponseTcpEvents        `json:"tcpEvents,omitempty"`
}
type GetNetworkApplicationDrilldownResponseEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Datetime         *string                    `json:"datetime,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetNetworkApplicationDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                          `json:"-"`
	Events               *[]GetNetworkApplicationDrilldownResponseEventsItem `json:"events,omitempty"`
	InvestigationQueries *map[string]string                                  `json:"investigationQueries,omitempty"`
}
type GetCPUDrilldownResponseCPUUsage struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetCPUDrilldownResponseDevices struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetCPUDrilldownResponseInterrupts struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetCPUDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage         `json:"-"`
	CPUUsage         *GetCPUDrilldownResponseCPUUsage   `json:"cpuUsage,omitempty"`
	Devices          *GetCPUDrilldownResponseDevices    `json:"devices,omitempty"`
	Interrupts       *GetCPUDrilldownResponseInterrupts `json:"interrupts,omitempty"`
}
type GetMemoryDrilldownResponseMemoryUsage struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetMemoryDrilldownResponseDevices struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetMemoryDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage             `json:"-"`
	MemoryUsage      *GetMemoryDrilldownResponseMemoryUsage `json:"memoryUsage,omitempty"`
	Devices          *GetMemoryDrilldownResponseDevices     `json:"devices,omitempty"`
}
type GetDiskPerformanceDrilldownResponseEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Datetime         *string                    `json:"datetime,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetDiskPerformanceDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                       `json:"-"`
	Events               *[]GetDiskPerformanceDrilldownResponseEventsItem `json:"events,omitempty"`
	InvestigationQueries *map[string]string                               `json:"investigationQueries,omitempty"`
}
type GetDriveSpaceDrilldownResponseEventsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Datetime         *string                    `json:"datetime,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetDriveSpaceDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage                  `json:"-"`
	Events               *[]GetDriveSpaceDrilldownResponseEventsItem `json:"events,omitempty"`
	InvestigationQueries *map[string]string                          `json:"investigationQueries,omitempty"`
}
type GetGPUDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetNPUDrilldownResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}

type TimeRange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
type AlertsRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	AlertID   string `json:"alertId"`
}
type ApplicationRequest struct {
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	ApplicationName string `json:"applicationName"`
	DeviceName      string `json:"deviceName"`
}

type GetUserInteractionsDrilldownResponseUserDetailsDataFullName struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseUserDetailsDataDepartment struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseUserDetailsDataType struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseUserDetailsData struct {
	AdditionalFields map[string]json.RawMessage                                     `json:"-"`
	FullName         *GetUserInteractionsDrilldownResponseUserDetailsDataFullName   `json:"fullName,omitempty"`
	Department       *GetUserInteractionsDrilldownResponseUserDetailsDataDepartment `json:"department,omitempty"`
	Type             *GetUserInteractionsDrilldownResponseUserDetailsDataType       `json:"type,omitempty"`
}
type GetUserInteractionsDrilldownResponseUserDetails struct {
	AdditionalFields map[string]json.RawMessage                           `json:"-"`
	Data             *GetUserInteractionsDrilldownResponseUserDetailsData `json:"data,omitempty"`
}
type GetUserInteractionsDrilldownResponseInteractionDataDuration struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseInteractionData struct {
	AdditionalFields map[string]json.RawMessage                                   `json:"-"`
	Duration         *GetUserInteractionsDrilldownResponseInteractionDataDuration `json:"duration,omitempty"`
}
type GetUserInteractionsDrilldownResponseInteraction struct {
	AdditionalFields     map[string]json.RawMessage                           `json:"-"`
	Data                 *GetUserInteractionsDrilldownResponseInteractionData `json:"data,omitempty"`
	InvestigationQueries *map[string]string                                   `json:"investigationQueries,omitempty"`
}
type GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTime struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseLifeCyclesEventsItemType struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopVisible struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopReady struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseLifeCyclesEventsItemRemoteSessionStartupDuration struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseLifeCyclesEventsItemProfileLoadDuration struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseLifeCyclesEventsItemGpoLoadDuration struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseLifeCyclesEventsItemLogonScriptDuration struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Label            *string                    `json:"label,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Tooltip          *string                    `json:"tooltip,omitempty"`
}
type GetUserInteractionsDrilldownResponseLifeCyclesEventsItem struct {
	AdditionalFields             map[string]json.RawMessage                                                            `json:"-"`
	Time                         *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTime                         `json:"time,omitempty"`
	Type                         *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemType                         `json:"type,omitempty"`
	TimeUntilDesktopVisible      *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopVisible      `json:"timeUntilDesktopVisible,omitempty"`
	TimeUntilDesktopReady        *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemTimeUntilDesktopReady        `json:"timeUntilDesktopReady,omitempty"`
	RemoteSessionStartupDuration *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemRemoteSessionStartupDuration `json:"remoteSessionStartupDuration,omitempty"`
	ProfileLoadDuration          *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemProfileLoadDuration          `json:"profileLoadDuration,omitempty"`
	GpoLoadDuration              *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemGpoLoadDuration              `json:"gpoLoadDuration,omitempty"`
	LogonScriptDuration          *GetUserInteractionsDrilldownResponseLifeCyclesEventsItemLogonScriptDuration          `json:"logonScriptDuration,omitempty"`
}
type GetUserInteractionsDrilldownResponseLifeCycles struct {
	AdditionalFields     map[string]json.RawMessage                                  `json:"-"`
	Events               *[]GetUserInteractionsDrilldownResponseLifeCyclesEventsItem `json:"events,omitempty"`
	InvestigationQueries *map[string]string                                          `json:"investigationQueries,omitempty"`
}
type GetUserInteractionsDrilldownResponse struct {
	AdditionalFields map[string]json.RawMessage                       `json:"-"`
	UserDetails      *GetUserInteractionsDrilldownResponseUserDetails `json:"userDetails,omitempty"`
	Interaction      *GetUserInteractionsDrilldownResponseInteraction `json:"interaction,omitempty"`
	LifeCycles       *GetUserInteractionsDrilldownResponseLifeCycles  `json:"lifeCycles,omitempty"`
}

type GetRoundTripTimeDrilldownResponseRtt struct {
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetRoundTripTimeDrilldownResponse struct {
	AdditionalFields map[string]json.RawMessage            `json:"-"`
	Rtt              *GetRoundTripTimeDrilldownResponseRtt `json:"rtt,omitempty"`
}

type GetNetworkLatencyDrilldownResponseNetworkLatency struct {
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	InvestigationQueries *map[string]string         `json:"investigationQueries,omitempty"`
}
type GetNetworkLatencyDrilldownResponse struct {
	AdditionalFields map[string]json.RawMessage                        `json:"-"`
	NetworkLatency   *GetNetworkLatencyDrilldownResponseNetworkLatency `json:"networkLatency,omitempty"`
}
