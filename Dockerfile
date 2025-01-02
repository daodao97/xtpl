FROM golang:1.22 AS builder

ARG BUILD_VERSION

WORKDIR /build

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=${BUILD_VERSION}" -o app ./cmd/main.go

FROM alpine:latest AS final

WORKDIR /app
COPY --from=builder /build/app /app/
COPY *.yaml /app/
COPY openapi.json /app/openapi.json

RUN apk update && \
    apk add --no-cache sudo tzdata

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT ["/app/app"]