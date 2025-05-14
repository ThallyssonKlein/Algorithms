package group_anagrams

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "Exemplo 1",
			input:    []string{"act", "pots", "tops", "cat", "stop", "hat"},
			expected: [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
		{
			name:     "Exemplo 2 - String única",
			input:    []string{"x"},
			expected: [][]string{{"x"}},
		},
		{
			name:     "Exemplo 3 - String vazia",
			input:    []string{""},
			expected: [][]string{{""}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GroupAnagrams(tt.input)

			// Ordenar os resultados para comparação consistente
			sortGroups(result)
			sortGroups(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GroupAnagrams() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// sortGroups ordena os grupos e as strings dentro dos grupos para comparação consistente
func sortGroups(groups [][]string) {
	for i := range groups {
		// Ordenar strings dentro do grupo
		sort.Strings(groups[i])
	}
	// Ordenar grupos
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) == 0 {
			return true
		}
		if len(groups[j]) == 0 {
			return false
		}
		return groups[i][0] < groups[j][0]
	})
}
