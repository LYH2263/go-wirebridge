package buffer

import "sync"

// Pool 字节切片池。
type Pool struct {
	sz   int
	pool sync.Pool
}

// NewPool 构造固定容量池。
func NewPool(size int) *Pool {
	if size < 16 {
		size = 16
	}
	return &Pool{
		sz: size,
		pool: sync.Pool{
			New: func() any {
				b := make([]byte, size)
				return &b
			},
		},
	}
}

// Get 取一块。
func (p *Pool) Get() []byte {
	b := p.pool.Get().(*[]byte)
	return (*b)[:p.sz]
}

// Put 归还。
func (p *Pool) Put(b []byte) {
	if cap(b) < p.sz {
		return
	}
	bb := b[:p.sz]
	p.pool.Put(&bb)
}
