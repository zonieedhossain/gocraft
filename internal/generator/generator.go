package generator

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"runtime"
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
		mainTemplate = "main_fiber.go.tmpl"
	case "echo":
		mainTemplate = "main_echo.go.tmpl"
	case "gin":
		mainTemplate = "main_gin.go.tmpl"
	default:
		return fmt.Errorf("unsupported web framework: %s", opts.Web)
	}

	// Load template
	_ = renderTemplate("main.go.tmpl", filepath.Join(base, "main.go"), opts)
	_ = renderTemplate("env.tmpl", filepath.Join(base, ".env"), opts)
	_ = renderTemplate("Dockerfile.tmpl", filepath.Join(base, "Dockerfile"), opts)
	_ = renderTemplate("go.mod.tmpl", filepath.Join(base, "go.mod"), opts)
	_ = renderTemplate("README.md.tmpl", filepath.Join(base, "README.md"), opts)
	_ = renderTemplate("Makefile.tmpl", filepath.Join(base, "Makefile"), opts)
	_ = renderTemplate(mainTemplate, filepath.Join(base, "cmd", "main.go"), opts)
	_ = renderTemplate("routes.go.tmpl", filepath.Join(base, "internal", "routes", "routes.go"), opts)
	_ = renderTemplate("hello.go.tmpl", filepath.Join(base, "internal", "handlers", "hello.go"), opts)

	return nil
}

func renderTemplate(templateName, outputPath string, data any) error {
	// Resolve the current file location (generator.go)
	_, currentFile, _, _ := runtime.Caller(0)

	// Base: /path/to/gocraft/internal/generator
	basePath := filepath.Join(filepath.Dir(currentFile), "..", "templates")
	fullTemplatePath := filepath.Join(basePath, templateName)

	// 👇 Debug print
	fmt.Println("🛠 Template path:", fullTemplatePath)

	tmpl, err := template.ParseFiles(fullTemplatePath)
	if err != nil {
		fmt.Println("❌ Parse error:", err)
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
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

	// ✅ Get module path from go.mod
	modulePath, err := getModulePathFromGoMod()
	if err != nil {
		return fmt.Errorf("failed to detect module path: %w", err)
	}

	data := struct {
		Name       string
		NameTitle  string
		Web        string
		ModulePath string
	}{
		Name:       name,
		NameTitle:  capitalize(name),
		Web:        web,
		ModulePath: modulePath,
	}

	base := "."

	// Generate handler
	err = renderTemplate("module_handler.go.tmpl",
		filepath.Join(base, "internal", "handlers", name+".go"), data)
	if err != nil {
		return err
	}

	// Generate routes
	err = renderTemplate("module_routes.go.tmpl",
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

func getModulePathFromGoMod() (string, error) {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", fmt.Errorf("module path not found in go.mod")
}
