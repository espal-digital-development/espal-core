// Code generated by espal-store-synthesizer. DO NOT EDIT.
package emailtemplate_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/emailtemplate"
)

func TestAttachmentTable(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	if attachment.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestAttachmentTableAlias(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	if attachment.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestAttachmentIsUpdated(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	attachment.IsUpdated()
}

func TestAttachmentID(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	attachment.ID()
}

func TestAttachmentCreatedByID(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := "testValue"
	attachment.SetCreatedByID(testValue)
	if testValue != attachment.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestAttachmentUpdatedByID(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := "testValue"
	attachment.SetUpdatedByID(&testValue)
	if &testValue != attachment.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestAttachmentCreatedAt(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := time.Now()
	attachment.SetCreatedAt(testValue)
	if testValue != attachment.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestAttachmentUpdatedAt(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := time.Now()
	attachment.SetUpdatedAt(&testValue)
	if &testValue != attachment.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestAttachmentCreatedByFirstName(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := "testValue"
	attachment.SetCreatedByFirstName(&testValue)
	if &testValue != attachment.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestAttachmentCreatedBySurname(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := "testValue"
	attachment.SetCreatedBySurname(&testValue)
	if &testValue != attachment.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestAttachmentUpdatedByFirstName(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := "testValue"
	attachment.SetUpdatedByFirstName(&testValue)
	if &testValue != attachment.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestAttachmentUpdatedBySurname(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := "testValue"
	attachment.SetUpdatedBySurname(&testValue)
	if &testValue != attachment.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestAttachmentEmailTemplateID(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := "testValue"
	attachment.SetEmailTemplateID(testValue)
	if testValue != attachment.EmailTemplateID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestAttachmentFilePath(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := "testValue"
	attachment.SetFilePath(testValue)
	if testValue != attachment.FilePath() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestAttachmentLanguage(t *testing.T) {
	attachment := emailtemplate.NewAttachmentEntity()
	testValue := uint16(65000)
	attachment.SetLanguage(&testValue)
	if &testValue != attachment.Language() {
		t.Fatal("Getter did not return the Set value")
	}
}
