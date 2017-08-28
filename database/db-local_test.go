package database

import (
	"reflect"
	"testing"

	errors "github.com/pkg/errors"
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

func TestUpdate_happypath_success(t *testing.T) {
	data := MemDb{db: newConnection()}
	data.db = populateSKUs()

	var expected []sku
	expected = append(expected, sku{"1", "polo", "CK", 2500})
	expected = append(expected, sku{"2", "pen", "bic", 50})
	expected = append(expected, sku{"3", "watch", "timex", 25})

	err := data.Update("1", 2500)
	if err != nil {
		t.Errorf("Actual: %v\nExpected: %v", data.db, expected)
	}
	if !reflect.DeepEqual(data.db, expected) {
		t.Errorf("Actual: %v\nExpected: %v", data.db, expected)
	}
}

func TestUpdate_idDoesntExist_noIDfound(t *testing.T) {
	data := MemDb{db: newConnection()}
	data.db = populateSKUs()

	err := data.Update("5", 2500)
	expected := errors.New("No matching product id found")
	if err.Error() != expected.Error() {
		t.Errorf("Actual: %v\nExpected: %v", err, expected)
	}
}
