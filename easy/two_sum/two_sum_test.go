package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Caso 1: Números consecutivos",
			nums:     []int{3, 4, 5, 6},
			target:   7,
			expected: []int{0, 1},
		},
		{
			name:     "Caso 2: Números não consecutivos",
			nums:     []int{4, 5, 6},
			target:   10,
			expected: []int{0, 2},
		},
		{
			name:     "Caso 3: Números iguais",
			nums:     []int{5, 5},
			target:   10,
			expected: []int{0, 1},
		},
		{
			name:     "Caso 4: Números negativos",
			nums:     []int{-1, -2, -3, -4},
			target:   -7,
			expected: []int{2, 3},
		},
		{
			name:     "Caso 5: Array vazio",
			nums:     []int{},
			target:   0,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := TwoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("TwoSum(%v, %d) = %v; esperado %v", tc.nums, tc.target, result, tc.expected)
			}
		})
	}
} 