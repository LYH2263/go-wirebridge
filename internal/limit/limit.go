package limit

import "sync"

// Limiter 帧长限制。
type Limiter struct {
	mu  sync.RWMutex
	max int
}

// New 构造。
func New(max int) *Limiter {
	if max <= 0 {
		max = 1 << 20
	}
	return &Limiter{max: max}
}

// SetMax 更新上限。
func (l *Limiter) SetMax(max int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if max > 0 {
		l.max = max
	}
}

// Max 当前上限。
func (l *Limiter) Max() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.max
}
