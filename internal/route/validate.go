package route

import (
	"fmt"
	"strings"
)

// Validate 校验单条元数据。
func Validate(m Meta) error {
	if m.Name == "" {
		return fmt.Errorf("route: name required")
	}
	if strings.Contains(m.Name, " ") && strings.TrimSpace(m.Name) != m.Name {
		return fmt.Errorf("route: name has surrounding spaces")
	}
	for _, tag := range m.Tags {
		if tag == "" {
			return fmt.Errorf("route: empty tag")
		}
	}
	return nil
}
