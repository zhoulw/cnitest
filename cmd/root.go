package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run hugo...")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// 在执行每个命令之前调用此函数
		// 可以在这里进行一些初始化操作
		// 例如加载配置文件
		fmt.Println("preRun called·····")
	},
}

var subCmd = &cobra.Command{}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 本地标志，它只适用于该指令命令
	var Source string
	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")

	// 父命令的本地标志

	// 必选标志
	var Region string
	rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	rootCmd.MarkFlagRequired("region")
}
