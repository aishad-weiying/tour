package cmd

import (
	"github.com/spf13/cobra"
	"time"
	"tour/internal/timer"
	"strings"
	"strconv"
	"fmt"
)

var calculateTime string // 用户输入的时间
var duration string // duration 类型的时间

var calculateTimeCmd = &cobra.Command{
	Use:"calc",// time的calc子命令，用来做时间计算
	Short:"计算需要的时间",
	Long:"计算需要的时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTime time.Time
		var err error
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" { // 如果传入的时间为空，则获取当前时间
			currentTime = timer.GetNowTime()
		}else {
			if !strings.Contains(calculateTime," ") {// 如果 calculateTime 中不存在空格，说明时间格式为2006-01-02
				layout = "2006-01-02"
			}
			// 如果包含空格，说明时间格式为2006-01-02 15:04:05
			currentTime ,err = time.Parse(layout,calculateTime)
			if err != nil {
				// 说明传入的是时间戳类型的时间
				// 将字符串转换成整型
				r, _ := strconv.Atoi(calculateTime)
				// 将时间戳类型转换为时间类型
				currentTime = time.Unix(int64(r),0)
			}
		}
		calc,err := timer.GetCalculateTime(currentTime,duration)
		if err != nil {
			fmt.Println("时间转换失败：",err)
		}
		fmt.Println("结果：",calc.Format(layout),"时间戳：",calc.Unix())
	},
}

func init()  {
	calculateTimeCmd.Flags().StringVarP(&calculateTime,"calculate","c","","需要计算的时间，有效单位为时间戳或者格式化之后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration,"duration","d","","持续时间")
}
