package main

import (
	"fmt"
	"strconv"
)

type DeleteCommand struct {
	taskId int
}

func (d *DeleteCommand) initialize(commandLineArguments []string) error {
	if len(commandLineArguments) != 2 {
		return fmt.Errorf("delete command requires an id only")
	}

	d.taskId, _ = strconv.Atoi(commandLineArguments[1])

	return nil
}

func (d *DeleteCommand) execute(tasks *Tasks) error {
	return tasks.removeTask(d.taskId)
}
