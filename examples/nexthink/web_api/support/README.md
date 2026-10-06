# Support browser API examples

Use `NEXTHINK_API=web` and browser authentication through the SDK root client. Each method folder has the exact inputs and a runnable example.

For device calls, set `NEXTHINK_DEVICE_ID` to the Collector UID (`device.collector.uid`), not the separate device UID (`device.uid`). The UI routes label this parameter `deviceID`, but supplying the device UID can return empty results instead of an error. Obtain both identifiers with `devices | list device.uid, device.collector.uid, device.name` and select the intended device.

These private UI contracts may vary by tenant version. Analytics endpoints expose reads; lifecycle operations belong to their configuration resources. UTC headers are the default; services accept `WithTimeZone` for the browser time context.

- [GetProfile](GetProfile/README.md)
- [GetPlatform](GetPlatform/README.md)
- [ListUsers](ListUsers/README.md)
- [Search](Search/README.md)
