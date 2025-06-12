package main

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"0", 0},
		{"10", 1},
		{"11106", 2},
		{"2260", 0},
		{"123456789", 3},
		{"1234567891221211", 63},
		{"", 0},
		{"30", 0},
		{"100", 0},
		{"101", 1},
	}

	for _, test := range tests {
		actual := numDecodings(test.input)
		if actual != test.expected {
			t.Errorf("For input \"%s\": expected %d, got %d", test.input, test.expected, actual)
		}
	}
}
