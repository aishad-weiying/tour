package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: "支持多种单词格式的转换",
	Version: "1.0.1",
}

func Execute() error {
	return rootCmd.Execute()
}

func init()  {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}