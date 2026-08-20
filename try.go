package wirebridge

import (
	"context"
	"time"
)

// TryServe 带超时的 Serve。
func (b *Bridge) TryServe(raw []byte, timeout time.Duration) (Frame, error) {
	if timeout <= 0 {
		timeout = b.defaultTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// 仍走 ServeFrameContext（plant 忽略取消）
	return b.ServeFrameContext(ctx, raw)
}
