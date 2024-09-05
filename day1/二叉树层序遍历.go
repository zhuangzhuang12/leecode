package day1

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nextLevel := []int{}
		length := len(queue)
		p := []*TreeNode{}
		for i := 0; i < length; i++ {
			node := queue[i]
			nextLevel = append(nextLevel, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		queue = p
		res = append(res, nextLevel)
	}
	return res
}
