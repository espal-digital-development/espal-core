// Code generated by espal-store-synthesizer. DO NOT EDIT.
package shop_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/shop"
)

func TestTranslationTable(t *testing.T) {
	tt := shop.NewTranslationEntity()
	if tt.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestTranslationTableAlias(t *testing.T) {
	tt := shop.NewTranslationEntity()
	if tt.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestTranslationIsUpdated(t *testing.T) {
	tt := shop.NewTranslationEntity()
	tt.IsUpdated()
}

func TestTranslationID(t *testing.T) {
	tt := shop.NewTranslationEntity()
	tt.ID()
}

func TestTranslationCreatedByID(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := "testValue"
	tt.SetCreatedByID(testValue)
	if testValue != tt.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationUpdatedByID(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := "testValue"
	tt.SetUpdatedByID(&testValue)
	if &testValue != tt.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationCreatedAt(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := time.Now()
	tt.SetCreatedAt(testValue)
	if testValue != tt.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationUpdatedAt(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := time.Now()
	tt.SetUpdatedAt(&testValue)
	if &testValue != tt.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationCreatedByFirstName(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := "testValue"
	tt.SetCreatedByFirstName(&testValue)
	if &testValue != tt.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationCreatedBySurname(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := "testValue"
	tt.SetCreatedBySurname(&testValue)
	if &testValue != tt.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationUpdatedByFirstName(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := "testValue"
	tt.SetUpdatedByFirstName(&testValue)
	if &testValue != tt.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationUpdatedBySurname(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := "testValue"
	tt.SetUpdatedBySurname(&testValue)
	if &testValue != tt.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationLanguage(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := uint16(65000)
	tt.SetLanguage(testValue)
	if testValue != tt.Language() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationField(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := uint16(65000)
	tt.SetField(testValue)
	if testValue != tt.Field() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationValue(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := "testValue"
	tt.SetValue(testValue)
	if testValue != tt.Value() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationShopID(t *testing.T) {
	tt := shop.NewTranslationEntity()
	testValue := "testValue"
	tt.SetShopID(testValue)
	if testValue != tt.ShopID() {
		t.Fatal("Getter did not return the Set value")
	}
}
