# Autopilot.CreateApproval

See the [resource guide](../README.md) for authentication and validation limits.

This operation changes tenant configuration. Use only an explicitly intended test target. The sample values are synthetic; adapt them before execution.

`_id` identifies the target action; it is required in the create payload and is not a server-generated approval UUID.

```sh
export NEXTHINK_REQUEST_FILE="$PWD/examples/nexthink/web_api/autopilot/CreateApproval/request.example.json"
go run ./examples/nexthink/web_api/autopilot/CreateApproval
```
