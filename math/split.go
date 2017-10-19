package math

func Split(b int64) (slice_int []int) {
	for ; b != 0; b = b / 10 {
		slice_int = append(slice_int, int(b%10))
	}
	for i, count := 0, len(slice_int)/2; i < count; i++ {
		slice_int[i], slice_int[count-i-1] = slice_int[count-i-1], slice_int[i]
	}
	return
}
