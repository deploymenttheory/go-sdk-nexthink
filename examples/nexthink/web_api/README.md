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

`NEXTHINK_NQL_MODE` is optional. `NEXTHINK_MENU` is a menu API ID from `GetMenu` (for example `bus-menu`), not a display name. Custom Fields Get requires `NEXTHINK_CUSTOM_FIELD_TYPE=MANUAL` or `COMPUTED`. Software Metering Create returns a boolean; obtain its UUID through List. Monitor Update requires the content ID, revision, and the separate monitor UUID returned by Get. Campaign Create saves a draft; these examples do not publish it.

Use `ListOptions` to paginate Applications and Campaigns when the first page is insufficient. GraphQL examples print partial data before returning an error. Inspect `graphql.GraphQLErrors` using `errors.As` in application code.

Collector update configuration, legacy device settings, and telemetry remain less validated than the management lifecycles. They accept caller-supplied JSON rather than a fabricated successful payload. The legacy settings route returned 403 in this lab. See the [coverage audit](../../../docs/web-api-coverage.md) for the distinction between compiled examples and successful live validation.
