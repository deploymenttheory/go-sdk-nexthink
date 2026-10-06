# Autopilot examples

Use browser-token authentication: set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE_URL`, and `NEXTHINK_BROWSER_TOKEN`. Run examples from the repository root. Each operation below has its own instructions and request JSON where needed.

These are UI contracts from Autopilot/Spark bundles. The lab session returns HTTP 401 for cockpit configuration/settings; source and unit verification do not establish feature entitlement. Settings updates, approval changes, domain replacement and agent input updates change tenant configuration. CreateTicket can create a real external ticket. Review each request before running any mutation. No mutation in this package was executed during discovery. Downloads return the exact CSV as base64-encoded `data` in the JSON output.

| Operation | Inputs |
|---|---|
| [GetConfiguration](./GetConfiguration/README.md) | none |
| [UpdateWebSearch](./UpdateWebSearch/README.md) | request JSON |
| [UpdateFilesystemTool](./UpdateFilesystemTool/README.md) | request JSON |
| [UpdateAgentName](./UpdateAgentName/README.md) | request JSON |
| [GetSettings](./GetSettings/README.md) | none |
| [SaveSettings](./SaveSettings/README.md) | request JSON |
| [GetWebSearchDomains](./GetWebSearchDomains/README.md) | none |
| [ReplaceWebSearchDomains](./ReplaceWebSearchDomains/README.md) | request JSON |
| [GetApproval](./GetApproval/README.md) | resource ID |
| [CreateApproval](./CreateApproval/README.md) | request JSON |
| [UpdateApproval](./UpdateApproval/README.md) | request JSON, resource ID |
| [ListCalls](./ListCalls/README.md) | none |
| [GetKnowledgeArticleCount](./GetKnowledgeArticleCount/README.md) | none |
| [ListKnowledgeConnectors](./ListKnowledgeConnectors/README.md) | none |
| [CreateTicket](./CreateTicket/README.md) | request JSON |
| [GetConversation](./GetConversation/README.md) | resource ID |
| [GetRecommendationConversationIDs](./GetRecommendationConversationIDs/README.md) | resource ID |
| [UpdateAgentActionInputs](./UpdateAgentActionInputs/README.md) | request JSON, resource ID |
| [DownloadCategorization](./DownloadCategorization/README.md) | none |
| [GetAgentAction](./GetAgentAction/README.md) | resource ID |
| [GetAgentActionInputs](./GetAgentActionInputs/README.md) | resource ID |
