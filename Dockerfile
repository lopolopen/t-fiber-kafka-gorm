ARG TARGETOS=linux
ARG TARGETARCH=amd64
ARG BUILD_MODE=release

FROM golang:1.25.0-alpine AS builder
WORKDIR /app
# Use a reliable GOPROXY; change if needed
ENV GOPROXY=https://goproxy.cn,direct
# RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM builder AS build-debug
ARG TARGETOS
ARG TARGETARCH
ARG COMMIT_SHA
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -gcflags="all=-N -l" -tags=debug -ldflags="-X main.commitSHA=${COMMIT_SHA}" -o main ./cmd/api

FROM builder AS build-release
ARG TARGETOS
ARG TARGETARCH
ARG COMMIT_SHA
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w -X main.commitSHA=${COMMIT_SHA}" -o main ./cmd/api

FROM build-${BUILD_MODE} AS build-final

FROM alpine:3.23.4
WORKDIR /app
# Create a non-root user and ensure permissions for /app
RUN addgroup -S app && adduser -S app -G app && mkdir -p /app/etc && chown -R app:app /app
COPY --from=build-final /app/main .
COPY ./cmd/api/etc/config.yaml ./etc/config.yaml
USER app
EXPOSE 8080
CMD ["./main"]