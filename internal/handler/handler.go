package handler

import "context"

// Frame Handler 侧帧。
type Frame struct {
	Flags   uint8
	Opcode  uint16
	Payload []byte
}

// Ctx Handler 上下文。
type Ctx interface {
	Context() context.Context
	Opcode() uint16
}

// Handler 处理一帧。
type Handler interface {
	Handle(ctx context.Context, c Ctx, in Frame) (Frame, error)
}

// Func 函数适配。
type Func func(ctx context.Context, c Ctx, in Frame) (Frame, error)

type funcHandler struct{ fn Func }

func (f funcHandler) Handle(ctx context.Context, c Ctx, in Frame) (Frame, error) {
	return f.fn(ctx, c, in)
}

// Adapt 包装函数。
func Adapt(fn Func) Handler {
	if fn == nil {
		return nil
	}
	return funcHandler{fn: fn}
}
