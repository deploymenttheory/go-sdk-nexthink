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
export NEXTHINK_API=public
export NEXTHINK_EXPORT_QUERY_ID="${NEXTHINK_EXPORT_QUERY_ID:-$NEXTHINK_QUERY_ID}"
umask 077
lab_example_output=$(mktemp -d "${TMPDIR:-/tmp}/nexthink-examples.XXXXXX")
lab_example_root=$(pwd)
lab_example_failures=0
for lab_example_source in examples/nexthink/_build_client/*/main.go examples/nexthink/public_api/nql/*/main.go examples/nexthink/public_api/remote_actions/{ListRemoteActions,GetRemoteActionDetails}/main.go examples/nexthink/public_api/workflows/{ListWorkflows,GetWorkflowDetails}/main.go; do
    lab_example_package=${lab_example_source%/main.go}
    lab_example_name=${lab_example_package//\//_}
    if go build -o "$lab_example_output/$lab_example_name" "./$lab_example_package" &&
       (cd "$lab_example_output" && "./$lab_example_name" > "$lab_example_name.log" 2>&1); then
        printf 'PASS %s\n' "$lab_example_package"
    else
        printf 'FAIL %s\n' "$lab_example_package"
        lab_example_failures=$((lab_example_failures + 1))
    fi
    cd "$lab_example_root"
done
printf 'Private logs and exports: %s\n' "$lab_example_output"
exit "$lab_example_failures"
