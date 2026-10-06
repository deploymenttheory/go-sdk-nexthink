package support_timeline

// EndpointGetAlertsAndErrors is used by the first-party browser UI.
const EndpointGetAlertsAndErrors = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/timeline/alertsAndErrors"

// EndpointGetPerformance is used by the first-party browser UI.
const EndpointGetPerformance = "/apigateway/atl/support-device-timeline-be/api/v2/device/{deviceID}/timeline/performance"

// EndpointGetConnectivity is used by the first-party browser UI.
const EndpointGetConnectivity = "/apigateway/atl/support-device-timeline-be/api/v5/device/{deviceID}/timeline/connectivity"

// EndpointGetApplicationConnectivity is used by the first-party browser UI.
const EndpointGetApplicationConnectivity = "/apigateway/atl/support-device-timeline-be/api/v5/device/{deviceID}/timeline/connectivityApplications"

// EndpointGetActivity is used by the first-party browser UI.
const EndpointGetActivity = "/apigateway/atl/support-device-timeline-be/api/v5/device/{deviceID}/timeline/activity"

// EndpointGetApplications is used by the first-party browser UI.
const EndpointGetApplications = "/apigateway/atl/support-device-timeline-be/api/v5/device/{deviceID}/timeline/appex"

// EndpointGetUserInteractions is used by the first-party browser UI.
const EndpointGetUserInteractions = "/apigateway/atl/support-device-timeline-be/api/v2/device/{deviceID}/timeline/usersInteractions"

// EndpointGetCollaboration is used by the first-party browser UI.
const EndpointGetCollaboration = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/timeline/collaboration"

// EndpointGetErrorsDrilldown is used by the first-party browser UI.
const EndpointGetErrorsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v2/device/{deviceID}/drilldowns/errors"

// EndpointGetFreezesDrilldown is used by the first-party browser UI.
const EndpointGetFreezesDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/freezes"

// EndpointGetAlertsDrilldown is used by the first-party browser UI.
const EndpointGetAlertsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v2/device/{deviceID}/drilldowns/alerts"

// EndpointGetActionsDrilldown is used by the first-party browser UI.
const EndpointGetActionsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/actions"

// EndpointGetSystemBootsDrilldown is used by the first-party browser UI.
const EndpointGetSystemBootsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/systemboots"

// EndpointGetSystemBootsAndSuspendsDrilldown is used by the first-party browser UI.
const EndpointGetSystemBootsAndSuspendsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v2/device/{deviceID}/drilldowns/systemboots"

// EndpointGetTeamsCallsDrilldown is used by the first-party browser UI.
const EndpointGetTeamsCallsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v2/device/{deviceID}/drilldowns/msteams"

// EndpointGetZoomCallsDrilldown is used by the first-party browser UI.
const EndpointGetZoomCallsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v2/device/{deviceID}/drilldowns/zoom"

// EndpointGetDesktopConnectivityDrilldown is used by the first-party browser UI.
const EndpointGetDesktopConnectivityDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/connectivity/applications/desktop"

// EndpointGetWebConnectivityDrilldown is used by the first-party browser UI.
const EndpointGetWebConnectivityDrilldown = "/apigateway/atl/support-device-timeline-be/api/v2/device/{deviceID}/drilldowns/connectivity/applications/web"

// EndpointGetDesktopApplicationsDrilldown is used by the first-party browser UI.
const EndpointGetDesktopApplicationsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/desktopapplications"

// EndpointGetWebApplicationsDrilldown is used by the first-party browser UI.
const EndpointGetWebApplicationsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v3/device/{deviceID}/drilldowns/webapplications"

// EndpointGetInstallationsDrilldown is used by the first-party browser UI.
const EndpointGetInstallationsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/installationevents"

// EndpointGetEthernetDrilldown is used by the first-party browser UI.
const EndpointGetEthernetDrilldown = "/apigateway/atl/support-device-timeline-be/api/v2/device/{deviceID}/drilldowns/ethernet"

// EndpointGetWiFiDrilldown is used by the first-party browser UI.
const EndpointGetWiFiDrilldown = "/apigateway/atl/support-device-timeline-be/api/v2/device/{deviceID}/drilldowns/wifi"

// EndpointGetConnectionsDrilldown is used by the first-party browser UI.
const EndpointGetConnectionsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/connectionevents"

// EndpointGetNetworkApplicationDrilldown is used by the first-party browser UI.
const EndpointGetNetworkApplicationDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/networkapp"

// EndpointGetCPUDrilldown is used by the first-party browser UI.
const EndpointGetCPUDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/cpu"

// EndpointGetMemoryDrilldown is used by the first-party browser UI.
const EndpointGetMemoryDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/memory"

// EndpointGetDiskPerformanceDrilldown is used by the first-party browser UI.
const EndpointGetDiskPerformanceDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/diskperformance"

// EndpointGetDriveSpaceDrilldown is used by the first-party browser UI.
const EndpointGetDriveSpaceDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/drivespace"

// EndpointGetGPUDrilldown is used by the first-party browser UI.
const EndpointGetGPUDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/gpu/{slot}"

// EndpointGetNPUDrilldown is used by the first-party browser UI.
const EndpointGetNPUDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/drilldowns/npu"

// EndpointGetUserInteractionsDrilldown is used by the first-party browser UI.
const EndpointGetUserInteractionsDrilldown = "/apigateway/atl/support-device-timeline-be/api/v3/device/{deviceID}/drilldowns/user/{userID}/userinteractions"

// EndpointGetRoundTripTimeDrilldown is used by the first-party browser UI.
const EndpointGetRoundTripTimeDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/user/{userID}/drilldowns/rtt"

// EndpointGetNetworkLatencyDrilldown is used by the first-party browser UI.
const EndpointGetNetworkLatencyDrilldown = "/apigateway/atl/support-device-timeline-be/api/v1/device/{deviceID}/user/{userID}/drilldowns/networklatency"
