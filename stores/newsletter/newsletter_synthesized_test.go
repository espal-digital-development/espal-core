// Code generated by espal-store-synthesizer. DO NOT EDIT.
package newsletter_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/newsletter"
)

func TestNewsletterTable(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	if newsletter.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestNewsletterTableAlias(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	if newsletter.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestNewsletterIsUpdated(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	newsletter.IsUpdated()
}

func TestNewsletterID(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	newsletter.ID()
}

func TestNewsletterCreatedByID(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	newsletter.SetCreatedByID(testValue)
	if testValue != newsletter.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterUpdatedByID(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	newsletter.SetUpdatedByID(&testValue)
	if &testValue != newsletter.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterCreatedAt(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := time.Now()
	newsletter.SetCreatedAt(testValue)
	if testValue != newsletter.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterUpdatedAt(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := time.Now()
	newsletter.SetUpdatedAt(&testValue)
	if &testValue != newsletter.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterCreatedByFirstName(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	newsletter.SetCreatedByFirstName(&testValue)
	if &testValue != newsletter.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterCreatedBySurname(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	newsletter.SetCreatedBySurname(&testValue)
	if &testValue != newsletter.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterUpdatedByFirstName(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	newsletter.SetUpdatedByFirstName(&testValue)
	if &testValue != newsletter.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterUpdatedBySurname(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	newsletter.SetUpdatedBySurname(&testValue)
	if &testValue != newsletter.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterDomainID(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	newsletter.SetDomainID(testValue)
	if testValue != newsletter.DomainID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterActive(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := true
	newsletter.SetActive(testValue)
	if testValue != newsletter.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterSendAt(t *testing.T) {
	newsletter := newsletter.NewNewsletterEntity()
	testValue := time.Now()
	newsletter.SetSendAt(&testValue)
	if &testValue != newsletter.SendAt() {
		t.Fatal("Getter did not return the Set value")
	}
}
