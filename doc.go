// Package wirebridge 提供长度前缀二进制帧的编解码、按 opcode 路由到 Handler，
// 以及可选回写与路由表热更新/持久化。配套 cmd/wired 管理服务与 web 管理页。
//
// 帧格式（大端）：
//
//	[4 byte length][1 byte flags][2 byte opcode][payload…]
//
// length 覆盖 flags 起至 payload 末尾的字节数；最大帧长由 Bridge 选项限制。
package wirebridge
