package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// rootCmd is initialize command line
var rootCmd = &cobra.Command{
	Use:   "Sys Forge Api",
	Short: "an Nova Admin service",
	Long:  "Top level command for api service",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// it would be start up websocket server
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
