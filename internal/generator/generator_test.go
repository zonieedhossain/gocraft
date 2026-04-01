package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		opts Options
	}{
		{
			name: "Fiber-GORM-Postgres",
			opts: Options{ProjectName: "fiber_gorm", Web: "fiber", DB: "postgres", ORM: "gorm", Auth: true, Docker: true, ModulePath: "github.com/test/fiber_gorm"},
		},
		{
			name: "Echo-Bun-MySQL",
			opts: Options{ProjectName: "echo_bun", Web: "echo", DB: "mysql", ORM: "bun", Auth: true, Docker: true, ModulePath: "github.com/test/echo_bun"},
		},
		{
			name: "Gin-Sqlc-SQLite",
			opts: Options{ProjectName: "gin_sqlc", Web: "gin", DB: "sqlite", ORM: "sqlc", Auth: true, Docker: true, ModulePath: "github.com/test/gin_sqlc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Generate(tt.opts)
			if err != nil {
				t.Fatalf("Generate failed: %v", err)
			}

			// Core files
			checkFile(t, tt.opts.ProjectName, "go.mod")
			checkFile(t, tt.opts.ProjectName, "cmd/server/main.go")
			
			// Optional files
			if tt.opts.Auth {
				checkFile(t, tt.opts.ProjectName, "internal/middleware/auth.go")
			}
			if tt.opts.Docker {
				checkFile(t, tt.opts.ProjectName, "Dockerfile")
				checkFile(t, tt.opts.ProjectName, "docker-compose.yml")
			}
			if tt.opts.ORM == "sqlc" {
				checkFile(t, tt.opts.ProjectName, "sqlc.yaml")
				checkFile(t, tt.opts.ProjectName, "internal/repository/schema.sql")
				checkFile(t, tt.opts.ProjectName, "internal/repository/sqlc_generated.go")
			}

			// Cleanup
			_ = os.RemoveAll(tt.opts.ProjectName)
		})
	}
}

func checkFile(t *testing.T, projectName, subPath string) {
	path := filepath.Join(projectName, subPath)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("expected file %s to exist", path)
	}
}
