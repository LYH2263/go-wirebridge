package frame

import "sync"

// BufferPool 可复用编码缓冲。
type BufferPool struct {
	pool sync.Pool
}

// NewBufferPool 构造。
func NewBufferPool(defaultSize int) *BufferPool {
	if defaultSize < 64 {
		defaultSize = 64
	}
	return &BufferPool{
		pool: sync.Pool{
			New: func() any {
				b := make([]byte, 0, defaultSize)
				return &b
			},
		},
	}
}

// Get 取缓冲。
func (p *BufferPool) Get() *[]byte {
	return p.pool.Get().(*[]byte)
}

// Put 归还。
func (p *BufferPool) Put(b *[]byte) {
	if b == nil {
		return
	}
	*b = (*b)[:0]
	p.pool.Put(b)
}
