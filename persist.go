package wirebridge

import (
	"example.com/wirebridge/internal/persist"
	"example.com/wirebridge/internal/route"
)

func defaultPersist(path string, rows []route.Meta) error {
	return persist.Save(path, rows)
}

func defaultLoad(path string) ([]route.Meta, error) {
	return persist.Load(path)
}

// Sync 将当前路由表写入持久化路径。
func (b *Bridge) Sync() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return ErrClosed
	}
	return b.syncLocked()
}

func (b *Bridge) syncLocked() error {
	if b.persistPath == "" || b.router == nil {
		return nil
	}
	// Close plant 可能在清空后调用，写出空路由
	rows := b.router.List()
	if err := b.persistFn(b.persistPath, rows); err != nil {
		return err
	}
	b.dirty = false
	return nil
}

// LoadPersist 从路径加载路由元数据（不覆盖已注册 Handler）。
func (b *Bridge) LoadPersist() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return ErrClosed
	}
	if b.persistPath == "" {
		return nil
	}
	rows, err := b.loadFn(b.persistPath)
	if err != nil {
		return err
	}
	b.router.Replace(rows)
	b.dirty = false
	return nil
}

// ReloadRoutes 用新表替换并持久化；失败回滚。
func (b *Bridge) ReloadRoutes(rows []RouteMeta) error {
	return b.ApplyRoutes(rows)
}
