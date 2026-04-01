/*
Copyright © 2025 Zonieed_Hossain

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocraft",
	Short: "gocraft is a world-class CLI for scaffolding Go microservices",
	Long: `gocraft is a production-grade tool designed to scaffold 
Go microservices following Clean Architecture principles. It supports multiple 
web frameworks, ORMs, and databases out of the box.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Root flags if any can be added here
}
