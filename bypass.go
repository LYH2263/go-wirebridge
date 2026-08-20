package wirebridge

import (
	"example.com/wirebridge/internal/bypass"
)

func writeBypass(path string, f Frame) error {
	return bypass.Append(path, bypass.Entry{
		Opcode:  uint16(f.Opcode),
		Flags:   uint8(f.Flags),
		Payload: f.Payload,
	})
}
