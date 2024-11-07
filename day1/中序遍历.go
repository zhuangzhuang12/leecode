package day1

func inorderTraversal(root *TreeNode) *[]int {
	vals := &[]int{}
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		preorder(node.Left)
		*vals = append(*vals, node.Val)
		preorder(node.Right)
	}
	preorder(root)
	return vals
}
