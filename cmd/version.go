package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of hey",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hey " + version)
	},
}
