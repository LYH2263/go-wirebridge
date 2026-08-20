package wirebridge

import (
	"example.com/wirebridge/internal/bypass"
)

func writeBypass(path string, f Frame) error {
	err := bypass.Append(path, bypass.Entry{
		Opcode:  uint16(f.Opcode),
		Flags:   uint8(f.Flags),
		Payload: f.Payload,
	})
	if err != nil {
		// BUG: 旁路写失败未调用 AbortWrite 收尾
		return err
	}
	return nil
}
