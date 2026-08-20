package wirebridge_test

import (
	"testing"

	"example.com/wirebridge"
	"example.com/wirebridge/internal/handler"
	"example.com/wirebridge/internal/persist"
)

func TestBug06_PersistFailureDoesNotApply(t *testing.T) {
	store := persist.NewMemoryStore()
	b := wirebridge.New(wirebridge.WithMemoryPersist(store))
	defer b.Close()
	if err := b.Register(2, "echo", handler.Echo(), "old"); err != nil {
		t.Fatal(err)
	}
	before, _ := b.ListRoutes()
	store.SetFail(true)
	err := b.ApplyRoutes([]wirebridge.RouteMeta{{
		Opcode: 3, Name: "stats", Enabled: true, Tags: []string{"new"},
	}})
	if err == nil {
		t.Fatal("expected persist failure")
	}
	after, _ := b.ListRoutes()
	if len(after) != len(before) || after[0].Name != "echo" {
		t.Fatalf("routes half-applied on persist fail: %+v", after)
	}
}
