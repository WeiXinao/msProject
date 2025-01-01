package cmds

import (
	"github.com/spf13/cobra"
)

var MessageCmd = &cobra.Command{
	Use:   "msg",
	Short: "生成数据库对应于 proto 语言的模型",
}
