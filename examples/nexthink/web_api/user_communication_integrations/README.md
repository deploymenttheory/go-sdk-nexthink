# UserCommunicationIntegrations examples

Use the SDK single entry point with `NEXTHINK_API=web` and browser authentication as described in [the web API guide](../README.md). Run examples from the repository root.

| Example | HTTP operation | Inputs |
| --- | --- | --- |
| [List](List/main.go) | `GET /apigateway/user-communication-integrations/api/v1/integrations` | None |
| [Get](Get/main.go) | `GET /apigateway/user-communication-integrations/api/v1/integrations/{id}` | `NEXTHINK_CONTENT_ID` |
| [Create](Create/main.go) | `POST /apigateway/user-communication-integrations/api/v1/integrations` | `NEXTHINK_REQUEST_FILE` (IntegrationInput) |
| [Update](Update/main.go) | `PUT /apigateway/user-communication-integrations/api/v1/integrations/{id}` | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE` (IntegrationInput) |
| [Delete](Delete/main.go) | `DELETE /apigateway/user-communication-integrations/api/v1/integrations/{id}` | `NEXTHINK_CONTENT_ID` |
| [ListAzureConnectors](ListAzureConnectors/main.go) | `GET /apigateway/user-communication-integrations/api/v1/integrations/externals/azureconnectors` | None |

```sh
go run ./examples/nexthink/web_api/user_communication_integrations/List
```

Create and Update take `{"name":"Fixture channel","content":{"azureTenantId":"tenant-id","welcomeMessage":"Hello"}}`, or `{"name":"","content":{"azureConnectorId":"connector-id"}}`. The latter matches the UI connector selection mode. Delete removes the selected integration. These writes change the employee communication integration configuration; use only an explicitly chosen test integration. The lab had no existing integrations or Azure connector choices, so read validation covered empty lists; mutation contracts are source-backed and unit-tested.
