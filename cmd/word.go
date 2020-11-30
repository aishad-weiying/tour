package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"tour/internal/work"
	"log"
)

// 声明转换模式
const (
	MODE_UPPER                         = iota + 1
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWRE_CAMLCASE
	MODE_CMALCASE_TO_UNDERSCORE
)

var desc = strings.Join([]string{
	"该子命令支持各种格式的单词转换，模式如下",
	"1: 全部单词转换为小写",
	"2: 全部单词转换为大写",
	"3: 下划线单词转换为大写驼峰单词",
	"4: 下划线单词转换为小写驼峰单词",
	"5: 驼峰单词转换为下划线单词",
},"\n")

var wordCmd = &cobra.Command{
	Use:   "word",        // 子命令
	Short: "单词格式转换",      // 短格式的帮助信息，显示在 命令 -h 中
	Long:  desc, // 长格式的帮助信息，显示在 命令 子命令 -h中
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_LOWER:
			content = work.ToLower(str)
		case MODE_UPPER:
			content  = work.ToUpper(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = work.UndersocoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWRE_CAMLCASE:
			content = work.UndersocoreToLowerCamelCase(str)
		case MODE_CMALCASE_TO_UNDERSCORE:
			content = work.CamelCaseToUndersocore(str)
		default:
			log.Println("暂不支持该转换模式，请运行 help word 查看帮助文档")
		}
		log.Println("输出结果",content)
	},                    // 这行这个子命令的生活要执行的函数
}

var str string // 要转换的字符串
var mode int8	// 要转换的模式

func init() {
	wordCmd.Flags().StringVarP(&str,"str","s","","请输入要转换的字符串")
	wordCmd.Flags().Int8VarP(&mode,"mode","m",0,"请输入要转换的模式")
}