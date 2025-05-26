package cmd

import (
	"fmt"

	"github.com/oxidnova/novadm/backend/internal"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version for this service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:\t", internal.Version)
		fmt.Println("Building:\t", internal.GOVersion)
		fmt.Println("Commit Sha:\t", internal.Commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
