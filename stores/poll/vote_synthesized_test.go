// Code generated by espal-store-synthesizer. DO NOT EDIT.
package poll_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/poll"
)

func TestVoteTable(t *testing.T) {
	v := poll.NewVoteEntity()
	if v.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestVoteTableAlias(t *testing.T) {
	v := poll.NewVoteEntity()
	if v.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestVoteIsUpdated(t *testing.T) {
	v := poll.NewVoteEntity()
	v.IsUpdated()
}

func TestVoteID(t *testing.T) {
	v := poll.NewVoteEntity()
	v.ID()
}

func TestVoteCreatedByID(t *testing.T) {
	v := poll.NewVoteEntity()
	testValue := "testValue"
	v.SetCreatedByID(testValue)
	if testValue != v.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVoteUpdatedByID(t *testing.T) {
	v := poll.NewVoteEntity()
	testValue := "testValue"
	v.SetUpdatedByID(&testValue)
	if &testValue != v.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVoteCreatedAt(t *testing.T) {
	v := poll.NewVoteEntity()
	testValue := time.Now()
	v.SetCreatedAt(testValue)
	if testValue != v.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVoteUpdatedAt(t *testing.T) {
	v := poll.NewVoteEntity()
	testValue := time.Now()
	v.SetUpdatedAt(&testValue)
	if &testValue != v.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVoteCreatedByFirstName(t *testing.T) {
	v := poll.NewVoteEntity()
	testValue := "testValue"
	v.SetCreatedByFirstName(&testValue)
	if &testValue != v.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVoteCreatedBySurname(t *testing.T) {
	v := poll.NewVoteEntity()
	testValue := "testValue"
	v.SetCreatedBySurname(&testValue)
	if &testValue != v.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVoteUpdatedByFirstName(t *testing.T) {
	v := poll.NewVoteEntity()
	testValue := "testValue"
	v.SetUpdatedByFirstName(&testValue)
	if &testValue != v.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVoteUpdatedBySurname(t *testing.T) {
	v := poll.NewVoteEntity()
	testValue := "testValue"
	v.SetUpdatedBySurname(&testValue)
	if &testValue != v.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestVotePollOptionID(t *testing.T) {
	v := poll.NewVoteEntity()
	testValue := "testValue"
	v.SetPollOptionID(testValue)
	if testValue != v.PollOptionID() {
		t.Fatal("Getter did not return the Set value")
	}
}
