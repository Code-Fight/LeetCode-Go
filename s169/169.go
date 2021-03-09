package s169

/*

给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

思路：
hashtable+打擂台 O(n)

还有其他的利用数学思路的解法，但是感觉不通用，然后还有摩尔投票的解法
摩尔投票的意思：遇到相同的就加1，遇到不同的就减1，减到0就重新换个数开始计数，总能找到最多的那个

*/
/*
//hashtable 打擂台的解法
func majorityElement(nums []int) int {
	all := make(map[int]int)
	max :=0
	maxVal :=0
	for _,v:= range nums{

		// 每次遍历到 就直接加 1
		all[v]+=1
		// 如果当前元素 跟 目前最大的值一样，就继续投票
		if v ==maxVal{
			max++
		}else {
			// 否则 需要判断 当前 v 出现的次数，是否大于max
			// 大于就将maxval 换成v 出现的次数max 换成all[v]即可
			if all[v]>max{
				max = all[v]
				maxVal = v
			}
		}

	}
	return maxVal
}
*/
//摩尔投票
func majorityElement(nums []int) int {
	max :=0
	maxVal :=0
	for _,v:= range nums{

		if max == 0{
			maxVal = v
			max = 1
		}else {
			if maxVal == v {
				max++
			}else {
				max--
			}
		}


	}
	return maxVal
}
