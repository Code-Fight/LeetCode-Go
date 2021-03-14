package s53

import "testing"

func Test53(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	t.Log(maxSubArray(nums))
}
