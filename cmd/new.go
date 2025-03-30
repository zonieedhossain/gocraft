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

// CLI Flags
var (
	webFramework  string
	dbType        string
	ormType       string
	useAuth       bool
	useDocker     bool
	githubUser    string
	gitlabUser    string
	bitbucketUser string
	customPath    string
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Go project with selected stack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Validate VCS/module-path flags
		if githubUser == "" && gitlabUser == "" && bitbucketUser == "" && customPath == "" {
			fmt.Println("❌ Please provide a module path using one of the following flags:")
			fmt.Println("   --github, --gitlab, --bitbucket, or --module-path")
			os.Exit(1)
		}

		// Construct module path
		var modulePath string
		switch {
		case customPath != "":
			modulePath = customPath
		case githubUser != "":
			modulePath = fmt.Sprintf("github.com/%s/%s", githubUser, projectName)
		case gitlabUser != "":
			modulePath = fmt.Sprintf("gitlab.com/%s/%s", gitlabUser, projectName)
		case bitbucketUser != "":
			modulePath = fmt.Sprintf("bitbucket.org/%s/%s", bitbucketUser, projectName)
		}

		// Build generator options
		opts := generator.Options{
			ProjectName: projectName,
			Web:         webFramework,
			DB:          dbType,
			ORM:         ormType,
			Auth:        useAuth,
			Docker:      useDocker,
			ModulePath:  modulePath,
		}

		// Run generator
		if err := generator.Generate(opts); err != nil {
			fmt.Println("❌ Failed to generate project:", err)
		} else {
			fmt.Println("✅ Project scaffolded successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Core stack options
	newCmd.Flags().StringVar(&webFramework, "web", "fiber", "Web framework (fiber, echo, gin)")
	newCmd.Flags().StringVar(&dbType, "db", "postgres", "Database type (postgres, mysql, sqlite)")
	newCmd.Flags().StringVar(&ormType, "orm", "bun", "ORM to use (bun, gorm, sqlc)")
	newCmd.Flags().BoolVar(&useAuth, "auth", false, "Include auth module")
	newCmd.Flags().BoolVar(&useDocker, "docker", false, "Include Docker support")

	// Module path
	newCmd.Flags().StringVar(&githubUser, "github", "", "GitHub username")
	newCmd.Flags().StringVar(&gitlabUser, "gitlab", "", "GitLab username")
	newCmd.Flags().StringVar(&bitbucketUser, "bitbucket", "", "Bitbucket username")
	newCmd.Flags().StringVar(&customPath, "module-path", "", "Custom full module path (e.g. vcs.mycorp.dev/team/app)")
}
