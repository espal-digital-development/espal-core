// Code generated by espal-store-synthesizer. DO NOT EDIT.
package site_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/site"
)

func TestUserTable(t *testing.T) {
	u := site.NewUserEntity()
	if u.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestUserTableAlias(t *testing.T) {
	u := site.NewUserEntity()
	if u.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestUserIsUpdated(t *testing.T) {
	u := site.NewUserEntity()
	u.IsUpdated()
}

func TestUserID(t *testing.T) {
	u := site.NewUserEntity()
	u.ID()
}

func TestUserCreatedByID(t *testing.T) {
	u := site.NewUserEntity()
	testValue := "testValue"
	u.SetCreatedByID(testValue)
	if testValue != u.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUserUpdatedByID(t *testing.T) {
	u := site.NewUserEntity()
	testValue := "testValue"
	u.SetUpdatedByID(&testValue)
	if &testValue != u.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUserCreatedAt(t *testing.T) {
	u := site.NewUserEntity()
	testValue := time.Now()
	u.SetCreatedAt(testValue)
	if testValue != u.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUserUpdatedAt(t *testing.T) {
	u := site.NewUserEntity()
	testValue := time.Now()
	u.SetUpdatedAt(&testValue)
	if &testValue != u.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUserCreatedByFirstName(t *testing.T) {
	u := site.NewUserEntity()
	testValue := "testValue"
	u.SetCreatedByFirstName(&testValue)
	if &testValue != u.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUserCreatedBySurname(t *testing.T) {
	u := site.NewUserEntity()
	testValue := "testValue"
	u.SetCreatedBySurname(&testValue)
	if &testValue != u.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUserUpdatedByFirstName(t *testing.T) {
	u := site.NewUserEntity()
	testValue := "testValue"
	u.SetUpdatedByFirstName(&testValue)
	if &testValue != u.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUserUpdatedBySurname(t *testing.T) {
	u := site.NewUserEntity()
	testValue := "testValue"
	u.SetUpdatedBySurname(&testValue)
	if &testValue != u.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUserSiteID(t *testing.T) {
	u := site.NewUserEntity()
	testValue := "testValue"
	u.SetSiteID(testValue)
	if testValue != u.SiteID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUserUserID(t *testing.T) {
	u := site.NewUserEntity()
	testValue := "testValue"
	u.SetUserID(testValue)
	if testValue != u.UserID() {
		t.Fatal("Getter did not return the Set value")
	}
}
