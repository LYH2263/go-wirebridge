package errors

import "errors"

// 哨兵与公共包对齐的别名空间，便于 internal 引用。
var (
	ErrClosed    = errors.New("wirebridge: closed")
	ErrTooLarge  = errors.New("limit: frame too large")
	ErrTruncated = errors.New("limit: frame truncated")
)
