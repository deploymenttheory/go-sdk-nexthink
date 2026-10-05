# Assets examples

Use the shared [authentication setup](../README.md). From the repository root, run `go run ./examples/nexthink/web_api/assets/List`. Write examples require `NEXTHINK_REQUEST_FILE` containing an explicit request; replace the synthetic values with your intended lab target.

| Method | Inputs beyond authentication | Sample |
| --- | --- | --- |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Create/request.example.json) |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Delete/request.example.json) |
| [GetSignedURL](GetSignedURL/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [List](List/main.go) | None | — |
| [Update](Update/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE` | [JSON](Update/request.example.json) |

Asset metadata comes from List; GetSignedURL is the UI's read/download operation. It returns a temporary signed URL that must be kept private. The SDK does not download it or forward the browser bearer token to it.

Create and Update accept `UploadRequest`: a filename, MIME type, and raw file bytes. JSON input represents `data` as base64; Go's JSON decoder turns that into bytes and the SDK then constructs one `data:<mediaType>;base64,...` body. The HTTP body is text/plain with `x-nxt-asset-name`, not JSON or multipart. Update and Delete return HTTP metadata without a JSON result. Update changes the file bytes and revision but does not rename the asset, even when the filename header differs.

Curl and all five examples passed using a synthetic SVG. Replacement bytes were downloaded and compared exactly. The temporary asset was deleted.
