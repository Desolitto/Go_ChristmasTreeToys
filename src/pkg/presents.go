package pkg

import (
	"container/heap"
	"fmt"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (h PresentHeap) Len() int {
	return len(h)
}

func (h PresentHeap) Less(i, j int) bool {
	if h[i].Value == h[j].Value {
		return h[i].Size < h[j].Size
	}
	return h[i].Value > h[j].Value
}

func (h PresentHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PresentHeap) Push(x interface{}) {
	*h = append(*h, x.(Present))
}

func (h *PresentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func GetNCoolestPresents(presents []Present, n int) ([]Present, error) {
	if n < 0 || n > len(presents) {
		return nil, fmt.Errorf("invalid value of n: %d", n)
	}
	h := &PresentHeap{}
	heap.Init(h)
	for _, p := range presents {
		heap.Push(h, p)
	}
	result := make([]Present, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, heap.Pop(h).(Present))
	}
	return result, nil
}

// Одномерный массив
func GrabPresentsOptimized(presents []Present, size int) ([]Present, error) {
	if size < 1 {
		return nil, fmt.Errorf("wrong disk size")
	}
	n := len(presents)
	if n == 0 {
		return nil, fmt.Errorf("no Presents")
	}

	dp := make([]int, size+1)

	for _, present := range presents {
		for j := size; j >= present.Size; j-- {
			dp[j] = max(dp[j], dp[j-present.Size]+present.Value)
		}
	}
	if dp[size] == 0 {
		return nil, fmt.Errorf("no presents can be taken with the given size")
	}
	result := []Present{}
	w := size
	for i := n - 1; i >= 0 && w > 0; i-- {
		if presents[i].Size <= w && dp[w] != dp[w-presents[i].Size] {
			result = append(result, presents[i])
			w -= presents[i].Size
		}
	}

	return result, nil
}

// Двумерный массив
func GrabPresents(presents []Present, size int) ([]Present, error) {
	if size < 1 {
		return nil, fmt.Errorf("wrong disk size")
	}
	n := len(presents)
	if n == 0 {
		return nil, fmt.Errorf("no Presents")
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, size+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= size; j++ {
			if presents[i-1].Size <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-presents[i-1].Size]+presents[i-1].Value)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	if dp[n][size] == 0 {
		return nil, fmt.Errorf("no presents can be taken with the given size")
	}
	result := []Present{}
	w := size
	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			result = append(result, presents[i-1])
			w -= presents[i-1].Size
		}
	}

	return result, nil
}
