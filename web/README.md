# xtpl-web

Generated from the gossr `minimal` Vue template.

```bash
pnpm install
pnpm dev       # Vite on 127.0.0.1:3333
pnpm typecheck
pnpm build     # dist/client + dist/server/server.js
```

The generated `embed.go` exposes `Dist`. Wire it into the host Go application:

```go
type homePayload map[string]any

func (payload homePayload) AsMap() map[string]any { return payload }

data := gossr.NewDataEngine()
data.GET("/", gossr.WrapSSR(func(c *gin.Context) (gossr.SSRPayload, error) {
    return homePayload{
        "message": "Hello from Go",
        "path": c.Request.URL.Path,
    }, nil
}))

if err := gossr.SsrWithOptions(router, web.Dist, gossr.Options{
    DataEngine: data,
}); err != nil {
    log.Fatal(err)
}
```

The integration lives in `server.go`. In development, run `make web_dev` and `make dev` in separate terminals, then open `http://127.0.0.1:4001`. The Vite proxy intentionally keeps `changeOrigin: false` so gossr's same-origin authorization accepts `/_ssr/data` requests.

Project policy belongs in the host application: authentication through `SessionResolver`/middleware, data permissions in Go handlers, and application-specific locale rules. The generated files only implement the rendering protocol.
