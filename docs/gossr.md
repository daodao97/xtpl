# gossr integration

The public site uses `github.com/daodao97/gossr`; the existing `/_` admin UI remains a separate SPA.

SSR is optional for projects created from this template. See the “选择前端模式” section in the root README for instructions on removing gossr or replacing `web/` with a standalone Vue SPA.

## Development

Build the SSR bundle once, then start Vite and the Go host in separate terminals:

```bash
make web
make web_dev
```

```bash
make dev
```

Open <http://127.0.0.1:4001>. The Go host proxies public frontend requests to Vite while keeping `/_ssr/data` and existing API/admin routes in Gin.

## Production

```bash
make build
./build/server --bind :4001
```

The Docker build compiles both `adminui/ui` and `web/dist` before building the Go binary. The latter contains client assets and the Goja-compatible `server.js` bundle embedded by `web/embed.go`.

Page payload routes are registered in `web/server.go`. Authentication and authorization must remain in the host application; gossr does not provide a default identity implementation.
