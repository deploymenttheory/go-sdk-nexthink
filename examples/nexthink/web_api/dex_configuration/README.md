
### Additional management operations

These examples use browser authentication. Request-based examples load `NEXTHINK_REQUEST_FILE`; copy the corresponding `request.example.json` and replace synthetic values. Mutation examples change configuration and should use explicitly selected targets.

- [GetAccount](GetAccount/main.go): `getAccount`.
- [UpdateApplications](UpdateApplications/main.go): `mutateApplications`.
- [GetScoreMetrics](GetScoreMetrics/main.go): `getEcScoreMetrics`.
- [UpdateScoreMetrics](UpdateScoreMetrics/main.go): `mutateEcScoreMetrics`.
- [GetVDIOptIn](GetVDIOptIn/main.go): `getOptInVdi`.
- [OptInVDI](OptInVDI/main.go): `optInVdi`.
- [GetMemoryMetricsOptIn](GetMemoryMetricsOptIn/main.go): `getOptInMemoryMetrics`.
- [OptInMemoryMetrics](OptInMemoryMetrics/main.go): `optInMemoryMetrics`.
- [GetCampaign](GetCampaign/main.go): `getConfigCampaign`.
- [EnableCampaign](EnableCampaign/main.go): `enableCampaign`.

Read examples were compared with curl in the lab. Configuration writes and opt-ins were tested against synthetic response fixtures and the observed UI request builders; they were not applied to tenant-wide settings. Score metric updates are sparse patches: omit a threshold to leave it alone, or pass JSON `null` to reset it to the default. `selected: false` and explicit threshold resets are preserved.

- [GetApplications](GetApplications/main.go): DEX application selection and URL/binary patterns on `/dex-ec/graphql`; distinct from the metering application query.
