package valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Exemplo 1 - Anagramas válidos",
			s:        "racecar",
			t:        "carrace",
			expected: true,
		},
		{
			name:     "Exemplo 2 - Não são anagramas",
			s:        "jar",
			t:        "jam",
			expected: false,
		},
		{
			name:     "Strings vazias",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "Strings de tamanhos diferentes",
			s:        "abc",
			t:        "abcd",
			expected: false,
		},
		{
			name:     "Mesmas letras, ordem diferente",
			s:        "listen",
			t:        "silent",
			expected: true,
		},
		{
			name:     "Mesma string",
			s:        "hello",
			t:        "hello",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("IsAnagram(%q, %q) = %v; esperado %v",
					tt.s, tt.t, result, tt.expected)
			}
		})
	}
} 