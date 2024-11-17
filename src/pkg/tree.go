package pkg

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func CountToys(node *TreeNode) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.HasToy {
		count = 1
	}
	count += CountToys(node.Left)
	count += CountToys(node.Right)
	return count
}

func AreToysBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return CountToys(root.Left) == CountToys(root.Right)
}

func UnrollGarland(root *TreeNode) []bool {
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
