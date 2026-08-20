package wirebridge

import (
	"time"

	"example.com/wirebridge/internal/clock"
	"example.com/wirebridge/internal/persist"
)

// Option 构造选项。
type Option func(*Bridge)

// WithMaxFrame 设置最大帧长（含 header）。
func WithMaxFrame(n int) Option {
	return func(b *Bridge) {
		if n > 0 {
			b.maxFrame = n
		}
	}
}

// WithPersistPath 设置路由快照路径。
func WithPersistPath(path string) Option {
	return func(b *Bridge) {
		b.persistPath = path
	}
}

// WithClock 注入时钟（测试用）。
func WithClock(c clock.Clock) Option {
	return func(b *Bridge) {
		if c != nil {
			b.clk = c
		}
	}
}

// WithBypassLog 设置旁路试帧日志路径。
func WithBypassLog(path string) Option {
	return func(b *Bridge) {
		b.bypassLog = path
	}
}

// WithDefaultTimeout Handler 默认等待上限（供 Wait 类 Handler 使用）。
func WithDefaultTimeout(d time.Duration) Option {
	return func(b *Bridge) {
		if d > 0 {
			b.defaultTimeout = d
		}
	}
}

// WithName 设置 Bridge 显示名。
func WithName(name string) Option {
	return func(b *Bridge) {
		b.name = name
	}
}

// WithMemoryPersist 使用内存持久化（测试可 SetFail）。
func WithMemoryPersist(store *persist.MemoryStore) Option {
	return func(b *Bridge) {
		if store == nil {
			return
		}
		b.persistPath = "memory://"
		b.persistFn = store.Save
		b.loadFn = store.Load
	}
}
