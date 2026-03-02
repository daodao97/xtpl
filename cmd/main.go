package main

import (
	"log/slog"
	"os"

	"xproxy/admin"
	"xproxy/api"
	"xproxy/conf"
	"xproxy/dao"
	"xproxy/job"

	"github.com/daodao97/xgo/xapp"
	"github.com/daodao97/xgo/xlog"
	"github.com/daodao97/xgo/xredis"
	"github.com/daodao97/xgo/xrequest"
	"github.com/gin-gonic/gin"
)

var Version string

func init() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "dev" {
		xlog.SetLogger(xlog.StdoutTextPretty(xlog.WithLevel(slog.LevelDebug)))
		xrequest.SetRequestDebug(true)
	} else {
		xlog.SetLogger(xlog.StdoutJson(xlog.WithLevel(slog.LevelInfo)))
		xrequest.SetRequestDebug(false)
	}
}

func main() {
	app := xapp.NewApp().
		AddStartup(
			conf.Init,
			func() error {
				return xredis.Inits(conf.Get().Redis)
			},
			dao.Init,
		).
		AddServer(xapp.NewGinHttpServer(xapp.Args.Bind, h))

	if os.Getenv("ENABLE_CRON") == "true" {
		app.AddServer(job.NewCronServer())
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func h() *gin.Engine {
	r := xapp.NewGin(xapp.WithPrintReqeustLog(false))
	api.Setup(r)
	admin.SetupRouter(r)

	return r
}
