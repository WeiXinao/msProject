package ormx

import (
	"errors"

	"xorm.io/xorm"
)

type DBConn interface {
	Rollback() error
	Commit() error
}

type Transaction interface {
	Tx(fn func(session any) error) error
	EngineTx(fn func(engine *xorm.Engine) error) error
}

var	ErrTypeConvert    = errors.New("类型转换错误")