package s437

/*
给定一个二叉树，它的每个结点都存放着一个整数值。

找出路径和等于给定数值的路径总数。

路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11


思路：
暴力解法：每个子节点都可以作为起点去寻找sum，所以 递归每个子节点
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func pathSum(root *TreeNode, sum int) int {
	if root ==nil{
		return 0
	}

	//寻找当前节点是否存在sum
	cur := pathSumLoop(root,sum)
	//以当前节点的左子节点 为起点去寻找sum
	curLeft:= pathSum(root.Left,sum)
	//以当前节点的右子节点 为起点去寻找sum
	curRight:=pathSum(root.Right,sum)

	return cur + curLeft+curRight
}



func pathSumLoop(node *TreeNode, sum int) int {

	if node == nil{
		return 0
	}
	cur :=0

	// 因为不限制起点和结束的节点，那么必须在找到一个节点之后，继续往后走，如果当前节点不是sum那么cur就是0
	if node.Val ==sum{
		cur = 1
	}
	if node.Left!=nil{
		cur +=pathSumLoop(node.Left,sum - node.Val)
	}

	if node.Right !=nil{
		cur +=pathSumLoop(node.Right,sum - node.Val)
	}

	return cur
}
