package session

import (
	"crypto/rand"
	"encoding/hex"
)

// NewID 生成随机会话 ID。
func NewID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}
