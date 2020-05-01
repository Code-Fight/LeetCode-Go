package s108

/*
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5


思路：注意高度平衡，并且提到了是一个升序的数组，那么直接中间元素为root即可，然后分别为每个子节点为根继续进行后续的查找即可。
*/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) ==0{
		return nil
	}

	mid :=len(nums)/2
	left := nums[:mid]
	right := nums[mid+1:]

	return &TreeNode{nums[mid],sortedArrayToBST(left),sortedArrayToBST(right)}

}
