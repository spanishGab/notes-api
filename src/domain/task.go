package domain

import (
	"fmt"
	"regexp"
)

type Task struct {
	title       string
	description string
	status      string
	isLocked    bool
	color       string
}

func New(title string, description string, status string, color string) (*Task, error) {
	if title == "" && description == "" {
		return nil, fmt.Errorf("a title or a description is required")
	}

	var isStatusValid = isStatusCreatable(status)

	if !isStatusValid {
		return nil, fmt.Errorf("invalid status")
	}

	var defaultColor string = color
	if color == "" {
		defaultColor = "#ffffff"
	}

	var task *Task = &Task{
		title: title,
		description: description,
		status: status,
		isLocked: false,
		color: defaultColor,
	}

	return task, nil
}

func (task *Task) SetTitle(title string) (error) {
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	task.title = title
	return nil
}

func (task *Task) SetDescription(description string) (error) {
	if description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	task.description = description
	return nil
}

func (task *Task) Lock() {
	if !task.isLocked {
		task.isLocked = true
	}
	// todo: error
}

func (task *Task) Unlock() {
	if task.isLocked {
		task.isLocked = false
	}
	// todo: error
}

func (task *Task) SetColor(color string) (error) {
	match, err := regexp.MatchString("^#[0-9a-fA-F]{6}$", color)

	if err != nil {
		return fmt.Errorf("error matching regex: %v", err)
	}

	if !match {
		return fmt.Errorf("color must be a valid 6-character hexadecimal string starting with #")
	}

	task.color = color
	return nil
}

func (task *Task) SetStatus(status string) error {
	if task.canChangeTodoStatus(status) {
		task.status = status
		return nil
	}

	if task.canChangeInProgressStatus(status) {
		task.status = status
		return nil
	}

	if task.canChangeDoneStatus(status) {
		task.status = status
		return nil
	}

	if !task.canChangeCanceledStatus() {
		return fmt.Errorf("can not allowed to change status from %s to %s", task.status, status)
	}

	return fmt.Errorf("can not allowed to change status from %s to %s", task.status, status)
}

func (task *Task) canChangeTodoStatus(status string) bool {
	if task.status == "TODO" {
		if status == "IN_PROGRESS" || status == "CANCELED" {
			return true
		}
	}
	return false
}

func (task *Task) canChangeInProgressStatus(status string) bool {
	if task.status == "IN_PROGRESS" {
		if status == "TODO" || status == "DONE" || status == "CANCELED" {
			return true
		}
	}
	return false
}

func (task *Task) canChangeDoneStatus(status string) (bool) {
	if (task.status == "DONE") {
		if (status == "TODO" || status == "IN_PROGRESS") {
			return true
		}
	}
	return false
}

func (task *Task) canChangeCanceledStatus() bool {
	return task.status != "CANCELED"
}

func isStatusCreatable(status string) (bool) {
	if status == "TODO" || status == "IN_PROGRESS" {
		return true
	}
	return false
}