package s101
/*
题目：
给定一个二叉树，检查它是否是镜像对称的。

 

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
 

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

思路：
首先要明确什么是镜像，根据题意可知：子节点全部为null 并且 当前节点相等就是镜像
使用递归思路，那么递归结束的时机就是，节点是否为null，都为null就是镜像，其中一个不为null就不是镜像

递归过程中判断是为镜像的标准就是，当前节点val相同，并且子节点为镜像
*/

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root,root)
}


func isMirror(left *TreeNode, right *TreeNode) bool{

	// 两个几点都为nil 说明相同
	if left == nil && right==nil{
		return true
	}
	if left ==nil || right ==nil {
		return false
	}

	// 同时判断左右节点必须都为true 才可以为true
	return isMirror(left.Left,right.Right) && isMirror(left.Right,right.Left) && left.Val ==right.Val
}
