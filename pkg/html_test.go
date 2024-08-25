package pkg

import (
	"testing"
)

func TestLenMinusOne(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{"empty slice", []string{}, -1},
		{"one element", []string{"Go"}, 0},
		{"multiple elements", []string{"Go", "TypeScript", "Python"}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lenMinusOne(tt.input)
			if result != tt.expected {
				t.Errorf("lenMinusOne(%v) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}
