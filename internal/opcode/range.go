package opcode

// InRange 检查是否落在 [lo, hi]。
func InRange(op, lo, hi uint16) bool {
	return op >= lo && op <= hi
}

// Clamp 限制到范围。
func Clamp(op, lo, hi uint16) uint16 {
	if op < lo {
		return lo
	}
	if op > hi {
		return hi
	}
	return op
}
