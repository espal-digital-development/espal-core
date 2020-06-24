// Code generated by espal-store-synthesizer. DO NOT EDIT.
package poll_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/poll"
)

func TestOptionTable(t *testing.T) {
	option := poll.NewOptionEntity()
	if option.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestOptionTableAlias(t *testing.T) {
	option := poll.NewOptionEntity()
	if option.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestOptionIsUpdated(t *testing.T) {
	option := poll.NewOptionEntity()
	option.IsUpdated()
}

func TestOptionID(t *testing.T) {
	option := poll.NewOptionEntity()
	option.ID()
}

func TestOptionCreatedByID(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := "testValue"
	option.SetCreatedByID(testValue)
	if testValue != option.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionUpdatedByID(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := "testValue"
	option.SetUpdatedByID(&testValue)
	if &testValue != option.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionCreatedAt(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := time.Now()
	option.SetCreatedAt(testValue)
	if testValue != option.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionUpdatedAt(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := time.Now()
	option.SetUpdatedAt(&testValue)
	if &testValue != option.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionCreatedByFirstName(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := "testValue"
	option.SetCreatedByFirstName(&testValue)
	if &testValue != option.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionCreatedBySurname(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := "testValue"
	option.SetCreatedBySurname(&testValue)
	if &testValue != option.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionUpdatedByFirstName(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := "testValue"
	option.SetUpdatedByFirstName(&testValue)
	if &testValue != option.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionUpdatedBySurname(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := "testValue"
	option.SetUpdatedBySurname(&testValue)
	if &testValue != option.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionActive(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := true
	option.SetActive(testValue)
	if testValue != option.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionSorting(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := uint(1e9)
	option.SetSorting(testValue)
	if testValue != option.Sorting() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionPollID(t *testing.T) {
	option := poll.NewOptionEntity()
	testValue := "testValue"
	option.SetPollID(testValue)
	if testValue != option.PollID() {
		t.Fatal("Getter did not return the Set value")
	}
}
