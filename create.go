package main

import (
	"time"
)

type AddCommand struct {
	description string
}

func (addCommand *AddCommand) CheckCommandLineArguments(commandLineArguments []string) {
	if len(commandLineArguments) == 1 || len(commandLineArguments) > 2 {
		return
	}

	addCommand.description = commandLineArguments[1]
}

func generateTaskID(tasks []Task) int {

	var id int = 0
	for _, task := range tasks {
		if task.Id >= id {
			id = task.Id + 1
		}
	}
	return id
}

func (addCommand AddCommand) executeCommand(tasks *[]Task) {
	var newTask Task
	newTask.Id = generateTaskID(*tasks)
	newTask.Description = addCommand.description
	newTask.CreationTime = time.Now()
	newTask.UpdateTime = time.Now()
	newTask.Status = ToDo
	*tasks = append(*tasks, newTask)
}
