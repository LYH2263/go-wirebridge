# 二进制帧桥接

长度前缀帧编解码、按 opcode 路由到 Handler、可选回写；管理页发送试帧、查看路由表。


## 环境

- 镜像：`benzhi.Dockerfile` 基于 `golang:1.22`（官方多架构）
- `go.mod` 语言版本：go 1.22
- 容器内使用镜像自带工具链即可

## 标准命令

```bash
go build ./...
go test ./... -count=1
go vet ./...
```

## 构建评测镜像（须双架构）

验证请用 `bash -c`（勿用 `bash -lc`）。

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh go-wirebridge linux/amd64
docker run --platform linux/amd64 --rm go-wirebridge:latest bash -c 'go build ./...'

./build_benzhi_docker.sh go-wirebridge linux/arm64
docker run --platform linux/arm64 --rm go-wirebridge:latest bash -c 'go build ./...'
```

构建阶段已 `go mod download`；容器内编译不应再出现 `downloading ...`。
