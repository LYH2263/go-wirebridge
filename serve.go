package wirebridge

import (
	"context"

	"example.com/wirebridge/internal/buffer"
	"example.com/wirebridge/internal/frame"
	"example.com/wirebridge/internal/handler"
)

// ServeFrame 解码并路由一帧（Background ctx）。
func (b *Bridge) ServeFrame(raw []byte) (Frame, error) {
	return b.ServeFrameContext(context.Background(), raw)
}

// ServeFrameContext 带取消的服务入口。
func (b *Bridge) ServeFrameContext(ctx context.Context, raw []byte) (Frame, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	// 入口即检查取消，避免忽略调用方 ctx
	if err := ctx.Err(); err != nil {
		return Frame{}, err
	}

	b.mu.Lock()
	closed := b.closed
	router := b.router
	registry := b.registry
	limiter := b.limiter
	maxFrame := b.maxFrame
	bypassLog := b.bypassLog
	b.mu.Unlock()

	// 关闭态立即返回明确错误，不再进路由表：
	// 热更新 Close 已把 router/registry 置空，继续往下会 nil 解引用。
	if closed || router == nil {
		b.bumpClosed()
		return Frame{}, ErrClosed
	}

	decoded, err := frame.Decode(raw, maxFrame, limiter)
	if err != nil {
		b.bumpDecodeFail()
		return Frame{}, err
	}
	// Payload 必须独立：后续 Handler / 调用方不得污染读缓冲
	payload := buffer.CloneBytes(decoded.Payload)
	in := Frame{
		Flags:   Flags(decoded.Flags),
		Opcode:  Opcode(decoded.Opcode),
		Payload: payload,
	}

	if in.Flags&FlagBypass != 0 && bypassLog != "" {
		if err := writeBypass(bypassLog, in); err != nil {
			return Frame{}, err
		}
		b.bumpBypass()
	}

	meta, ok := router.Get(uint16(in.Opcode))
	if !ok || !meta.Enabled {
		b.bumpRouteMiss()
		return Frame{}, ErrNoRoute
	}

	h := registry.Get(uint16(in.Opcode))
	if h == nil {
		b.bumpRouteMiss()
		return Frame{}, ErrNilHandler
	}

	sctx := &serveCtx{ctx: ctx, op: in.Opcode, bridge: b}
	out, err := handler.Invoke(ctx, h, sctx, handler.Frame{
		Flags:   uint8(in.Flags),
		Opcode:  uint16(in.Opcode),
		Payload: buffer.CloneBytes(in.Payload),
	})
	if err != nil {
		b.bumpError()
		return Frame{}, err
	}
	b.bumpServed()
	result := Frame{
		Flags:   Flags(out.Flags),
		Opcode:  Opcode(out.Opcode),
		Payload: buffer.CloneBytes(out.Payload),
	}
	if result.Flags&FlagReply != 0 {
		b.bumpReply()
	}
	return result, nil
}

type serveCtx struct {
	ctx    context.Context
	op     Opcode
	bridge *Bridge
}

func (s *serveCtx) Context() context.Context { return s.ctx }
func (s *serveCtx) Opcode() uint16           { return uint16(s.op) }
func (s *serveCtx) Bridge() *Bridge          { return s.bridge }
