package wirebridge

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func validateMeta(r RouteMeta) error {
	if r.Name == "" {
		return fmt.Errorf("wirebridge: empty route name for opcode %d", r.Opcode)
	}
	if !utf8.ValidString(r.Name) {
		return fmt.Errorf("wirebridge: invalid route name encoding")
	}
	if strings.ContainsAny(r.Name, "\x00\n\r") {
		return fmt.Errorf("wirebridge: route name contains illegal chars")
	}
	for _, t := range r.Tags {
		if strings.TrimSpace(t) == "" {
			return fmt.Errorf("wirebridge: empty tag")
		}
	}
	return nil
}

// ValidateFrame 轻量校验已解码帧。
func ValidateFrame(f Frame) error {
	if f.Flags&^(FlagReply|FlagError|FlagMore|FlagBypass) != 0 {
		return ErrInvalidFrame
	}
	return nil
}
