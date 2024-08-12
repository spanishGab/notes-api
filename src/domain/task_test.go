package domain

import (
	"reflect"
	"testing"
)

func setupTask() *Task {
	return &Task{
		title:       "Any task",
		description: "Any description",
		status:      "TODO",
		isLocked:    false,
		color:       "#000000",
	}
}

func TestTask(t *testing.T) {
	t.Run("should create a new task", func(t *testing.T) {
		task, _ := New("Any task", "Any description", "TODO", "#000000")

		mockTask := setupTask()

		if !reflect.DeepEqual(task, mockTask) {
			t.Errorf("expected '%v' but got '%v'", mockTask, task)
		}
	})

	t.Run("should return an error if title and description are empty", func(t *testing.T) {
		_, err := New("", "", "TODO", "#000000")

		expected := "a title or a description is required"

		if err != nil && err.Error() != expected {
			t.Errorf("expected '%s' but got '%v'", expected, err)
		}
	})

	t.Run("should return an error if status is invalid", func(t *testing.T) {
		_, err := New("", "any description", "ANY_STATUS", "#000000")

		expected := "invalid status"

		if err != nil && err.Error() != expected {
			t.Errorf("expected '%s' but got '%v'", expected, err)
		}
	})

	t.Run("should set a color '#ffffff' if a color is not passed", func(t *testing.T) {
		task, _ := New("any task", "any description", "TODO", "")

		var expectedTask *Task = &Task{
			title:       "any task",
			description: "any description",
			status:      "TODO",
			isLocked:    false,
			color:       "#ffffff",
		}

		if !reflect.DeepEqual(task, expectedTask) {
			t.Errorf("expected '%v' but got '%v'", expectedTask, task)
		}
	})

	t.Run("should set a new task title", func(t *testing.T) {
		expected := "First Task"

		mockTask := setupTask()
		mockTask.SetTitle(expected)

		if mockTask.title != expected {
			t.Errorf("expected '%s' but got '%s'", expected, mockTask.title)
		}
	})

	t.Run("should set a new task description", func(t *testing.T) {
		expected := "first description"

		mockTask := setupTask()
		mockTask.SetDescription(expected)

		if mockTask.description != expected {
			t.Errorf("expected '%s' but got '%s'", expected, mockTask.description)
		}
	}) 

	t.Run("should set isLocked to true", func(t *testing.T) {
		expected := true;

		mockTask := setupTask()
		mockTask.Lock()

		if mockTask.isLocked != expected {
			t.Errorf("expected '%t' but got '%t'", expected, mockTask.isLocked)
		}
	}) 

	t.Run("should set isLocked to false", func(t *testing.T) {
		expected := false;

		mockTask := setupTask()
		mockTask.Lock()

		mockTask.Unlock()

		if mockTask.isLocked != expected {
			t.Errorf("expected '%t' but got '%t'", expected, mockTask.isLocked)
		}
	}) 

	t.Run("should set a color to '#c3c3c3'", func(t *testing.T) {
		expected := "#c3c3c3"

		mockTask := setupTask()
		mockTask.setColor(expected)

		if mockTask.color != expected {
			t.Errorf("expected '%s' but got '%s'", expected, mockTask.color)
		}
	})

	t.Run("should change status from 'TODO' to 'IN_PROGRESS'", func(t *testing.T) {
		expected := "IN_PROGRESS"

		mockTask := setupTask()
		mockTask.SetStatus(expected)

		if mockTask.status != expected {
			t.Errorf("expected '%s' but got '%s'", expected, mockTask.status)
		}
	})

	t.Run("should change status from 'TODO' to 'CANCELED'", func(t *testing.T) {
		expected := "CANCELED"

		mockTask := setupTask()
		mockTask.SetStatus(expected)

		if mockTask.status != expected {
			t.Errorf("expected '%s' but got '%s'", expected, mockTask.status)
		}
	})

	t.Run("should not change the 'TODO' status", func(t *testing.T) {
		newStatus := "DONE"
		expected := "can not allowed to change status from TODO to DONE"

		mockTask := setupTask()
		err := mockTask.SetStatus(newStatus)

		if err != nil && err.Error() != expected {
			t.Errorf("expected '%s' but got '%v'", expected, err)
		}
	})

	t.Run("should change status from 'IN_PROGRESS' to 'TODO'", func(t *testing.T) {
		expected := "TODO"

		var mockTask *Task = &Task{
			title:       "any task",
			description: "any description",
			status:      "IN_PROGRESS",
			isLocked:    false,
			color:       "#ffffff",
		}
		mockTask.SetStatus(expected)

		if mockTask.status != expected {
			t.Errorf("expected '%s' but got '%s'", expected, mockTask.status)
		}
	})

	t.Run("should change status from 'IN_PROGRESS' to 'DONE'", func(t *testing.T) {
		expected := "DONE"

		var mockTask *Task = &Task{
			title:       "any task",
			description: "any description",
			status:      "IN_PROGRESS",
			isLocked:    false,
			color:       "#ffffff",
		}
		mockTask.SetStatus(expected)

		if mockTask.status != expected {
			t.Errorf("expected '%s' but got '%s'", expected, mockTask.status)
		}
	})

	t.Run("should change status from 'IN_PROGRESS' to 'CANCELED'", func(t *testing.T) {
		expected := "CANCELED"

		var mockTask *Task = &Task{
			title:       "any task",
			description: "any description",
			status:      "IN_PROGRESS",
			isLocked:    false,
			color:       "#ffffff",
		}
		mockTask.SetStatus(expected)

		if mockTask.status != expected {
			t.Errorf("expected '%s' but got '%s'", expected, mockTask.status)
		}
	})

	t.Run("should not change the 'IN_PROGRESS' status", func(t *testing.T) {
		newStatus := "IN_PROGRESS"
		expected := "can not allowed to change status from IN_PROGRESS to IN_PROGRESS"

		var mockTask *Task = &Task{
			title:       "any task",
			description: "any description",
			status:      "IN_PROGRESS",
			isLocked:    false,
			color:       "#ffffff",
		}
		err := mockTask.SetStatus(newStatus)

		if err != nil && err.Error() != expected {
			t.Errorf("expected '%s' but got '%v'", expected, err)
		}
	})

	t.Run("should not change the 'DONE' status", func(t *testing.T) {
		newStatus := "IN_PROGRESS"
		expected := "can not allowed to change status from DONE to IN_PROGRESS"

		var mockTask *Task = &Task{
			title:       "any task",
			description: "any description",
			status:      "DONE",
			isLocked:    false,
			color:       "#ffffff",
		}
		err := mockTask.SetStatus(newStatus)

		if err != nil && err.Error() != expected {
			t.Errorf("expected '%s' but got '%v'", expected, err)
		}
	})

	t.Run("should not change the 'CANCELED' status", func(t *testing.T) {
		newStatus := "IN_PROGRESS"
		expected := "can not allowed to change status from CANCELED to IN_PROGRESS"

		var mockTask *Task = &Task{
			title:       "any task",
			description: "any description",
			status:      "CANCELED",
			isLocked:    false,
			color:       "#ffffff",
		}
		err := mockTask.SetStatus(newStatus)

		if err != nil && err.Error() != expected {
			t.Errorf("expected '%s' but got '%v'", expected, err)
		}
	})

	t.Run("should not change the 'ANY' status", func(t *testing.T) {
		newStatus := "ANY"
		expected := "can not allowed to change status from ANY to ANY"

		var mockTask *Task = &Task{
			title:       "any task",
			description: "any description",
			status:      "ANY",
			isLocked:    false,
			color:       "#ffffff",
		}
		err := mockTask.SetStatus(newStatus)

		if err != nil && err.Error() != expected {
			t.Errorf("expected '%s' but got '%v'", expected, err)
		}
	})
}
