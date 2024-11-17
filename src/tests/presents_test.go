package tests

import (
	"sort"
	"testing"
	"tree-toys/src/pkg"
)

func TestGetNCoolestPresents(t *testing.T) {
	tests := []struct {
		presents []pkg.Present
		n        int
		expected []pkg.Present
		err      bool
	}{
		{
			presents: []pkg.Present{{Value: 10, Size: 5}, {Value: 20, Size: 2}, {Value: 20, Size: 1}},
			n:        2,
			expected: []pkg.Present{{Value: 20, Size: 1}, {Value: 20, Size: 2}},
			err:      false,
		},
		{
			presents: []pkg.Present{{Value: 15, Size: 4}, {Value: 10, Size: 5}, {Value: 20, Size: 2}},
			n:        1,
			expected: []pkg.Present{{Value: 20, Size: 2}},
			err:      false,
		},
		{
			presents: []pkg.Present{{Value: 5, Size: 1}, {Value: 5, Size: 2}},
			n:        2,
			expected: []pkg.Present{{Value: 5, Size: 1}, {Value: 5, Size: 2}},
			err:      false,
		},
		{
			presents: []pkg.Present{{Value: 5, Size: 1}},
			n:        0,
			expected: []pkg.Present{},
			err:      false,
		},
		{
			presents: []pkg.Present{{Value: 5, Size: 1}},
			n:        -1,
			expected: nil,
			err:      true,
		},
		{
			presents: []pkg.Present{{Value: 5, Size: 1}},
			n:        2,
			expected: []pkg.Present{},
			err:      true,
		},
		{
			presents: []pkg.Present{{Value: 20, Size: 5}, {Value: 20, Size: 2}, {Value: 20, Size: 1}},
			n:        3,
			expected: []pkg.Present{{Value: 20, Size: 1}, {Value: 20, Size: 2}, {Value: 20, Size: 5}},
			err:      false,
		},
	}

	for _, test := range tests {
		result, err := pkg.GetNCoolestPresents(test.presents, test.n)

		if (err != nil) != test.err {
			t.Errorf("expected error: %v, got: %v", test.err, err)
		}
		if !test.err && !equalPresents(result, test.expected) {
			t.Errorf("expected: %v, got: %v", test.expected, result)
		}
	}
}
func sortPresents(presents []pkg.Present) {
	sort.Slice(presents, func(i, j int) bool {
		if presents[i].Value != presents[j].Value {
			return presents[i].Value > presents[j].Value
		}
		return presents[i].Size < presents[j].Size
	})
}

func TestGrabPresents(t *testing.T) {
	tests := []struct {
		presents []pkg.Present // Список подарков
		size     int           // Максимальный размер
		expected []pkg.Present // Ожидаемый результат
		err      bool          // Ожидаемая ошибка
	}{
		{
			presents: []pkg.Present{{Value: 60, Size: 10}, {Value: 100, Size: 20}, {Value: 120, Size: 30}},
			size:     50,
			expected: []pkg.Present{{Value: 100, Size: 20}, {Value: 120, Size: 30}},
			err:      false,
		},
		{
			presents: []pkg.Present{{Value: 10, Size: 5}, {Value: 15, Size: 10}, {Value: 40, Size: 20}},
			size:     15,
			expected: []pkg.Present{{Value: 15, Size: 10}, {Value: 10, Size: 5}},
			err:      false,
		},
		{
			presents: []pkg.Present{{Value: 20, Size: 5}, {Value: 30, Size: 10}, {Value: 50, Size: 15}},
			size:     5,
			expected: []pkg.Present{{Value: 20, Size: 5}},
			err:      false,
		},
		{
			presents: []pkg.Present{{Value: 10, Size: 5}, {Value: 15, Size: 10}, {Value: 40, Size: 20}},
			size:     1,
			expected: nil,
			err:      true,
		},
		{
			presents: []pkg.Present{{Value: 10, Size: 5}, {Value: 15, Size: 10}, {Value: 40, Size: 20}},
			size:     1,
			expected: nil,
			err:      true,
		},
		{
			presents: []pkg.Present{{Value: 10, Size: 5}},
			size:     -1,
			expected: nil,
			err:      true,
		},
		{
			presents: []pkg.Present{},
			size:     10,
			expected: []pkg.Present{},
			err:      true,
		},
	}
	for i, test := range tests {
		result, err := pkg.GrabPresents(test.presents, test.size)

		if (err != nil) != test.err {
			t.Errorf("test %d: expected error: %v, got: %v", i, test.err, err)
		}

		sortPresents(result)
		sortPresents(test.expected)

		if len(result) != len(test.expected) {
			t.Errorf("test %d: length mismatch: expected %d, got %d", i, len(test.expected), len(result))
			t.Logf("result: %v, expected: %v", result, test.expected)
			continue
		}

		for j := 0; j < len(result); j++ {
			if result[j] != test.expected[j] {
				t.Errorf("test %d: at index %d, expected %v, got %v", i, j, test.expected[j], result[j])
			}
		}
	}
}

func equalPresents(a, b []pkg.Present) bool {
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
