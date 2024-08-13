package controllers

import (
	"notes-api/src/domain"

	"fmt"

	"github.com/google/uuid"
)

var TasksDB = make(map[uuid.UUID]*domain.Task)

type NewTask struct {
	Title          string
	Description    string
	Color          string
}

type TaskController struct {}

func parseUUID(id string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("invalid id")
	}
	return parsedID, nil
}

func (controller *TaskController) CreateTask(task NewTask) (*domain.Task, error) {
	newTask, err := domain.New(task.Title, task.Description, "TODO", task.Color)
	if err != nil {
		return nil, err
	}

	id := uuid.New()
	TasksDB[id] = newTask

	return newTask, nil
}

func (controller *TaskController) GetTask(id string) (*domain.Task, error) {
	parsedID, parseIDErr := parseUUID(id)
	if parseIDErr != nil {
		return nil, parseIDErr
	}

	task := TasksDB[parsedID]
	if (task == nil) {
		return nil, fmt.Errorf("task not found")
	}
	return task, nil
}

func (controller *TaskController) UpdateTask(id string, title string, description string, status string, isLocked bool, color string) (*domain.Task, error) {
	parsedID, parseIDErr := parseUUID(id)
	if parseIDErr != nil {
		return nil, parseIDErr
	}

	task := TasksDB[parsedID]
	if task == nil {
		return nil, fmt.Errorf("task not found")
	}

	if title == "" && description == "" && status == "" && color == "" {
		return nil, fmt.Errorf("no data to update")
	}

	if title != "" {
		task.SetTitle(title)
	}

	if description != "" {
		task.SetDescription(description)
	}

	if status != "" {
		task.SetStatus(status)
	}

	if isLocked {
			task.Lock()
		} else {
			task.Unlock()
		}

	if color != "" {
		task.SetColor(color)
	}

	TasksDB[parsedID] = task
	return task, nil
}
