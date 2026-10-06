# Azure AD credentials examples

Use the SDK single entry point with `NEXTHINK_API=web` and browser authentication as described in [the web API guide](../README.md).

[CheckCredentials](CheckCredentials/main.go) sends `POST /apigateway/azuread-fetcher/check-credentials` as form data. Set `NEXTHINK_REQUEST_FILE` to a private JSON file containing `{"tenant_id":"...","client_id":"...","ms_national_cloud":"GLOBAL","client_secret":"..."}`. The UI uses GLOBAL or US_L4 (GCC) for the cloud value. Omit client_secret to check the saved credentials, matching the UI's unchanged-secret flow. Request JSON is converted to form data without trimming credential characters.

```sh
go run ./examples/nexthink/web_api/azure_ad_credentials/CheckCredentials
```

The method returns HTTP response metadata. Unsuccessful HTTP status codes retain that metadata and return an error. This validates third-party credentials; it does not authenticate an SDK client. No Azure credentials were supplied for live validation, so this contract is source-backed and unit-tested. Configuration and secret storage use LegacyConnectors.
