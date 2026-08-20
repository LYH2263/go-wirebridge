package wirebridge

// Opcode 业务操作码。
type Opcode uint16

// Flags 帧标志位。
type Flags uint8

const (
	FlagNone   Flags = 0
	FlagReply  Flags = 1 << 0
	FlagError  Flags = 1 << 1
	FlagMore   Flags = 1 << 2
	FlagBypass Flags = 1 << 3
)

// Frame 解码后的业务帧（Payload 为独立拷贝）。
type Frame struct {
	Flags   Flags
	Opcode  Opcode
	Payload []byte
}

// RouteMeta 路由元数据（管理页/热更新用）。
type RouteMeta struct {
	Opcode      Opcode
	Name        string
	Description string
	Tags        []string
	Enabled     bool
	Priority    int
}

// HandlerFunc 适配函数为 Handler。
type HandlerFunc func(ctx ServeCtx, in Frame) (Frame, error)

// ServeCtx 传递给 Handler 的上下文视图。
type ServeCtx interface {
	Context() interface {
		Done() <-chan struct{}
		Err() error
	}
	Opcode() Opcode
	Bridge() *Bridge
}

// Stats 运行计数。
type Stats struct {
	Served       uint64
	Replies      uint64
	Errors       uint64
	Rejected     uint64
	RouteMiss    uint64
	DecodeFail   uint64
	BypassLogged uint64
}
