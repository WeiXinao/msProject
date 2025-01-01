package cmds

import (
	"github.com/spf13/cobra"
)

var ModelCmd = &cobra.Command{
	Use:   "model",
	Short: "生成 go 语言的数据库模型",
}
