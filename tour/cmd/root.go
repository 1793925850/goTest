package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 每一个子命令都存在一个命令树中
	rootCmd.AddCommand(wordCmd)
}
