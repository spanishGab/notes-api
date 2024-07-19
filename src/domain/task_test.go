package domain

import (
	"reflect"
	"testing"
)

func TestSetTitle(t *testing.T) {
	t.Run("", func(t *testing.T) {
		task, _ := New("First Task", "Any description", "TODO", "#000000")


		expectedTask := &Task{
			title:       "First Task",
			description: "Any description",
			status:      "TODO",
			isLocked:    false,
			color:       "#000000",
		}

		if !reflect.DeepEqual(task, expectedTask) {
			t.Errorf("expected '%v' but got '%v'", expectedTask, task)
		}
	})

	t.Run("should return an error if title and description are empty", func(t *testing.T) {
		_, err := New("", "", "TODO", "#000000")

		expected := "a title or a description is required"

		if err.Error() != expected {
			t.Errorf("expected '%s' but got '%v'", expected, err)
		}
	})

	t.Run("y", func(t *testing.T) {
		_, err := New("", "any description", "ANY_STATUS", "#000000")

		expected := "invalid status"
		if err == nil {
			t.Errorf("expected an error but got nil")
		} else if err.Error() != expected {
			t.Errorf("expected '%s' but got '%v'", expected, err)
		}
	})

	t.Run("should set a new task title", func(t *testing.T) {
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
	})
}
