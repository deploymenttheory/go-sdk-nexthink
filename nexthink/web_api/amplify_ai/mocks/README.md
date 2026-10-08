# Fixture provenance

Request fields are traced from the publicly distributed Nexthink Amplify extension version 1.34.0. The extension calls `nxai/amplify-ai/v1` using browser credentials. These fixtures use synthetic resolution, device and ticket identifiers.

Success response fixtures are controlled transport examples based on fields consumed by the extension, **not captured successful lab responses**. Responses remain `json.RawMessage` so the SDK does not assert an unverified server schema. Empty `steps` or `sections` can represent pending work; HTTP 200 alone does not prove a completed analysis.

On 2026-10-06 both read routes returned HTTP 404 on the exact extension API hostname and portal gateway. The user's effective AI-insights claim was false. No live analysis generation, action execution, ticket mutation or AI telemetry was attempted. Positive live acceptance remains outstanding for these eleven operations.
