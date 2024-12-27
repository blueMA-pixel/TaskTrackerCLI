package main

import (
	"fmt"
	"time"
)

type AddCommand struct {
	description string
}

func (addCommand *AddCommand) initialize(commandLineArguments []string) error {
	if len(commandLineArguments) == 1 || len(commandLineArguments) > 2 {
		return fmt.Errorf("add command requires a description only")
	}

	addCommand.description = commandLineArguments[1]

	return nil
}

func (addCommand *AddCommand) execute(tasks *Tasks) error {
	var newTask Task
	newTask.Description = addCommand.description
	newTask.CreationTime = time.Now()
	newTask.UpdateTime = time.Now()
	newTask.Status = ToDo
	return tasks.addTask(newTask)
}
