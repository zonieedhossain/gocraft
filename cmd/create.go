package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zonieedhossain/gocraft/internal/generator"
)

var (
	projectName string
	withFolders string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Generate a new project structure",
	RunE: func(cmd *cobra.Command, args []string) error {
		if projectName == "" {
			return fmt.Errorf("project name is required")
		}

		folders := strings.Split(withFolders, ",")
		return generator.GenerateFolders(projectName, folders)
	},
}

func init() {
	createCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project")
	createCmd.Flags().StringVarP(&withFolders, "with", "w", "", "Comma-separated folders to include")
	rootCmd.AddCommand(createCmd)
}
