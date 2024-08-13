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


func (controller *TaskController) CreateTask(task NewTask) (*domain.Task, error) {
	newTask, err := domain.New(task.Title, task.Description, "TODO", task.Color)
	if err != nil {
		return nil, err
	}

	id := uuid.New()
	TasksDB[id] = newTask

	return newTask, nil
}
