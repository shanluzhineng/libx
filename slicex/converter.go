package slicex

// conver []T to []interface{}
func ToInterfaceSlice[T any](list []T) []interface{} {
	result := make([]interface{}, 0)
	if len(list) <= 0 {
		return result
	}
	for index := range list {
		result = append(result, list[index])
	}
	return result
}
