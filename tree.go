package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func countToys(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + countToys(node.Left) + countToys(node.Right)
}

func areToysBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftCount := countToys(root.Left)
	rightCount := countToys(root.Right)

	return leftCount == rightCount
}

func main() {
	root := &TreeNode{
		HasToy: true,
		Left: &TreeNode{
			HasToy: true,
		},
		Right: &TreeNode{
			HasToy: true,
		},
	}
	fmt.Println(areToysBalanced(root))
}
