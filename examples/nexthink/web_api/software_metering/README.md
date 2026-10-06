# SoftwareMetering examples

Configure `NEXTHINK_API=web`, instance, region, and browser authentication as described in the [example index](../README.md). Run examples from the repository root.

| Example | Required inputs |
| --- | --- |
| [List](List/main.go) | None (first page where paginated) |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` matching [Create_input.json](../../../../nexthink/web_api/software_metering/mocks/Create_input.json) |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` matching [Update_input.json](../../../../nexthink/web_api/software_metering/mocks/Update_input.json) and `NEXTHINK_CONTENT_ID` |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` matching [Delete_input.json](../../../../nexthink/web_api/software_metering/mocks/Delete_input.json) |

```sh
go run ./examples/nexthink/web_api/software_metering/List
```

Create returns a boolean; use List to obtain the new UUID. Thresholds must cover the selected applications. Hybrid applications use WEB and DESKTOP thresholds.

Replace synthetic IDs and revisions with the object you intend to manage. Create and Update write the supplied object; Delete removes it. GraphQL examples print partial data before reporting errors so returned identifiers remain available.

The following browser UI analytics operations also have runnable examples. Each uses `NEXTHINK_REQUEST_FILE` with the linked JSON.

| Example | Request |
| --- | --- |
| [AutoConfigureMetering](AutoConfigureMetering/main.go) | [AutoConfigureMeteringRequest](AutoConfigureMetering/request.example.json) |
| [GetApplications](GetApplications/main.go) | [GetApplicationsRequest](GetApplications/request.example.json) |
| [GetConfigurationByApplicationUUID](GetConfigurationByApplicationUUID/main.go) | [GetConfigurationByApplicationUUIDRequest](GetConfigurationByApplicationUUID/request.example.json) |
| [GetConfigurationDetails](GetConfigurationDetails/main.go) | [GetConfigurationDetailsRequest](GetConfigurationDetails/request.example.json) |
| [GetEmployeesTable](GetEmployeesTable/main.go) | [GetEmployeesTableRequest](GetEmployeesTable/request.example.json) |
| [GetPackages](GetPackages/main.go) | [GetPackagesRequest](GetPackages/request.example.json) |
| [GetUsageBreakdown](GetUsageBreakdown/main.go) | [GetUsageBreakdownRequest](GetUsageBreakdown/request.example.json) |
| [GetUsageByLicenseEndpoint](GetUsageByLicenseEndpoint/main.go) | [GetUsageByLicenseEndpointRequest](GetUsageByLicenseEndpoint/request.example.json) |
| [GetUsageByLicenseEndpointCount](GetUsageByLicenseEndpointCount/main.go) | [GetUsageByLicenseEndpointCountRequest](GetUsageByLicenseEndpointCount/request.example.json) |
| [GetUsageDistribution](GetUsageDistribution/main.go) | [GetUsageDistributionRequest](GetUsageDistribution/request.example.json) |
| [GetUsageDistributionByCategory](GetUsageDistributionByCategory/main.go) | [GetUsageDistributionByCategoryRequest](GetUsageDistributionByCategory/request.example.json) |
| [GetUsageOverview](GetUsageOverview/main.go) | [GetUsageOverviewRequest](GetUsageOverview/request.example.json) |
| [GetConfigurationUsageOverview](GetConfigurationUsageOverview/main.go) | [GetConfigurationUsageOverviewRequest](GetConfigurationUsageOverview/request.example.json) |

`GetUsageOverview` selects an application within a configuration; `GetConfigurationUsageOverview` covers the complete configuration and supports filters. `GetApplications` accepts an empty object to list available applications. `GetPackages` allows an empty search string. Relative time ranges use server enums such as `PAST_90_DAYS`; pagination includes `offset` and `size`.

`AutoConfigureMetering` creates a metering configuration when required and can return `ALREADY_CONFIGURED` on repetition. Select an intended lab application before running it. It was validated on a disposable application, followed by configuration and application cleanup. All analytics reads were exercised; `GetUsageDistribution` returned a backend GraphQL execution error on the empty lab fixture and the SDK preserved that error. Populated success fixtures cover this response shape independently.
