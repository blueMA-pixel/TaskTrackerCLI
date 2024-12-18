package main

import (
	"strconv"
)

type DeleteCommand struct {
	taskId int
}

func (d *DeleteCommand) initialize(commandLineArguments []string) error {
	if len(commandLineArguments) != 2 {
		return nil
	}

	d.taskId, _ = strconv.Atoi(commandLineArguments[1])

	return nil
}

func (d *DeleteCommand) execute(tasks *Tasks) error {
	return tasks.removeTask(d.taskId)
}
