package wirebridge_test

import (
	"os"
	"path/filepath"
	"testing"

	"example.com/wirebridge"
	"example.com/wirebridge/internal/handler"
)

func TestBug09_BypassLogFileClosed(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bypass.jsonl")
	b := wirebridge.New(wirebridge.WithBypassLog(path))
	defer b.Close()
	if err := b.Register(1, "echo", handler.Echo()); err != nil {
		t.Fatal(err)
	}
	raw, err := b.Encode(wirebridge.Frame{
		Flags:   wirebridge.FlagBypass,
		Opcode:  1,
		Payload: []byte("x"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := b.ServeFrame(raw); err != nil {
		t.Fatal(err)
	}
	// Windows：句柄未关则 Remove 失败
	if err := os.Remove(path); err != nil {
		t.Fatalf("bypass log handle still held: %v", err)
	}
}
