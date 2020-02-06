package db

import "github.com/astaxie/beego/orm"

// DB is the wrapper of orm.Ormer.
type DB struct {
	orm.Ormer
	withTx bool
}
