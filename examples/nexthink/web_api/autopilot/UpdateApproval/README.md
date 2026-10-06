# Autopilot.UpdateApproval

See the [resource guide](../README.md) for authentication and validation limits.

This operation changes tenant configuration. Use only an explicitly intended test target. The sample values are synthetic; adapt them before execution.

`NEXTHINK_RESOURCE_ID` is the action identifier used by its approval record. Preserve the current revision when updating.

```sh
export NEXTHINK_REQUEST_FILE="$PWD/examples/nexthink/web_api/autopilot/UpdateApproval/request.example.json"
export NEXTHINK_RESOURCE_ID="replace-with-lab-id"
go run ./examples/nexthink/web_api/autopilot/UpdateApproval
```
