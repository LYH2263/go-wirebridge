package wirebridge

// SnapshotStats 返回计数快照。
func (b *Bridge) SnapshotStats() Stats {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.stats
}

func (b *Bridge) bumpServed() {
	b.mu.Lock()
	b.stats.Served++
	b.mu.Unlock()
}

func (b *Bridge) bumpReply() {
	b.mu.Lock()
	b.stats.Replies++
	b.mu.Unlock()
}

func (b *Bridge) bumpError() {
	b.mu.Lock()
	b.stats.Errors++
	b.mu.Unlock()
}

func (b *Bridge) bumpRouteMiss() {
	b.mu.Lock()
	b.stats.RouteMiss++
	b.mu.Unlock()
}

func (b *Bridge) bumpDecodeFail() {
	b.mu.Lock()
	b.stats.DecodeFail++
	b.mu.Unlock()
}

func (b *Bridge) bumpBypass() {
	b.mu.Lock()
	b.stats.BypassLogged++
	b.mu.Unlock()
}

func (b *Bridge) bumpRejected() {
	b.mu.Lock()
	b.stats.Rejected++
	b.mu.Unlock()
}
