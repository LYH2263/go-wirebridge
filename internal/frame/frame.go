package frame

// Frame 内部帧表示。
type Frame struct {
	Flags   uint8
	Opcode  uint16
	Payload []byte
}

const (
	LenSize    = 4
	FlagsSize  = 1
	OpcodeSize = 2
	HeaderSize = FlagsSize + OpcodeSize // length 覆盖的头部
)

// Overhead 外层 length 前缀 + 内层头。
func Overhead() int { return LenSize + HeaderSize }
