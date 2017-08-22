package logic

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rue-brettadcock/storefront/mocks"
)

var l logic

func TestAddProductSKU_IDalreadyExists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mDB := mocks.NewMockSKUDataAccess(mockCtrl)
	l.mydb = mDB
	mDB.EXPECT().Get(100).Return("ID EXISTS")

	sku := SKU{"100", "polo", "ralph lauren", "10"}
	expected := errors.New("Product id already exists")
	actual := l.AddProductSKU(sku)

	if actual != expected {
		t.Errorf("Actual: %s\nExpected: %s", actual, expected)
	}
}

func TestAddProductSKU_quantityLessThan1(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mDB := mocks.NewMockSKUDataAccess(mockCtrl)
	l.mydb = mDB
	mDB.EXPECT().Get(25).Return("")

	sku := SKU{"25", "longboard", "landyachtz", "0"}
	expected := errors.New("Quantity must be at least 1")
	actual := l.AddProductSKU(sku)

	if actual != expected {
		t.Errorf("Actual: %s\nExpected: %s", actual, expected)
	}
}

func TestAddProductSKU_NegativeID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mDB := mocks.NewMockSKUDataAccess(mockCtrl)
	l.mydb = mDB
	mDB.EXPECT().Get(-1).Return("")

	sku := SKU{"-1", "longboard", "landyachtz", "1"}
	expected := errors.New("ID must be positive")
	actual := l.AddProductSKU(sku)

	if actual != expected {
		t.Errorf("Actual: %s\nExpected: %s", actual, expected)
	}
}

func TestAddProductSKU_ValidInput(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mDB := mocks.NewMockSKUDataAccess(mockCtrl)
	l.mydb = mDB
	mDB.EXPECT().Get(25).Return("")
	mDB.EXPECT().Insert("25", "longboard", "landyachtz", "1").Return(nil)

	sku := SKU{"25", "longboard", "landyachtz", "1"}
	actual := l.AddProductSKU(sku)

	if actual != nil {
		t.Errorf("Actual: %s\nExpected: nil", actual)
	}
}

func TestUpdateProductQuantity_IDdoesntExist(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mDB := mocks.NewMockSKUDataAccess(mockCtrl)
	l.mydb = mDB
	mDB.EXPECT().Get(25).Return("")

	sku := SKU{"25", "", "", "13"}
	expected := errors.New("Product id doesn't exist")
	actual := l.UpdateProductQuantity(sku)

	if actual != expected {
		t.Errorf("Actual: %s\nExpected: %s", actual, expected)
	}
}

func TestUpdateProductQuantity_SuccessfulUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mDB := mocks.NewMockSKUDataAccess(mockCtrl)
	l.mydb = mDB
	mDB.EXPECT().Get(25).Return("ID EXISTS")
	mDB.EXPECT().Update(25, 13).Return(nil)

	sku := SKU{"25", "", "", "13"}
	actual := l.UpdateProductQuantity(sku)

	if actual != nil {
		t.Errorf("Actual: %s\nExpected: nil", actual)
	}
}

func TestDeleteID_IDdoesntExist(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mDB := mocks.NewMockSKUDataAccess(mockCtrl)
	l.mydb = mDB
	mDB.EXPECT().Get(1000).Return("")

	sku := SKU{"1000", "", "", "13"}
	expected := errors.New("Product id doesn't exist")
	actual := l.DeleteID(sku)

	if actual != expected {
		t.Errorf("Actual: %s\nExpected: %s", actual, expected)
	}
}

func TestDeleteID_SuccessfulDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mDB := mocks.NewMockSKUDataAccess(mockCtrl)
	l.mydb = mDB
	mDB.EXPECT().Get(1000).Return("ID EXISTS")
	mDB.EXPECT().Delete(1000).Return(nil)

	sku := SKU{"1000", "longboard", "boosted", "13"}
	actual := l.DeleteID(sku)

	if actual != nil {
		t.Errorf("Actual: %s\nExpected: nil", actual)
	}
}
