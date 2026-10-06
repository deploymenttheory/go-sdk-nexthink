# Autopilot.CreateTicket

See the [resource guide](../README.md) for authentication and validation limits.

This operation changes tenant configuration and may create a real external ticket. Use only an explicitly intended test target. The sample values are synthetic; adapt them before execution.

A successful HTTP response may still contain `status: "FAILED"`; inspect the result. Set either a categorization row index (zero is valid) or category/subcategory as appropriate to the current settings.

```sh
export NEXTHINK_REQUEST_FILE="$PWD/examples/nexthink/web_api/autopilot/CreateTicket/request.example.json"
go run ./examples/nexthink/web_api/autopilot/CreateTicket
```
