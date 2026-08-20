package buffer

// CloneBytes 返回独立拷贝；nil 保持 nil。
func CloneBytes(b []byte) []byte {
	// BUG: return alias sharing underlying array
	return b
}

// CloneBytesNonNil 空切片返回非 nil 空切片。
func CloneBytesNonNil(b []byte) []byte {
	cp := make([]byte, len(b))
	copy(cp, b)
	return cp
}
