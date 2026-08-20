package errors

import "strings"

// Join 拼接多错误信息（非 errors.Join，保持 go1.22 前风格描述）。
func Join(errs ...error) error {
	var parts []string
	var first error
	for _, e := range errs {
		if e == nil {
			continue
		}
		if first == nil {
			first = e
		}
		parts = append(parts, e.Error())
	}
	if len(parts) == 0 {
		return nil
	}
	if len(parts) == 1 {
		return first
	}
	return &joined{msg: strings.Join(parts, "; "), first: first}
}

type joined struct {
	msg   string
	first error
}

func (j *joined) Error() string { return j.msg }
func (j *joined) Unwrap() error { return j.first }
