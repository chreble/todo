package task

import (
	"testing"
)

var tasks Tasks

func init() {
	tasks = Tasks{}
}

func TestCreate(t *testing.T) {
	text := "Foo"
	task := tasks.Create(Task{Text: text, Completed: false})

	if task.Text != text {
		t.Errorf("expected text %q, got %q", text, task.Text)
	}
	if task.Completed {
		t.Errorf("a new instance should not return as completed")
	}
}
