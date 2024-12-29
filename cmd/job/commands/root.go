package commands

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "athena-job",
	Short: "athena-job",
	Long:  `athena-job is a tool that helps you create and manage jobs.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("hello,athena-job")
	},
}

// Execute 执行命令行操作
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("exec job err: ", err)
		os.Exit(1)
	}
}
