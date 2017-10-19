package strings

func Len(src string) (n int) {
	return len([]rune(src))
}
