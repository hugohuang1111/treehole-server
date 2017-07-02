package utils

// GetIntFromMap get int from map
func GetIntFromMap(m map[string]interface{}, key string) int {
	v := GetInterfaceFromMap(m, key)
	if nil == v {
		return 0
	}

	if i, ok := v.(int); ok {
		return i
	}

	return 0
}

// GetStringFromMap get string from map
func GetStringFromMap(m map[string]interface{}, key string) string {
	v := GetInterfaceFromMap(m, key)
	if nil == v {
		return ""
	}

	if s, ok := v.(string); ok {
		return s
	}

	return ""
}

// GetInterfaceFromMap get interface from map
func GetInterfaceFromMap(m map[string]interface{}, key string) interface{} {
	if v, ok := m[key]; ok {
		return v
	}

	return nil
}
