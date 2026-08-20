package metrics

import "sync/atomic"

// Counter 原子计数。
type Counter struct{ v uint64 }

func (c *Counter) Add(n uint64) { atomic.AddUint64(&c.v, n) }
func (c *Counter) Inc()         { c.Add(1) }
func (c *Counter) Get() uint64  { return atomic.LoadUint64(&c.v) }
func (c *Counter) Reset()       { atomic.StoreUint64(&c.v, 0) }
