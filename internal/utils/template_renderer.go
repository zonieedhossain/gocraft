package utils

import (
	"os"
	"path/filepath"
	"text/template"
)

// RenderTemplate renders a template file with provided data and writes to output path
func RenderTemplate(tplPath, outPath string, data any) error {
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return err
	}

	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return tpl.Execute(outFile, data)
}
