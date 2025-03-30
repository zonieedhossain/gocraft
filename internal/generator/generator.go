package generator

import (
	"html/template"
	"os"
	"path/filepath"
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

	// Load main.go template
	_ = renderTemplate("internal/templates/main.go.tmpl", filepath.Join(base, "main.go"), opts)
	_ = renderTemplate("internal/templates/env.tmpl", filepath.Join(base, ".env"), opts)
	_ = renderTemplate("internal/templates/Dockerfile.tmpl", filepath.Join(base, "Dockerfile"), opts)
	_ = renderTemplate("internal/templates/go.mod.tmpl", filepath.Join(base, "go.mod"), opts)
	_ = renderTemplate("internal/templates/README.md.tmpl", filepath.Join(base, "README.md"), opts)

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
