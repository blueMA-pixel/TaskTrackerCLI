package main

import "fmt"

type ListCommand struct {
	filterPresent bool
	filterStatus  Status
}

func (l *ListCommand) initialize(commandLineArguments []string) error {

	if len(commandLineArguments) == 1 {
		return nil // todo return error here
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

	return nil
}

func (l *ListCommand) execute(tasks *Tasks) error {

	for _, task := range *tasks {
		if !l.filterPresent || task.Status == l.filterStatus {
			fmt.Println(task)
		}
	}

	return nil
}
