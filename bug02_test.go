package wirebridge_test

import (
	"testing"

	"example.com/wirebridge"
	"example.com/wirebridge/internal/handler"
)

func TestBug02_ListRoutesTagsShared(t *testing.T) {
	b := wirebridge.New()
	defer b.Close()
	if err := b.Register(7, "admin", handler.Echo(), "prod", "v1"); err != nil {
		t.Fatal(err)
	}
	rows, err := b.ListRoutes()
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("want 1 route, got %d", len(rows))
	}
	rows[0].Tags[0] = "mutated"
	got, ok := b.GetRoute(7)
	if !ok {
		t.Fatal("missing route")
	}
	if got.Tags[0] == "mutated" {
		t.Fatal("ListRoutes Tags aliases internal route metadata")
	}
	if got.Tags[0] != "prod" {
		t.Fatalf("tags corrupted: %v", got.Tags)
	}
}
