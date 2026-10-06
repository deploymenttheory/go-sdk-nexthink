# Library examples

Configure `NEXTHINK_API=web`, instance, region and browser authentication as described in the [example index](../README.md). Run from the repository root. For examples with inputs, copy the linked JSON, replace fixture identifiers and set `NEXTHINK_REQUEST_FILE` to your local file.

| Example | Inputs |
| --- | --- |
| [ListContents](ListContents/main.go) | None |
| [ListPacks](ListPacks/main.go) | None |
| [GetContent](GetContent/main.go) | [Request JSON](GetContent/request.example.json) |
| [GetCustomContent](GetCustomContent/main.go) | [Request JSON](GetCustomContent/request.example.json) |
| [GetCreateCopyInfo](GetCreateCopyInfo/main.go) | [Request JSON](GetCreateCopyInfo/request.example.json) |
| [InstallContent](InstallContent/main.go) | [Request JSON](InstallContent/request.example.json) |
| [InstallCustomContent](InstallCustomContent/main.go) | [Request JSON](InstallCustomContent/request.example.json) |
| [InstallPack](InstallPack/main.go) | [Request JSON](InstallPack/request.example.json) |
| [InstallCustomPack](InstallCustomPack/main.go) | [Request JSON](InstallCustomPack/request.example.json) |
| [UpdateContent](UpdateContent/main.go) | [Request JSON](UpdateContent/request.example.json) |
| [UpdatePack](UpdatePack/main.go) | [Request JSON](UpdatePack/request.example.json) |
| [UpdateCustomContent](UpdateCustomContent/main.go) | [Request JSON](UpdateCustomContent/request.example.json) |
| [GetDependenciesStatus](GetDependenciesStatus/main.go) | [Request JSON](GetDependenciesStatus/request.example.json) |
| [InstallDependencies](InstallDependencies/main.go) | [Request JSON](InstallDependencies/request.example.json) |
| [GetPack](GetPack/main.go) | [Request JSON](GetPack/request.example.json) |
| [ImportCustomPack](ImportCustomPack/main.go) | [Request JSON](ImportCustomPack/request.example.json) |
| [GetPackInstallationStatus](GetPackInstallationStatus/main.go) | [Request JSON](GetPackInstallationStatus/request.example.json) |
| [GetContentInstallationStatus](GetContentInstallationStatus/main.go) | [Request JSON](GetContentInstallationStatus/request.example.json) |
| [DeleteCustomPack](DeleteCustomPack/main.go) | [Request JSON](DeleteCustomPack/request.example.json) |
| [GetLocale](GetLocale/main.go) | None |

```sh
go run ./examples/nexthink/web_api/library/ListContents
NEXTHINK_REQUEST_FILE=/path/to/request.json go run ./examples/nexthink/web_api/library/GetContent
```

Use catalog `fileName` values for standard content and packs. `GetPack` leaves `packUUID` empty for a standard pack; the optional UUID resolves custom packs. Custom content uses both `contentID` and `resourceName`. The service applies the required `nx-caller-service` header to custom installation and status requests.

Install, import and update examples change tenant content; `InstallPack` can install multiple resources and `DeleteCustomPack` removes a custom pack. Custom installations start asynchronous work: poll `GetPackInstallationStatus` or `GetContentInstallationStatus` until `isInstalling` is false, then inspect status and errors. Resource-dependent import/update results are retained as raw JSON.

Seven read examples were exercised. Six matched curl responses exactly; catalog ordering and some built-in versions varied between HTTP snapshots. All 810 captured catalog records independently round-tripped through the SDK response type without field loss. After correcting shared response buffering, the live SDK result also matched its complete 2,164,916-byte HTTP body exactly. Custom pack writes and installation polling success were validated with UI source contracts and synthetic unit fixtures because no disposable custom-pack fixture was available. No built-in content was installed during validation. The unused UI `contents/todo/config` helper returned 404 and is excluded.
