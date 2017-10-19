package strings

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Rand(n int, slice_string ...string) string {
	if len(slice_string) == 1 {
		var sep string
		if strings.Contains(slice_string[0], ",") {
			sep = ","
		}
		slice_string = strings.Split(slice_string[0], sep)
	}
	var new_slice_string = make([]string, n)
	for i, slice_n := 0, len(slice_string); i < n; i++ {
		new_slice_string[i] = slice_string[rand.Intn(slice_n)]
	}
	return strings.Join(new_slice_string, "")
}
