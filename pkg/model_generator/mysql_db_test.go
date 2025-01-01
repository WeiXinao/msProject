package main

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestXormDao_ShowTableCols(t *testing.T) {
	db, err := InitMysqlDB("root:123456@tcp(192.168.5.4:3307)/ms_project?charset=utf8")
	require.NoError(t, err)
	dao := NewXormDao(db)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	table, err := dao.ShowTableCols(ctx, "ms_project_menu")
	cancel()
	require.NoError(t, err)
	t.Logf("%+v", table)

	model, err := table.ToModel()
	require.NoError(t, err)
	t.Logf("%+v", model)
}
