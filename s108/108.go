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


思路：
1. 首先高度平衡，那么左右子树的高度差不能超过1，所以根节点必须为整个数列的中间的值
2. 如果该数列已经排序好了，那么就直接使用，如果没有排序需要先排序，这里排序好了
3. 找到了中间节点，然后 将中间节点的左右分别构建子树即可
4. 左右子树同样需要平衡，不能超过1的高度差，那么继续找左右子树的中间节点，
5. 然后一直递归下去即可

注意高度平衡，并且提到了是一个升序的数组，那么直接中间元素为root即可，然后分别为每个子节点为根继续进行后续的查找即可。
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
