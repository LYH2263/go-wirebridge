package route

import (
	"fmt"
	"sort"
	"sync"
)

// Table opcode → Meta。
type Table struct {
	mu   sync.RWMutex
	byOp map[uint16]Meta
}

// NewTable 空表。
func NewTable() *Table {
	return &Table{byOp: make(map[uint16]Meta)}
}

// Upsert 写入拷贝。
func (t *Table) Upsert(m Meta) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.byOp == nil {
		t.byOp = make(map[uint16]Meta)
	}
	if m.Name == "" {
		return fmt.Errorf("route: empty name for opcode %d", m.Opcode)
	}
	t.byOp[m.Opcode] = CloneMeta(m)
	return nil
}

// Remove 删除。
func (t *Table) Remove(op uint16) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, ok := t.byOp[op]; !ok {
		return false
	}
	delete(t.byOp, op)
	return true
}

// Get 返回拷贝。
func (t *Table) Get(op uint16) (Meta, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	m, ok := t.byOp[op]
	if !ok {
		return Meta{}, false
	}
	return CloneMeta(m), true
}

// List 返回全部拷贝，按 priority 降序。
func (t *Table) List() []Meta {
	t.mu.RLock()
	defer t.mu.RUnlock()
	out := make([]Meta, 0, len(t.byOp))
	for _, m := range t.byOp {
		out = append(out, CloneMeta(m))
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Priority == out[j].Priority {
			return out[i].Opcode < out[j].Opcode
		}
		return out[i].Priority > out[j].Priority
	})
	return out
}

// Replace 整体替换。
func (t *Table) Replace(rows []Meta) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.byOp = make(map[uint16]Meta, len(rows))
	for _, m := range rows {
		t.byOp[m.Opcode] = CloneMeta(m)
	}
}

// Len 条数。
func (t *Table) Len() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.byOp)
}
