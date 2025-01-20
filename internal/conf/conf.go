package conf

import (
	"fmt"
	"log/slog"

	"github.com/daodao97/xgo/xapp"
	"github.com/daodao97/xgo/xdb"
	"github.com/daodao97/xgo/xredis"

	"github.com/daodao97/xgo/xlog"
)

type config struct {
	AdminPath string           `yaml:"admin_path"`
	JwtSecret string           `yaml:"jwt_secret"`
	Database  []xdb.Config     `yaml:"database" envPrefix:"DATABASE"`
	Redis     []xredis.Options `yaml:"redis" envPrefix:"REDIS"`
}

func (c *config) Print() {
	xlog.Debug("load config", slog.Any("config", fmt.Sprintf("%+v", c)))
}

var _c *config

func Get() *config {
	return _c
}

func InitConf() error {
	_c = &config{
		AdminPath: "/_",
		JwtSecret: "change_me",
	}

	if err := xapp.InitConf(_c); err != nil {
		return err
	}
	_c.Print()

	return nil
}
