package s322

import "testing"

func Test322(t *testing.T) {

	t.Log(coinChange([]int{1, 2, 5}, 11))

	t.Log(coinChange([]int{2}, 3))
}
