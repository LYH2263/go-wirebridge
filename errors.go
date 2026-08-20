package wirebridge

import (
	"errors"

	"example.com/wirebridge/internal/limit"
)

var (
	ErrClosed       = errors.New("wirebridge: closed")
	ErrNilHandler   = errors.New("wirebridge: nil handler")
	ErrNoRoute      = errors.New("wirebridge: no route")
	ErrTooLarge     = limit.ErrTooLarge
	ErrTruncated    = limit.ErrTruncated
	ErrBadLength    = errors.New("wirebridge: bad length")
	ErrInvalidFrame = errors.New("wirebridge: invalid frame")
	ErrPersist      = errors.New("wirebridge: persist failed")
	ErrDisabled     = errors.New("wirebridge: route disabled")
	ErrCanceled     = errors.New("wirebridge: canceled")
)
