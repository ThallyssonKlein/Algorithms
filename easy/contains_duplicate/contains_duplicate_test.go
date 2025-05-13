package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Array com duplicata",
			nums:     []int{1, 2, 3, 3},
			expected: true,
		},
		{
			name:     "Array sem duplicata",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "Array vazio",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "Array com um elemento",
			nums:     []int{1},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("ContainsDuplicate(%v) = %v; esperado %v", tt.nums, result, tt.expected)
			}
		})
	}
} 