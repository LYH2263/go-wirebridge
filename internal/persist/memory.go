package persist

import (
	"errors"
	"sync"

	"example.com/wirebridge/internal/route"
)

// MemoryStore 内存持久化（测试用，可强制失败）。
type MemoryStore struct {
	mu   sync.Mutex
	rows []route.Meta
	fail bool
}

// NewMemoryStore 构造。
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

// SetFail 下次 Save 失败。
func (m *MemoryStore) SetFail(v bool) {
	m.mu.Lock()
	m.fail = v
	m.mu.Unlock()
}

// Save 实现。
func (m *MemoryStore) Save(_ string, rows []route.Meta) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.fail {
		return errors.New("persist: forced failure")
	}
	m.rows = make([]route.Meta, 0, len(rows))
	for _, r := range rows {
		m.rows = append(m.rows, route.CloneMeta(r))
	}
	return nil
}

// Load 实现。
func (m *MemoryStore) Load(_ string) ([]route.Meta, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]route.Meta, 0, len(m.rows))
	for _, r := range m.rows {
		out = append(out, route.CloneMeta(r))
	}
	return out, nil
}

// Snapshot 当前行。
func (m *MemoryStore) Snapshot() []route.Meta {
	rows, _ := m.Load("")
	return rows
}
