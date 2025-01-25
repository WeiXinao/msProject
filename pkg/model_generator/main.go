package main

import (
	"context"
	_ "embed"
	"github.com/WeiXinao/msProject/pkg/model_generator/cmds"
	"github.com/spf13/cobra"
	"os"
	"text/template"
	"time"
)

var (
	//go:embed templates/model.tmpl
	modelTmpl     string
	modelTmplName string = "model.tmpl"
	//go:embed templates/message.tmpl
	messageTmpl     string
	messageTmplName string = "message.tmpl"
)

func main() {
	cmds.ModelCmd.Run = func(cmd *cobra.Command, args []string) {
		dsn, err := cmd.Flags().GetString("dsn")
		if err != nil {
			panic(err)
		}
		table, err := cmd.Flags().GetString("table")
		if err != nil {
			panic(err)
		}
		dst, err := cmd.Flags().GetString("dst")
		if err != nil {
			panic(err)
		}
		err = GenModel(dsn, table, dst)
		if err != nil {
			panic(err)
		}
	}

	cmds.MessageCmd.Run = func(cmd *cobra.Command, args []string) {
		dsn, err := cmd.Flags().GetString("dsn")
		if err != nil {
			panic(err)
		}
		table, err := cmd.Flags().GetString("table")
		if err != nil {
			panic(err)
		}
		dst, err := cmd.Flags().GetString("dst")
		if err != nil {
			panic(err)
		}
		err = GenMessage(dsn, table, dst)
		if err != nil {
			panic(err)
		}
	}
	cmds.Execute()
}

func GenMessage(dsn, tableName, dst string) error {
	table, err := GetTableInfo(dsn, tableName)
	if err != nil {
		return err
	}
	message, err := table.ToMessage()
	if err != nil {
		return err
	}
	err = ParseTemplate(messageTmpl, messageTmplName, dst, message)
	if err != nil {
		panic(err)
	}
	return err
}

func GenModel(dsn, tableName, dst string) error {
	table, err := GetTableInfo(dsn, tableName)
	if err != nil {
		return err
	}
	model, err := table.ToModel()
	if err != nil {
		return err
	}
	err = ParseTemplate(modelTmpl, modelTmplName, dst, model)
	if err != nil {
		panic(err)
	}
	return err
}

func GetTableInfo(dsn, tableName string) (Table, error) {
	db, err := InitMysqlDB(dsn)
	if err != nil {
		panic(err)
	}
	dao := NewXormDao(db)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	table, err := dao.ShowTableCols(ctx, tableName)
	cancel()
	if err != nil {
		return Table{}, err
	}
	return table, nil
}

func ParseTemplate(tmpl, tmplName, dst string, data any) error {
	fm := template.FuncMap{
		"LeadingSpace": func(s string) string {
			if s != "" {
				return " " + s
			}
			return s
		},
		"AddOne": func(a int) int {
			return a + 1
		},
	}
	t := template.New(tmplName)
	t.Funcs(fm)
	parser, err := t.Parse(tmpl)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(dst, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	return parser.Execute(file, data)
}
