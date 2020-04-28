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
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]<intervals[j][0]
	})
	merged :=[][]int{}
	for _,v:=range intervals{
		if len(merged)==0 || v[0]>merged[len(merged)-1][1]{
			merged = append(merged, v)
		}else {
			lastVal :=merged[len(merged)-1][1]
			//这里需要取一个两者的大值来作为右数值
			merged[len(merged)-1][1] =int(math.Max(float64(lastVal),float64(v[1])) )
		}

	}
	return merged
}
