package gotests

import "testing"

func TestSubtraction(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -5, -3, -2},
		{"mixed numbers", -5, 3, -8},
		{"zero values", 0, 0, 0},
		{"with zero", 10, 0, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtraction(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtraction(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
