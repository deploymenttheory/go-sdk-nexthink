# Headless local-password authentication

This example lists application configurations using a local Nexthink username and password. The SDK launches its own headless Chromium browser when authentication is needed. It does not require desktop Chrome, an existing browser session, a display, or an exported session file.

The account must permit **password-only local login**. SSO, MFA, password-change requirements, and locked accounts return an authentication error; the SDK does not prompt a person or change tenant policy. API client IDs and secrets are separate credentials and cannot be used here.

## Install the browser runtime

Use the same Playwright Go version as this checkout's [go.mod](../../../../go.mod). Install its driver and Chromium explicitly before running the program:

```sh
# macOS and Windows; also Linux with browser dependencies already installed:
go run github.com/mxschmitt/playwright-go/cmd/playwright@v0.6201.1 install chromium

# Linux runner/image preparation, including required system libraries:
go run github.com/mxschmitt/playwright-go/cmd/playwright@v0.6201.1 install --with-deps chromium
```

The Linux dependency installation may require administrator access. The SDK does not install or download browsers during requests. Unset `DEBUGP` and `PWDEBUG` when using password authentication: any nonempty value, including `0`, returns `password.ErrBrowserUnavailable` before browser startup. These Playwright debugging modes can expose protocol credentials or invoke interactive behavior. Install as the runner user, or set `PLAYWRIGHT_DRIVER_PATH` and `PLAYWRIGHT_BROWSERS_PATH` consistently during installation and execution. For shared installations, ensure the runtime user can execute the driver's `node` binary; the Dockerfile explicitly sets that permission after root installation. A preinstalled Chromium executable can be selected with `NEXTHINK_BROWSER_EXECUTABLE_PATH`; it must be compatible with the pinned driver, which is still required. The managed Chromium installation is the default.

## Run

Inject the account credentials using your secret manager. The placeholders below name the required values; do not store real credentials in scripts or commit them to the repository.

```sh
export NEXTHINK_INSTANCE=your-instance
export NEXTHINK_REGION=eu
export NEXTHINK_USERNAME='your-local-account'
export NEXTHINK_PASSWORD='your-local-password'
go run ./examples/nexthink/_build_client/headless_password
```

The program uses `nexthink.NewClient` with `BrowserCredentials.UsernamePassword`. Its result output contains the HTTP status and application counts; the SDK also emits redacted transport diagnostics. It then closes the SDK-owned authentication provider. The two-minute request context includes the first login; an API-only 30-second context may be too short for browser startup and authentication.

For the equivalent environment-based constructor, set these additional variables and use `nexthink.NewClientFromEnv()`:

```sh
export NEXTHINK_API=web
export NEXTHINK_WEB_AUTH=password
export NEXTHINK_LOGIN_TIMEOUT=90s
```

`NEXTHINK_LOGIN_TIMEOUT` applies to `NewClientFromEnv`. The Go example explicitly sets `LoginTimeout: 90*time.Second`. Both constructors accept `NEXTHINK_BROWSER_PROXY` / `BrowserProxy` and `NEXTHINK_BROWSER_EXECUTABLE_PATH` / `BrowserExecutablePath` as appropriate. `BrowserProxy` applies to browser login and direct token renewal, separately from the API HTTP transport's `WithProxy` option. Browser TLS uses the browser/runner trust store; Go TLS callbacks and `WithTransport` do not configure Chromium. Certificate verification is not disabled.

## Linux container

Build from the repository root. The image installs the matching driver, Chromium and its dependencies during the build, and runs the example as a non-root user. It contains no tenant credentials.

```sh
docker build -f examples/nexthink/_build_client/headless_password/Dockerfile -t nexthink-password .
docker run --rm --init --shm-size=1g \
  -e NEXTHINK_INSTANCE -e NEXTHINK_REGION \
  -e NEXTHINK_USERNAME -e NEXTHINK_PASSWORD \
  nexthink-password
```

Use an ephemeral container and inject secrets at execution time, never with Docker build arguments. No session directory needs to be mounted. See [Playwright's container guidance](https://playwright.dev/docs/docker) for runner-specific process and shared-memory requirements.

## GitHub Actions with injected secrets

The following is an explicit, manually invoked **live** read-only job. Create the four named repository/environment secrets first. Do not attach a live credential job to untrusted pull requests. This is separate from the SDK's [fixture workflow](../../../../.github/workflows/browser-auth.yml), which uses only local test pages and no tenant secrets.

```yaml
name: Read Nexthink applications
on:
  workflow_dispatch:
permissions:
  contents: read
jobs:
  applications:
    runs-on: ubuntu-latest
    timeout-minutes: 10
    steps:
      - uses: actions/checkout@de0fac2e4500dabe0009e67214ff5f5447ce83dd
      - uses: actions/setup-go@4a3601121dd01d1626a1e23e37211e3254c1c06c
        with:
          go-version-file: go.mod
      - name: Prepare pinned browser runtime
        run: go run github.com/mxschmitt/playwright-go/cmd/playwright@v0.6201.1 install --with-deps chromium
      - name: List application configurations
        env:
          NEXTHINK_INSTANCE: ${{ secrets.NEXTHINK_INSTANCE }}
          NEXTHINK_REGION: ${{ secrets.NEXTHINK_REGION }}
          NEXTHINK_USERNAME: ${{ secrets.NEXTHINK_USERNAME }}
          NEXTHINK_PASSWORD: ${{ secrets.NEXTHINK_PASSWORD }}
        run: go run ./examples/nexthink/_build_client/headless_password
```

## Token lifetime and failures

Tokens and any issued refresh credentials remain in process memory. Each login uses an isolated browser context that is closed after acquisition or failure; no reusable profile or session-state file is saved. The provider shares one authentication attempt between concurrent callers and honors cancellation. Call `defer c.Close()` for SDK-owned providers; caller-supplied providers remain the caller's responsibility.

A valid cached access token is reused. Renewal uses a supported refresh exchange when available, otherwise a new headless password login. Failed credentials or an unsupported challenge stop further automatic credential submissions for that client. Correct the configuration and construct a new client after resolving the account condition. Temporary network errors do not trigger repeated password submissions within one attempt.

A Web API HTTP 401 invalidates the cached token for the next request. The failed API operation is not automatically replayed, so write operations are not duplicated by authentication recovery. HTTP 403 remains a permission error. Do not log credentials, tokens, or browser traffic; the SDK's authentication diagnostics omit secret values.

Controlled browser fixtures establish form handling, challenge rejection and cleanup behavior. They do not establish that every tenant allows unattended login or grants refresh tokens. See the [lab validation report](../../../../docs/lab-validation.md) for observed live behavior.

With the same local-account environment variables configured, explicitly run the read-only live test and wait for the server-issued access token to expire:

```sh
NEXTHINK_LIVE_PASSWORD_TEST=1 NEXTHINK_LIVE_WAIT_FOR_EXPIRY=1 \
  go test -v ./nexthink/auth/password -run '^TestLivePasswordAuthentication$' -count=1 -timeout 15m
```

This requires curl and the installed browser runtime. It compares SDK and curl responses before and after expiry without printing token values or response contents. It bounds the expiry observation to ten minutes and is never enabled by the fixture workflow.
