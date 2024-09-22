package main

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

	cur := []*TreeNode{root}
	for len(cur) > 0 {
		var val []int
		var next []*TreeNode
		for _, v := range cur {
			val = append(val, v.Val)
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		res = append(res, val)
		cur = next
	}

	return res
}

func main() {

}
