package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	// Execute 使用 args（默认情况下为 os.args[1:]），并在命令树中运行，为命令找到合适的匹配项，然后找到相应的标志。
	return rootCmd.Execute()
}

func init() {
	// 每一个子命令都储存在一个命令树中
	// rootCmd 相当于根命令，内容是空的，子命令都存在它下面
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
	rootCmd.AddCommand(jsonCmd)
}
