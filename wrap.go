package wirebridge

import (
	"fmt"

	ierr "example.com/wirebridge/internal/errors"
)

// WrapTooLarge 包装过长错误，保留 ErrTooLarge 链。
func WrapTooLarge(detail error) error {
	return ierr.Wrap(ErrTooLarge, detail)
}

// WrapTruncated 包装截断错误。
func WrapTruncated(detail error) error {
	return ierr.Wrap(ErrTruncated, detail)
}

// FormatServeError 管理页展示用。
func FormatServeError(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%v", err)
}
