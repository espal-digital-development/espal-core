// Code generated by espal-store-synthesizer. DO NOT EDIT.
package group_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/tax/group"
)

func TestGroupTable(t *testing.T) {
	group := group.NewGroupEntity()
	if group.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestGroupTableAlias(t *testing.T) {
	group := group.NewGroupEntity()
	if group.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestGroupIsUpdated(t *testing.T) {
	group := group.NewGroupEntity()
	group.IsUpdated()
}

func TestGroupID(t *testing.T) {
	group := group.NewGroupEntity()
	group.ID()
}

func TestGroupCreatedByID(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := "testValue"
	group.SetCreatedByID(testValue)
	if testValue != group.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupUpdatedByID(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := "testValue"
	group.SetUpdatedByID(&testValue)
	if &testValue != group.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupCreatedAt(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := time.Now()
	group.SetCreatedAt(testValue)
	if testValue != group.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupUpdatedAt(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := time.Now()
	group.SetUpdatedAt(&testValue)
	if &testValue != group.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupCreatedByFirstName(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := "testValue"
	group.SetCreatedByFirstName(&testValue)
	if &testValue != group.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupCreatedBySurname(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := "testValue"
	group.SetCreatedBySurname(&testValue)
	if &testValue != group.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupUpdatedByFirstName(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := "testValue"
	group.SetUpdatedByFirstName(&testValue)
	if &testValue != group.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupUpdatedBySurname(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := "testValue"
	group.SetUpdatedBySurname(&testValue)
	if &testValue != group.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupActive(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := true
	group.SetActive(testValue)
	if testValue != group.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupSorting(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := uint(1e9)
	group.SetSorting(testValue)
	if testValue != group.Sorting() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupCode(t *testing.T) {
	group := group.NewGroupEntity()
	testValue := "testValue"
	group.SetCode(testValue)
	if testValue != group.Code() {
		t.Fatal("Getter did not return the Set value")
	}
}
