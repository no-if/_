package math

import (
	"math"
)

func Join(slice_int []int) (b int64) {
	for i, count := 0, len(slice_int); i < count; i++ {
		b = b + int64(math.Pow10(count-i-1))*int64(slice_int[i])
	}
	return b
}
