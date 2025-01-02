package dao

import (
	"github.com/daodao97/xgo/xdb"
	_ "github.com/go-sql-driver/mysql"
)

var OperatorModel xdb.Model

func Init() {
	OperatorModel = xdb.New(
		"operator",
		xdb.WithFakeDelKey("is_deleted"),
	)
}
