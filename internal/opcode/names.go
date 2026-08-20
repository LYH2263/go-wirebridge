package opcode

// 常用 opcode 常量（示例协议）。
const (
	Ping    uint16 = 1
	Echo    uint16 = 2
	Stats   uint16 = 3
	Admin   uint16 = 10
	Bypass  uint16 = 20
	Custom0 uint16 = 100
)

// Name 人类可读名。
func Name(op uint16) string {
	switch op {
	case Ping:
		return "ping"
	case Echo:
		return "echo"
	case Stats:
		return "stats"
	case Admin:
		return "admin"
	case Bypass:
		return "bypass"
	default:
		return "custom"
	}
}

// IsReserved 是否保留段。
func IsReserved(op uint16) bool {
	return op < 100
}
