// Code generated by espal-store-synthesizer. DO NOT EDIT.
package menu_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/menu"
)

func TestMenuTable(t *testing.T) {
	m := menu.NewMenuEntity()
	if m.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestMenuTableAlias(t *testing.T) {
	m := menu.NewMenuEntity()
	if m.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestMenuIsUpdated(t *testing.T) {
	m := menu.NewMenuEntity()
	m.IsUpdated()
}

func TestMenuID(t *testing.T) {
	m := menu.NewMenuEntity()
	m.ID()
}

func TestMenuCreatedByID(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := "testValue"
	m.SetCreatedByID(testValue)
	if testValue != m.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuUpdatedByID(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := "testValue"
	m.SetUpdatedByID(&testValue)
	if &testValue != m.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuCreatedAt(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := time.Now()
	m.SetCreatedAt(testValue)
	if testValue != m.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuUpdatedAt(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := time.Now()
	m.SetUpdatedAt(&testValue)
	if &testValue != m.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuCreatedByFirstName(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := "testValue"
	m.SetCreatedByFirstName(&testValue)
	if &testValue != m.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuCreatedBySurname(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := "testValue"
	m.SetCreatedBySurname(&testValue)
	if &testValue != m.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuUpdatedByFirstName(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := "testValue"
	m.SetUpdatedByFirstName(&testValue)
	if &testValue != m.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuUpdatedBySurname(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := "testValue"
	m.SetUpdatedBySurname(&testValue)
	if &testValue != m.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuActive(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := true
	m.SetActive(testValue)
	if testValue != m.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuSorting(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := uint(1e9)
	m.SetSorting(testValue)
	if testValue != m.Sorting() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuSlugID(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := "testValue"
	m.SetSlugID(&testValue)
	if &testValue != m.SlugID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuExternalLink(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := "testValue"
	m.SetExternalLink(&testValue)
	if &testValue != m.ExternalLink() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuInternalLink(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := "testValue"
	m.SetInternalLink(&testValue)
	if &testValue != m.InternalLink() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestMenuParentID(t *testing.T) {
	m := menu.NewMenuEntity()
	testValue := "testValue"
	m.SetParentID(&testValue)
	if &testValue != m.ParentID() {
		t.Fatal("Getter did not return the Set value")
	}
}
