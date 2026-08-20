package handler

import (
	"context"
	"time"
)

// Wait 在 Handler 内可中断等待；必须响应 ctx 取消。
func Wait(ctx context.Context, d time.Duration) error {
	// BUG: 不听 ctx，固定 Sleep
	if d <= 0 {
		return nil
	}
	time.Sleep(d)
	return nil
}

// WaitPoll 轮询直到 ready 或 ctx 取消。
func WaitPoll(ctx context.Context, interval time.Duration, ready func() bool) error {
	if interval <= 0 {
		interval = 10 * time.Millisecond
	}
	if ready != nil && ready() {
		return nil
	}
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			if ready != nil && ready() {
				return nil
			}
		}
	}
}
