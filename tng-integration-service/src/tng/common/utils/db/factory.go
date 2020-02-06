package db

import "github.com/astaxie/beego/orm"

// Factory is DB transaction factory.
type Factory interface {
	GetDB(withTx ...bool) *DB
	GetDBByAlias(alias string, withTx ...bool) *DB
	Commit(db *DB)
	Rollback(db *DB)
}

type factory struct {
}

// NewDBFactory returns a new Factory instance.
func NewDBFactory() Factory {
	return &factory{}
}

// GetDB returns a new instance of DB.
// withTx = true => returns a DB with Tx.
func (f *factory) GetDB(withTx ...bool) *DB {
	isTx := false
	if len(withTx) > 0 {
		isTx = withTx[0]
	}
	o := orm.NewOrm()
	if isTx {
		_ = o.Begin()
	}
	return &DB{o, isTx}
}

// GetDBByAlias returns a new instance of DB by alias.
// withTx = true => returns a DB with Tx.
func (f *factory) GetDBByAlias(alias string, withTx ...bool) *DB {
	isTx := false
	if len(withTx) > 0 {
		isTx = withTx[0]
	}
	o := orm.NewOrm()
	err := o.Using(alias)
	if err != nil {
		return nil
	}
	if isTx {
		_ = o.Begin()
	}
	return &DB{o, isTx}
}

func (f *factory) Commit(db *DB) {
	if db.withTx {
		_ = db.Commit()
	}
}

func (f *factory) Rollback(db *DB) {
	if db.withTx {
		_ = db.Rollback()
	}
}
