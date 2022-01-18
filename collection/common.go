package collection_tools

// StringIn todo 泛型
func StringIn(elements []string, target string) bool {
	if elements == nil || len(elements) == 0 {
		return false
	}

	for _, ele := range elements {
		if ele == target {
			return true
		}
	}

	return false
}
