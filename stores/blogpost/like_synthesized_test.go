// Code generated by espal-store-synthesizer. DO NOT EDIT.
package blogpost_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/blogpost"
)

func TestLikeTable(t *testing.T) {
	l := blogpost.NewLikeEntity()
	if l.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestLikeTableAlias(t *testing.T) {
	l := blogpost.NewLikeEntity()
	if l.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestLikeIsUpdated(t *testing.T) {
	l := blogpost.NewLikeEntity()
	l.IsUpdated()
}

func TestLikeID(t *testing.T) {
	l := blogpost.NewLikeEntity()
	l.ID()
}

func TestLikeCreatedByID(t *testing.T) {
	l := blogpost.NewLikeEntity()
	testValue := "testValue"
	l.SetCreatedByID(testValue)
	if testValue != l.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLikeUpdatedByID(t *testing.T) {
	l := blogpost.NewLikeEntity()
	testValue := "testValue"
	l.SetUpdatedByID(&testValue)
	if &testValue != l.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLikeCreatedAt(t *testing.T) {
	l := blogpost.NewLikeEntity()
	testValue := time.Now()
	l.SetCreatedAt(testValue)
	if testValue != l.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLikeUpdatedAt(t *testing.T) {
	l := blogpost.NewLikeEntity()
	testValue := time.Now()
	l.SetUpdatedAt(&testValue)
	if &testValue != l.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLikeCreatedByFirstName(t *testing.T) {
	l := blogpost.NewLikeEntity()
	testValue := "testValue"
	l.SetCreatedByFirstName(&testValue)
	if &testValue != l.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLikeCreatedBySurname(t *testing.T) {
	l := blogpost.NewLikeEntity()
	testValue := "testValue"
	l.SetCreatedBySurname(&testValue)
	if &testValue != l.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLikeUpdatedByFirstName(t *testing.T) {
	l := blogpost.NewLikeEntity()
	testValue := "testValue"
	l.SetUpdatedByFirstName(&testValue)
	if &testValue != l.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLikeUpdatedBySurname(t *testing.T) {
	l := blogpost.NewLikeEntity()
	testValue := "testValue"
	l.SetUpdatedBySurname(&testValue)
	if &testValue != l.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLikeBlogPostID(t *testing.T) {
	l := blogpost.NewLikeEntity()
	testValue := "testValue"
	l.SetBlogPostID(testValue)
	if testValue != l.BlogPostID() {
		t.Fatal("Getter did not return the Set value")
	}
}
