# Amplify UpdateConfiguration

Configure the root SDK web client with `NEXTHINK_API=web`, instance, region and web authentication. Manage Amplify permission is required.

Read `GetConfiguration` first and save its current contents. Copy `request.example.json`, edit it, and set `NEXTHINK_REQUEST_FILE`. The example uses a nonmatching `.invalid` test URL and leaves usage reporting disabled.

Set `NEXTHINK_CONTENT_ID` and `NEXTHINK_REVISION_NUMBER` from the latest configuration document. The server uses the revision for concurrency control; read again after every successful update. This request replaces the entire application list and reporting flag, so preserve unrelated applications.

Add or edit an application by modifying its entry. Delete by omitting its entry. Reorder by changing array order. To remove all applications, send `"itsmConfigList": []`. An omitted or null list is rejected by the SDK. `enableUsageDataReporting` is always sent, including false.

```sh
go run ./examples/nexthink/web_api/amplify/UpdateConfiguration
```
