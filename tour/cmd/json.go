package cmd

import (
	"log"

	"tour/internal/json2struct"

	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json转换和处理",
	Long:  "json转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var json2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "json转换",
	Long:  "json转换",
	Run: func(cmd *cobra.Command, args []string) {
		// 创建解析器实例
		parser, err := json2struct.NewParser(str)

		if err != nil {
			// 如果创建失败，就记入日志并退出程序
			log.Fatalf("json2struct.NewParser err: %v", err)
		}

		// 使用解析器进行格式转换
		content := parser.Json2Struct()

		// 输出，并记入日志
		log.Printf("输出结构： %s", content)
	},
}

func init() {
	jsonCmd.AddCommand(json2structCmd)

	json2structCmd.Flags().StringVarP(&str, "str", "s", "", "请输入json字符串")
}
