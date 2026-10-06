# Runnable web API examples

Every exported resource method has an example calling `client.WebAPI` through `nexthink.NewClientFromEnv`. The coverage test checks both API families so new methods cannot silently omit examples.

From the repository root:

```sh
export NEXTHINK_API=web
export NEXTHINK_INSTANCE=your-instance
export NEXTHINK_REGION=eu
export NEXTHINK_WEB_AUTH=chrome
go run ./examples/nexthink/web_api/workflows/List
```

Chrome must already be signed into that instance. Token authentication is also supported; see [authentication](../../../docs/web-api.md). Write examples require an explicit `NEXTHINK_REQUEST_FILE`. Replace synthetic IDs, names and revisions with the dedicated object you intend to manage. The JSON files below are SDK inputs, not GraphQL request envelopes unless the example is `GraphQL.Execute`.

| Resource/method | Example inputs beyond authentication | JSON input |
| --- | --- | --- |
| [connectors.StartTest](connectors/StartTest/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](connectors/StartTest/request.example.json) |
| [connectors.GetTest](connectors/GetTest/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [workflows.ListConnectorDefinitions](workflows/ListConnectorDefinitions/main.go) | None | — |
| [workflows.ListConnectorCredentials](workflows/ListConnectorCredentials/main.go) | None | — |
| [knowledge_bases.List](knowledge_bases/List/main.go) | None | — |
| [knowledge_bases.GetContents](knowledge_bases/GetContents/main.go) | None | — |
| [knowledge_bases.Create](knowledge_bases/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](knowledge_bases/Create/request.example.json) |
| [knowledge_bases.Delete](knowledge_bases/Delete/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [knowledge_bases.GetDownloadURL](knowledge_bases/GetDownloadURL/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [knowledge_bases.UploadFile](knowledge_bases/UploadFile/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](knowledge_bases/UploadFile/request.example.json) |
| [knowledge_bases.StartMultipartUpload](knowledge_bases/StartMultipartUpload/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](knowledge_bases/StartMultipartUpload/request.example.json) |
| [knowledge_bases.UploadPart](knowledge_bases/UploadPart/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](knowledge_bases/UploadPart/request.example.json) |
| [knowledge_bases.CompleteMultipartUpload](knowledge_bases/CompleteMultipartUpload/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](knowledge_bases/CompleteMultipartUpload/request.example.json) |
| [legacy_connectors.List](legacy_connectors/List/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](legacy_connectors/List/request.example.json) |
| [legacy_connectors.Get](legacy_connectors/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [legacy_connectors.Create](legacy_connectors/Create/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_CONTENT_ID` | [JSON](legacy_connectors/Create/request.example.json) |
| [legacy_connectors.Update](legacy_connectors/Update/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_CONTENT_ID` | [JSON](legacy_connectors/Update/request.example.json) |
| [legacy_connectors.Delete](legacy_connectors/Delete/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [legacy_connectors.SaveSecrets](legacy_connectors/SaveSecrets/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_CONTENT_ID` | [JSON](legacy_connectors/SaveSecrets/request.example.json) |
| [webhooks.List](webhooks/List/main.go) | None | — |
| [webhooks.Get](webhooks/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [webhooks.Create](webhooks/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](webhooks/Create/request.example.json) |
| [webhooks.Update](webhooks/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](webhooks/Update/request.example.json) |
| [webhooks.Delete](webhooks/Delete/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [webhooks.GetAvailability](webhooks/GetAvailability/main.go) | None | — |
| [webhooks.Test](webhooks/Test/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](webhooks/Test/request.example.json) |
| [data_exporters.List](data_exporters/List/main.go) | None | — |
| [data_exporters.Get](data_exporters/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [data_exporters.Create](data_exporters/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](data_exporters/Create/request.example.json) |
| [data_exporters.Update](data_exporters/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](data_exporters/Update/request.example.json) |
| [data_exporters.Delete](data_exporters/Delete/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [data_exporters.GetCustomerInfo](data_exporters/GetCustomerInfo/main.go) | None | — |
| [data_exporters.ListStatuses](data_exporters/ListStatuses/main.go) | None | — |
| [data_exporters.GetPlaceholders](data_exporters/GetPlaceholders/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](data_exporters/GetPlaceholders/request.example.json) |
| [data_exporters.StartTest](data_exporters/StartTest/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](data_exporters/StartTest/request.example.json) |
| [data_exporters.GetTest](data_exporters/GetTest/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_EXECUTION_ID` | — |
| [dashboards.CreateWidget](dashboards/CreateWidget/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/CreateWidget/request.example.json) |
| [dashboards.UpdateWidget](dashboards/UpdateWidget/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/UpdateWidget/request.example.json) |
| [dashboards.DeleteWidget](dashboards/DeleteWidget/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/DeleteWidget/request.example.json) |
| [dashboards.CreateFilter](dashboards/CreateFilter/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/CreateFilter/request.example.json) |
| [dashboards.UpdateFilter](dashboards/UpdateFilter/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/UpdateFilter/request.example.json) |
| [dashboards.DeleteFilter](dashboards/DeleteFilter/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/DeleteFilter/request.example.json) |
| [dashboards.CreateTab](dashboards/CreateTab/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/CreateTab/request.example.json) |
| [dashboards.UpdateTab](dashboards/UpdateTab/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/UpdateTab/request.example.json) |
| [dashboards.UpdateTabs](dashboards/UpdateTabs/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/UpdateTabs/request.example.json) |
| [dashboards.DeleteTab](dashboards/DeleteTab/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/DeleteTab/request.example.json) |
| [dashboards.UpdateLayout](dashboards/UpdateLayout/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/UpdateLayout/request.example.json) |
| [dashboards.Export](dashboards/Export/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/Export/request.example.json) |
| [dashboards.Duplicate](dashboards/Duplicate/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/Duplicate/request.example.json) |
| [dashboards.Import](dashboards/Import/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/Import/request.example.json) |
| [checklists.Export](checklists/Export/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [checklists.Import](checklists/Import/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](checklists/Import/request.example.json) |
| [checklists.ListGroupedFields](checklists/ListGroupedFields/main.go) | None | — |
| [monitors.Export](monitors/Export/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [monitors.ExportLibrary](monitors/ExportLibrary/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [monitors.Import](monitors/Import/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](monitors/Import/request.example.json) |
| [connectors.List](connectors/List/main.go) | None | — |
| [connectors.Get](connectors/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [connectors.Create](connectors/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](connectors/Create/request.example.json) |
| [connectors.Update](connectors/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](connectors/Update/request.example.json) |
| [connectors.Delete](connectors/Delete/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [connectors.ListTemplates](connectors/ListTemplates/main.go) | None | — |
| [connectors.GetTemplate](connectors/GetTemplate/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [connectors.ListManualCustomFields](connectors/ListManualCustomFields/main.go) | `NEXTHINK_DATA_MODEL_OBJECT` | — |
| [connector_credentials.List](connector_credentials/List/main.go) | None | — |
| [connector_credentials.ListIDs](connector_credentials/ListIDs/main.go) | None | — |
| [connector_credentials.Get](connector_credentials/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [connector_credentials.Create](connector_credentials/Create/main.go) | `NEXTHINK_CONTENT_ID` and `NEXTHINK_REQUEST_FILE` | [JSON](connector_credentials/Create/request.example.json) |
| [connector_credentials.Update](connector_credentials/Update/main.go) | `NEXTHINK_CONTENT_ID` and `NEXTHINK_REQUEST_FILE` | [JSON](connector_credentials/Update/request.example.json) |
| [connector_credentials.Delete](connector_credentials/Delete/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [applications.List](applications/List/main.go) | None | — |
| [applications.Get](applications/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [applications.Create](applications/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/applications/mocks/Create_input.json) |
| [applications.Update](applications/Update/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/applications/mocks/Update_input.json) |
| [applications.Delete](applications/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/applications/mocks/Delete_input.json) |
| [campaigns.Create](campaigns/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/campaigns/mocks/Create_input.json) |
| [campaigns.Update](campaigns/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/campaigns/mocks/Update_input.json) |
| [campaigns.Delete](campaigns/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/campaigns/mocks/Delete_input.json) |
| [campaigns.Get](campaigns/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [campaigns.List](campaigns/List/main.go) | None | — |
| [collector_management.GetDownloadLinks](collector_management/GetDownloadLinks/main.go) | None | — |
| [collector_management.GetUpdateConfiguration](collector_management/GetUpdateConfiguration/main.go) | None | — |
| [collector_management.GetPlatforms](collector_management/GetPlatforms/main.go) | None | — |
| [collector_management.GetVersions](collector_management/GetVersions/main.go) | None | — |
| [collector_management.GetGroups](collector_management/GetGroups/main.go) | None | — |
| [collector_management.GetTargetVersions](collector_management/GetTargetVersions/main.go) | None | — |
| [collector_management.SetUpdateConfiguration](collector_management/SetUpdateConfiguration/main.go) | `NEXTHINK_REQUEST_FILE` | Caller-supplied payload; see validation limits below |
| [content_administration.GetConfiguration](content_administration/GetConfiguration/main.go) | `NEXTHINK_CONTENT_CONFIGURATION` | — |
| [content_administration.List](content_administration/List/main.go) | None | — |
| [custom_fields.Create](custom_fields/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/custom_fields/mocks/Create_input.json) |
| [custom_fields.Update](custom_fields/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/custom_fields/mocks/Update_input.json) |
| [custom_fields.Delete](custom_fields/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/custom_fields/mocks/Delete_input.json) |
| [custom_fields.Get](custom_fields/Get/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_CUSTOM_FIELD_TYPE` | — |
| [custom_fields.List](custom_fields/List/main.go) | None | — |
| [device_configuration.GetProfiles](device_configuration/GetProfiles/main.go) | None | — |
| [device_configuration.SetProfiles](device_configuration/SetProfiles/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](device_configuration/SetProfiles/request.example.json) |
| [device_configuration.GetSettings](device_configuration/GetSettings/main.go) | None | — |
| [device_configuration.SetSettings](device_configuration/SetSettings/main.go) | `NEXTHINK_REQUEST_FILE` | Caller-supplied payload; see validation limits below |
| [graphql.Execute](graphql/Execute/main.go) | `NEXTHINK_GRAPHQL_OPERATION`, `NEXTHINK_REQUEST_FILE` | [JSON](graphql/Execute/request.example.json) |
| [license.GetFeatureStatus](license/GetFeatureStatus/main.go) | None | — |
| [monitors.Create](monitors/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/monitors/mocks/Create_input.json) |
| [monitors.Update](monitors/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/monitors/mocks/Update_input.json) |
| [monitors.Delete](monitors/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/monitors/mocks/Delete_input.json) |
| [monitors.Get](monitors/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [monitors.List](monitors/List/main.go) | None | — |
| [nql_editor.Validate](nql_editor/Validate/main.go) | None | [JSON](nql_editor/Validate/request.example.json) |
| [nql_editor.Complete](nql_editor/Complete/main.go) | `NEXTHINK_NQL_MODE`, `NEXTHINK_REQUEST_FILE` | [JSON](nql_editor/Complete/request.example.json) |
| [nql_editor.Hover](nql_editor/Hover/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](nql_editor/Hover/request.example.json) |
| [nql_editor.Resolve](nql_editor/Resolve/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](nql_editor/Resolve/request.example.json) |
| [nql_editor.GetHighlighting](nql_editor/GetHighlighting/main.go) | None | — |
| [nql_queries.Get](nql_queries/Get/main.go) | `NEXTHINK_QUERY_CONTENT_ID` | — |
| [nql_queries.Create](nql_queries/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](nql_queries/Create/request.example.json) |
| [nql_queries.Update](nql_queries/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](nql_queries/Update/request.example.json) |
| [nql_queries.Delete](nql_queries/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](nql_queries/Delete/request.example.json) |
| [nql_queries.List](nql_queries/List/main.go) | None | — |
| [product_shell.GetMenu](product_shell/GetMenu/main.go) | None | — |
| [product_shell.GetUser](product_shell/GetUser/main.go) | None | — |
| [product_shell.GetModules](product_shell/GetModules/main.go) | None | — |
| [product_shell.GetConfiguration](product_shell/GetConfiguration/main.go) | None | — |
| [product_shell.GetFlag](product_shell/GetFlag/main.go) | `NEXTHINK_FLAG` | — |
| [product_shell.GetDynamicMenu](product_shell/GetDynamicMenu/main.go) | `NEXTHINK_MENU` | — |
| [product_shell.ValidateClaims](product_shell/ValidateClaims/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](product_shell/ValidateClaims/request.example.json) |
| [product_shell.PostTelemetry](product_shell/PostTelemetry/main.go) | `NEXTHINK_REQUEST_FILE` | Caller-supplied payload; see validation limits below |
| [remote_actions.Get](remote_actions/Get/main.go) | `NEXTHINK_REMOTE_ACTION_UUID` | — |
| [remote_actions.GetForView](remote_actions/GetForView/main.go) | `NEXTHINK_REMOTE_ACTION_UUID` | — |
| [remote_actions.GetContentVolume](remote_actions/GetContentVolume/main.go) | None | — |
| [remote_actions.InspectBashScript](remote_actions/InspectBashScript/main.go) | `NEXTHINK_SCRIPT_FILE` | — |
| [remote_actions.InspectPowerShellScript](remote_actions/InspectPowerShellScript/main.go) | `NEXTHINK_SCRIPT_FILE` | — |
| [remote_actions.GetPowerShellSignature](remote_actions/GetPowerShellSignature/main.go) | `NEXTHINK_SCRIPT_FILE` | — |
| [remote_actions.Create](remote_actions/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](remote_actions/Create/request.example.json) |
| [remote_actions.Update](remote_actions/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](remote_actions/Update/request.example.json) |
| [remote_actions.Delete](remote_actions/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](remote_actions/Delete/request.example.json) |
| [remote_actions.List](remote_actions/List/main.go) | None | — |
| [rule_based_custom_fields.List](rule_based_custom_fields/List/main.go) | None | — |
| [rule_based_custom_fields.Get](rule_based_custom_fields/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [rule_based_custom_fields.Create](rule_based_custom_fields/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/rule_based_custom_fields/mocks/Create_input.json) |
| [rule_based_custom_fields.Update](rule_based_custom_fields/Update/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/rule_based_custom_fields/mocks/Update_input.json) |
| [rule_based_custom_fields.Delete](rule_based_custom_fields/Delete/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/rule_based_custom_fields/mocks/Delete_input.json) |
| [software_metering.List](software_metering/List/main.go) | None | — |
| [software_metering.Get](software_metering/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [software_metering.Create](software_metering/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/software_metering/mocks/Create_input.json) |
| [software_metering.Update](software_metering/Update/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/software_metering/mocks/Update_input.json) |
| [software_metering.Delete](software_metering/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/software_metering/mocks/Delete_input.json) |
| [workflows.List](workflows/List/main.go) | None | — |
| [workflows.Get](workflows/Get/main.go) | `NEXTHINK_WORKFLOW_UUID` | — |
| [workflows.Export](workflows/Export/main.go) | `NEXTHINK_WORKFLOW_UUID` | — |
| [workflows.Create](workflows/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](workflows/Create/request.example.json) |
| [workflows.Update](workflows/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](workflows/Update/request.example.json) |
| [workflows.Delete](workflows/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](workflows/Delete/request.example.json) |
| [writing_assistant.Get](writing_assistant/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [writing_assistant.Create](writing_assistant/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/writing_assistant/mocks/Create_input.json) |
| [writing_assistant.Update](writing_assistant/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/writing_assistant/mocks/Update_input.json) |
| [writing_assistant.Delete](writing_assistant/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](../../../nexthink/web_api/writing_assistant/mocks/Delete_input.json) |
| [writing_assistant.List](writing_assistant/List/main.go) | None | — |

| [assets.Create](assets/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](assets/Create/request.example.json) |
| [assets.Delete](assets/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](assets/Delete/request.example.json) |
| [assets.GetSignedURL](assets/GetSignedURL/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [assets.List](assets/List/main.go) | None | — |
| [assets.Update](assets/Update/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE` | [JSON](assets/Update/request.example.json) |
| [checklists.Create](checklists/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](checklists/Create/request.example.json) |
| [checklists.Delete](checklists/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](checklists/Delete/request.example.json) |
| [checklists.Get](checklists/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [checklists.List](checklists/List/main.go) | None | — |
| [checklists.Update](checklists/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](checklists/Update/request.example.json) |
| [dashboards.Create](dashboards/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/Create/request.example.json) |
| [dashboards.Delete](dashboards/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/Delete/request.example.json) |
| [dashboards.Get](dashboards/Get/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_PRODUCT_AREA` | — |
| [dashboards.List](dashboards/List/main.go) | None | — |
| [dashboards.Update](dashboards/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](dashboards/Update/request.example.json) |
| [ratings.Create](ratings/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](ratings/Create/request.example.json) |
| [ratings.Delete](ratings/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](ratings/Delete/request.example.json) |
| [ratings.Get](ratings/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [ratings.List](ratings/List/main.go) | None | — |
| [ratings.Update](ratings/Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](ratings/Update/request.example.json) |
| [investigations.Create](investigations/Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](investigations/Create/request.example.json) |
| [investigations.Delete](investigations/Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](investigations/Delete/request.example.json) |
| [investigations.Export](investigations/Export/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [investigations.Get](investigations/Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [investigations.Import](investigations/Import/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](investigations/Import/request.example.json) |
| [investigations.List](investigations/List/main.go) | None | — |
| [investigations.Update](investigations/Update/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE` | [JSON](investigations/Update/request.example.json) |

`NEXTHINK_NQL_MODE` is optional. `NEXTHINK_MENU` is a menu API ID from `GetMenu` (for example `bus-menu`), not a display name. Custom Fields Get requires `NEXTHINK_CUSTOM_FIELD_TYPE=MANUAL` or `COMPUTED`. Software Metering Create returns a boolean; obtain its UUID through List. Monitor Update requires the content ID, revision, and the separate monitor UUID returned by Get. Campaign Create saves a draft; these examples do not publish it.

Saved NQL query Create returns a server-generated content ID. Use the returned ID for Get, Update and Delete; the server did not retain the content ID supplied on creation in the lab.

Use `ListOptions` to paginate Applications and Campaigns when the first page is insufficient. GraphQL examples print partial data before returning an error. Inspect `graphql.GraphQLErrors` using `errors.As` in application code.

Workflow Export writes the opaque export to stdout and HTTP metadata to stderr. Redirect stdout to retain the export, for example `go run ./examples/nexthink/web_api/workflows/Export > workflow-export.txt`.

Collector update configuration, legacy device settings, and telemetry remain less validated than the management lifecycles. They accept caller-supplied JSON rather than a fabricated successful payload. The legacy settings route returned 403 in this lab. See the [coverage audit](../../../docs/web-api-coverage.md) for the distinction between compiled examples and successful live validation.

## Analytics, metadata and library additions

These services use the same root client and browser authentication. Resource guides list every operation, request sample and live-validation limitation. Most analytics calls require filters and a time range; library installation and configuration writes require explicit request files.

| Resource | Guide |
| --- | --- |
| `application_experience` | [Examples and inputs](application_experience/README.md) |
| `software_metering` | [Examples and inputs](software_metering/README.md) |
| `alert_hub` | [Examples and inputs](alert_hub/README.md) |
| `diagnostics` | [Examples and inputs](diagnostics/README.md) |
| `benchmark` | [Examples and inputs](benchmark/README.md) |
| `dex_scores` | [Examples and inputs](dex_scores/README.md) |
| `dex_configuration` | [Examples and inputs](dex_configuration/README.md) |
| `cci_insights` | [Examples and inputs](cci_insights/README.md) |
| `network_insights` | [Examples and inputs](network_insights/README.md) |
| `data_exploration` | [Examples and inputs](data_exploration/README.md) |
| `library` | [Examples and inputs](library/README.md) |
| `workflows` | [Examples and inputs](workflows/README.md) |
| `remote_actions` | [Examples and inputs](remote_actions/README.md) |
| `campaigns` | [Examples and inputs](campaigns/README.md) |
| `monitors` | [Examples and inputs](monitors/README.md) |
| `ratings` | [Examples and inputs](ratings/README.md) |
| `checklists` | [Examples and inputs](checklists/README.md) |
| `investigations` | [Examples and inputs](investigations/README.md) |
| `dashboards` | [Examples and inputs](dashboards/README.md) |
