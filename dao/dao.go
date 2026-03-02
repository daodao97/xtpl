package dao

import (
	"xproxy/conf"

	xcache "github.com/daodao97/xgo/cache"
	"github.com/daodao97/xgo/xdb"
	"github.com/daodao97/xgo/xredis"
	_ "github.com/go-sql-driver/mysql"
)

func Init() error {
	err := xdb.Inits(conf.Get().Database)
	if err != nil {
		return err
	}
	if rdb := xredis.Get(); rdb != nil {
		xdb.SetCache(xcache.NewRedis(rdb, xcache.WithPrefix("xdb")))
	}

	ProjectUser = xdb.New("project_user", xdb.WithCacheKey("id"))

	return nil
}
