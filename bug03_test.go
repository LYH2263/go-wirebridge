package wirebridge_test

import (
	"encoding/binary"
	"errors"
	"testing"

	"example.com/wirebridge"
	"example.com/wirebridge/internal/handler"
)

func TestBug03_ServeAfterCloseNoPanic(t *testing.T) {
	b := wirebridge.New()
	_ = b.Register(2, "echo", handler.Echo())
	_ = b.Close()
	raw := make([]byte, 7)
	binary.BigEndian.PutUint32(raw[0:4], 3)
	raw[4] = 0
	binary.BigEndian.PutUint16(raw[5:7], 2)
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("ServeFrame after Close panicked: %v", r)
		}
	}()
	_, err := b.ServeFrame(raw)
	if !errors.Is(err, wirebridge.ErrClosed) {
		t.Fatalf("want ErrClosed, got %v", err)
	}
}
