package s1

/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

你可以按任意顺序返回答案。


思路：
首先，为了快速定位需要的数字是否存在，那么最好把所有的num放到一个map中，这样再挨个找即可
但是，这样需要遍历两次

可以，一边遍历，一边往map中加，如果map中有这个值，直接返回map的val，
map中val保存的是num对应的索引

如果找不到，就将该元素加到map中，给后面的元素使用。
*/

func twoSum(nums []int, target int) []int {
	readerMap := make(map[int]int)
	ret := []int{}
	for i := 0 ;i<len(nums) ; i++{
		if v,ok :=readerMap[target-nums[i]];ok{
			ret = append(ret,[]int{v,i}...)
			return ret
		}
		readerMap[nums[i]] = i
	}
	return ret
}
