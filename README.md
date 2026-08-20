# go-wirebridge

长度前缀二进制帧桥接：Encode/Decode、按 opcode 路由 Handler、路由热更新与持久化。

## 运行管理服务

```bash
go run ./cmd/wired -addr :8101 -web web
```

浏览器打开 http://127.0.0.1:8101 。

## 库用法

```go
b := wirebridge.New(wirebridge.WithMaxFrame(1 << 20))
defer b.Close()
_ = b.RegisterFunc(2, "echo", ...)
out, err := b.ServeFrame(raw)
```
