package s40

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

思路：
主要是回溯的思想，参考39题的回溯模板

这里是不重复的，所以首先得做一个排序，如果不能排序，那么可以做一个HashTable，已经计算过的就排除掉
防止出现重复的重复计算的问题

然后 需要注意 剪枝的问题，如果大于当前target 那么就不要继续找了，以及重复元素的过滤

*/

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates)==0{
		return nil
	}
	sort.Ints(candidates)
	ret :=[][]int{}
	path :=[]int{}
	_dfs(&ret ,path,candidates,target)
	return ret
}

func _dfs(ret *[][]int,path []int, candidates []int, target int)  {
	if target ==0{
		temp := make([]int,len(path))
		copy(temp,path)
		*ret = append(*ret, temp)
		return
	}
	if target>0{
		for i,v :=range candidates{

			//判断下面的值是否值得继续递归，如果都大了，就没必要了
			if target< v{
				return
			}
			//过滤重复的元素
			if i>0 && v == candidates[i-1]{
				continue
			}

			path = append(path, v)

			_dfs(ret,path,candidates[i+1:],target-v)
			path = path[:len(path)-1]

		}
	}
}
