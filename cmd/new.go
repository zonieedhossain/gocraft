/*
Copyright © 2025 Zonieed Hossain
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zonieedhossain/gocraft/internal/generator"
)

// CLI Flags
var (
	webFramework string
	dbType       string
	ormType      string
	useAuth      bool
	useDocker    bool
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Go project with selected stack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		opts := generator.Options{
			ProjectName: projectName,
			Web:         webFramework,
			DB:          dbType,
			ORM:         ormType,
			Auth:        useAuth,
			Docker:      useDocker,
		}

		err := generator.Generate(opts)
		if err != nil {
			fmt.Println("❌ Failed to generate project:", err)
		} else {
			fmt.Println("✅ Project scaffolded successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVar(&webFramework, "web", "fiber", "Web framework (fiber, echo, gin)")
	newCmd.Flags().StringVar(&dbType, "db", "postgres", "Database type (postgres, mysql, sqlite)")
	newCmd.Flags().StringVar(&ormType, "orm", "bun", "ORM to use (bun, gorm, sqlc)")
	newCmd.Flags().BoolVar(&useAuth, "auth", false, "Include auth module")
	newCmd.Flags().BoolVar(&useDocker, "docker", false, "Include Docker support")
}
