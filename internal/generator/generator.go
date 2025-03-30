package generator

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
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

	var createdFiles []string
	create := func(templateName, outputPath string) error {
		path, err := renderTemplate(templateName, outputPath, opts)
		if err != nil {
			rollbackNewProject(base, createdFiles)
			return err
		}
		createdFiles = append(createdFiles, path)
		return nil
	}

	var mainTemplate string
	switch opts.Web {
	case "fiber":
		mainTemplate = "main_fiber.go.tmpl"
	case "echo":
		mainTemplate = "main_echo.go.tmpl"
	case "gin":
		mainTemplate = "main_gin.go.tmpl"
	case "graphql":
		mainTemplate = "main_graphql.go.tmpl"
	default:
		return fmt.Errorf("unsupported web framework: %s", opts.Web)
	}

	if err := create("main.go.tmpl", filepath.Join(base, "main.go")); err != nil {
		return err
	}
	if err := create("env.tmpl", filepath.Join(base, ".env")); err != nil {
		return err
	}
	if err := create("Dockerfile.tmpl", filepath.Join(base, "Dockerfile")); err != nil {
		return err
	}
	if err := create("go.mod.tmpl", filepath.Join(base, "go.mod")); err != nil {
		return err
	}
	if err := create("README.md.tmpl", filepath.Join(base, "README.md")); err != nil {
		return err
	}
	if err := create("Makefile.tmpl", filepath.Join(base, "Makefile")); err != nil {
		return err
	}
	if err := create(mainTemplate, filepath.Join(base, "cmd", "main.go")); err != nil {
		return err
	}
	if err := create("routes.go.tmpl", filepath.Join(base, "internal", "routes", "routes.go")); err != nil {
		return err
	}
	if err := create("hello.go.tmpl", filepath.Join(base, "internal", "handlers", "hello.go")); err != nil {
		return err
	}

	err := runGoModTidy(opts.ProjectName)
	if err != nil {
		fmt.Println("⚠️ go mod tidy failed:", err)
	}

	return nil
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

	var createdFiles []string
	create := func(templateName, outputPath string) error {
		path, err := renderTemplate(templateName, outputPath, data)
		if err != nil {
			rollbackFiles(createdFiles)
			return err
		}
		createdFiles = append(createdFiles, path)
		return nil
	}

	base := "."
	if err := create("module_handler.go.tmpl", filepath.Join(base, "internal", "handlers", name+".go")); err != nil {
		return err
	}
	if err := create("module_routes.go.tmpl", filepath.Join(base, "internal", "routes", name+"_routes.go")); err != nil {
		return err
	}

	if web == "graphql" {
		if err := create("module_graphql_schema.go.tmpl", filepath.Join(base, "internal", "handlers", name+"_schema.go")); err != nil {
			return err
		}
	}

	err = runGoModTidy(".")
	if err != nil {
		fmt.Println("⚠️ go mod tidy failed:", err)
	}

	return nil
}

func renderTemplate(templateName, outputPath string, data any) (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get caller")
	}

	templateDir := filepath.Join(filepath.Dir(currentFile), "..", "templates")
	fullTemplatePath := filepath.Join(templateDir, templateName)

	tmpl, err := template.ParseFiles(fullTemplatePath)
	if err != nil {
		return "", fmt.Errorf("parse error: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return "", fmt.Errorf("file create error: %w", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return "", fmt.Errorf("template exec error: %w", err)
	}

	return outputPath, nil
}

func runGoModTidy(path string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
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
	case strings.Contains(str, "github.com/graphql-go/graphql"):
		return "graphql", nil
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

func rollbackFiles(files []string) {
	fmt.Println("\n⚠️ Rolling back generated files...")
	for _, file := range files {
		_ = os.Remove(file)
		fmt.Println("🗑️ Deleted:", file)
	}
}

func rollbackNewProject(projectName string, files []string) {
	rollbackFiles(files)
	_ = os.RemoveAll(projectName)
	fmt.Println("🧹 Deleted folder:", projectName)
}
