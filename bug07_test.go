package wirebridge_test

import (
	"context"
	"encoding/binary"
	"testing"
	"time"

	"example.com/wirebridge"
	"example.com/wirebridge/internal/handler"
)

func TestBug07_ServeFrameContextHonorsCancel(t *testing.T) {
	b := wirebridge.New()
	defer b.Close()
	if err := b.Register(2, "slow", handler.DelayEcho(2*time.Second)); err != nil {
		t.Fatal(err)
	}
	raw := make([]byte, 7)
	binary.BigEndian.PutUint32(raw[0:4], 3)
	binary.BigEndian.PutUint16(raw[5:7], 2)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	start := time.Now()
	_, err := b.ServeFrameContext(ctx, raw)
	if err == nil {
		t.Fatal("ServeFrameContext ignored canceled context")
	}
	if time.Since(start) > 500*time.Millisecond {
		t.Fatalf("blocked too long ignoring cancel: %v", time.Since(start))
	}
}
