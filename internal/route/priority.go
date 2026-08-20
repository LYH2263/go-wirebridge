package route

// SortByPriority 原地按 priority 降序、opcode 升序。
func SortByPriority(rows []Meta) {
	for i := 0; i < len(rows); i++ {
		for j := i + 1; j < len(rows); j++ {
			swap := false
			if rows[j].Priority > rows[i].Priority {
				swap = true
			} else if rows[j].Priority == rows[i].Priority && rows[j].Opcode < rows[i].Opcode {
				swap = true
			}
			if swap {
				rows[i], rows[j] = rows[j], rows[i]
			}
		}
	}
}
