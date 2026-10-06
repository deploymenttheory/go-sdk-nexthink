# ApplicationExperience.GetApplicationInsights

Set `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome`, `NEXTHINK_INSTANCE`, and `NEXTHINK_REGION` with an authenticated Chrome portal session.

Set `NEXTHINK_REQUEST_FILE` to a JSON file with a real configured application ID:

```json
{
  "applicationId": "00000000-0000-0000-0000-000000000000",
  "currentTimeframe": "from 2026-10-01 to 2026-10-05",
  "previousTimeframe": "from 2026-09-26 to 2026-09-30",
  "timezone": "UTC",
  "breakdowns": {
    "adoption": [
      "user.ad.department"
    ]
  },
  "type": "adoption"
}
```

The lab has no configured application fixture; this operation has source and unit coverage but no successful live response yet.

```sh
go run ./examples/nexthink/web_api/application_experience/GetApplicationInsights
```
