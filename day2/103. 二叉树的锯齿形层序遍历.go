package day2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	treeList := []*TreeNode{root}
	if root == nil {
		return res
	}

	for level := 0; len(treeList) > 0; level++ {
		q := treeList
		treeList = nil
		vals := []int{}
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				treeList = append(treeList, node.Left)
			}
			if node.Right != nil {
				treeList = append(treeList, node.Right)
			}
		}
		if level%2 == 1 {
			for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
				vals[i], vals[j] = vals[j], vals[i]
			}
		}
		res = append(res, vals)
	}
	return res

}
