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
		// 写失败须删除半写入文件收尾；收尾错误不覆盖原始写失败原因
		_ = bypass.AbortWrite(path)
		return err
	}
	return nil
}
