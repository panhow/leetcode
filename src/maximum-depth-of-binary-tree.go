package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		3,
		&TreeNode{9, nil, nil},
		&TreeNode{
			20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}},
	}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	var (
		res  = 0
		walk func(curr *TreeNode, depth int)
	)
	walk = func(curr *TreeNode, depth int) {
		if curr == nil {
			return
		}
		if depth > res {
			res = depth
		}
		walk(curr.Left, depth+1)
		walk(curr.Right, depth+1)
	}
	walk(root, 1)
	return res
}
