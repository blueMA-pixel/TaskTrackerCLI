package main

import "fmt"

type ListCommand struct {
	filterStatus Status
}

func (l *ListCommand) initialize(commandLineArguments []string) error {

	if len(commandLineArguments) == 1 {
		return fmt.Errorf("list Command Could not be initialized")
	}

	switch commandLineArguments[1] {
	case "done":
		l.filterStatus = Done
	case "to-do":
		l.filterStatus = ToDo
	case "in-progress":
		l.filterStatus = InProgress
	default:
		return fmt.Errorf("no such status filter")
	}

	return nil
}

func (l *ListCommand) execute(tasks *Tasks) error {

	for _, task := range *tasks {
		if l.filterStatus != Invalid && task.Status == l.filterStatus {
			fmt.Println(task)
		}
	}

	return nil
}
