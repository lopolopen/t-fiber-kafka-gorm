FROM --platform=$BUILDPLATFORM golang:1.25.0-alpine AS builder

WORKDIR /app

# ENV GOPROXY=https://goproxy.io,direct
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -ldflags="-s -w" -o main ./cmd/api

FROM alpine:latest

# RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/cmd/api/etc ./etc

CMD ["./main", "-f", "etc/config.yaml"]