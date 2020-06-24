// Code generated by espal-store-synthesizer. DO NOT EDIT.
package product_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/product"
)

func TestModelTable(t *testing.T) {
	model := product.NewModelEntity()
	if model.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestModelTableAlias(t *testing.T) {
	model := product.NewModelEntity()
	if model.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestModelIsUpdated(t *testing.T) {
	model := product.NewModelEntity()
	model.IsUpdated()
}

func TestModelID(t *testing.T) {
	model := product.NewModelEntity()
	model.ID()
}

func TestModelCreatedByID(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetCreatedByID(testValue)
	if testValue != model.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelUpdatedByID(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetUpdatedByID(&testValue)
	if &testValue != model.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelCreatedAt(t *testing.T) {
	model := product.NewModelEntity()
	testValue := time.Now()
	model.SetCreatedAt(testValue)
	if testValue != model.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelUpdatedAt(t *testing.T) {
	model := product.NewModelEntity()
	testValue := time.Now()
	model.SetUpdatedAt(&testValue)
	if &testValue != model.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelCreatedByFirstName(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetCreatedByFirstName(&testValue)
	if &testValue != model.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelCreatedBySurname(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetCreatedBySurname(&testValue)
	if &testValue != model.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelUpdatedByFirstName(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetUpdatedByFirstName(&testValue)
	if &testValue != model.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelUpdatedBySurname(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetUpdatedBySurname(&testValue)
	if &testValue != model.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelActive(t *testing.T) {
	model := product.NewModelEntity()
	testValue := true
	model.SetActive(testValue)
	if testValue != model.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelSorting(t *testing.T) {
	model := product.NewModelEntity()
	testValue := uint(1e9)
	model.SetSorting(testValue)
	if testValue != model.Sorting() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelKey(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetKey(&testValue)
	if &testValue != model.Key() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelTaxGroupID(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetTaxGroupID(testValue)
	if testValue != model.TaxGroupID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelNameRepresentationID(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetNameRepresentationID(&testValue)
	if &testValue != model.NameRepresentationID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelDescriptionRepresentationID(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetDescriptionRepresentationID(&testValue)
	if &testValue != model.DescriptionRepresentationID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelImageRepresentationID(t *testing.T) {
	model := product.NewModelEntity()
	testValue := "testValue"
	model.SetImageRepresentationID(&testValue)
	if &testValue != model.ImageRepresentationID() {
		t.Fatal("Getter did not return the Set value")
	}
}
