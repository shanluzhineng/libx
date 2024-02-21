package slicex

// chunk a slice to
func ChunkSlice[SV any](input []SV, chunkSize int) [][]SV {
	var result [][]SV

	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize

		if end > len(input) {
			end = len(input)
		}

		result = append(result, input[i:end])
	}

	return result
}
