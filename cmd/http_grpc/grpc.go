/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package http_grpc

import (
	"fmt"

	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(http_grpc *cobra.Command, args []string) {
		port, _ := http_grpc.Flags().GetInt("port")

		if port != 0 {
			fmt.Println("Port number is provided")
		} else {
			fmt.Println("grpc called")
		}
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	grpcCmd.PersistentFlags().Int("port", 9090, "A port that the grpc server can be served")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
