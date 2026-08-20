package wirebridge

import (
	"sync"
	"time"

	"example.com/wirebridge/internal/clock"
	"example.com/wirebridge/internal/handler"
	"example.com/wirebridge/internal/limit"
	"example.com/wirebridge/internal/route"
)

const (
	defaultMaxFrame  = 1 << 20
	defaultTimeout   = 5 * time.Second
	headerSize       = 7 // flags(1)+opcode(2) 计入 length；外层 4 字节 length 另计
	lengthPrefixSize = 4
)

// Bridge 二进制帧桥接门面。零值不可用，须 New。
type Bridge struct {
	mu sync.Mutex

	closed bool
	name   string
	clk    clock.Clock

	router   *route.Table
	registry *handler.Registry
	limiter  *limit.Limiter

	maxFrame       int
	persistPath    string
	bypassLog      string
	defaultTimeout time.Duration
	dirty          bool
	persistFn      func(path string, rows []route.Meta) error
	loadFn         func(path string) ([]route.Meta, error)

	stats Stats
}

// New 构造 Bridge。
func New(opts ...Option) *Bridge {
	b := &Bridge{
		clk:            clock.Real{},
		router:         route.NewTable(),
		registry:       handler.NewRegistry(),
		limiter:        limit.New(defaultMaxFrame),
		maxFrame:       defaultMaxFrame,
		defaultTimeout: defaultTimeout,
		name:           "wirebridge",
	}
	for _, o := range opts {
		if o != nil {
			o(b)
		}
	}
	if b.clk == nil {
		b.clk = clock.Real{}
	}
	if b.router == nil {
		b.router = route.NewTable()
	}
	if b.registry == nil {
		b.registry = handler.NewRegistry()
	}
	if b.limiter == nil {
		b.limiter = limit.New(b.maxFrame)
	}
	b.limiter.SetMax(b.maxFrame)
	if b.persistFn == nil {
		b.persistFn = defaultPersist
	}
	if b.loadFn == nil {
		b.loadFn = defaultLoad
	}
	return b
}

// Name 返回显示名。
func (b *Bridge) Name() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.name
}

// MaxFrame 当前最大帧长。
func (b *Bridge) MaxFrame() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.maxFrame
}

// IsClosed 是否已关闭。
func (b *Bridge) IsClosed() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.closed
}
