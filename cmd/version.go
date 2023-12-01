package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "print the version og cnitest",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CniTest version is v0.0.1 ")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
