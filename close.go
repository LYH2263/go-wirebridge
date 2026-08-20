package wirebridge

// Close 关闭 Bridge：先把脏路由 Sync 落盘，再标记关闭并释放路由表引用。
// 顺序不得颠倒，否则会把空表写进持久化文件。
func (b *Bridge) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return nil
	}
	var first error
	b.closed = true
	// 先把脏路由落盘，再释放内存路由表；顺序颠倒会把空表写进持久化文件。
	if b.dirty && b.persistPath != "" && b.router != nil {
		if err := b.syncLocked(); err != nil && first == nil {
			first = err
		}
	}
	if b.router != nil {
		b.router.Replace(nil)
	}
	if b.registry != nil {
		b.registry.Clear()
	}
	return first
}
