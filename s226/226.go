package s226

/*
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

思路：
这个题，想明白一个点，怎么翻转，代码贼简单。
怎么翻转：
要翻转一个二叉树，其实就是把所有的节点的子树逐层（从低向上）翻转一下，就可以实现了。
一定要从低往上翻转
*/


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}

	// 不是nil 就继续往下走，深度优先
	if root.Left!=nil{
		invertTree(root.Left)
	}

	if root.Right!=nil{
		invertTree(root.Right)
	}

	// 翻转左右节点 哪怕是nil
	root.Left,root.Right = root.Right,root.Left

	return root
}
