package handler

import (
	"context"
	"testing"
	"time"
)

func TestBug08_HandlerWaitHonorsContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	start := time.Now()
	err := Wait(ctx, 2*time.Second)
	if err == nil {
		t.Fatal("Wait ignored canceled ctx")
	}
	if time.Since(start) > 500*time.Millisecond {
		t.Fatalf("Wait blocked too long: %v", time.Since(start))
	}
}
