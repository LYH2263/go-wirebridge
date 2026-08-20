package frame

// SizeOf 计算编码后总长。
func SizeOf(payloadLen int) int {
	return LenSize + HeaderSize + payloadLen
}

// PayloadCapacity 给定 maxFrame 时 payload 最大长度。
func PayloadCapacity(maxFrame int) int {
	if maxFrame <= LenSize+HeaderSize {
		return 0
	}
	return maxFrame - LenSize - HeaderSize
}
