// Code generated by espal-store-synthesizer. DO NOT EDIT.
package product_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/product"
)

func TestVariantTable(t *testing.T) {
	v := product.NewVariantEntity()
	if v.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestVariantTableAlias(t *testing.T) {
	v := product.NewVariantEntity()
	if v.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestVariantIsUpdated(t *testing.T) {
	v := product.NewVariantEntity()
	v.IsUpdated()
}

func TestVariantID(t *testing.T) {
	v := product.NewVariantEntity()
	v.ID()
}

func TestVariantCreatedByID(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetCreatedByID(testValue)
	if testValue != v.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantUpdatedByID(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetUpdatedByID(&testValue)
	if &testValue != v.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantCreatedAt(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := time.Now()
	v.SetCreatedAt(testValue)
	if testValue != v.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantUpdatedAt(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := time.Now()
	v.SetUpdatedAt(&testValue)
	if &testValue != v.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantCreatedByFirstName(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetCreatedByFirstName(&testValue)
	if &testValue != v.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantCreatedBySurname(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetCreatedBySurname(&testValue)
	if &testValue != v.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantUpdatedByFirstName(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetUpdatedByFirstName(&testValue)
	if &testValue != v.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantUpdatedBySurname(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetUpdatedBySurname(&testValue)
	if &testValue != v.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantModelID(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetModelID(testValue)
	if testValue != v.ModelID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantActive(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := true
	v.SetActive(testValue)
	if testValue != v.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantKey(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetKey(&testValue)
	if &testValue != v.Key() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantSorting(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := uint(1e9)
	v.SetSorting(testValue)
	if testValue != v.Sorting() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantTaxGroupID(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetTaxGroupID(testValue)
	if testValue != v.TaxGroupID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantNameRepresentationID(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetNameRepresentationID(&testValue)
	if &testValue != v.NameRepresentationID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantDescriptionRepresentationID(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetDescriptionRepresentationID(&testValue)
	if &testValue != v.DescriptionRepresentationID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVariantImageRepresentationID(t *testing.T) {
	v := product.NewVariantEntity()
	testValue := "testValue"
	v.SetImageRepresentationID(&testValue)
	if &testValue != v.ImageRepresentationID() {
		t.Fatal("Getter did not return the Set value")
	}
}
