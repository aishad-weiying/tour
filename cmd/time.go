package cmd

import (
	"github.com/spf13/cobra"
)

var timeCmd  = &cobra.Command{
	Use:"time",//time 子命令
	Short:"时间格式处理",
	Long:"时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init()  {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
}

