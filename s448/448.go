package s448

/*
给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。

找到所有在 [1, n] 范围之间没有出现在数组中的数字。

您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。

示例:

输入:
[4,3,2,7,8,2,3,1]

输出:
[5,6]

思路：
最简单的版本就是做一个HashTable 然后就能快速的找出来了，但是需要浪费一点空间，并且

优化的办法：
遍历输入数组的每个元素一次。
我们将把 |nums[i]|-1 索引位置的元素标记为负数。即 nums[|nums[i] |- 1] \times -1nums[∣nums[i]∣−1]×−1 。
然后遍历数组，若当前数组元素 nums[i] 为负数，说明我们在数组中存在数字 i+1。
https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/solution/zhao-dao-suo-you-shu-zu-zhong-xiao-shi-de-shu-zi-2/

返回的时候，为了省内存，可以直接在原地操作。多建议一个哨兵index就可以了。

*/

import "math"

func findDisappearedNumbers(nums []int) []int {

	for i:=0;i<len(nums);i++{

		curr := int(math.Abs(float64(nums[i])))

		nums[curr-1] = int(math.Abs(float64(nums[curr-1])) * -1)

	}
	splitIndex :=0
	for i:=0;i<len(nums);i++{
		if nums[i]>0{
			nums[splitIndex] = i+1
			splitIndex++
		}
	}
	return nums[:splitIndex]
}
