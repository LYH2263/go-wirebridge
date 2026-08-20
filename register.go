package wirebridge

import (
	"example.com/wirebridge/internal/handler"
	"example.com/wirebridge/internal/route"
)

// Register 注册 opcode 对应 Handler，并写入路由元数据。
func (b *Bridge) Register(op Opcode, name string, h handler.Handler, tags ...string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return ErrClosed
	}
	if h == nil {
		return ErrNilHandler
	}
	meta := route.Meta{
		Opcode:  uint16(op),
		Name:    name,
		Tags:    append([]string(nil), tags...),
		Enabled: true,
	}
	if err := b.router.Upsert(meta); err != nil {
		return err
	}
	b.registry.Set(uint16(op), h)
	b.dirty = true
	return nil
}

// RegisterFunc 注册函数 Handler。
func (b *Bridge) RegisterFunc(op Opcode, name string, fn handler.Func, tags ...string) error {
	return b.Register(op, name, handler.Adapt(fn), tags...)
}

// Unregister 移除路由与 Handler。
func (b *Bridge) Unregister(op Opcode) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return ErrClosed
	}
	b.router.Remove(uint16(op))
	b.registry.Delete(uint16(op))
	b.dirty = true
	return nil
}

// EnableRoute 启用/禁用路由（Handler 仍保留）。
func (b *Bridge) EnableRoute(op Opcode, enabled bool) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return ErrClosed
	}
	m, ok := b.router.Get(uint16(op))
	if !ok {
		return ErrNoRoute
	}
	m.Enabled = enabled
	_ = b.router.Upsert(m)
	b.dirty = true
	return nil
}
