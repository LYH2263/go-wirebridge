package wirebridge

// Healthy 简单健康检查。
func (b *Bridge) Healthy() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return !b.closed && b.router != nil && b.registry != nil
}
