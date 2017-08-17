package service

import (
	"testing"
)

var p Presentation

func TestFormatPath(t *testing.T) {
	tests := []struct {
		input  string
		path   string
		prefix string
	}{
		{"/addSKU/id=4?name=brett?vend=rue?amt=2", "id=4;name=brett;vend=rue;amt=2", "/addSKU"},
		{"/updateSKU/id=4?amt=10", "id=4;amt=10", "/updateSKU"},
		{"/deleteSKU/id=1000", "id=1000", "/deleteSKU"},
		{"/printSKUs/", "", "/printSKUs"},
	}

	for _, c := range tests {
		path, prefix := formatPath(c.input)
		if path != c.path {
			t.Errorf("Actual: %s\nExpected: %s", path, c.path)
		}
		if prefix != c.prefix {
			t.Errorf("Actual: %s\nExpected: %s", prefix, c.prefix)
		}
	}
}
