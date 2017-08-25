package database

import (
	"reflect"
	"testing"
)

func populateSKUs() []sku {
	var res []sku
	res = append(res, sku{"1", "polo", "CK", 100})
	res = append(res, sku{"2", "pen", "bic", 50})
	res = append(res, sku{"3", "watch", "timex", 25})
	return res
}

func TestDelete_multipleElementSlice(t *testing.T) {
	data := MemDb{db: newConnection()}
	data.db = populateSKUs()

	err := data.Delete("1")
	if err != nil {
		t.Error(err)
	}

	var empty []sku
	expected := append(empty, sku{"2", "pen", "bic", 50})
	expected = append(expected, sku{"3", "watch", "timex", 25})
	if !reflect.DeepEqual(data.db, expected) {
		t.Errorf("Actual: %v\nExpected: %v", data.db, expected)
	}
}

func TestDelete_emptySlice(t *testing.T) {
	data := MemDb{db: newConnection()}
	var empty []sku

	err := data.Delete("1")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(data.db, empty) {
		t.Errorf("Actual: %v\nExpected: %v", data.db, empty)
	}
}
