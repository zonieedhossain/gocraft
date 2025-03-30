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
	tmpl, err := template.ParseFiles("internal/templates/main.go.tmpl")
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(base, "cmd", "main.go"))
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Render the template
	err = tmpl.Execute(file, opts)
	if err != nil {
		return err
	}

	return nil
}
