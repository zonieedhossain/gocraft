package generator

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	ProjectName string
	Web         string
	DB          string
	ORM         string
	Auth        bool
	Docker      bool
	ModulePath  string
}

func Generate(opts Options) error {
	base := opts.ProjectName
	dirs := []string{
		"cmd",
		"internal/config",
		"internal/handlers",
		"internal/routes",
		"internal/db",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(base, dir), 0755)
		if err != nil {
			return err
		}
	}

	var mainTemplate string
	switch opts.Web {
	case "fiber":
		mainTemplate = "internal/templates/main_fiber.go.tmpl"
	case "echo":
		mainTemplate = "internal/templates/main_echo.go.tmpl"
	case "gin":
		mainTemplate = "internal/templates/main_gin.go.tmpl"
	default:
		return fmt.Errorf("unsupported web framework: %s", opts.Web)
	}

	// Load template
	_ = renderTemplate("internal/templates/main.go.tmpl", filepath.Join(base, "main.go"), opts)
	_ = renderTemplate("internal/templates/env.tmpl", filepath.Join(base, ".env"), opts)
	_ = renderTemplate("internal/templates/Dockerfile.tmpl", filepath.Join(base, "Dockerfile"), opts)
	_ = renderTemplate("internal/templates/go.mod.tmpl", filepath.Join(base, "go.mod"), opts)
	_ = renderTemplate("internal/templates/README.md.tmpl", filepath.Join(base, "README.md"), opts)
	_ = renderTemplate("internal/templates/Makefile.tmpl", filepath.Join(base, "Makefile"), opts)
	_ = renderTemplate(mainTemplate, filepath.Join(base, "cmd", "main.go"), opts)
	_ = renderTemplate("internal/templates/routes.go.tmpl", filepath.Join(base, "internal", "routes", "routes.go"), opts)
	_ = renderTemplate("internal/templates/hello.go.tmpl", filepath.Join(base, "internal", "handlers", "hello.go"), opts)

	return nil
}

func renderTemplate(templatePath, outputPath string, data any) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	return tmpl.Execute(file, data)
}

func GenerateModule(name, web string) error {
	if web == "" {
		detected, err := detectWebFramework()
		if err != nil {
			return fmt.Errorf("could not detect web framework: %w", err)
		}
		web = detected
		fmt.Println("🔍 Detected web framework:", web)
	}

	data := struct {
		Name      string
		NameTitle string
		Web       string
	}{
		Name:      name,
		NameTitle: capitalize(name),
		Web:       web,
	}

	base := "."

	// Generate handler
	err := renderTemplate("internal/templates/module_handler.go.tmpl",
		filepath.Join(base, "internal", "handlers", name+".go"), data)
	if err != nil {
		return err
	}

	// Generate routes
	err = renderTemplate("internal/templates/module_routes.go.tmpl",
		filepath.Join(base, "internal", "routes", name+"_routes.go"), data)
	if err != nil {
		return err
	}

	return nil
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

func detectWebFramework() (string, error) {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	str := string(content)

	switch {
	case strings.Contains(str, "github.com/gofiber/fiber"):
		return "fiber", nil
	case strings.Contains(str, "github.com/labstack/echo"):
		return "echo", nil
	case strings.Contains(str, "github.com/gin-gonic/gin"):
		return "gin", nil
	default:
		return "", fmt.Errorf("no known framework found in go.mod")
	}
}
