package main

import (
	"fmt"
	"io"
	"strconv"
)

type DeleteCommand struct {
	taskId int
	writer io.Writer
}

func (d *DeleteCommand) initialize(commandLineArguments []string, writer io.Writer) error {
	if len(commandLineArguments) != 2 {
		return fmt.Errorf("delete command requires an id only")
	}

	d.taskId, _ = strconv.Atoi(commandLineArguments[1])
	d.writer = writer
	return nil
}

func (d *DeleteCommand) execute(tasks *Tasks) error {
	return tasks.removeTask(d.taskId)
}
