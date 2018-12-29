package main

import "fmt"

/**
Given a binary tree, find the leftmost value in the last row of the tree.

Example 1:
Input:

    2
   / \
  1   3

Output:
1
Example 2:
Input:

        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7

Output:
7
Note: You may assume the tree (i.e., the given root node) is not NULL.
*/

func main() {
	root := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(findBottomLeftValue(root))
}

func findBottomLeftValue(root *TreeNode) int {
	var (
		res  = 0
		resD = 0
		walk func(dep int, curr *TreeNode)
	)
	walk = func(dep int, curr *TreeNode) {
		if curr == nil {
			return
		}
		walk(dep+1, curr.Left)
		walk(dep+1, curr.Right)
		if dep > resD {
			resD = dep
			res = curr.Val
		}
	}
	walk(1, root)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
