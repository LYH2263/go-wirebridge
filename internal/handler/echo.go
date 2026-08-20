package handler

import (
	"context"
	"time"

	"example.com/wirebridge/internal/buffer"
)

// Echo 回显 Handler。
func Echo() Handler {
	return Adapt(func(ctx context.Context, c Ctx, in Frame) (Frame, error) {
		if err := ctx.Err(); err != nil {
			return Frame{}, err
		}
		return Frame{
			Flags:   in.Flags | 1, // reply
			Opcode:  in.Opcode,
			Payload: buffer.CloneBytes(in.Payload),
		}, nil
	})
}

// DelayEcho 先 Wait 再回显（用于 context 测试）。
func DelayEcho(d time.Duration) Handler {
	return Adapt(func(ctx context.Context, c Ctx, in Frame) (Frame, error) {
		if err := Wait(ctx, d); err != nil {
			return Frame{}, err
		}
		return Frame{
			Flags:   in.Flags | 1,
			Opcode:  in.Opcode,
			Payload: buffer.CloneBytes(in.Payload),
		}, nil
	})
}
