// Code generated by espal-store-synthesizer. DO NOT EDIT.
package emailtemplate_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/emailtemplate"
)

func TestEmailTemplateTable(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	if e.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestEmailTemplateTableAlias(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	if e.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestEmailTemplateIsUpdated(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	e.IsUpdated()
}

func TestEmailTemplateID(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	e.ID()
}

func TestEmailTemplateCreatedByID(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	testValue := "testValue"
	e.SetCreatedByID(testValue)
	if testValue != e.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestEmailTemplateUpdatedByID(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	testValue := "testValue"
	e.SetUpdatedByID(&testValue)
	if &testValue != e.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestEmailTemplateCreatedAt(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	testValue := time.Now()
	e.SetCreatedAt(testValue)
	if testValue != e.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestEmailTemplateUpdatedAt(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	testValue := time.Now()
	e.SetUpdatedAt(&testValue)
	if &testValue != e.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestEmailTemplateCreatedByFirstName(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	testValue := "testValue"
	e.SetCreatedByFirstName(&testValue)
	if &testValue != e.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestEmailTemplateCreatedBySurname(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	testValue := "testValue"
	e.SetCreatedBySurname(&testValue)
	if &testValue != e.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestEmailTemplateUpdatedByFirstName(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	testValue := "testValue"
	e.SetUpdatedByFirstName(&testValue)
	if &testValue != e.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestEmailTemplateUpdatedBySurname(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	testValue := "testValue"
	e.SetUpdatedBySurname(&testValue)
	if &testValue != e.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestEmailTemplateDomainID(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	testValue := "testValue"
	e.SetDomainID(testValue)
	if testValue != e.DomainID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestEmailTemplateActive(t *testing.T) {
	e := emailtemplate.NewEmailTemplateEntity()
	testValue := true
	e.SetActive(testValue)
	if testValue != e.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}
