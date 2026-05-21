FROM golang:1.25.0-alpine AS builder

WORKDIR /app

# Use a reliable GOPROXY; change if needed
ENV GOPROXY=https://goproxy.cn,direct

# RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Provide sane defaults so CI can pass different targets via --build-arg
ARG TARGETOS=linux
ARG TARGETARCH=amd64
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w" -o main ./cmd/api

FROM alpine:3.23.4

WORKDIR /app

ARG COMMIT_SHA
ENV COMMIT_SHA=${COMMIT_SHA}

# Create a non-root user and ensure permissions for /app
RUN addgroup -S app && adduser -S app -G app && mkdir -p /app/etc && chown -R app:app /app

COPY --from=builder /app/main .
COPY --from=builder /app/cmd/api/etc ./etc

USER app

EXPOSE 8080

CMD ["./main", "-f", "etc/config.yaml"]