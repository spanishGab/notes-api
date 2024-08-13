package main

import (
	"fmt"
	"notes-api/src/controllers"
)

/*
	run test:
		- go test
		- go test -cover ou go test -coverprofile=coverage.out
		- go tool cover -html=coverage.out

*/

func main() {
	fmt.Println(controllers.TasksDB)

	var controller = &controllers.TaskController{}

	var task = controllers.NewTask{
		Title: "New task", 
		Description: "My description", 
		Color: "#000000",
	}
	controller.CreateTask(task)
	fmt.Println("all tasks", controllers.TasksDB)
}
