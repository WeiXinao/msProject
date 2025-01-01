package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	MySQLDriverName = "mysql"
)

func InitMysqlDB(dsn string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(MySQLDriverName, dsn)
	return engine, err
}

type XormDao struct {
	db *xorm.Engine
}

func NewXormDao(engine *xorm.Engine) *XormDao {
	return &XormDao{
		db: engine,
	}
}

func (x *XormDao) ShowTableCols(ctx context.Context, tableName string) (Table, error) {
	results, err := x.db.Context(ctx).QueryString(fmt.Sprintf("SHOW FULL COLUMNS FROM %s", tableName))
	if err != nil {
		return Table{}, err
	}
	return QueryResultToTable(tableName, results), nil
}
