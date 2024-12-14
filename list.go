package main

import "fmt"

type ListCommand struct {
	filterPresent bool
	filterStatus  Status
}

func (l *ListCommand) CheckCommandLineArguments(commandLineArguments []string) {

	if len(commandLineArguments) == 1 {
		return
	}

	switch commandLineArguments[1] {
	case "done":
		l.filterStatus = Done
		l.filterPresent = true
	case "to-do":
		l.filterStatus = ToDo
		l.filterPresent = true
	case "in-progress":
		l.filterStatus = InProgress
		l.filterPresent = true
	default:
		fmt.Println("No such filter")
	}
}

func (l *ListCommand) executeCommand(tasks []Task) {

	for _, task := range tasks {
		if !l.filterPresent || task.Status == l.filterStatus {
			fmt.Println(task)
		}
	}
}
