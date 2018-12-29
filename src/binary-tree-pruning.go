package main

import "fmt"

/**

We are given the head node root of a binary tree, where additionally every node's value is either a 0 or a 1.

Return the same tree where every subtree (of the given tree) not containing a 1 has been removed.

(Recall that the subtree of a node X is X, plus every node that is a descendant of X.)

Example 1:
Input: [1,null,0,0,1]
Output: [1,null,0,null,1]

Explanation:
Only the red nodes satisfy the property "every subtree not containing a 1".
The diagram on the right represents the answer.

Example 2:
Input: [1,0,1,0,0,0,1]
Output: [1,null,1,null,1]

Example 3:
Input: [1,1,0,1,1,0,1,0]
Output: [1,1,0,1,1,null,1]
*/

func main() {
	root := &TreeNode{
		1,
		nil,
		&TreeNode{
			0,
			&TreeNode{0, nil, nil},
			&TreeNode{1, nil, nil}},
	}
	fmt.Println(pruneTree(root))
}

func pruneTree(root *TreeNode) *TreeNode {
	var (
		walk func(curr *TreeNode) *TreeNode
	)
	walk = func(curr *TreeNode) *TreeNode {
		if curr == nil {
			return nil
		}
		if curr.Left == nil && curr.Right == nil && curr.Val == 0 {
			return nil
		}
		curr.Left = walk(curr.Left)
		curr.Right = walk(curr.Right)
		return curr
	}
	walk(root)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
