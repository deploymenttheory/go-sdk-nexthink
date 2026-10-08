# AI Tools browser API

`client.WebAPI.AITools` exposes application AI tool LCRUD, Microsoft Copilot configuration, governance, shared module settings, insights and adoption-goal LCRUD. These endpoints require browser authentication and the corresponding AI Tools permissions. They were recovered from AI Dexterity UI 11.227.1, including its lazy-loaded administration, governance and goal chunks.

Application monitoring and governance share the same revisioned configuration document. Read the current tool, retain its existing monitoring and campaign configuration, and pass its `_rev` to `Update`. Custom tool NQL IDs begin with `#`; built-in tool IDs do not. `Delete` takes the current revision too. Goal dates use RFC3339 timestamps, as sent by the UI.

The configuration and response models retain web rules, binary names, Teams bot and network monitoring, credential references, license counts, experience campaigns, compliance policies, discovery metadata and goal metric/population/timeline fields. Module filters and optional goal progress remain JSON because their content is extensible. Credentials are references to an existing connector credential, not secrets.

## Scope and evidence

All 28 methods have runnable examples and JSON-backed success/error contract tests. Application-tool and goal LCRUD were validated with curl, then repeated using SDK examples with disposable fixtures and immediate cleanup. Successful list/get responses were compared as complete JSON documents, including discovered-tool metadata.

The lab has no module singleton: `GetModule` returns 404. Creating or updating that singleton would persist tenant-wide campaign configuration, so those writes have contract tests but were not run live. The existing built-in Copilot configuration was read; its writes and credential checks require a separate integration fixture. `GetGoalInsights` is feature-gated in the UI and returns HTTP 403 in this lab while goal configuration LCRUD succeeds.

Fixtures for unexecuted Copilot/module writes, credential validation and goal insights are modeled contract examples, not claimed live successes. Other success fixtures come from sanitized live payloads (lists are reduced while retaining complete item fields). Tests compare complete fixture JSON after decoding and re-encoding.

AI dashboard widgets reuse the existing `DataExploration`, `Dashboards`, `CCIBenchmarks`, `NQLEditor` and campaign services. Administration listing is available through `ContentAdministration` with content key `listing-aidex-tools`; no duplicate wrapper is added here. No module delete or separate governance CRUD route was evidenced in the UI.
