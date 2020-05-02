package s110

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

思路：
首先要确定什么是高度平衡二叉树：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

用递归，将整个树的求解结果，拆分为每个子树的结果，然后先获得子树的高度，判定是否符合高度，如果大于1那么直接return -1 实现快速剪枝来结束循环

*/


import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func isBalanced(root *TreeNode) bool {
	if root ==nil{
		return true
	}

	res :=balanceLoop(root.Left)

	if res==-1{
		return false
	}
	return true

}

func balanceLoop(nood *TreeNode) float64 {
	if nood ==nil {
		return 0
	}
	left :=balanceLoop(nood.Left)
	// 剪枝
	if left==-1	{
		return -1
	}
	right :=balanceLoop(nood.Right)
	if right ==-1{
		return -1
	}

	if math.Abs(left-right) >1 {
		return -1
	}else {
		// 如果该子树平衡 那么返回当前的高度 并加上1 就为父节点的高度
		return math.Max((left),(right))+1

	}

}
