package main

import (
	"testing"
)

func TestRPN(t *testing.T) {
	var cases = []struct {
		expected int
		input string
	}{
		{
			expected: 5,
			input: "2 3 + =",
		},
		{
			expected: 26,
			input: "2 3 * 4 5 * + =",
		},
		{
			expected: 15,
			input: "1 2 3 4 + * + =",
		},
		{
			expected: 21,
			input: "1 2 + 3 4 + * =",
		},
		{
			expected: 0,
			input: "1 2 + 3 4 + *",
		},
		{
			expected: 5555,
			input: "1234 4321 + =",
		},
		{
			expected: 24,
			input: "12 2 * =",
		},
		{
			expected: 6,
			input: "12 2 / = =",
		},
		{
			expected: 6,
			input: "13 2 / =",
		},
		{
			expected: 6,
			input: "13 2 / =",
		},
		{
			expected: 12,
			input: "2 13 2 / * =",
		},
		{
			expected: -1,
			input: "2 3 - =",
		},
		{
			expected: 1234,
			input: "1234 6 11 5 7 - / + * =",
		},
	}

	for _, item := range cases {
		result, err := RPN(item.input)
		if err != nil {
			t.Errorf("%v", err)
		}
		if result != item.expected {
			t.Errorf("results not match\nGot: %v\nExpected: %v", result, item.expected)
		}
	}

}