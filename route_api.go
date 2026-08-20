package wirebridge

import (
	"example.com/wirebridge/internal/route"
)

// ListRoutes 返回路由元数据拷贝（含 Tags 深拷贝）。
func (b *Bridge) ListRoutes() ([]RouteMeta, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return nil, ErrClosed
	}
	rows := b.router.List()
	out := make([]RouteMeta, 0, len(rows))
	for _, m := range rows {
		out = append(out, toPublic(m))
	}
	return out, nil
}

// GetRoute 按 opcode 取元数据拷贝。
func (b *Bridge) GetRoute(op Opcode) (RouteMeta, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed || b.router == nil {
		return RouteMeta{}, false
	}
	m, ok := b.router.Get(uint16(op))
	if !ok {
		return RouteMeta{}, false
	}
	return toPublic(m), true
}

// ApplyRoutes 校验并整体替换路由元数据；成功后持久化。
// 持久化失败时不得保留新表（须回滚）。
func (b *Bridge) ApplyRoutes(rows []RouteMeta) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return ErrClosed
	}
	internal := make([]route.Meta, 0, len(rows))
	for _, r := range rows {
		if err := validateMeta(r); err != nil {
			return err
		}
		internal = append(internal, toInternal(r))
	}
	old := b.router.List()
	b.router.Replace(internal)
	if b.persistPath != "" {
		if err := b.persistFn(b.persistPath, internal); err != nil {
			b.router.Replace(old)
			return err
		}
	}
	b.dirty = false
	return nil
}

func toPublic(m route.Meta) RouteMeta {
	return RouteMeta{
		Opcode:      Opcode(m.Opcode),
		Name:        m.Name,
		Description: m.Description,
		// BUG: Tags 与库存共享
		Tags:     m.Tags,
		Enabled:  m.Enabled,
		Priority: m.Priority,
	}
}

func toInternal(r RouteMeta) route.Meta {
	return route.Meta{
		Opcode:      uint16(r.Opcode),
		Name:        r.Name,
		Description: r.Description,
		Tags:        route.CloneStrings(r.Tags),
		Enabled:     r.Enabled,
		Priority:    r.Priority,
	}
}
