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

## Product-permission discovery additions

The following operations cover AI Tools/governance, Amplify, Workspace and device/user/product administration. Each links to a resource-local program and request guide. Authentication uses the same root client, including headless password authentication for CI. See the [acceptance report](../../../docs/acceptance/README.md) for live passes and unavailable features.

| Resource/method | Example inputs beyond authentication | JSON input |
| --- | --- | --- |
| [action_executions.GetDeviceActions](action_executions/GetDeviceActions/main.go) | `NEXTHINK_DEVICE_ID`; [guide](action_executions/GetDeviceActions/README.md) | — |
| [ai_tools.CheckCopilotCredentials](ai_tools/CheckCopilotCredentials/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](ai_tools/CheckCopilotCredentials/README.md) | [JSON](ai_tools/CheckCopilotCredentials/request.example.json) |
| [ai_tools.Create](ai_tools/Create/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](ai_tools/Create/README.md) | [JSON](ai_tools/Create/request.example.json) |
| [ai_tools.CreateCopilot](ai_tools/CreateCopilot/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](ai_tools/CreateCopilot/README.md) | [JSON](ai_tools/CreateCopilot/request.example.json) |
| [ai_tools.CreateGoal](ai_tools/CreateGoal/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](ai_tools/CreateGoal/README.md) | [JSON](ai_tools/CreateGoal/request.example.json) |
| [ai_tools.CreateModule](ai_tools/CreateModule/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](ai_tools/CreateModule/README.md) | [JSON](ai_tools/CreateModule/request.example.json) |
| [ai_tools.Delete](ai_tools/Delete/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REVISION`; [guide](ai_tools/Delete/README.md) | — |
| [ai_tools.DeleteCopilot](ai_tools/DeleteCopilot/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REVISION`; [guide](ai_tools/DeleteCopilot/README.md) | — |
| [ai_tools.DeleteGoal](ai_tools/DeleteGoal/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REVISION`; [guide](ai_tools/DeleteGoal/README.md) | — |
| [ai_tools.Get](ai_tools/Get/main.go) | `NEXTHINK_CONTENT_ID`; [guide](ai_tools/Get/README.md) | — |
| [ai_tools.GetCopilot](ai_tools/GetCopilot/main.go) | `NEXTHINK_CONTENT_ID`; [guide](ai_tools/GetCopilot/README.md) | — |
| [ai_tools.GetGoal](ai_tools/GetGoal/main.go) | `NEXTHINK_CONTENT_ID`; [guide](ai_tools/GetGoal/README.md) | — |
| [ai_tools.GetGoalInsights](ai_tools/GetGoalInsights/main.go) | `NEXTHINK_CONTENT_ID`; [guide](ai_tools/GetGoalInsights/README.md) | — |
| [ai_tools.GetGovernanceActiveUsers](ai_tools/GetGovernanceActiveUsers/main.go) | None; [guide](ai_tools/GetGovernanceActiveUsers/README.md) | — |
| [ai_tools.GetGovernanceDashboard](ai_tools/GetGovernanceDashboard/main.go) | None; [guide](ai_tools/GetGovernanceDashboard/README.md) | — |
| [ai_tools.GetGovernanceTrends](ai_tools/GetGovernanceTrends/main.go) | None; [guide](ai_tools/GetGovernanceTrends/README.md) | — |
| [ai_tools.GetLegacyTool](ai_tools/GetLegacyTool/main.go) | `NEXTHINK_CONTENT_ID`; [guide](ai_tools/GetLegacyTool/README.md) | — |
| [ai_tools.GetLicense](ai_tools/GetLicense/main.go) | None; [guide](ai_tools/GetLicense/README.md) | — |
| [ai_tools.GetModule](ai_tools/GetModule/main.go) | None; [guide](ai_tools/GetModule/README.md) | — |
| [ai_tools.GetOverviewInsights](ai_tools/GetOverviewInsights/main.go) | `NEXTHINK_LANGUAGE`; [guide](ai_tools/GetOverviewInsights/README.md) | — |
| [ai_tools.GetRedirectURLs](ai_tools/GetRedirectURLs/main.go) | None; [guide](ai_tools/GetRedirectURLs/README.md) | — |
| [ai_tools.GetToolInsights](ai_tools/GetToolInsights/main.go) | `NEXTHINK_LANGUAGE`, `NEXTHINK_REQUEST_FILE`; [guide](ai_tools/GetToolInsights/README.md) | [JSON](ai_tools/GetToolInsights/request.example.json) |
| [ai_tools.List](ai_tools/List/main.go) | None; [guide](ai_tools/List/README.md) | — |
| [ai_tools.ListGoals](ai_tools/ListGoals/main.go) | None; [guide](ai_tools/ListGoals/README.md) | — |
| [ai_tools.ListSystemTools](ai_tools/ListSystemTools/main.go) | None; [guide](ai_tools/ListSystemTools/README.md) | — |
| [ai_tools.Update](ai_tools/Update/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE`, `NEXTHINK_REVISION`; [guide](ai_tools/Update/README.md) | [JSON](ai_tools/Update/request.example.json) |
| [ai_tools.UpdateCopilot](ai_tools/UpdateCopilot/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE`, `NEXTHINK_REVISION`; [guide](ai_tools/UpdateCopilot/README.md) | [JSON](ai_tools/UpdateCopilot/request.example.json) |
| [ai_tools.UpdateGoal](ai_tools/UpdateGoal/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE`, `NEXTHINK_REVISION`; [guide](ai_tools/UpdateGoal/README.md) | [JSON](ai_tools/UpdateGoal/request.example.json) |
| [ai_tools.UpdateModule](ai_tools/UpdateModule/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE`, `NEXTHINK_REVISION`; [guide](ai_tools/UpdateModule/README.md) | [JSON](ai_tools/UpdateModule/request.example.json) |
| [amplify.GetConfiguration](amplify/GetConfiguration/main.go) | None; [guide](amplify/GetConfiguration/README.md) | — |
| [amplify.GetDevicePackages](amplify/GetDevicePackages/main.go) | `NEXTHINK_CONTENT_ID`; [guide](amplify/GetDevicePackages/README.md) | — |
| [amplify.GetDeviceProperties](amplify/GetDeviceProperties/main.go) | `NEXTHINK_CONTENT_ID`; [guide](amplify/GetDeviceProperties/README.md) | — |
| [amplify.GetDeviceUsers](amplify/GetDeviceUsers/main.go) | `NEXTHINK_CONTENT_ID`; [guide](amplify/GetDeviceUsers/README.md) | — |
| [amplify.GetUserDevices](amplify/GetUserDevices/main.go) | `NEXTHINK_CONTENT_ID`; [guide](amplify/GetUserDevices/README.md) | — |
| [amplify.GetUserProperties](amplify/GetUserProperties/main.go) | `NEXTHINK_CONTENT_ID`; [guide](amplify/GetUserProperties/README.md) | — |
| [amplify.PostInsights](amplify/PostInsights/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify/PostInsights/README.md) | [JSON](amplify/PostInsights/request.example.json) |
| [amplify.Search](amplify/Search/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify/Search/README.md) | [JSON](amplify/Search/request.example.json) |
| [amplify.SearchDevices](amplify/SearchDevices/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify/SearchDevices/README.md) | [JSON](amplify/SearchDevices/request.example.json) |
| [amplify_ai.ExecuteAction](amplify_ai/ExecuteAction/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify_ai/ExecuteAction/README.md) | [JSON](amplify_ai/ExecuteAction/request.example.json) |
| [amplify_ai.ExecuteUserAction](amplify_ai/ExecuteUserAction/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify_ai/ExecuteUserAction/README.md) | [JSON](amplify_ai/ExecuteUserAction/request.example.json) |
| [amplify_ai.GenerateOrFetchAnalysis](amplify_ai/GenerateOrFetchAnalysis/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify_ai/GenerateOrFetchAnalysis/README.md) | [JSON](amplify_ai/GenerateOrFetchAnalysis/request.example.json) |
| [amplify_ai.GetMockMetadata](amplify_ai/GetMockMetadata/main.go) | `NEXTHINK_CONTENT_ID`; [guide](amplify_ai/GetMockMetadata/README.md) | — |
| [amplify_ai.GetResolutionPlan](amplify_ai/GetResolutionPlan/main.go) | `NEXTHINK_CONTENT_ID`; [guide](amplify_ai/GetResolutionPlan/README.md) | — |
| [amplify_ai.PostMetric](amplify_ai/PostMetric/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify_ai/PostMetric/README.md) | [JSON](amplify_ai/PostMetric/request.example.json) |
| [amplify_ai.RefreshResolutionStep](amplify_ai/RefreshResolutionStep/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify_ai/RefreshResolutionStep/README.md) | [JSON](amplify_ai/RefreshResolutionStep/request.example.json) |
| [amplify_ai.ReportTicketRetrievalDuration](amplify_ai/ReportTicketRetrievalDuration/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify_ai/ReportTicketRetrievalDuration/README.md) | [JSON](amplify_ai/ReportTicketRetrievalDuration/request.example.json) |
| [amplify_ai.ResolveTicket](amplify_ai/ResolveTicket/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify_ai/ResolveTicket/README.md) | [JSON](amplify_ai/ResolveTicket/request.example.json) |
| [amplify_ai.SubmitFeedback](amplify_ai/SubmitFeedback/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify_ai/SubmitFeedback/README.md) | [JSON](amplify_ai/SubmitFeedback/request.example.json) |
| [amplify_ai.UpdateResolutionStep](amplify_ai/UpdateResolutionStep/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify_ai/UpdateResolutionStep/README.md) | [JSON](amplify_ai/UpdateResolutionStep/request.example.json) |
| [workspace.CancelConversation](workspace/CancelConversation/main.go) | `NEXTHINK_RESOURCE_ID`; [guide](workspace/CancelConversation/README.md) | — |
| [workspace.Chat](workspace/Chat/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](workspace/Chat/README.md) | [JSON](workspace/Chat/request.example.json) |
| [workspace.CreateConversationShare](workspace/CreateConversationShare/main.go) | `NEXTHINK_RESOURCE_ID`; [guide](workspace/CreateConversationShare/README.md) | — |
| [workspace.DeleteConversation](workspace/DeleteConversation/main.go) | `NEXTHINK_RESOURCE_ID`; [guide](workspace/DeleteConversation/README.md) | — |
| [workspace.DeleteConversationFile](workspace/DeleteConversationFile/main.go) | `NEXTHINK_FILE_ID`, `NEXTHINK_RESOURCE_ID`; [guide](workspace/DeleteConversationFile/README.md) | — |
| [workspace.GetConversation](workspace/GetConversation/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace/GetConversation/README.md) | [JSON](workspace/GetConversation/request.example.json) |
| [workspace.GetSharedConversation](workspace/GetSharedConversation/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace/GetSharedConversation/README.md) | [JSON](workspace/GetSharedConversation/request.example.json) |
| [workspace.ListConversations](workspace/ListConversations/main.go) | `NEXTHINK_AUTOMATION_ID`; [guide](workspace/ListConversations/README.md) | — |
| [workspace.MCPProxy](workspace/MCPProxy/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](workspace/MCPProxy/README.md) | [JSON](workspace/MCPProxy/request.example.json) |
| [workspace.MarkConversationRead](workspace/MarkConversationRead/main.go) | `NEXTHINK_RESOURCE_ID`; [guide](workspace/MarkConversationRead/README.md) | — |
| [workspace.UpdateConversation](workspace/UpdateConversation/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace/UpdateConversation/README.md) | [JSON](workspace/UpdateConversation/request.example.json) |
| [workspace.UploadConversationFile](workspace/UploadConversationFile/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace/UploadConversationFile/README.md) | [JSON](workspace/UploadConversationFile/request.example.json) |
| [workspace_agents.CheckSkillAvailability](workspace_agents/CheckSkillAvailability/main.go) | `NEXTHINK_SOURCE`; [guide](workspace_agents/CheckSkillAvailability/README.md) | — |
| [workspace_agents.CompleteSkillMultipartUpload](workspace_agents/CompleteSkillMultipartUpload/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace_agents/CompleteSkillMultipartUpload/README.md) | [JSON](workspace_agents/CompleteSkillMultipartUpload/request.example.json) |
| [workspace_agents.CreateSkill](workspace_agents/CreateSkill/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](workspace_agents/CreateSkill/README.md) | [JSON](workspace_agents/CreateSkill/request.example.json) |
| [workspace_agents.DeleteSkill](workspace_agents/DeleteSkill/main.go) | `NEXTHINK_RESOURCE_ID`; [guide](workspace_agents/DeleteSkill/README.md) | — |
| [workspace_agents.DeleteSkillFile](workspace_agents/DeleteSkillFile/main.go) | `NEXTHINK_FILE_ID`, `NEXTHINK_RESOURCE_ID`; [guide](workspace_agents/DeleteSkillFile/README.md) | — |
| [workspace_agents.GetSkill](workspace_agents/GetSkill/main.go) | `NEXTHINK_RESOURCE_ID`; [guide](workspace_agents/GetSkill/README.md) | — |
| [workspace_agents.ListSkills](workspace_agents/ListSkills/main.go) | `NEXTHINK_SOURCE`; [guide](workspace_agents/ListSkills/README.md) | — |
| [workspace_agents.StartSkillMultipartUpload](workspace_agents/StartSkillMultipartUpload/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace_agents/StartSkillMultipartUpload/README.md) | [JSON](workspace_agents/StartSkillMultipartUpload/request.example.json) |
| [workspace_agents.UpdateSkill](workspace_agents/UpdateSkill/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace_agents/UpdateSkill/README.md) | [JSON](workspace_agents/UpdateSkill/request.example.json) |
| [workspace_agents.UploadSkillFile](workspace_agents/UploadSkillFile/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace_agents/UploadSkillFile/README.md) | [JSON](workspace_agents/UploadSkillFile/request.example.json) |
| [workspace_agents.UploadSkillPart](workspace_agents/UploadSkillPart/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace_agents/UploadSkillPart/README.md) | [JSON](workspace_agents/UploadSkillPart/request.example.json) |
| [workspace_assignments.GetAssignment](workspace_assignments/GetAssignment/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace_assignments/GetAssignment/README.md) | [JSON](workspace_assignments/GetAssignment/request.example.json) |
| [workspace_assignments.GetUnreadAssignmentCount](workspace_assignments/GetUnreadAssignmentCount/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](workspace_assignments/GetUnreadAssignmentCount/README.md) | [JSON](workspace_assignments/GetUnreadAssignmentCount/request.example.json) |
| [workspace_assignments.ListAssignees](workspace_assignments/ListAssignees/main.go) | `NEXTHINK_RESOURCE_ID`; [guide](workspace_assignments/ListAssignees/README.md) | — |
| [workspace_assignments.ListAssignments](workspace_assignments/ListAssignments/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](workspace_assignments/ListAssignments/README.md) | [JSON](workspace_assignments/ListAssignments/request.example.json) |
| [workspace_assignments.MarkAssignmentRead](workspace_assignments/MarkAssignmentRead/main.go) | `NEXTHINK_RESOURCE_ID`; [guide](workspace_assignments/MarkAssignmentRead/README.md) | — |
| [workspace_assignments.UpdateAssignment](workspace_assignments/UpdateAssignment/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace_assignments/UpdateAssignment/README.md) | [JSON](workspace_assignments/UpdateAssignment/request.example.json) |
| [workspace_tasks.CreateTask](workspace_tasks/CreateTask/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](workspace_tasks/CreateTask/README.md) | [JSON](workspace_tasks/CreateTask/request.example.json) |
| [workspace_tasks.DeleteTask](workspace_tasks/DeleteTask/main.go) | `NEXTHINK_RESOURCE_ID`; [guide](workspace_tasks/DeleteTask/README.md) | — |
| [workspace_tasks.GetTask](workspace_tasks/GetTask/main.go) | `NEXTHINK_RESOURCE_ID`; [guide](workspace_tasks/GetTask/README.md) | — |
| [workspace_tasks.ListTasks](workspace_tasks/ListTasks/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](workspace_tasks/ListTasks/README.md) | [JSON](workspace_tasks/ListTasks/request.example.json) |
| [workspace_tasks.ReconcileTaskAgentAccess](workspace_tasks/ReconcileTaskAgentAccess/main.go) | None; [guide](workspace_tasks/ReconcileTaskAgentAccess/README.md) | — |
| [workspace_tasks.UpdateTask](workspace_tasks/UpdateTask/main.go) | `NEXTHINK_REQUEST_FILE`, `NEXTHINK_RESOURCE_ID`; [guide](workspace_tasks/UpdateTask/README.md) | [JSON](workspace_tasks/UpdateTask/request.example.json) |

| [amplify.CreateConfiguration](amplify/CreateConfiguration/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](amplify/CreateConfiguration/README.md) | [JSON](amplify/CreateConfiguration/request.example.json) |
| [amplify.UpdateConfiguration](amplify/UpdateConfiguration/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE`, `NEXTHINK_REVISION_NUMBER`; [guide](amplify/UpdateConfiguration/README.md) | [JSON](amplify/UpdateConfiguration/request.example.json) |
| [device_classification.CreateLocationType](device_classification/CreateLocationType/main.go) | `NEXTHINK_CSV_FILE`, `NEXTHINK_REQUEST_FILE`; [guide](device_classification/CreateLocationType/README.md) | [JSON](device_classification/CreateLocationType/request.example.json) |
| [device_classification.CreateOrganization](device_classification/CreateOrganization/main.go) | `NEXTHINK_CSV_FILE`, `NEXTHINK_REQUEST_FILE`; [guide](device_classification/CreateOrganization/README.md) | [JSON](device_classification/CreateOrganization/request.example.json) |
| [device_classification.CreateVPNEgress](device_classification/CreateVPNEgress/main.go) | `NEXTHINK_CSV_FILE`, `NEXTHINK_REQUEST_FILE`; [guide](device_classification/CreateVPNEgress/README.md) | [JSON](device_classification/CreateVPNEgress/request.example.json) |
| [device_classification.DeleteVPNEgress](device_classification/DeleteVPNEgress/main.go) | None; [guide](device_classification/DeleteVPNEgress/README.md) | — |
| [device_classification.DownloadLocationType](device_classification/DownloadLocationType/main.go) | `NEXTHINK_OUTPUT_FILE`; [guide](device_classification/DownloadLocationType/README.md) | — |
| [device_classification.DownloadOrganization](device_classification/DownloadOrganization/main.go) | `NEXTHINK_OUTPUT_FILE`; [guide](device_classification/DownloadOrganization/README.md) | — |
| [device_classification.DownloadVPNEgress](device_classification/DownloadVPNEgress/main.go) | `NEXTHINK_OUTPUT_FILE`; [guide](device_classification/DownloadVPNEgress/README.md) | — |
| [device_classification.GetGeoIP](device_classification/GetGeoIP/main.go) | None; [guide](device_classification/GetGeoIP/README.md) | — |
| [device_classification.GetLocationType](device_classification/GetLocationType/main.go) | None; [guide](device_classification/GetLocationType/README.md) | — |
| [device_classification.GetOrganization](device_classification/GetOrganization/main.go) | None; [guide](device_classification/GetOrganization/README.md) | — |
| [device_classification.GetVPNEgress](device_classification/GetVPNEgress/main.go) | None; [guide](device_classification/GetVPNEgress/README.md) | — |
| [device_classification.UpdateGeoIP](device_classification/UpdateGeoIP/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](device_classification/UpdateGeoIP/README.md) | [JSON](device_classification/UpdateGeoIP/request.example.json) |
| [device_classification.UpdateLocationType](device_classification/UpdateLocationType/main.go) | `NEXTHINK_CSV_FILE`, `NEXTHINK_REQUEST_FILE`; [guide](device_classification/UpdateLocationType/README.md) | [JSON](device_classification/UpdateLocationType/request.example.json) |
| [device_classification.UpdateOrganization](device_classification/UpdateOrganization/main.go) | `NEXTHINK_CSV_FILE`, `NEXTHINK_REQUEST_FILE`; [guide](device_classification/UpdateOrganization/README.md) | [JSON](device_classification/UpdateOrganization/request.example.json) |
| [device_classification.UpdateVPNEgress](device_classification/UpdateVPNEgress/main.go) | `NEXTHINK_CSV_FILE`, `NEXTHINK_REQUEST_FILE`; [guide](device_classification/UpdateVPNEgress/README.md) | [JSON](device_classification/UpdateVPNEgress/request.example.json) |
| [product_configuration.CreateInstance](product_configuration/CreateInstance/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](product_configuration/CreateInstance/README.md) | [JSON](product_configuration/CreateInstance/request.example.json) |
| [product_configuration.GetInstance](product_configuration/GetInstance/main.go) | `NEXTHINK_CONFIGURATION_KEY`; [guide](product_configuration/GetInstance/README.md) | — |
| [product_configuration.UpdateInstance](product_configuration/UpdateInstance/main.go) | `NEXTHINK_CONFIGURATION_KEY`, `NEXTHINK_REQUEST_FILE`; [guide](product_configuration/UpdateInstance/README.md) | [JSON](product_configuration/UpdateInstance/request.example.json) |
| [user_classification.List](user_classification/List/main.go) | None; [guide](user_classification/List/README.md) | — |
| [user_classification.Replace](user_classification/Replace/main.go) | `NEXTHINK_REQUEST_FILE`; [guide](user_classification/Replace/README.md) | [JSON](user_classification/Replace/request.example.json) |

## Existing resource examples

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

## Identity, support, execution and browser helper additions

Each guide lists the required inputs, operation examples and validation limits. `AccessManagement` and `LegacyAccessManagement` are available through the root client. Search and assistant chat return buffered events; they do not offer realtime callbacks or automatic SSE reconnection.

| Resource | Guide |
| --- | --- |
| `access_management` | [Examples and inputs](access_management/README.md) |
| `collaboration_comments` | [Examples and inputs](collaboration_comments/README.md) |
| `support` | [Examples and inputs](support/README.md) |
| `support_checklists` | [Examples and inputs](support_checklists/README.md) |
| `support_timeline` | [Examples and inputs](support_timeline/README.md) |
| `support_insights` | [Examples and inputs](support_insights/README.md) |
| `collaboration_tools` | [Examples and inputs](collaboration_tools/README.md) |
| `vdi` | [Examples and inputs](vdi/README.md) |
| `workflow_executions` | [Examples and inputs](workflow_executions/README.md) |
| `action_executions` | [Examples and inputs](action_executions/README.md) |
| `autopilot` | [Examples and inputs](autopilot/README.md) |
| `global_search` | [Examples and inputs](global_search/README.md) |
| `nlp_assistant` | [Examples and inputs](nlp_assistant/README.md) |
| `recommendations` | [Examples and inputs](recommendations/README.md) |
| `data_export` | [Examples and inputs](data_export/README.md) |
| `visual_editor` | [Examples and inputs](visual_editor/README.md) |
| `content_sharing` | [Examples and inputs](content_sharing/README.md) |
| `custom_field_values` | [Examples and inputs](custom_field_values/README.md) |
| `applications` | [Examples and inputs](applications/README.md) |
| `custom_fields` | [Examples and inputs](custom_fields/README.md) |
| `rule_based_custom_fields` | [Examples and inputs](rule_based_custom_fields/README.md) |

## Discovery lead additions

These guides cover the remaining source-confirmed lead contracts and their live-validation limits. Legacy portal operations require explicit session credentials.

| Resource | Guide |
| --- | --- |
| `teams_credentials` | [Examples and inputs](teams_credentials/README.md) |
| `zoom_notifications` | [Examples and inputs](zoom_notifications/README.md) |
| `azure_ad_credentials` | [Examples and inputs](azure_ad_credentials/README.md) |
| `user_communication_integrations` | [Examples and inputs](user_communication_integrations/README.md) |
| `legacy_connectors` | [Examples and inputs](legacy_connectors/README.md) |
| `query_builder` | [Examples and inputs](query_builder/README.md) |
| `cci_benchmarks` | [Examples and inputs](cci_benchmarks/README.md) |
| `appearance` | [Examples and inputs](appearance/README.md) |
| `mobile_tokens` | [Examples and inputs](mobile_tokens/README.md) |
| `snapshots` | [Examples and inputs](snapshots/README.md) |
| `ui_events` | [Examples and inputs](ui_events/README.md) |
| `observability` | [Examples and inputs](observability/README.md) |
| `collector_management` | [Examples and inputs](collector_management/README.md) |
| `global_search` | [Examples and inputs](global_search/README.md) |
