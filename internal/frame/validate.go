package frame

// ValidFlags 检查未知标志位。
func ValidFlags(flags uint8, allowMask uint8) bool {
	return flags&^allowMask == 0
}

// BodyLenOK 检查 bodyLen 与实际是否一致。
func BodyLenOK(bodyLen, payloadLen int) bool {
	return bodyLen == HeaderSize+payloadLen
}
