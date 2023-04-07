package Binary_Tree_Preorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	arr := []int{}
	traverse(root, &arr)
	return arr
}

func traverse(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	traverse(root.Left, arr)
	traverse(root.Right, arr)
}
