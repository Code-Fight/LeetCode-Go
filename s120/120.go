package s120

import (
	"math"
)

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。
相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

示例 1：
输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为11（即，2+3+5+1= 11）。

思路：
这里限定了只能从i或者i+1的方向移动，所以，跟斐波那契数列很像，只不过这里是二位的
依然是，先想怎么走
然后 构造dp公式和dp状态
*/

func minimumTotal(triangle [][]int) int {

	// 自底向上
	// DP公式：opt[i][j] = min(opt[i+1][j],opt[i+1][j+1])+triangle[i][j]
	// 可知，自底开始需要先初始化最后一层

	// 先初始化一个dp状态存储
	// !!!! 注意这里 其实可以不用开辟一个新的 opt  直接使用传过来的数组，这样就可以省去内存 因为 自底向上的计算，不会产生影响!!!!!!
	// !!!! 只需要屏蔽下面的代码，再把所有的 opt 换成 triangle 即可 !!!!
	opt := make([][]int, 0)
	for i := 0; i < len(triangle); i++ {
		temp := make([]int, 0)
		for j := 0; j < len(triangle[i]); j++ {
			temp = append(temp, triangle[i][j])
		}
		opt = append(opt, temp)
	}

	// 从底向上，因为最后一行也不需要计算，直接从倒数第二行开始计算
	for i := len(triangle) - 2; i >= 0; i-- {
		// 从每子元素的第一个开始计算
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = int(math.Min(float64(opt[i+1][j]), float64(opt[i+1][j+1]))) + triangle[i][j]
		}

	}

	return opt[0][0]

}
