package s104

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。


思路：
深度优先搜索，递归即可。
确定好停止条件也就是node为nil，然后层数，可以在最后一个节点时返回1，或者最后一个nil节点返回0 ，然后每层往上加一次
*/

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// 深度优先的递归
	// 就是只要节点不是nil 就往下走
	// 那么 只要遍历到nil的几点就返回0
	// 不是nil的节点，继续遍历
	// 同时，遍历完成会收到两个深度的返回值，
	// 取其中大的返回值即可
	if root ==nil{
		return 0
	}else {


		leftN :=maxDepth(root.Left)
		rightN:=maxDepth(root.Right)

		return int(math.Max(float64(rightN),float64(leftN))+1)

	}
}
