package commands

import (
	"log"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of athena",
	Long:  `All software has versions. This is athena`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("athena v1.0 -- HEAD")
	},
}
