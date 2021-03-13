package s120

import "testing"

func Test120(t *testing.T) {
	i := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}

	println(minimumTotal(i))
}
