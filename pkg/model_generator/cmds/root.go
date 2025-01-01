package cmds

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "modelgen",
	Short: "一个数据库模型生成器",
	Long:  "根据数据库表生成对应的golang, proto模型",
	Run: func(cmd *cobra.Command, args []string) {
		help, err := cmd.Flags().GetBool("help")
		if err != nil {
			panic(err)
		}
		if help {
			err := cmd.Help()
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	RootCmd.PersistentFlags().Bool("help", false, "help for modelgen")
	RootCmd.PersistentFlags().String("dsn", "", "数据库资源名称")
	RootCmd.PersistentFlags().String("table", "", "数据库名")
	RootCmd.PersistentFlags().String("dst", "", "目标地址")
	RootCmd.AddCommand(ModelCmd)
	RootCmd.AddCommand(MessageCmd)
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
