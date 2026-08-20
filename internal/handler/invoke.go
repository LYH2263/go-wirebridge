package handler

import (
	"context"
	"errors"
)

var ErrNilHandler = errors.New("handler: nil")

// Invoke 安全调用；nil Handler 返回 ErrNilHandler，不 panic。
func Invoke(ctx context.Context, h Handler, c Ctx, in Frame) (Frame, error) {
	if h == nil {
		return Frame{}, ErrNilHandler
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return Frame{}, err
	}
	return h.Handle(ctx, c, in)
}
