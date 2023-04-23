package mapx

// extract map[k]v from slice
// keySelectFunc cannot be nil
func FromSlice[k comparable, v any](source []v, keySelectFunc func(v) k) map[k]v {
	value := make(map[k]v)
	if len(source) <= 0 {
		return value
	}
	for i := range source {
		key := keySelectFunc(source[i])
		value[key] = source[i]
	}
	return value
}

// extract slice from map[k]v
func ToSlice[k comparable, v any](source map[k]v) []v {
	value := make([]v, 0)
	if len(source) <= 0 {
		return value
	}
	for _, eachItem := range source {
		value = append(value, eachItem)
	}
	return value
}
