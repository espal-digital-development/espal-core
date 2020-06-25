// Code generated by espal-store-synthesizer. DO NOT EDIT.
package group_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/tax/group"
)

func TestGroupTable(t *testing.T) {
	g := group.NewGroupEntity()
	if g.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestGroupTableAlias(t *testing.T) {
	g := group.NewGroupEntity()
	if g.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestGroupIsUpdated(t *testing.T) {
	g := group.NewGroupEntity()
	g.IsUpdated()
}

func TestGroupID(t *testing.T) {
	g := group.NewGroupEntity()
	g.ID()
}

func TestGroupCreatedByID(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := "testValue"
	g.SetCreatedByID(testValue)
	if testValue != g.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupUpdatedByID(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := "testValue"
	g.SetUpdatedByID(&testValue)
	if &testValue != g.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupCreatedAt(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := time.Now()
	g.SetCreatedAt(testValue)
	if testValue != g.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupUpdatedAt(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := time.Now()
	g.SetUpdatedAt(&testValue)
	if &testValue != g.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupCreatedByFirstName(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := "testValue"
	g.SetCreatedByFirstName(&testValue)
	if &testValue != g.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupCreatedBySurname(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := "testValue"
	g.SetCreatedBySurname(&testValue)
	if &testValue != g.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupUpdatedByFirstName(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := "testValue"
	g.SetUpdatedByFirstName(&testValue)
	if &testValue != g.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupUpdatedBySurname(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := "testValue"
	g.SetUpdatedBySurname(&testValue)
	if &testValue != g.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupActive(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := true
	g.SetActive(testValue)
	if testValue != g.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupSorting(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := uint(1e9)
	g.SetSorting(testValue)
	if testValue != g.Sorting() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGroupCode(t *testing.T) {
	g := group.NewGroupEntity()
	testValue := "testValue"
	g.SetCode(testValue)
	if testValue != g.Code() {
		t.Fatal("Getter did not return the Set value")
	}
}
