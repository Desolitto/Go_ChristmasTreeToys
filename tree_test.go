package tree

import "testing"

func TestBalancedTrue(t *testing.T) {
	tree := TreeNode{
		HasToy: false,
		Left: &TreeNode{
			HasToy: false,
			Left: &TreeNode{
				HasToy: false,
				Left:   nil,
				Right:  nil,
			},
			Right: &TreeNode{
				HasToy: true,
				Left:   nil,
				Right:  nil,
			},
		},
		Right: &TreeNode{
			HasToy: true,
			Left:   nil,
			Right:  nil,
		},
	}

	got := areToysBalanced(&tree)
	if got != true {
		t.Error("Got false, but true needed")
	}
}

func TestBalancedFalse(t *testing.T) {
	tree := TreeNode{
		HasToy: true,
		Left: &TreeNode{
			HasToy: true,
			Left:   nil,
			Right:  nil,
		},
		Right: &TreeNode{
			HasToy: false,
			Left:   nil,
			Right:  nil,
		},
	}

	got := areToysBalanced(&tree)
	if got != false {
		t.Error("Got true, but false needed")
	}
}

func TestUnroll(t *testing.T) {
	tree := TreeNode{
		HasToy: true,
		Left: &TreeNode{
			HasToy: true,
			Left: &TreeNode{
				HasToy: true,
				Left:   nil,
				Right:  nil,
			},
			Right: &TreeNode{
				HasToy: false,
				Left:   nil,
				Right:  nil,
			},
		},
		Right: &TreeNode{
			HasToy: false,
			Left: &TreeNode{
				HasToy: true,
				Left:   nil,
				Right:  nil,
			},
			Right: &TreeNode{
				HasToy: true,
				Left:   nil,
				Right:  nil,
			},
		},
	}

	got := unrollGarland(&tree)
	expect := []bool{true, true, false, true, true, false, true}
	if !equalSlices(got, expect) {
		t.Errorf("Got %v; expected %v", got, expect)
	}
}

// Вспомогательная функция для сравнения срезов
func equalSlices(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
