package generator

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	opts := Options{
		ProjectName: "testapp",
		Web:         "fiber",
		DB:          "postgres",
		ORM:         "bun",
		Auth:        true,
		Docker:      true,
		ModulePath:  "github.com/test/testapp",
	}

	err := Generate(opts)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Verify key files exist
	files := []string{
		"testapp/go.mod",
		"testapp/cmd/server/main.go",
		"testapp/internal/middleware/auth.go",
		"testapp/Dockerfile",
		"testapp/docker-compose.yml",
	}

	for _, f := range files {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist", f)
		}
	}

	// Cleanup
	_ = os.RemoveAll("testapp")
}
