package route

// CloneStrings 深拷贝字符串切片。
func CloneStrings(in []string) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	copy(out, in)
	return out
}

// CloneMeta 深拷贝元数据（含 Tags）。
func CloneMeta(m Meta) Meta {
	return Meta{
		Opcode:      m.Opcode,
		Name:        m.Name,
		Description: m.Description,
		Tags:        CloneStrings(m.Tags),
		Enabled:     m.Enabled,
		Priority:    m.Priority,
	}
}
