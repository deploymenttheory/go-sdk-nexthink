#!/usr/bin/env bash
# Read-only live checks. Run from the repository root with lab credentials in env.
set -euo pipefail
: "${NEXTHINK_CLIENT_ID:?required}"
: "${NEXTHINK_CLIENT_SECRET:?required}"
: "${NEXTHINK_INSTANCE:?required}"
: "${NEXTHINK_REGION:?required}"
: "${NEXTHINK_QUERY_ID:?a saved bounded query is required}"
: "${NEXTHINK_REMOTE_ACTION_ID:?required for detail example}"
: "${NEXTHINK_WORKFLOW_ID:?required for detail example}"
: "${NEXTHINK_EXPORT_ID:?required for export status example}"
: "${NEXTHINK_DOWNLOAD_URL:?a completed export download URL is required}"
: "${NEXTHINK_REQUEST_FILE:?JSON containing the saved bounded queryId is required}"
# Keep the request usable after each example enters its private working directory.
NEXTHINK_REQUEST_FILE="$(cd "$(dirname "$NEXTHINK_REQUEST_FILE")" && pwd)/$(basename "$NEXTHINK_REQUEST_FILE")"
export NEXTHINK_REQUEST_FILE
[[ -r "$NEXTHINK_REQUEST_FILE" ]] || { printf 'Request file is unreadable\n' >&2; exit 1; }
export NEXTHINK_API=public
export NEXTHINK_EXPORT_QUERY_ID="${NEXTHINK_EXPORT_QUERY_ID:-$NEXTHINK_QUERY_ID}"
umask 077
lab_example_output=$(mktemp -d "${TMPDIR:-/tmp}/nexthink-examples.XXXXXX")
lab_example_root=$(pwd)
lab_example_failures=0
for lab_example_source in examples/nexthink/_build_client/*/main.go examples/nexthink/public_api/nql/*/main.go examples/nexthink/public_api/remote_actions/{ListRemoteActions,GetRemoteActionDetails}/main.go examples/nexthink/public_api/workflows/{ListWorkflows,ListWorkflowsWithOptions,GetWorkflowDetails}/main.go; do
    # This harness configures public OAuth only. The password example has its
    # own browser runtime and local-account prerequisites.
    if [[ "$lab_example_source" == */headless_password/main.go ]]; then
        printf 'SKIP %s (run separately with local-account credentials)\n' "$lab_example_source"
        continue
    fi
    lab_example_package=${lab_example_source%/main.go}
    lab_example_name=${lab_example_package//\//_}
    lab_example_work="$lab_example_output/$lab_example_name.data"
    mkdir "$lab_example_work"
    if go build -o "$lab_example_output/$lab_example_name" "./$lab_example_package" &&
       (cd "$lab_example_work" && NEXTHINK_OUTPUT_FILE="$lab_example_work/output" "$lab_example_output/$lab_example_name" > example.log 2>&1); then
        printf 'PASS %s\n' "$lab_example_package"
    else
        printf 'FAIL %s\n' "$lab_example_package"
        lab_example_failures=$((lab_example_failures + 1))
    fi
    cd "$lab_example_root"
done
printf 'Private logs and exports: %s\n' "$lab_example_output"
exit "$lab_example_failures"
