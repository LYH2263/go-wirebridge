package handler

import "example.com/wirebridge/internal/buffer"

// ReplyOf 构造回复帧。
func ReplyOf(in Frame, payload []byte) Frame {
	return Frame{
		Flags:   in.Flags | 1,
		Opcode:  in.Opcode,
		Payload: buffer.CloneBytes(payload),
	}
}

// ErrorOf 构造错误标志帧。
func ErrorOf(in Frame, msg string) Frame {
	return Frame{
		Flags:   in.Flags | 2,
		Opcode:  in.Opcode,
		Payload: []byte(msg),
	}
}
