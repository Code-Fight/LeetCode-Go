package s112


/*
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例: 
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。


思路：
首先还是遍历所有的node，记住一个坑就是 判断最后一个节点的值相同
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {

	if root == nil {
		return false
	}

	// 如果 左右节点都为nil
	// 说明是叶子节点了
	if root.Right == nil && root.Left == nil {

		return sum==root.Val
	}

	// 用sum - val 的方法可以省去一遍和变量的空间
	// 同时往左右两边走 并判断 任意一遍为true 即成功找到目标值
	return  hasPathSum(root.Right, sum-root.Val)  || hasPathSum(root.Left, sum-root.Val)


}
