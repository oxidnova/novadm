package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/oxidnova/novadm/backend/svc"
	"github.com/spf13/cobra"
)

// is the exported server command line
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a specific server for api service",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var configPath string

func init() {
	serveCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "serve.yaml", "the path of yaml file, the config are loaded by environment variable.")

	signals = make(chan os.Signal)

	rootCmd.AddCommand(serveCmd)
}

func runServices(servers []svc.Service) {
	// 1. load config from file and env
	for _, server := range servers {
		if err := server.Load(configPath); err != nil {
			log.Fatalf("load %s server, err: %v", server.Name(), err)
		}
	}

	var stopFns []func(os.Signal)
	for _, server := range servers {
		stopFns = append(stopFns, server.Stop)
		go server.Run()
	}

	var stop = func(s os.Signal) {
		for _, fn := range stopFns {
			fn(s)
		}
	}

	embedNotifySignal(stop)
}

var signals chan os.Signal

func embedNotifySignal(fn func(os.Signal)) {
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sign := <-signals
	fn(sign)
}
