package lang

import (
	"math"
	"math/rand"
	"time"
)

// random byte array
func RandByteArray(length int) []int8 {
	datas := make([]int8, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		datas[i] = int8(rand.Intn(127) - 128)
	}
	return datas
}

// 在2个值之间随机一个值
func RandInt32(minValue int, maxValue int) int {
	rand.Seed(time.Now().UnixNano())
	randValue := rand.Intn(maxValue)
	if minValue < 0 {
		return randValue - int(math.Abs(float64(minValue)))
	}
	return randValue
}
