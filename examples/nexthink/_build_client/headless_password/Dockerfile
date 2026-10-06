# Build from the repository root using the command in README.md.
FROM golang:1.26.0-bookworm AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY nexthink ./nexthink
COPY examples/nexthink/_build_client/headless_password ./example
RUN CGO_ENABLED=0 go build -trimpath -o /out/nexthink-password ./example
RUN go build -o /out/playwright github.com/mxschmitt/playwright-go/cmd/playwright

FROM debian:bookworm-slim
ENV PLAYWRIGHT_DRIVER_PATH=/opt/playwright-driver \
    PLAYWRIGHT_BROWSERS_PATH=/opt/playwright-browsers
COPY --from=build /out/playwright /usr/local/bin/playwright
# Installation happens at image-build time, never during an API request.
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && playwright install --with-deps chromium \
    && rm -rf /var/lib/apt/lists/* \
    && useradd --create-home --uid 10001 sdk
# The driver installer gives Node owner-only execution permission.
# Make the root-installed runtime executable by the unprivileged SDK user.
RUN chmod 755 /opt/playwright-driver/node
COPY --from=build /out/nexthink-password /usr/local/bin/nexthink-password
USER sdk
WORKDIR /home/sdk
ENTRYPOINT ["/usr/local/bin/nexthink-password"]
