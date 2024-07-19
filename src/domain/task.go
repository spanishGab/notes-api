package domain

import (
	"fmt"
)

type Task struct {
	title       string
	description string
	status      string
	isLocked    bool
	color string
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

func (task *Task) SetTitle(title string) {
	task.title = title
}

func (task *Task) SetDescription(description string) {
	task.description = description
}

func (task *Task) Lock(isLocked bool) {
	if !task.isLocked {
		task.isLocked = isLocked
	}

	// todo: error
}

func (task *Task) Unlock(isUnlocked bool) {
	if task.isLocked {
		task.isLocked = isUnlocked
	}

	// todo: error
}

func (task *Task) setColor(color string) {
	task.color = color
}

func (task *Task) SetStatus(status string) error {
	if !task.canChangeTodoStatus(status) || 
		!task.canChangeInProgressStatus(status) ||
		!task.canChangeDoneStatus(status) ||
		!task.canChangeCanceledStatus(status) {
			return fmt.Errorf("can not allowed to change status from %s to %s", task.status, status)
	}

	task.status = status

	return nil
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

	// todo: add error
	return false
}

func (task *Task) canChangeDoneStatus(status string) bool {
	if task.status == "DONE" {
		return false
	}

	// todo: add error
	return false
}

func (task *Task) canChangeCanceledStatus(status string) bool {
	if task.status == "CANCELED" {
		return false
	}

	// todo: add error
	return false
}

func isStatusCreatable(status string) (bool) {
	if status == "TODO" || status == "IN_PROGRESS" {
		return true
	}
	return false
}