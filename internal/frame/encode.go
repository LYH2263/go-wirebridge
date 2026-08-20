package frame

import (
	"encoding/binary"

	"example.com/wirebridge/internal/limit"
)

// Encode 写出长度前缀帧。
func Encode(f Frame, maxFrame int, lim *limit.Limiter) ([]byte, error) {
	bodyLen := HeaderSize + len(f.Payload)
	total := LenSize + bodyLen
	if lim != nil {
		if err := lim.Check(total); err != nil {
			return nil, err
		}
	} else if maxFrame > 0 && total > maxFrame {
		return nil, limit.ErrTooLarge
	}
	out := make([]byte, total)
	binary.BigEndian.PutUint32(out[0:4], uint32(bodyLen))
	out[4] = f.Flags
	binary.BigEndian.PutUint16(out[5:7], f.Opcode)
	copy(out[7:], f.Payload)
	return out, nil
}

// EncodeInto 写入调用方缓冲；不够则返回所需长度。
func EncodeInto(dst []byte, f Frame) (int, error) {
	need := LenSize + HeaderSize + len(f.Payload)
	if len(dst) < need {
		return need, ErrShortBuffer
	}
	binary.BigEndian.PutUint32(dst[0:4], uint32(HeaderSize+len(f.Payload)))
	dst[4] = f.Flags
	binary.BigEndian.PutUint16(dst[5:7], f.Opcode)
	copy(dst[7:], f.Payload)
	return need, nil
}
