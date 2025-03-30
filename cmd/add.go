/*
Copyright © 2025 Zonieed Hossain
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zonieedhossain/gocraft/internal/generator"
	"os"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add module [name]",
	Short: "Add a new CRUD module to your project",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "module" {
			fmt.Println("❌ Invalid usage. Try: gocraft add module user")
			os.Exit(1)
		}
		moduleName := strings.ToLower(args[1])
		if webFramework == "" {
			fmt.Println("❌ Please provide the web framework with --web (fiber, echo, or gin)")
			os.Exit(1)
		}

		err := generator.GenerateModule(moduleName, webFramework)
		if err != nil {
			fmt.Println("❌ Failed to generate module:", err)
		} else {
			fmt.Println("✅ Module", moduleName, "added successfully!")
		}
	},
}

var addModuleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Add a new CRUD module to your project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]

		if webFramework == "" {
			fmt.Println("❌ Please provide the web framework with --web (fiber, echo, or gin)")
			os.Exit(1)
		}

		err := generator.GenerateModule(moduleName, webFramework)
		if err != nil {
			fmt.Println("❌ Failed to generate module:", err)
		} else {
			fmt.Println("✅ Module", moduleName, "added successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addModuleCmd)

	addModuleCmd.Flags().StringVar(&webFramework, "web", "", "Web framework used (fiber, echo, gin)")
}
