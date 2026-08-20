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
		_ = bypass.AbortWrite(path)
		return err
	}
	return nil
}

func bypassAbort(path string) error {
	return bypass.AbortWrite(path)
}
