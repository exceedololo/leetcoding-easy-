package Binary_Tree_Postorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	arr := []int{}
	traverse(root, &arr)
	return arr
}

func traverse(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, arr)
	traverse(root.Right, arr)
	*arr = append(*arr, root.Val)
}
