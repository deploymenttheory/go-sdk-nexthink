Browser session authentication is required. Set `NEXTHINK_API=web`, the instance URL, and browser-token authentication as described in the main examples README. Run each directory with `go run ./examples/nexthink/web_api/collaboration_comments/METHOD`.

Requests use `NEXTHINK_REQUEST_FILE`; each request-based example includes a synthetic `request.example.json`. Replace fixture values with your explicitly intended lab targets before running writes. Methods with path references require the matching `NEXTHINK_ID`, `NEXTHINK_DOCUMENT_ID`, `NEXTHINK_COMMENT_ID`, or `NEXTHINK_REPLY_ID` variable. Mention searches use `NEXTHINK_SEARCH`.

The RTC UI supplies the message UUID for creation and retries. `UpdateComment` wraps the message in `{operation:"editMessage",data:...}`; replies PATCH the message directly. Archive/unarchive use distinct operation values on the comment PATCH route. Acknowledgment methods preserve the server body in HTTP response metadata. The lab RTC routes returned403; no comment, reply, notification or mention was sent.

| Method | HTTP contract |
| --- | --- |
| `ResolveIdentifier` | `POST /identifier` |
| `ListComments` | `GET /documents/{documentID}/comments` |
| `CreateComment` | `POST /documents/{documentID}/comments` |
| `CreateReply` | `POST /documents/{documentID}/comments/{commentID}/replies` |
| `ListUserMentions` | `GET /user-mentions` |
| `ArchiveComment` | `PATCH /documents/{documentID}/comments/{commentID}` |
| `UnarchiveComment` | `PATCH /documents/{documentID}/comments/{commentID}` |
| `UpdateComment` | `PATCH /documents/{documentID}/comments/{commentID}` |
| `UpdateReply` | `PATCH /documents/{documentID}/comments/{commentID}/replies/{replyID}` |
| `DeleteComment` | `DELETE /documents/{documentID}/comments/{commentID}` |
| `DeleteReply` | `DELETE /documents/{documentID}/comments/{commentID}/replies/{replyID}` |
