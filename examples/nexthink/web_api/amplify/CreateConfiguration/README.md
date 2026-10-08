# Amplify CreateConfiguration

Configure the root SDK web client with `NEXTHINK_API=web`, instance, region and web authentication. Manage Amplify permission is required.

Read `GetConfiguration` first and save its current contents. Copy `request.example.json`, edit it, and set `NEXTHINK_REQUEST_FILE`. The example uses a nonmatching `.invalid` test URL and leaves usage reporting disabled.

Use this only when the collection is empty. The UI exposes no delete operation for the configuration document; removing its application entries uses UpdateConfiguration with an empty array.

```sh
go run ./examples/nexthink/web_api/amplify/CreateConfiguration
```
