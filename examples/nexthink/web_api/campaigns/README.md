# Campaigns examples

Configure `NEXTHINK_API=web`, instance, region, and browser authentication as described in the [example index](../README.md). Run examples from the repository root.

| Example | Required inputs |
| --- | --- |
| [List](List/main.go) | None (first page where paginated) |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` matching [Create_input.json](../../../../nexthink/web_api/campaigns/mocks/Create_input.json) |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` matching [Update_input.json](../../../../nexthink/web_api/campaigns/mocks/Update_input.json) |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` matching [Delete_input.json](../../../../nexthink/web_api/campaigns/mocks/Delete_input.json) |

```sh
go run ./examples/nexthink/web_api/campaigns/List
```

Create saves a draft. Update preserves the supplied status. Delete returns the deleted content ID. List supports pagination through ListOptions.

Replace synthetic IDs and revisions with the object you intend to manage. Create and Update write the supplied object; Delete removes it. GraphQL examples print partial data before reporting errors so returned identifiers remain available.

### Additional management operations

These examples use browser authentication. Request-based examples load `NEXTHINK_REQUEST_FILE`; copy the corresponding `request.example.json` and replace synthetic values. Mutation examples change configuration and should use explicitly selected targets.

- [GetBranding](GetBranding/main.go): `GetBranding`.
- [UpdateBranding](UpdateBranding/main.go): `updateBranding`.
- [SetStatus](SetStatus/main.go): `ChangeCampaignStatus`.
- [GetByNQLID](GetByNQLID/main.go): `FetchCampaignDocByNqlId`.
- [GetFromLibrary](GetFromLibrary/main.go): `LibraryContentByUuid`.
- [GetWithV6](GetWithV6/main.go): `campaignDocWithV6`.

`GetBranding` covers the equivalent `FetchBranding` and `GetBranding` UI queries. `UpdateBranding` uses the form's `doNotDisturbChoice` (for example `RATE_6_HOURS`), while reads return duration fields. Branding changes apply to the tenant.

`SetStatus` preserves server business errors: the lab rejects a transition to `DRAFT` and rejects retiring an unpublished campaign. The retirement example requires an already published target. The lab's new draft campaign had no legacy record for `GetWithV6`; that method is implemented from the observed UI contract and tested with synthetic legacy responses.
