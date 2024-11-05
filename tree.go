package tree

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func countToys(node *TreeNode) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.HasToy {
		count = 1
	}
	count += countToys(node.Left)
	count += countToys(node.Right)
	return count
}

func areToysBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return countToys(root.Left) == countToys(root.Right)
}

func unrollGarland(root *TreeNode) []bool {
	if root == nil {
		return []bool{}
	}

	var result []bool
	queue := []*TreeNode{root}
	level := false

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]bool, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[i]

			if level { // Четный уровень: слева направо
				currentLevel[levelSize-i-1] = node.HasToy
			} else { // Нечетный уровень: справа налево
				currentLevel[i] = node.HasToy
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}

		result = append(result, currentLevel...)
		queue = queue[levelSize:]
		level = !level
	}

	return result
}

// func unrollGarland(root *TreeNode) []bool {
// 	if root == nil {
// 		return []bool{}
// 	}
// 	var result []bool
// 	queue := []*TreeNode{root}
// 	level := 0
// 	for len(queue) > 0 {
// 		var nextLevel []*TreeNode
// 		if level%2 == 0 {
// 			for _, node := range queue {
// 				result = append(result, node.HasToy)
// 				if node.Right != nil {
// 					nextLevel = append(nextLevel, node.Right)
// 				}
// 				if node.Left != nil {
// 					nextLevel = append(nextLevel, node.Left)
// 				}
// 			}
// 		} else {
// 			for i := len(queue) - 1; i >= 0; i-- {
// 				node := queue[i]
// 				result = append(result, node.HasToy)
// 				if node.Right != nil {
// 					nextLevel = append(nextLevel, node.Right)
// 				}
// 				if node.Left != nil {
// 					nextLevel = append(nextLevel, node.Left)
// 				}

// 			}

// 		}
// 		queue = nextLevel
// 		level++
// 	}
// 	return result
// }

// func main() {
// 	root := &TreeNode{
// 		HasToy: true,
// 		Left: &TreeNode{
// 			HasToy: true,
// 		},
// 		Right: &TreeNode{
// 			HasToy: true,
// 		},
// 	}
// 	fmt.Println(areToysBalanced(root))
// }
