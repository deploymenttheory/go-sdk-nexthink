# SupportTimeline browser API examples

Use `NEXTHINK_API=web` and browser authentication through the SDK root client. Each method folder has the exact inputs and a runnable example.

For device calls, set `NEXTHINK_DEVICE_ID` to the Collector UID (`device.collector.uid`), not the separate device UID (`device.uid`). The UI routes label this parameter `deviceID`, but supplying the device UID can return empty results instead of an error. Obtain both identifiers with `devices | list device.uid, device.collector.uid, device.name` and select the intended device.

These private UI contracts may vary by tenant version. Analytics endpoints expose reads; lifecycle operations belong to their configuration resources. UTC headers are the default; services accept `WithTimeZone` for the browser time context.

- [GetAlertsAndErrors](GetAlertsAndErrors/README.md)
- [GetPerformance](GetPerformance/README.md)
- [GetConnectivity](GetConnectivity/README.md)
- [GetApplicationConnectivity](GetApplicationConnectivity/README.md)
- [GetActivity](GetActivity/README.md)
- [GetApplications](GetApplications/README.md)
- [GetUserInteractions](GetUserInteractions/README.md)
- [GetCollaboration](GetCollaboration/README.md)
- [GetErrorsDrilldown](GetErrorsDrilldown/README.md)
- [GetFreezesDrilldown](GetFreezesDrilldown/README.md)
- [GetAlertsDrilldown](GetAlertsDrilldown/README.md)
- [GetActionsDrilldown](GetActionsDrilldown/README.md)
- [GetSystemBootsDrilldown](GetSystemBootsDrilldown/README.md)
- [GetSystemBootsAndSuspendsDrilldown](GetSystemBootsAndSuspendsDrilldown/README.md)
- [GetTeamsCallsDrilldown](GetTeamsCallsDrilldown/README.md)
- [GetZoomCallsDrilldown](GetZoomCallsDrilldown/README.md)
- [GetDesktopConnectivityDrilldown](GetDesktopConnectivityDrilldown/README.md)
- [GetWebConnectivityDrilldown](GetWebConnectivityDrilldown/README.md)
- [GetDesktopApplicationsDrilldown](GetDesktopApplicationsDrilldown/README.md)
- [GetWebApplicationsDrilldown](GetWebApplicationsDrilldown/README.md)
- [GetInstallationsDrilldown](GetInstallationsDrilldown/README.md)
- [GetEthernetDrilldown](GetEthernetDrilldown/README.md)
- [GetWiFiDrilldown](GetWiFiDrilldown/README.md)
- [GetConnectionsDrilldown](GetConnectionsDrilldown/README.md)
- [GetNetworkApplicationDrilldown](GetNetworkApplicationDrilldown/README.md)
- [GetCPUDrilldown](GetCPUDrilldown/README.md)
- [GetMemoryDrilldown](GetMemoryDrilldown/README.md)
- [GetDiskPerformanceDrilldown](GetDiskPerformanceDrilldown/README.md)
- [GetDriveSpaceDrilldown](GetDriveSpaceDrilldown/README.md)
- [GetGPUDrilldown](GetGPUDrilldown/README.md)
- [GetNPUDrilldown](GetNPUDrilldown/README.md)
- [GetUserInteractionsDrilldown](GetUserInteractionsDrilldown/README.md)
- [GetRoundTripTimeDrilldown](GetRoundTripTimeDrilldown/README.md)
- [GetNetworkLatencyDrilldown](GetNetworkLatencyDrilldown/README.md)
