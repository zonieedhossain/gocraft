package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zonieedhossain/gocraft/internal/generator"
)

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

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Go microservice",
	Long: `Generate a production-ready Go microservice with Clean Architecture.
Supported Frameworks: fiber, echo, gin
Supported ORMs: bun, gorm, sqlc
Supported DBs: postgres, mysql, sqlite`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		modulePath := getModulePath(projectName)
		if modulePath == "" {
			fmt.Println("Error: module path is required.")
			fmt.Println("Use one of: --github, --gitlab, --bitbucket, or --module-path")
			os.Exit(1)
		}

		validateFlags()

		opts := generator.Options{
			ProjectName: projectName,
			Web:         strings.ToLower(webFramework),
			DB:          strings.ToLower(dbType),
			ORM:         strings.ToLower(ormType),
			Auth:        useAuth,
			Docker:      useDocker,
			ModulePath:  modulePath,
		}

		fmt.Printf("Scaffolding project '%s'...\n", projectName)
		if err := generator.Generate(opts); err != nil {
			fmt.Printf("Failed to generate project: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\nProject scaffolded successfully!")
		fmt.Printf("cd %s\n", projectName)
		fmt.Println("go run cmd/server/main.go")
	},
}

func getModulePath(projectName string) string {
	if customPath != "" {
		return customPath
	}
	if githubUser != "" {
		return fmt.Sprintf("github.com/%s/%s", githubUser, projectName)
	}
	if gitlabUser != "" {
		return fmt.Sprintf("gitlab.com/%s/%s", gitlabUser, projectName)
	}
	if bitbucketUser != "" {
		return fmt.Sprintf("bitbucket.org/%s/%s", bitbucketUser, projectName)
	}
	return ""
}

func validateFlags() {
	validWeb := map[string]bool{"fiber": true, "echo": true, "gin": true}
	if !validWeb[strings.ToLower(webFramework)] {
		fmt.Printf("Invalid web framework: %s. Supported: fiber, echo, gin\n", webFramework)
		os.Exit(1)
	}

	validDB := map[string]bool{"postgres": true, "mysql": true, "sqlite": true}
	if !validDB[strings.ToLower(dbType)] {
		fmt.Printf("Invalid database: %s. Supported: postgres, mysql, sqlite\n", dbType)
		os.Exit(1)
	}

	validORM := map[string]bool{"bun": true, "gorm": true, "sqlc": true}
	if !validORM[strings.ToLower(ormType)] {
		fmt.Printf("Invalid ORM: %s. Supported: bun, gorm, sqlc\n", ormType)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVar(&webFramework, "web", "fiber", "Web framework (fiber, echo, gin)")
	newCmd.Flags().StringVar(&dbType, "db", "postgres", "Database type (postgres, mysql, sqlite)")
	newCmd.Flags().StringVar(&ormType, "orm", "bun", "ORM (bun, gorm, sqlc)")
	newCmd.Flags().BoolVar(&useAuth, "auth", false, "Include JWT authentication")
	newCmd.Flags().BoolVar(&useDocker, "docker", false, "Include Docker & Docker Compose")

	newCmd.Flags().StringVar(&githubUser, "github", "", "GitHub username for module path")
	newCmd.Flags().StringVar(&gitlabUser, "gitlab", "", "GitLab username for module path")
	newCmd.Flags().StringVar(&bitbucketUser, "bitbucket", "", "Bitbucket username for module path")
	newCmd.Flags().StringVar(&customPath, "module-path", "", "Custom full module path")
}
