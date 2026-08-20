package handler

import "sync"

// Registry opcode → Handler。
type Registry struct {
	mu sync.RWMutex
	m  map[uint16]Handler
}

// NewRegistry 构造。
func NewRegistry() *Registry {
	return &Registry{m: make(map[uint16]Handler)}
}

// Set 注册；h 为 nil 时删除。
func (r *Registry) Set(op uint16, h Handler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.m == nil {
		r.m = make(map[uint16]Handler)
	}
	if h == nil {
		delete(r.m, op)
		return
	}
	r.m[op] = h
}

// Get 查找。
func (r *Registry) Get(op uint16) Handler {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.m[op]
}

// Delete 删除。
func (r *Registry) Delete(op uint16) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.m, op)
}

// Clear 清空。
func (r *Registry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.m = make(map[uint16]Handler)
}

// Len 数量。
func (r *Registry) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.m)
}
