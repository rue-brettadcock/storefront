package service

import (
	"testing"
)

var p Presentation

func TestFormatPath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"/addSKU/id=4?name=brett?vend=rue?amt=2", "id=4;name=brett;vend=rue;amt=2"},
		{"/updateSKU/id=4?amt=10", "id=4;amt=10"},
		{"/deleteSKU/id=1000", "id=1000"},
		{"/printSKUs/", ""},
	}

	for _, c := range tests {
		actual := formatPath(c.input)
		if actual != c.expected {
			t.Errorf("Actual: %s\nExpected: %s", actual, c.expected)
		}
	}
}
