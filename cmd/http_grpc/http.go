/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package http_grpc

import (
	"github.com/abdukhashimov/go_api/internal/config"
	"github.com/abdukhashimov/go_api/internal/pkg/logger"
	"github.com/abdukhashimov/go_api/internal/transport/handlers"
	"github.com/abdukhashimov/go_api/pkg/logger/factory"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(http_grpc *cobra.Command, args []string) {
		portNum, _ := http_grpc.Flags().GetString("port")

		cfg := config.Load()

		log, err := factory.Build(&cfg.Logging)
		if err != nil {
			panic(err)
		}

		logger.SetLogger(log)

		server := handlers.NewServer(cfg)
		server.Run(portNum)
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.PersistentFlags().String("port", "8080", "A port number that is used to run http server")
}
