package generator

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestScaffoldCompilation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	// Check if go is installed
	_, err := exec.LookPath("go")
	if err != nil {
		t.Skip("skipping integration test because go binary not found")
	}

	tests := []struct {
		name string
		opts Options
	}{
		{
			name: "Fiber-GORM-Postgres-Compilation",
			opts: Options{ProjectName: "comp_fiber", Web: "fiber", DB: "postgres", ORM: "gorm", Auth: true, Docker: true, ModulePath: "github.com/test/comp_fiber"},
		},
		{
			name: "Echo-Bun-MySQL-Compilation",
			opts: Options{ProjectName: "comp_echo", Web: "echo", DB: "mysql", ORM: "bun", Auth: true, Docker: true, ModulePath: "github.com/test/comp_echo"},
		},
		{
			name: "Gin-Sqlc-SQLite-Compilation",
			opts: Options{ProjectName: "comp_gin", Web: "gin", DB: "sqlite", ORM: "sqlc", Auth: true, Docker: true, ModulePath: "github.com/test/comp_gin"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate
			err := Generate(tt.opts)
			if err != nil {
				t.Fatalf("Generate failed: %v", err)
			}
			defer os.RemoveAll(tt.opts.ProjectName)

			// Tidy
			cmdTidy := exec.Command("go", "mod", "tidy")
			cmdTidy.Dir = tt.opts.ProjectName
			if out, err := cmdTidy.CombinedOutput(); err != nil {
				t.Fatalf("go mod tidy failed: %v\nOutput: %s", err, string(out))
			}

			// Build
			mainPath := filepath.Join("cmd", "server", "main.go")
			cmdBuild := exec.Command("go", "build", "-o", "server_bin", mainPath)
			cmdBuild.Dir = tt.opts.ProjectName
			if out, err := cmdBuild.CombinedOutput(); err != nil {
				t.Fatalf("go build failed: %v\nOutput: %s", err, string(out))
			}
		})
	}
}
