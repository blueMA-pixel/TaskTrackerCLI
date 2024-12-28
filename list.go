package main

import (
	"fmt"
	"io"
)

type ListCommand struct {
	filterStatus Status
	writer       io.Writer
}

func (l *ListCommand) initialize(commandLineArguments []string, writer io.Writer) error {

	if len(commandLineArguments) < 1 {
		return fmt.Errorf("list Command Could not be initialized")
	}

	l.writer = writer

	if len(commandLineArguments) == 1 {
		l.filterStatus = Invalid
		return nil
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
		if l.filterStatus == Invalid || task.Status == l.filterStatus {
			fmt.Fprint(l.writer, task)
		}
	}

	return nil
}
