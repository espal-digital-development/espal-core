// Code generated by espal-store-synthesizer. DO NOT EDIT.
package blogpost_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/blogpost"
)

func TestCommentTable(t *testing.T) {
	c := blogpost.NewCommentEntity()
	if c.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestCommentTableAlias(t *testing.T) {
	c := blogpost.NewCommentEntity()
	if c.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestCommentIsUpdated(t *testing.T) {
	c := blogpost.NewCommentEntity()
	c.IsUpdated()
}

func TestCommentID(t *testing.T) {
	c := blogpost.NewCommentEntity()
	c.ID()
}

func TestCommentCreatedByID(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := "testValue"
	c.SetCreatedByID(testValue)
	if testValue != c.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestCommentUpdatedByID(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := "testValue"
	c.SetUpdatedByID(&testValue)
	if &testValue != c.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestCommentCreatedAt(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := time.Now()
	c.SetCreatedAt(testValue)
	if testValue != c.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestCommentUpdatedAt(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := time.Now()
	c.SetUpdatedAt(&testValue)
	if &testValue != c.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestCommentCreatedByFirstName(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := "testValue"
	c.SetCreatedByFirstName(&testValue)
	if &testValue != c.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestCommentCreatedBySurname(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := "testValue"
	c.SetCreatedBySurname(&testValue)
	if &testValue != c.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestCommentUpdatedByFirstName(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := "testValue"
	c.SetUpdatedByFirstName(&testValue)
	if &testValue != c.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestCommentUpdatedBySurname(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := "testValue"
	c.SetUpdatedBySurname(&testValue)
	if &testValue != c.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestCommentBlogPostID(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := "testValue"
	c.SetBlogPostID(testValue)
	if testValue != c.BlogPostID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestCommentTitle(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := "testValue"
	c.SetTitle(&testValue)
	if &testValue != c.Title() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestCommentMessage(t *testing.T) {
	c := blogpost.NewCommentEntity()
	testValue := "testValue"
	c.SetMessage(testValue)
	if testValue != c.Message() {
		t.Fatal("Getter did not return the Set value")
	}
}
