package logic

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rue-brettadcock/storefront/database"
)

var l Logic

func init() {
	l = Logic{mydb: nil}
	l.mydb = database.New()

}

func TestAddProductSKU_idAlreadyExists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := NewMockDB(mockCtrl)
	expected := errors.New("Product id already exists")
	actual := l.AddProductSKUs(1, "polo", "ralph lauren", 10)

	if actual != expected {

	}
}
