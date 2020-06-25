// Code generated by espal-store-synthesizer. DO NOT EDIT.
package download_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/download"
)

func TestDownloadTable(t *testing.T) {
	d := download.NewDownloadEntity()
	if d.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestDownloadTableAlias(t *testing.T) {
	d := download.NewDownloadEntity()
	if d.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestDownloadIsUpdated(t *testing.T) {
	d := download.NewDownloadEntity()
	d.IsUpdated()
}

func TestDownloadID(t *testing.T) {
	d := download.NewDownloadEntity()
	d.ID()
}

func TestDownloadCreatedByID(t *testing.T) {
	d := download.NewDownloadEntity()
	testValue := "testValue"
	d.SetCreatedByID(testValue)
	if testValue != d.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDownloadUpdatedByID(t *testing.T) {
	d := download.NewDownloadEntity()
	testValue := "testValue"
	d.SetUpdatedByID(&testValue)
	if &testValue != d.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDownloadCreatedAt(t *testing.T) {
	d := download.NewDownloadEntity()
	testValue := time.Now()
	d.SetCreatedAt(testValue)
	if testValue != d.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDownloadUpdatedAt(t *testing.T) {
	d := download.NewDownloadEntity()
	testValue := time.Now()
	d.SetUpdatedAt(&testValue)
	if &testValue != d.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDownloadCreatedByFirstName(t *testing.T) {
	d := download.NewDownloadEntity()
	testValue := "testValue"
	d.SetCreatedByFirstName(&testValue)
	if &testValue != d.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDownloadCreatedBySurname(t *testing.T) {
	d := download.NewDownloadEntity()
	testValue := "testValue"
	d.SetCreatedBySurname(&testValue)
	if &testValue != d.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDownloadUpdatedByFirstName(t *testing.T) {
	d := download.NewDownloadEntity()
	testValue := "testValue"
	d.SetUpdatedByFirstName(&testValue)
	if &testValue != d.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDownloadUpdatedBySurname(t *testing.T) {
	d := download.NewDownloadEntity()
	testValue := "testValue"
	d.SetUpdatedBySurname(&testValue)
	if &testValue != d.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDownloadActive(t *testing.T) {
	d := download.NewDownloadEntity()
	testValue := true
	d.SetActive(testValue)
	if testValue != d.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDownloadFilePath(t *testing.T) {
	d := download.NewDownloadEntity()
	testValue := "testValue"
	d.SetFilePath(testValue)
	if testValue != d.FilePath() {
		t.Fatal("Getter did not return the Set value")
	}
}
