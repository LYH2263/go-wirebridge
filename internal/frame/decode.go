package frame

import (
	"encoding/binary"
	"fmt"

	"example.com/wirebridge/internal/buffer"
	"example.com/wirebridge/internal/limit"
)

// Decode 解析长度前缀帧。返回的 Payload 必须是独立拷贝，不得别名 raw。
func Decode(raw []byte, maxFrame int, lim *limit.Limiter) (Frame, error) {
	if len(raw) < LenSize {
		return Frame{}, wrapTruncated(fmt.Errorf("need %d length bytes, got %d", LenSize, len(raw)))
	}
	bodyLen := int(binary.BigEndian.Uint32(raw[0:4]))
	if bodyLen < HeaderSize {
		return Frame{}, ErrBadLength
	}
	total := LenSize + bodyLen
	if lim != nil {
		if err := lim.Check(total); err != nil {
			return Frame{}, err
		}
	} else if maxFrame > 0 && total > maxFrame {
		return Frame{}, wrapTooLarge(fmt.Errorf("total %d > max %d", total, maxFrame))
	}
	if len(raw) < total {
		return Frame{}, wrapTruncated(fmt.Errorf("need %d bytes, got %d", total, len(raw)))
	}
	flags := raw[4]
	opcode := binary.BigEndian.Uint16(raw[5:7])
	payloadRaw := raw[7:total]
	// 关键：拷贝 payload，避免与底层读缓冲共享
	payload := buffer.CloneBytes(payloadRaw)
	return Frame{Flags: flags, Opcode: opcode, Payload: payload}, nil
}

// PeekOpcode 不拷贝 payload，仅窥视 opcode。
func PeekOpcode(raw []byte) (uint16, error) {
	if len(raw) < LenSize+HeaderSize {
		return 0, ErrTruncated
	}
	bodyLen := int(binary.BigEndian.Uint32(raw[0:4]))
	if bodyLen < HeaderSize {
		return 0, ErrBadLength
	}
	return binary.BigEndian.Uint16(raw[5:7]), nil
}
