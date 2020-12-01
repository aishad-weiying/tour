package cmd

import (
	"github.com/spf13/cobra"
	"tour/internal/timer"
	"fmt"
	"time"
)

var nowTimeCmd = &cobra.Command{
	Use:"now",// time的now子命令，用户获取当前时间
	Short:"获取当前时间",
	Long:"获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime	:= timer.GetNowTime()
		fmt.Println("当前时间：",nowTime.Format(time.ANSIC ),"当前时间戳：",nowTime.Unix())
	},
}