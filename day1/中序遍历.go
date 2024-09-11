package day1

func inorderTraversal(root *TreeNode) *[]int {
	res := &[]int{}
	treverse(root, res)
	return res
}

func treverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	treverse(root.Left, res)
	*res = append(*res, root.Val)
	treverse(root.Right, res)
}
