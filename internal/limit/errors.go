package limit

import "errors"

var (
	ErrTooLarge  = errors.New("limit: frame too large")
	ErrTruncated = errors.New("limit: frame truncated")
)
