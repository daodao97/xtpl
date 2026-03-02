FROM node:20-alpine AS frontend-builder

WORKDIR /build/adminui

# 安装 pnpm
RUN corepack enable && corepack prepare pnpm@latest --activate

# 先复制 package.json 和 lock 文件，利用缓存
COPY adminui/package.json adminui/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

# 复制源码并构建
COPY adminui/ ./
RUN pnpm build

FROM golang:1.25-alpine AS builder

ARG BUILD_VERSION

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 复制前端构建产物到 adminui/ui 目录
COPY --from=frontend-builder /build/adminui/ui ./adminui/ui

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w -X main.Version=${BUILD_VERSION}" \
    -trimpath \
    -o app ./cmd/main.go

FROM alpine:3.19 AS final

WORKDIR /app
COPY --from=builder /build/app /app/
COPY *.yaml /app/

RUN apk update && \
    apk add --no-cache tzdata && \
    rm -rf /var/cache/apk/*

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT ["/app/app"]