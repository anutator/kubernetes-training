package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "probes_example_server",
	Short: "A simple client-server application for demonstrating Kubernetes probes",
}

var clientCmd = &cobra.Command{
	Use:   "http",
	Short: "Simple HTTP client for the probes-example-server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := client(args[0])
		if err != nil {
			fmt.Printf("Error running client: %v\n", err)
		}
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the probes-example-server",
	Run: func(cmd *cobra.Command, args []string) {
		err := server()
		if err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	},
}

func main() {
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.Execute()
}
