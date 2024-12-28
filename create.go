package main

import (
	"fmt"
	"io"
	"time"
)

type AddCommand struct {
	description string
	writer      io.Writer
}

func (a *AddCommand) initialize(commandLineArguments []string, writer io.Writer) error {
	if len(commandLineArguments) == 1 || len(commandLineArguments) > 2 {
		return fmt.Errorf("add command requires a description only")
	}

	a.description = commandLineArguments[1]

	a.writer = writer
	return nil
}

func (a *AddCommand) execute(tasks *Tasks) error {
	var newTask Task
	newTask.Description = a.description
	newTask.CreationTime = time.Now()
	newTask.UpdateTime = time.Now()
	newTask.Status = ToDo
	return tasks.addTask(newTask)
}
