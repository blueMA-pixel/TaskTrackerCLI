package main

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type UpdateCommand struct {
	id             int
	newDescription string
	writer         io.Writer
}

func (u *UpdateCommand) initialize(commandLineArguments []string, writer io.Writer) error {
	if len(commandLineArguments) < 2 || len(commandLineArguments) > 3 {
		return fmt.Errorf("update command requires an id and a new description")
	}

	u.id, _ = strconv.Atoi(commandLineArguments[1])
	u.newDescription = commandLineArguments[2]

	u.writer = writer
	return nil
}

func (u *UpdateCommand) execute(tasks *Tasks) error {

	task, _, err := tasks.findTask(u.id)

	if err != nil {
		return err
	}

	(*task).Description = u.newDescription
	(*task).UpdateTime = time.Now()

	return nil
}
