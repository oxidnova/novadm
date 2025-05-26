package cmd

import (
	"github.com/oxidnova/novadm/backend/svc"
	"github.com/oxidnova/novadm/backend/svc/api"
	"github.com/spf13/cobra"
)

// allCmd is the exported all services command line
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Start all servers for novadm service",
	Run:   runAll,
}

func runAll(cmd *cobra.Command, args []string) {
	servers := []svc.Service{
		&api.Server{},
	}

	runServices(servers)
}

func init() {
	serveCmd.AddCommand(allCmd)
}
