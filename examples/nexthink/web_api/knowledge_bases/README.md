# KnowledgeBases examples

## Additional browser operations

Use `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome` (or a browser token), and the lab instance/region. Supply `NEXTHINK_REQUEST_FILE` for examples with a request file, `NEXTHINK_CONTENT_ID` for ID arguments, and `NEXTHINK_EXECUTION_ID` for execution polling. Requests are synthetic templates; replace identifiers with your intended targets.

- [List](List/main.go)
- [GetContents](GetContents/main.go)
- [Create](Create/main.go) — [request](Create/request.example.json)
- [Delete](Delete/main.go)
- [GetDownloadURL](GetDownloadURL/main.go)
- [UploadFile](UploadFile/main.go) — [request](UploadFile/request.example.json)
- [StartMultipartUpload](StartMultipartUpload/main.go) — [request](StartMultipartUpload/request.example.json)
- [UploadPart](UploadPart/main.go) — [request](UploadPart/request.example.json)
- [CompleteMultipartUpload](CompleteMultipartUpload/main.go) — [request](CompleteMultipartUpload/request.example.json)

Create/Update/Import/Delete and upload methods write data. Test/StartTest contacts the configured destination or starts a server test; review the target first. Re-fetch revisions between dashboard mutations.

`UploadFile.Data` is the original CSV bytes (`[]byte`, represented as base64 in the example input JSON). The SDK encodes those bytes once for the wire. CSV columns are `number,kb_knowledge_base,text,short_description,kb_category`; number, knowledge base and text must have values. Registration is asynchronous.

For multipart upload, base64-encode the entire file once, then split the encoded string into chunks (the UI uses 6,990,508 characters). Pass each slice as `EncodedChunk`, retain the returned part number/etag/checksum, and complete in increasing part order. Do not independently base64-encode arbitrary raw-byte chunks. Download URLs are base64-encoded signed URLs; use `DecodeURL()` and download separately without forwarding the browser token.

Deletion can race asynchronous ingestion. The lab returned HTTP 500 with `KNOWLEDGE_BASE_PERSISTENCE_ERROR` when deleting immediately after registration; later explicit deletion of the same owned fixture succeeded with HTTP 204. Keep the created `contentId` until cleanup is confirmed. The SDK returns that server error and response metadata; a failed cleanup must not be counted as successful deletion.
