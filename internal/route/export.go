package route

import "sort"

// ExportMap 导出为 opcode→name。
func (t *Table) ExportMap() map[uint16]string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	out := make(map[uint16]string, len(t.byOp))
	for op, m := range t.byOp {
		out[op] = m.Name
	}
	return out
}

// Opcodes 排序后的 opcode 列表。
func (t *Table) Opcodes() []uint16 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	out := make([]uint16, 0, len(t.byOp))
	for op := range t.byOp {
		out = append(out, op)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}
