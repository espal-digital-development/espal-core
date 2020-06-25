// Code generated by espal-store-synthesizer. DO NOT EDIT.
package newsletter_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/newsletter"
)

func TestNewsletterTable(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	if n.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestNewsletterTableAlias(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	if n.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestNewsletterIsUpdated(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	n.IsUpdated()
}

func TestNewsletterID(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	n.ID()
}

func TestNewsletterCreatedByID(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	n.SetCreatedByID(testValue)
	if testValue != n.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterUpdatedByID(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	n.SetUpdatedByID(&testValue)
	if &testValue != n.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterCreatedAt(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := time.Now()
	n.SetCreatedAt(testValue)
	if testValue != n.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterUpdatedAt(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := time.Now()
	n.SetUpdatedAt(&testValue)
	if &testValue != n.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterCreatedByFirstName(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	n.SetCreatedByFirstName(&testValue)
	if &testValue != n.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterCreatedBySurname(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	n.SetCreatedBySurname(&testValue)
	if &testValue != n.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterUpdatedByFirstName(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	n.SetUpdatedByFirstName(&testValue)
	if &testValue != n.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterUpdatedBySurname(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	n.SetUpdatedBySurname(&testValue)
	if &testValue != n.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterDomainID(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := "testValue"
	n.SetDomainID(testValue)
	if testValue != n.DomainID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterActive(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := true
	n.SetActive(testValue)
	if testValue != n.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNewsletterSendAt(t *testing.T) {
	n := newsletter.NewNewsletterEntity()
	testValue := time.Now()
	n.SetSendAt(&testValue)
	if &testValue != n.SendAt() {
		t.Fatal("Getter did not return the Set value")
	}
}
