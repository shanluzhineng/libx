package slicex

// 将一个切片分割成指定数量的切片,每个子切片的长度都是chunkSize
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

// 将一个切片均匀地分配到多个子切片中，子切片的数量由参数 partitionCount 决定
// 如果原始切片的长度小于 partitionCount，则每个子切片只包含一个元素
func PartitionSlice[SV any](input []SV, partitionCount int) [][]SV {
	var partitionSlice [][]SV
	count := 0
	if len(input) < partitionCount {
		count = len(input)
	} else {
		count = partitionCount
	}
	partitionSlice = make([][]SV, count)

	for i := 0; i < count; i++ {
		partitionSlice[i] = make([]SV, 0)
	}

	for i, v := range input {
		index := i % count
		partitionSlice[index] = append(partitionSlice[index], v)
	}

	return partitionSlice
}
