// Code generated by espal-store-synthesizer. DO NOT EDIT.
package session_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/session"
)

func TestSessionTable(t *testing.T) {
	s := session.NewSessionEntity()
	if s.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestSessionTableAlias(t *testing.T) {
	s := session.NewSessionEntity()
	if s.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestSessionIsUpdated(t *testing.T) {
	s := session.NewSessionEntity()
	s.IsUpdated()
}

func TestSessionID(t *testing.T) {
	s := session.NewSessionEntity()
	s.ID()
}

func TestSessionCreatedByID(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := "testValue"
	s.SetCreatedByID(&testValue)
	if &testValue != s.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSessionUpdatedByID(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := "testValue"
	s.SetUpdatedByID(&testValue)
	if &testValue != s.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSessionCreatedAt(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := time.Now()
	s.SetCreatedAt(testValue)
	if testValue != s.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSessionUpdatedAt(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := time.Now()
	s.SetUpdatedAt(&testValue)
	if &testValue != s.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSessionCreatedByFirstName(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := "testValue"
	s.SetCreatedByFirstName(&testValue)
	if &testValue != s.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSessionCreatedBySurname(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := "testValue"
	s.SetCreatedBySurname(&testValue)
	if &testValue != s.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSessionUpdatedByFirstName(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := "testValue"
	s.SetUpdatedByFirstName(&testValue)
	if &testValue != s.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSessionUpdatedBySurname(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := "testValue"
	s.SetUpdatedBySurname(&testValue)
	if &testValue != s.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSessionTimeout(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := time.Second * 8
	s.SetTimeout(testValue)
	if testValue != s.Timeout() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSessionHash(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := "testValue"
	s.SetHash(testValue)
	if testValue != s.Hash() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSessionData(t *testing.T) {
	s := session.NewSessionEntity()
	testValue := []byte("testData")
	s.SetData(testValue)
	if !bytes.Equal(testValue, s.Data()) {
		t.Fatal("Getter did not return the Set value")
	}
}
