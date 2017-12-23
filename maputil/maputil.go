package maputil

// EqualStringString returns true fi the two maps are identical (same length, same values, same keys)
func EqualStringString(m1 map[string]string, m2 map[string]string) bool {
	if m1 == nil && m2 == nil {
		return true
	}

	if m1 == nil || m2 == nil {
		return false
	}

	if len(m1) != len(m2) {
		return false
	}

	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
