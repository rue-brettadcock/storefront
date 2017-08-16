package logic

import (
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

	expected := "Product id already exists"
	actual := l.AddProductSKU(100, "polo", "ralph lauren", 10)

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

	expected := "Quantity must be at least 1"
	actual := l.AddProductSKU(25, "longboard", "landyachtz", 0)

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

	expected := "ID must be positive"
	actual := l.AddProductSKU(-1, "longboard", "landyachtz", 1)

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
	mDB.EXPECT().Insert(25, "longboard", "landyachtz", 1).Return(nil)

	expected := "Product successfully added to database"
	actual := l.AddProductSKU(25, "longboard", "landyachtz", 1)

	if actual != expected {
		t.Errorf("Actual: %s\nExpected: %s", actual, expected)
	}
}

func TestUpdateProductQuantity_IDdoesntExist(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mDB := mocks.NewMockSKUDataAccess(mockCtrl)
	l.mydb = mDB
	mDB.EXPECT().Get(25).Return("")

	expected := "Product id doesn't exist"
	actual := l.UpdateProductQuantity(25, 13)

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
	expected := "SKU successfully updated"
	actual := l.UpdateProductQuantity(25, 13)

	if actual != expected {
		t.Errorf("Actual: %s\nExpected: %s", actual, expected)
	}
}

func TestDeleteID_IDdoesntExist(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mDB := mocks.NewMockSKUDataAccess(mockCtrl)
	l.mydb = mDB
	mDB.EXPECT().Get(1000).Return("")

	expected := "Product id doesn't exist"
	actual := l.DeleteID(1000)

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

	expected := "Product successfully deleted"
	actual := l.DeleteID(1000)

	if actual != expected {
		t.Errorf("Actual: %s\nExpected: %s", actual, expected)
	}
}
