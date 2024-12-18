package main

import (
	"strconv"
	"time"
)

type UpdateCommand struct {
	id             int
	newDescription string
}

func (u *UpdateCommand) initialize(commandLineArguments []string) error {
	if len(commandLineArguments) < 2 || len(commandLineArguments) > 3 {
		return nil
	}

	u.id, _ = strconv.Atoi(commandLineArguments[1])
	u.newDescription = commandLineArguments[2]

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
