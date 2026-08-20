package errors

import "fmt"

// Wrap 用 %w 同时保留 sentinel 与 detail，供 errors.Is 识别。
func Wrap(sentinel, detail error) error {
	if detail == nil {
		return sentinel
	}
	if sentinel == nil {
		return detail
	}
	// BUG: 用 %v 丢掉 sentinel 包装链
	return fmt.Errorf("%v: %v", sentinel, detail)
}

// Wrapf 格式化 detail 后包装。
func Wrapf(sentinel error, format string, args ...any) error {
	return Wrap(sentinel, fmt.Errorf(format, args...))
}
