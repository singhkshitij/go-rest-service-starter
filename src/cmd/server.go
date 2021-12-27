package cmd

import (
	"log"

	"github.com/singhkshitij/golang-rest-service-starter/src/api"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "server",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := api.NewServer()
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	},
}
