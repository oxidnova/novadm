package cmd

import (
	"github.com/oxidnova/novadm/backend/svc"
	"github.com/oxidnova/novadm/backend/svc/api"
	"github.com/spf13/cobra"
)

// apiCmd is the exported http RESTFul server command line
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start a RESTFul server for api service",
	Run:   runAuth,
}

func runAuth(cmd *cobra.Command, args []string) {
	servers := []svc.Service{
		&api.Server{},
	}

	runServices(servers)
}

func init() {
	serveCmd.AddCommand(apiCmd)
}
