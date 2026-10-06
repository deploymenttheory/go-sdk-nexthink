# Autopilot.ReplaceWebSearchDomains

See the [resource guide](../README.md) for authentication and validation limits.

This operation changes tenant configuration. Use only an explicitly intended test target. The sample values are synthetic; adapt them before execution.

Read the current settings first and preserve their revision and desired fields. This request replaces the supplied configuration; an empty domain list clears domains.

```sh
export NEXTHINK_REQUEST_FILE="$PWD/examples/nexthink/web_api/autopilot/ReplaceWebSearchDomains/request.example.json"
go run ./examples/nexthink/web_api/autopilot/ReplaceWebSearchDomains
```
