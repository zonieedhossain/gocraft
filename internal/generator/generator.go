package generator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/zonieedhossain/gocraft/internal/templates"
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
	// 1. Create directory structure
	dirs := []string{
		"cmd/server",
		"internal/domain",
		"internal/usecase",
		"internal/repository",
		"internal/delivery/http",
		"internal/infrastructure",
		"internal/middleware",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(opts.ProjectName, dir), 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// 2. Define files to generate
	files := []struct {
		tmplPath string
		outPath  string
	}{
		{"project/common/env.tmpl", ".env"},
		{"project/common/Makefile.tmpl", "Makefile"},
		{"project/common/README.md.tmpl", "README.md"},
		{"project/common/go.mod.tmpl", "go.mod"},
		{"project/cmd/server/main.go.tmpl", "cmd/server/main.go"},
		{"project/internal/domain/entity.tmpl", "internal/domain/user.go"},
		{"project/internal/usecase/usecase.tmpl", "internal/usecase/user_usecase.go"},
		{"project/internal/infrastructure/db.tmpl", "internal/infrastructure/db.go"},
		{"project/internal/repository/repository.tmpl", "internal/repository/user_repository.go"},
		{"project/internal/delivery/http/handler.tmpl", "internal/delivery/http/handler.go"},
	}

	if opts.ORM == "sqlc" {
		files = append(files, struct {
			tmplPath string
			outPath  string
		}{"project/common/sqlc.yaml.tmpl", "sqlc.yaml"})
		files = append(files, struct {
			tmplPath string
			outPath  string
		}{"project/internal/repository/schema.sql.tmpl", "internal/repository/schema.sql"})
		files = append(files, struct {
			tmplPath string
			outPath  string
		}{"project/internal/repository/query.sql.tmpl", "internal/repository/query.sql"})
		files = append(files, struct {
			tmplPath string
			outPath  string
		}{"project/internal/repository/sqlc_generated.tmpl", "internal/repository/sqlc_generated.go"})
	}

	if opts.Docker {
		files = append(files, struct {
			tmplPath string
			outPath  string
		}{"project/common/Dockerfile.tmpl", "Dockerfile"})
		files = append(files, struct {
			tmplPath string
			outPath  string
		}{"project/common/docker-compose.tmpl", "docker-compose.yml"})
	}

	if opts.Auth {
		files = append(files, struct {
			tmplPath string
			outPath  string
		}{"project/internal/middleware/auth.tmpl", "internal/middleware/auth.go"})
	}

	// 3. Render and write files
	for _, f := range files {
		if err := renderAndWrite(f.tmplPath, filepath.Join(opts.ProjectName, f.outPath), opts); err != nil {
			return err
		}
	}

	// 4. Run go mod tidy
	fmt.Println("Running 'go mod tidy' in the new project...")
	if err := runCommand(opts.ProjectName, "go", "mod", "tidy"); err != nil {
		fmt.Printf("Warning: 'go mod tidy' failed: %v. You may need to run it manually.\n", err)
	}

	return nil
}

func renderAndWrite(tmplPath, outPath string, opts Options) error {
	tmplData, err := templates.ProjectTemplates.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", tmplPath, err)
	}

	tmpl, err := template.New(tmplPath).Parse(string(tmplData))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", tmplPath, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, opts); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", tmplPath, err)
	}

	if err := os.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outPath, err)
	}

	return nil
}

func runCommand(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
