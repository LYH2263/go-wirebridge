package bypass

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBug09_BypassLogFileClosed(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bypass.jsonl")
	if err := Append(path, Entry{Opcode: 2, Flags: 8, Payload: []byte("x")}); err != nil {
		t.Fatal(err)
	}
	// Windows：句柄未关则 Remove 失败
	if err := os.Remove(path); err != nil {
		t.Fatalf("bypass log handle still held: %v", err)
	}
}
