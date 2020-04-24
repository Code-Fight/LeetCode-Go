package s39
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
​
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择

*/

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) ==0{
		return nil
	}

	ret := [][]int{}
	path :=[]int{}

	_dfs(&ret,path,candidates,target)
	return ret
}

func _dfs(ret *[][]int,path []int,candidates []int, target int)  {
	if target == 0{
		tmp := make([]int,len(path))
		copy(tmp,path)
		*ret = append(*ret,tmp)
		return
	}

	if target >0 {
		for _,v :=range candidates{
			//需要做一下剪枝，如果当前的值比队列中的最后一个值大，则跳过
			if len(path)>0 && v >path[len(path)-1]{
				continue
			}

			path = append(path, v)
			_dfs(ret,path,candidates,target-v)
			path = path[:len(path)-1]
		}
	}

}
