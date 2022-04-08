FROM --platform=${TARGETPLATFORM:-linux/amd64} ghcr.io/openfaas/of-watchdog:0.9.3 as watchdog
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.17-alpine3.15 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

# Required to enable Go modules
RUN apk add --no-cache git

# Allows you to add additional packages via build-arg
ARG CGO_ENABLED=0
ARG GO111MODULE="on"

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

ENV CGO_ENABLED=0

RUN mkdir -p /go/src/handler
WORKDIR /go/src/handler
COPY . .
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go test ./... -cover
RUN CGO_ENABLED=${CGO_ENABLED} GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build --ldflags "-s -w" -a -installsuffix cgo -o handler .

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3.14
RUN apk --no-cache add ca-certificates \
    && addgroup -S app && adduser -S -g app app

RUN mkdir -p /home/app \
    && chown app /home/app

WORKDIR /home/app

COPY --from=builder --chown=app /go/src/handler/handler    .
COPY --from=builder --chown=app /go/src/handler/migrations    /home/app/migrations
COPY --from=builder --chown=app /usr/bin/fwatchdog         .

USER app

ENV fprocess="./handler"
ENV mode="http"
ENV upstream_url="http://127.0.0.1:8082"
ENV prefix_logs="false"

CMD ["./fwatchdog"]