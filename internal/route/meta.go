package route

// Meta 路由元数据。
type Meta struct {
	Opcode      uint16
	Name        string
	Description string
	Tags        []string
	Enabled     bool
	Priority    int
}
