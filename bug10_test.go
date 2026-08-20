package wirebridge_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"example.com/wirebridge"
	"example.com/wirebridge/internal/handler"
)

func TestBug10_CloseSyncsBeforeDropRoutes(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "routes.json")
	b := wirebridge.New(wirebridge.WithPersistPath(path))
	if err := b.Register(2, "echo", handler.Echo(), "keep"); err != nil {
		t.Fatal(err)
	}
	if err := b.Close(); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Routes []struct {
			Name string `json:"Name"`
		} `json:"routes"`
	}
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	if len(doc.Routes) == 0 {
		t.Fatal("Close synced empty routes; dropped table before Sync")
	}
	if doc.Routes[0].Name != "echo" {
		t.Fatalf("unexpected persist: %+v", doc.Routes)
	}
}
