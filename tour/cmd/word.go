package cmd

import (
	"log"
	"strings"

	"tour/internal/word"

	"github.com/spf13/cobra"
)

const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下划线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

var (
	str  string
	mode int8
	desc = strings.Join([]string{
		"该子命令支持各种单词格式转换，模式如下：",
		"1：全部转大写",
		"2：全部转小写",
		"3：下划线转大写驼峰",
		"4：下划线转小写驼峰",
		"5：驼峰转下划线",
	}, "\n")

	// wordCmd 子命令的具体内容
	wordCmd = &cobra.Command{
		Use:   "word",   // 子命令的“关键词”
		Short: "单词格式转换", // 工具集中每个可用命令的简短描述
		Long:  desc,     // 工具集中某个子命令的详细描述
		Run: func(cmd *cobra.Command, args []string) { // 子命令的执行函数
			var content string

			switch mode {
			case ModeUpper:
				content = word.ToUpper(str)
			case ModeLower:
				content = word.ToLower(str)
			case ModeUnderscoreToUpperCamelCase:
				content = word.UnderscoreToUpperCamelCase(str)
			case ModeUnderscoreToLowerCamelCase:
				content = word.UnderscoreToLowerCamelCase(str)
			case ModeCamelCaseToUnderscore:
				content = word.CamelCaseToUnderscore(str)
			default:
				log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
			}

			log.Printf("输出结果：%s", content)
		},
	}
)

// word 子命令的行参数的设置和初始化
func init() {
	// 给子命令加入条件标签
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")   // 表示该标签接收的输入为 string
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式") // 表示该标签接收的输入为 int
}
