package s55

/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
示例 2:

输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。


思路：看图片
*/

import "math"

func canJump(nums []int) bool {
	maxJump :=nums[0]

	for i :=0; i<len(nums) ;i++{
		// 结束条件，只要最远能跳到的距离能和数组长度相等
		// 说明能到
		if maxJump == len(nums){
			return true
		}

		// 如果当前位置 比 最远能到的距离大
		// 说明前面所有的点加起来都到不了这个点
		if i > maxJump{
			return false
		}

		// 比较当前的点所能达到的距离(i 加上 i最大能调的距离 num[i])
		// 与之前最远能跳的距离
		maxJump =int( math.Max(float64(i+nums[i]),float64(maxJump)))
	}

	return false
}
