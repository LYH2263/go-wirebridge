package frame

import (
	"errors"
	"fmt"

	ierr "example.com/wirebridge/internal/errors"
	"example.com/wirebridge/internal/limit"
)

var (
	ErrTruncated   = limit.ErrTruncated
	ErrTooLarge    = limit.ErrTooLarge
	ErrBadLength   = errors.New("frame: bad length")
	ErrShortBuffer = errors.New("frame: short buffer")
)

func wrapTruncated(detail error) error {
	return ierr.Wrap(ErrTruncated, detail)
}

func wrapTooLarge(detail error) error {
	return ierr.Wrap(ErrTooLarge, detail)
}

// Describe 人类可读错误。
func Describe(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("frame error: %v", err)
}
