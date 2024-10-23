package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		expectErr  bool
	}{
		{"2 + 2", 4, false},
		{"2 - 2", 0, false},
		{"2 * 3", 6, false},
		{"6 / 2", 3, false},
		{"2 + 3 * 4", 14, false},
		{"(2 + 3) * 4", 20, false},
		{"2 + (3 * 4)", 14, false},
		{"(2 + 3) * (4 - 1)", 15, false},
		{"2 + 2 * (3 - 1)", 6, false},
		{"(2 + 2 * 3", 0, true}, // unmatched opening parenthesis
		{"2 / 0", 0, true},      // division by zero
		{"fluffy", 0, true},     // unknown token
		{"", 0, true},           // empty expression
	}

	for _, test := range tests {
		result, err := Calc(test.expression)
		if test.expectErr {
			if err == nil {
				t.Errorf("expected an error for expression '%s', got none", test.expression)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error for expression '%s': %v", test.expression, err)
			}
			if result != test.expected {
				t.Errorf("for expression '%s': expected %v, got %v", test.expression, test.expected, result)
			}
		}
	}
}
