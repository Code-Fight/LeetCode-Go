package s39

import "sort"

/*

给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]


思路：直接用回溯 ，然后 回溯的模板如下
https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie/hui-su-suan-fa-xiang-jie-xiu-ding-ban、
回溯一般为了快 都要进行剪枝

result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择

*/

func combinationSum(candidates []int, target int) [][]int {
	// 先排序，为了剪枝
	sort.Ints(candidates)
	res := [][]int{}
	dfs(candidates, nil, target, 0, &res)//深度优先
	return res
}

func dfs(candidates, nums []int, target, left int, res *[][]int) {
	// 结束递归
	if target == 0 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := left; i < len(candidates); i++ {// left限定，形成分支
		if target < candidates[i] { //剪枝
			return
		}
		// 直接组合append(nums, candidates[i]) 就不需要撤销了
		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i, res)//分支
	}
}
