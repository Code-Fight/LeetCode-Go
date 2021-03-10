package s56

import (
	"math"
	"sort"
)
/*
56. 合并区间
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

思路：
首先进行排序，注意这里使用了sort库
然后，再两个区间一一对比，
对比的时候注意：当前节点左数值和已经合并的最后一个节点的右数值对比，然后如果需要合并，需要判断下，当前节点的右数值是否大于已经合并后的右数值
如果大于要覆盖掉
*/


func merge(intervals [][]int) [][]int {

	// sort
	// 按照区间的第一个元素进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]<intervals[j][0]
	})

	merged :=[][]int{}


	for _,v:=range intervals {

		// 处理边界，第一个区间直接进merged
		// 如果当前区间的左值（最小值）大于 已经合并过的（merged）的最大值，那直接加进数组中
		// 否则，说明当前 v 和 merged 的最后一个有区间的重合
		if len(merged) == 0 || v[0] > merged[len(merged)-1][1] {
			merged = append(merged, v)
		} else {
			// 有区间重合的元素，需要处理区间的右值，为当前元素的右值，或merged最后一个元素的右值
			// 所以，比较大小 并 赋值即可
			lastVal := merged[len(merged)-1][1]
			//这里需要取一个两者的大值来作为右数值
			merged[len(merged)-1][1] = int(math.Max(float64(lastVal), float64(v[1])))
		}

	}
	return merged
}
