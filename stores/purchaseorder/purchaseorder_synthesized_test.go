// Code generated by espal-store-synthesizer. DO NOT EDIT.
package purchaseorder_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/purchaseorder"
)

func TestPurchaseOrderTable(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	if p.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestPurchaseOrderTableAlias(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	if p.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestPurchaseOrderIsUpdated(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	p.IsUpdated()
}

func TestPurchaseOrderID(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	p.ID()
}

func TestPurchaseOrderCreatedByID(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := "testValue"
	p.SetCreatedByID(testValue)
	if testValue != p.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderUpdatedByID(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := "testValue"
	p.SetUpdatedByID(&testValue)
	if &testValue != p.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderCreatedAt(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := time.Now()
	p.SetCreatedAt(testValue)
	if testValue != p.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderUpdatedAt(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := time.Now()
	p.SetUpdatedAt(&testValue)
	if &testValue != p.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderCreatedByFirstName(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := "testValue"
	p.SetCreatedByFirstName(&testValue)
	if &testValue != p.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderCreatedBySurname(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := "testValue"
	p.SetCreatedBySurname(&testValue)
	if &testValue != p.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderUpdatedByFirstName(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := "testValue"
	p.SetUpdatedByFirstName(&testValue)
	if &testValue != p.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderUpdatedBySurname(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := "testValue"
	p.SetUpdatedBySurname(&testValue)
	if &testValue != p.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderSupplierID(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := "testValue"
	p.SetSupplierID(testValue)
	if testValue != p.SupplierID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderCurrency(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := uint16(65000)
	p.SetCurrency(testValue)
	if testValue != p.Currency() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderComments(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := "testValue"
	p.SetComments(&testValue)
	if &testValue != p.Comments() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderSellingPartyAutograph(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := "testValue"
	p.SetSellingPartyAutograph(&testValue)
	if &testValue != p.SellingPartyAutograph() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPurchaseOrderBuyingPartyAutograph(t *testing.T) {
	p := purchaseorder.NewPurchaseOrderEntity()
	testValue := "testValue"
	p.SetBuyingPartyAutograph(&testValue)
	if &testValue != p.BuyingPartyAutograph() {
		t.Fatal("Getter did not return the Set value")
	}
}
