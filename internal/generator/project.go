package generator

import (
	"fmt"
	"github.com/zonieedhossain/gocraft/internal/utils"
)

type ProjectConfig struct {
	ModuleName string
	Framework  string
	AppName    string
}

func GenerateBaseProject(cfg ProjectConfig) error {
	fmt.Println("🚀 Generating base project structure...")

	files := []struct {
		Template string
		Output   string
	}{
		{"templates/go.mod.tpl", "go.mod"},
		{"templates/.env.tpl", ".env"},
		{"templates/Makefile.tpl", "Makefile"},
		{"templates/config/config.go.tpl", "internal/config/config.go"},
		{"templates/internal/server.go.tpl", "internal/server.go"},
		{"templates/router.tpl", "internal/routes/router.go"},
		{fmt.Sprintf("templates/%s.tpl", cfg.Framework), "cmd/main.go"},
	}

	for _, file := range files {
		err := utils.RenderTemplate(file.Template, file.Output, cfg)
		if err != nil {
			return fmt.Errorf("failed to render %s: %w", file.Template, err)
		}
		fmt.Printf("✅ %s generated\n", file.Output)
	}

	return nil
}
