package conf

import (
	"github.com/daodao97/xgo/xapp"
	"github.com/daodao97/xgo/xdb"
	"github.com/daodao97/xgo/xredis"
)

type Conf struct {
	Database []xdb.Config     `json:"database" yaml:"database" envPrefix:"DATABASE"`
	Redis    []xredis.Options `json:"redis" yaml:"redis" envPrefix:"REDIS"`
}

var ConfInstance *Conf

func Init() error {
	ConfInstance = &Conf{}

	err := xapp.InitConf(ConfInstance)
	if err != nil {
		return err
	}

	return nil
}

func Get() *Conf {
	return ConfInstance
}
