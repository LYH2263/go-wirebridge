package buffer

// AppendCopy 追加并保证底层独立（若 cap 共享则拷贝）。
func AppendCopy(dst, src []byte) []byte {
	if len(src) == 0 {
		return dst
	}
	out := make([]byte, len(dst)+len(src))
	copy(out, dst)
	copy(out[len(dst):], src)
	return out
}

// Equal 常量时间近似的字节比较（长度不同直接 false）。
func Equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var v byte
	for i := range a {
		v |= a[i] ^ b[i]
	}
	return v == 0
}
