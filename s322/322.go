package s322

import (
	"fmt"
)

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

你可以认为每种硬币的数量是无限的。

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

思路：
这个题目跟走楼梯是类似的，从0走到11层有多少种方法，再从这些方法中找出走的次数最少的。
那子问题就是 f(n) =min(f(n-1),f(n-2),f(n-5))+1  +1的原因是这里统计的是走的次数
那么状态就是dp[n] 为到当前价格最少的步数，因为 coins 是可变的，所以咱们从第0层开始走
所以  base dp  = 0
dp递推公式：f(n) = min(f(n-1),f(n-2),f(n-5))+1
*/

func coinChange(coins []int, amount int) int {

	// 因为从第0层开始走，所以dp 要多一位
	dp := make([]int, amount+1)
	// maxVal 是为了方便 coins 里面的每一个元素都互相比较而设置一个值
	maxVal := amount + 1

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// 第0层肯定是走上来肯定是 0
	// base 状态
	// 这里不是 maxVal 的原因是：这是一个特殊的层，也是一个基础层，这个曾需要作为第一层的基础，从第0层走上去第一层最短的路径就是
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		// 这里的给一个 maxVal ，是为了保证下面 分别对比 coins 里面的每一个硬币时，不让初始值产生影响
		dp[i] = maxVal
		for j := 0; j < len(coins); j++ {
			// 这里的 min 要和比较  dp[i] 的原因是 coins 里面是多个，
			// 通过上面的for j 的循环，将 coins 里面所有对象都对比的一遍
			// dp[i-coins[j] 代表的就是从 coins[j] 第 j 个元素走上来需要步数
			// 这里的判断是防止数组越界，还有如果硬币金额大于当前层数需要的金额，那就变成负的金额，毫无意义
			if coins[j] <= i {
				dp[i] = min(dp[i-coins[j]]+1, dp[i])
			}
		}
	}
	// 因为所有的dp状态的初始值都是math.MaxInt64
	// 所以，如果能走到这一步，那么说明个肯定不是这个值了
	// 否则，依然是默认的最大值
	if dp[amount] == maxVal {
		fmt.Println(dp[amount])
		return -1
	}

	return dp[amount]

}
