package metrics

// Snapshot 一组计数。
type Snapshot struct {
	Served     uint64 `json:"served"`
	Errors     uint64 `json:"errors"`
	RouteMiss  uint64 `json:"route_miss"`
	DecodeFail uint64 `json:"decode_fail"`
}

// Collector 收集器。
type Collector struct {
	Served     Counter
	Errors     Counter
	RouteMiss  Counter
	DecodeFail Counter
}

// Snap 快照。
func (c *Collector) Snap() Snapshot {
	return Snapshot{
		Served:     c.Served.Get(),
		Errors:     c.Errors.Get(),
		RouteMiss:  c.RouteMiss.Get(),
		DecodeFail: c.DecodeFail.Get(),
	}
}
