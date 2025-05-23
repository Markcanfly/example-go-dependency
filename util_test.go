package examplegodependency

import (
	"testing"
)

func TestSpongeBobCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "HeLlO"},
		{"world", "WoRlD"},
		{"gopher", "GoPhEr"},
	}

	for _, test := range tests {
		result := SpongeBobCase(test.input)
		if result != test.expected {
			t.Errorf("SpongeBobCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
