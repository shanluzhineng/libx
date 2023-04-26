package mapx

// extract map[k]v from []v
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

// extract []v from map[k]v
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

// extract []string from map[k]v
func ToSliceStringWith[k comparable, v any](source map[k]v, valueFunc func(k) string) []string {
	value := make([]string, 0)
	if len(source) <= 0 {
		return value
	}
	for eachKey := range source {
		value = append(value, valueFunc(eachKey))
	}
	return value
}

// extract []k from map[k]v
func ExtractMapKeys[k comparable, v any](m map[k]v) []k {
	result := make([]k, 0)
	for eachKey := range m {
		result = append(result, eachKey)
	}
	return result
}

func ToInterfaceMap[T comparable, V any](v map[T]V) map[T]interface{} {
	result := make(map[T]interface{}, 0)
	if len(v) <= 0 {
		return result
	}
	for eachKey, eachValue := range v {
		result[eachKey] = eachValue
	}
	return result
}
