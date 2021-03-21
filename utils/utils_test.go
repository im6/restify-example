package utils

import "testing"

func TestIsHexColor(t *testing.T) {
	var testCases = []struct {
		input    string
		expected bool
	}{
		{"a3b", true},
		{"aob", false},
		{"abd192", true},
		{"000CDE", true},
	}
	for _, oneTestCase := range testCases {
		if output := IsHexColor(oneTestCase.input); output != oneTestCase.expected {
			t.Errorf("test failed: input %v, expect %v, received %v", oneTestCase.input, oneTestCase.expected, output)
		}
	}
}
