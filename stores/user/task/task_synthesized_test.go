// Code generated by espal-store-synthesizer. DO NOT EDIT.
package task_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/user/task"
)

func TestTaskTable(t *testing.T) {
	tt := task.NewTaskEntity()
	if tt.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestTaskTableAlias(t *testing.T) {
	tt := task.NewTaskEntity()
	if tt.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestTaskIsUpdated(t *testing.T) {
	tt := task.NewTaskEntity()
	tt.IsUpdated()
}

func TestTaskID(t *testing.T) {
	tt := task.NewTaskEntity()
	tt.ID()
}

func TestTaskCreatedByID(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := "testValue"
	tt.SetCreatedByID(testValue)
	if testValue != tt.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskUpdatedByID(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := "testValue"
	tt.SetUpdatedByID(&testValue)
	if &testValue != tt.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskCreatedAt(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := time.Now()
	tt.SetCreatedAt(testValue)
	if testValue != tt.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskUpdatedAt(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := time.Now()
	tt.SetUpdatedAt(&testValue)
	if &testValue != tt.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskCreatedByFirstName(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := "testValue"
	tt.SetCreatedByFirstName(&testValue)
	if &testValue != tt.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskCreatedBySurname(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := "testValue"
	tt.SetCreatedBySurname(&testValue)
	if &testValue != tt.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskUpdatedByFirstName(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := "testValue"
	tt.SetUpdatedByFirstName(&testValue)
	if &testValue != tt.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskUpdatedBySurname(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := "testValue"
	tt.SetUpdatedBySurname(&testValue)
	if &testValue != tt.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskIssuedByID(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := "testValue"
	tt.SetIssuedByID(testValue)
	if testValue != tt.IssuedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskAssignedToID(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := "testValue"
	tt.SetAssignedToID(&testValue)
	if &testValue != tt.AssignedToID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskDescription(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := "testValue"
	tt.SetDescription(testValue)
	if testValue != tt.Description() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskCompletedNotes(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := "testValue"
	tt.SetCompletedNotes(&testValue)
	if &testValue != tt.CompletedNotes() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTaskCompletedAt(t *testing.T) {
	tt := task.NewTaskEntity()
	testValue := time.Now()
	tt.SetCompletedAt(&testValue)
	if &testValue != tt.CompletedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}
