# xtpl

`xtpl` 是一个用于快速创建 Go Web 项目的模板，内置：

- Gin + xapp 后端启动、配置、数据库、Redis 与定时任务骨架
- `/_` 管理后台及独立的 `adminui` Vue SPA
- 可选的 gossr + Goja 公共 Vue SSR 站点
- Makefile、Air、Docker 与 Compose 构建配置

## 从模板创建项目

推荐先在 GitHub 使用 **Use this template** 创建新仓库，再克隆新仓库。也可以直接复制：

```bash
git clone git@github.com:daodao97/xtpl.git my-app
cd my-app
rm -rf .git
git init
```

将默认 Go module `xproxy` 替换为正式 module。脚本会同时更新所有 Go import，并在项目同级目录创建带时间戳的完整备份：

```bash
./update_mod_refs.sh . github.com/your-org/my-app
```

确认新项目正常后，可以删除脚本生成的备份目录并提交第一次变更。

## 初始化检查清单

开始业务开发前至少完成以下修改：

1. 修改 `conf.dev.yaml` 中的 MySQL DSN、数据库名和 Redis 地址；不要提交生产密码。
2. 修改 `admin/server.go` 中的后台标题、Logo 和默认头像。
3. **替换 `admin/server.go` 中的 JWT Secret**，生产环境应从安全配置或环境变量读取。
4. 修改 `web/server.go` 的首页 payload 和 `web/index.html` 的页面标题。
5. 按部署环境检查 `compose.yaml` 的端口、数据库与 Redis 网络地址。

gossr 不提供默认身份认证。需要登录态时，由宿主应用通过 Gin 中间件和 `gossr.Options.SessionResolver` 接入；所有敏感 API 仍必须在 Go 后端鉴权。

## 选择前端模式

公共站点的 SSR 是模板默认能力，不是强制依赖。新项目可以按场景选择：

| 模式 | 适合场景 | 建议 |
| --- | --- | --- |
| gossr SSR | 公开网站、SEO、动态 Head、首屏内容 | 保留默认 `web/` |
| 纯 Vue SPA + Go API | 登录后系统、工具、无需 SEO 的业务前端 | 移除 gossr，前端独立构建和部署 |
| 仅管理后台/API | 内部服务、管理工具、暂时没有公共站点 | 删除 `web/`，保留 `adminui` |

### 不使用 SSR：仅保留管理后台和 API

执行以下调整：

1. 删除 `cmd/main.go` 中的 `xproxy/web` import 和 `web.SetupRouter(r)` 调用。
2. 删除整个 `web/` 目录。
3. 从 Makefile 删除 `web`、`web_dev` 目标；让 `run`、`build` 只依赖 `admin`，并移除 `dev` 对 `web` 的依赖和 SSR 环境变量。
4. 从 Dockerfile 删除 `web-builder` 阶段，以及复制 `/build/web/dist` 的语句。
5. 从 `.air.toml` 删除 `DEV_MODE`、`DEV_SERVER_URL`。
6. 整理 Go 依赖：

```bash
go mod tidy
```

`go mod tidy` 会自动移除 gossr、goja 及仅由 SSR 使用的间接依赖，不需要手工编辑 `go.sum`。

### 不使用 SSR：创建纯 Vue SPA

先按上一节移除默认 `web/` 和 gossr，再使用标准 Vite 创建前端。建议放在 `web/`，保持目录语义一致：

```bash
pnpm create vite web --template vue-ts
pnpm --dir web install
```

开发时在 Vite 中把 `/api` 代理到 Go 服务：

```ts
// web/vite.config.ts
export default defineConfig({
  server: {
    port: 3333,
    proxy: {
      '/api': 'http://127.0.0.1:4001',
    },
  },
})
```

分别启动前后端：

```bash
pnpm --dir web dev
go run ./cmd --app-env dev --bind :4001
```

生产环境推荐把 `pnpm --dir web build` 生成的静态文件部署到 CDN、对象存储或 Nginx，让 Go 服务只提供 `/api`。这样路由 fallback、缓存和压缩都由静态服务器负责，后端不需要模拟 SSR。

如果必须使用同一个容器，可以在 Dockerfile 中保留一个 SPA builder，再由 Nginx sidecar 或网关服务其 `dist/`；不建议把 SPA 强行接回 gossr，因为这会重新引入不需要的 SSR 运行时。

无论选择 SSR 还是 SPA，浏览器中的路由守卫都不能代替后端鉴权，敏感数据接口必须由 Gin 中间件或具体 handler 验证用户权限。

## 环境要求

- Go 1.25.8+
- Node.js `^20.19.0` 或 `>=22.12.0`
- pnpm 10.20.0
- MySQL
- Redis
- 可选：Docker 与 Docker Compose

安装两个前端的依赖：

```bash
make install
```

## 本地开发

先确保 `conf.dev.yaml` 指向可用的 MySQL 和 Redis。

### 公共 SSR 站点

终端一，启动 Vite：

```bash
make web_dev
```

终端二，启动 Go 服务：

```bash
make dev
```

访问 <http://127.0.0.1:4001>。`make dev` 会先构建一次服务端 SSR bundle，然后由 Go 服务代理浏览器资源到 Vite。

### 管理后台

另开终端启动管理端 Vite：

```bash
make admin_dev
```

默认管理端开发端口为 `3001`。生产构建后的管理后台由 Go 服务挂载在 `/_`。

### 路由边界

| 路径 | 用途 |
| --- | --- |
| `/` | 公共 Vue SSR 站点 |
| `/_ssr/data/*` | gossr 页面数据通道 |
| `/_` | 管理后台 |
| `/api/*` | 宿主业务 API，可在 `api/setup.go` 注册 |

明确注册的 API 和管理端路由优先，gossr 只接管剩余的公共页面请求。

## 添加公共页面

页面采用文件路由：

1. 在 `web/src/pages/` 新建 Vue 页面。
2. 在 `web/server.go` 的独立 `DataEngine` 中注册对应的 payload handler。
3. 页面通过 `useSsrData<T>()` 读取服务端数据。
4. 不需要页面数据的静态路由可设置 `meta.ssrData: false`。

修改普通页面导航时可以直接使用 `RouterLink`。更完整的 gossr 接入说明见 [docs/gossr.md](docs/gossr.md)。

## 测试与构建

```bash
make typecheck
go test ./...
go vet ./...
```

生产构建会依次构建 adminui、公共 SSR 前端和 Go 二进制：

```bash
make build
./build/server --bind :4001
```

SSR 集成测试需要先存在 `web/dist/server/server.js`；执行 `make web` 后，`go test ./web -v` 会进行真实 Goja SSR 渲染验证。

## Docker

```bash
docker compose up --build
```

默认映射到 <http://127.0.0.1:8080>。容器中的 `127.0.0.1` 指向容器自身，部署前需要将数据库和 Redis 地址改成 Compose 服务名或外部可访问地址。

Docker 构建阶段会分别生成：

- `adminui/ui`：管理后台静态资源
- `web/dist/client`：公共站点客户端资源
- `web/dist/server/server.js`：Goja SSR bundle
- `/app/app`：包含上述产物的最终 Go 二进制

## 项目结构

```text
.
├── cmd/            # 应用入口
├── api/            # 业务 API 注册
├── admin/          # 管理端路由、模型与配置
├── adminui/        # 管理后台 Vue SPA
├── web/            # 公共 Vue SSR 站点及 gossr 接入
├── conf/           # 配置结构与加载
├── dao/            # 数据访问层
├── job/            # 定时任务
├── docs/           # 扩展文档
└── Makefile        # 常用开发与构建命令
```
