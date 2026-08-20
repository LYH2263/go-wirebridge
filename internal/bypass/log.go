package bypass

import (
	"encoding/json"
	"os"
	"time"

	"example.com/wirebridge/internal/buffer"
)

// Entry 旁路日志条目。
type Entry struct {
	Opcode  uint16    `json:"opcode"`
	Flags   uint8     `json:"flags"`
	Payload []byte    `json:"payload"`
	At      time.Time `json:"at"`
}

// Append 追加一行 JSON；必须关闭文件句柄。
func Append(path string, e Entry) error {
	if e.At.IsZero() {
		e.At = time.Now().UTC()
	}
	e.Payload = buffer.CloneBytes(e.Payload)
	raw, err := json.Marshal(e)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	// BUG: 成功与写失败路径均未 Close
	if _, err := f.Write(append(raw, '\n')); err != nil {
		return err
	}
	return f.Sync()
}

// AbortWrite 旁路写失败时的收尾（删除半写入文件）。
func AbortWrite(path string) error {
	// BUG: plant 不收尾
	_ = path
	return nil
}
