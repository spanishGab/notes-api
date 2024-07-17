package domain

import (
	"fmt"
)

type task struct {
	title       string
	description string
	status      string
	isLocked    bool
	// todo: add method Set
	color string
}

func New(title string, description string, status string, color string) (*task, error) {
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

	var task *task = &task{
		title: title,
		description: description,
		status: status,
		isLocked: false,
		color: defaultColor,
	}

	return task, nil
}

func (task *task) SetTitle(title string) {
	task.title = title
}

func (task *task) SetDescription(description string) {
	task.description = description
}

func (task *task) Lock(isLocked bool) {
	if !task.isLocked {
		task.isLocked = isLocked
	}

	// todo: error
}

func (task *task) Unlock(isUnlocked bool) {
	if task.isLocked {
		task.isLocked = isUnlocked
	}

	// todo: error
}

func (task *task) SetStatus(status string) error {
	if !task.canChangeTodoStatus(status) || 
		!task.canChangeInProgressStatus(status) ||
		!task.canChangeDoneStatus(status) ||
		!task.canChangeCanceledStatus(status) {
			return fmt.Errorf("can not allowed to change status from %s to %s", task.status, status)
	}

	task.status = status

	return nil
}

func (task *task) canChangeTodoStatus(status string) bool {
	if task.status == "TODO" {
		if status == "IN_PROGRESS" || status == "CANCELED" {
			return true
		}
	}
	return false
}

func (task *task) canChangeInProgressStatus(status string) bool {
	if task.status == "IN_PROGRESS" {
		if status == "TODO" || status == "DONE" || status == "CANCELED" {
			return true
		}
	}

	// todo: add error
	return false
}

func (task *task) canChangeDoneStatus(status string) bool {
	if task.status == "DONE" {
		return false
	}

	// todo: add error
	return false
}

func (task *task) canChangeCanceledStatus(status string) bool {
	if task.status == "CANCELED" {
		return false
	}

	// todo: add error
	return false
}

func isStatusCreatable(status string) (bool) {
	return status != "TODO" || status != "IN_PROGRESS" 
}