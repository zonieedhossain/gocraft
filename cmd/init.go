package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zonieedhossain/gocraft/internal/generator"
	"github.com/zonieedhossain/gocraft/internal/utils"
	"os"
	"path/filepath"
	"strings"
)

var moduleName string
var framework string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project with selected stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		if framework == "" {
			fmt.Print("Choose your web framework [fiber/echo/gin]: ")
			f, _ := reader.ReadString('\n')
			framework = strings.TrimSpace(f)
		}
		if framework != "fiber" && framework != "echo" && framework != "gin" {
			return fmt.Errorf("invalid framework: %s", framework)
		}

		if moduleName == "" {
			cwd, _ := os.Getwd()
			fmt.Printf("Enter module name (e.g., github.com/yourname/%s): ", filepath.Base(cwd))
			m, _ := reader.ReadString('\n')
			moduleName = strings.TrimSpace(m)
		}

		state := utils.GocraftState{
			ModuleName: moduleName,
			Framework:  framework,
		}
		if err := utils.SaveState(state); err != nil {
			return err
		}

		return generator.GenerateBaseProject(generator.ProjectConfig{
			ModuleName: moduleName,
			Framework:  framework,
			AppName:    filepath.Base(moduleName),
		})
	},
}

func init() {
	initCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Go module name (e.g., github.com/you/project)")
	initCmd.Flags().StringVarP(&framework, "framework", "f", "", "Web framework (fiber, echo, gin)")
}
