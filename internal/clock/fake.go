package clock

import (
	"sync"
	"time"
)

// Fake 可控时钟。
type Fake struct {
	mu   sync.Mutex
	now  time.Time
	wait time.Duration
}

// NewFake 构造。
func NewFake(start time.Time) *Fake {
	return &Fake{now: start}
}

func (f *Fake) Now() time.Time {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.now
}

func (f *Fake) Since(t time.Time) time.Duration {
	return f.Now().Sub(t)
}

func (f *Fake) Sleep(d time.Duration) {
	f.mu.Lock()
	f.now = f.now.Add(d)
	f.wait += d
	f.mu.Unlock()
}

// Advance 推进。
func (f *Fake) Advance(d time.Duration) {
	f.Sleep(d)
}

// Waited 累计 Sleep。
func (f *Fake) Waited() time.Duration {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.wait
}
