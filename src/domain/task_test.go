package domain

import (
	"testing"
)

func TestSetTitle(t *testing.T) {
	var task *Task = &Task{
		title:       "Any task",
		description: "Any description",
		status:      "TODO",
		isLocked:    false,
		color:       "#000000",
	}

	expected := "First Task"
	task.SetTitle(expected)

	if task.title != expected {
		t.Errorf("expected '%s' but got '%s'", expected, task.title)
	}
}
