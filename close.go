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
	if b.dirty && b.persistPath != "" && b.router != nil {
		if err := b.syncLocked(); err != nil && first == nil {
			first = err
		}
	}
	b.closed = true
	// 保持 router/registry 非 nil 语义由 Serve 的 closed 检查拦截；
	// 清空路由内容前已 Sync。
	if b.router != nil {
		b.router.Replace(nil)
	}
	if b.registry != nil {
		b.registry.Clear()
	}
	return first
}
