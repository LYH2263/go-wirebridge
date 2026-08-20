package wirebridge

import (
	"example.com/wirebridge/internal/frame"
)

// Encode 将 Frame 编码为长度前缀字节。
func (b *Bridge) Encode(f Frame) ([]byte, error) {
	b.mu.Lock()
	maxFrame := b.maxFrame
	limiter := b.limiter
	closed := b.closed
	b.mu.Unlock()
	if closed {
		return nil, ErrClosed
	}
	return frame.Encode(frame.Frame{
		Flags:   uint8(f.Flags),
		Opcode:  uint16(f.Opcode),
		Payload: f.Payload,
	}, maxFrame, limiter)
}

// Decode 解码原始字节为 Frame（Payload 独立拷贝）。
func (b *Bridge) Decode(raw []byte) (Frame, error) {
	b.mu.Lock()
	maxFrame := b.maxFrame
	limiter := b.limiter
	closed := b.closed
	b.mu.Unlock()
	if closed {
		return Frame{}, ErrClosed
	}
	d, err := frame.Decode(raw, maxFrame, limiter)
	if err != nil {
		return Frame{}, err
	}
	return Frame{
		Flags:   Flags(d.Flags),
		Opcode:  Opcode(d.Opcode),
		Payload: append([]byte(nil), d.Payload...),
	}, nil
}
