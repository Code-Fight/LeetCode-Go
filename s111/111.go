package s111

/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

思路：
这个题直接遍历找最后一个，注意一下剪枝就可以了。
然后坑在：如果是如下这种情况，所以判定的情况，
改为先判断node的left和right是否都为nil 如果都是nil 就说明到根了，否则只走不为nil的一侧
    3
   / \
  9  20
 /     \
15     7

*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
var minD int
func minDepth(root *TreeNode) int {
	if root ==nil{
		return 0
	}
	minD =-1
	minDepthLoop(1,root)
	return minD
}

func minDepthLoop(curDepth int, node *TreeNode)  {
	// 左右都为nil 说明到了叶子节点
	// 然后 判断当前的深度是否比最小的深度小
	// 如果小，说明当前路径更短 就保存
	// 否则，直接return
	if node.Right ==nil  && node.Left ==nil{
		if minD ==-1 || minD >curDepth{
			minD = curDepth
		}
		return
	}
	// 剪枝
	// 如果当前深度大于 最小深度 就没有继续查找下去的必要
	if minD!=-1 && curDepth>minD{
		return
	}

	curDepth ++
	if node.Left !=nil{
		minDepthLoop(curDepth ,node.Left)
	}
	if node.Right !=nil{
		minDepthLoop(curDepth ,node.Right)

	}
}
