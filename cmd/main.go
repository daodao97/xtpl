package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/daodao97/xgo/utils"
	"github.com/daodao97/xgo/xapp"
	"github.com/daodao97/xgo/xdb"
	"github.com/daodao97/xgo/xlog"
	"github.com/daodao97/xgo/xredis"

	"xtpl/internal/admin"
	"xtpl/internal/api"
	"xtpl/internal/conf"
	"xtpl/internal/dao"
)

var Version string

func init() {
	if !utils.IsGoRun() {
		xlog.SetLogger(xlog.StdoutJson(xlog.WithLevel(slog.LevelDebug)))
	}
}

func main() {
	app := xapp.NewApp().
		AddStartup(
			conf.InitConf,
			func() error {
				return xdb.Inits(conf.Get().Database)
			},
			func() error {
				return xredis.Inits(conf.Get().Redis)
			},
		).
		AfterStarted(func() {
			xlog.Debug("version", xlog.String("version", Version))
			dao.Init()
		}).
		AddServer(xapp.NewHttp(xapp.Args.Bind, h))

	if err := app.Run(); err != nil {
		fmt.Printf("Application error: %v\n", err)
	}
}

func h() http.Handler {
	e := xapp.NewGin()
	defer func() {
		xapp.GenerateOpenAPIDoc(
			e,
			xapp.WithInfo("xtpl", Version),
			xapp.WithBearerAuth(),
		)
	}()

	api.SetupRouter(e)
	admin.SetupRouter(e)
	return e.Handler()
}
