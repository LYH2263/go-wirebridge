package limit

import (
	"fmt"

	ierr "example.com/wirebridge/internal/errors"
)

// Check 校验 total 字节是否超限；超限时以 %w 包装 ErrTooLarge。
func (l *Limiter) Check(total int) error {
	if l == nil {
		return nil
	}
	max := l.Max()
	if total < 0 {
		return ierr.Wrap(ErrTooLarge, fmt.Errorf("negative size %d", total))
	}
	if max > 0 && total > max {
		return ierr.Wrap(ErrTooLarge, fmt.Errorf("size %d exceeds max %d", total, max))
	}
	return nil
}

// CheckTruncated 当 got < need 时返回包装后的 ErrTruncated。
func CheckTruncated(need, got int) error {
	if got >= need {
		return nil
	}
	return ierr.Wrap(ErrTruncated, fmt.Errorf("need %d got %d", need, got))
}
