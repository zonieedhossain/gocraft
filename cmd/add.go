/*
Copyright © 2025 Zonieed Hossain
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zonieedhossain/gocraft/internal/generator"
)

// addModuleCmd handles: gocraft add module [name]
var addModuleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Add a new CRUD module to your project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]

		err := generator.GenerateModule(moduleName, webFramework)
		if err != nil {
			fmt.Println("❌ Failed to generate module:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Module", moduleName, "added successfully!")
	},
}

// addCmd is just the parent namespace: gocraft add ...
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add resources like modules to an existing project",
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addModuleCmd)

	addModuleCmd.Flags().StringVar(&webFramework, "web", "", "Web framework used (fiber, echo, gin)")
}
