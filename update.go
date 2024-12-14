package main

import (
	"strconv"
	"time"
)

type UpdateCommand struct {
	id             int
	newDescription string
}

func (u *UpdateCommand) CheckCommandLineArguments(commandLineArguments []string) {
	if len(commandLineArguments) < 2 || len(commandLineArguments) > 3 {
		return
	}

	u.id, _ = strconv.Atoi(commandLineArguments[1])
	u.newDescription = commandLineArguments[2]
}

func (u *UpdateCommand) executeCommand(tasks *[]Task) {
	var taskIndex = findSliceIndex(*tasks, u.id)

	if taskIndex < 0 || taskIndex > len(*tasks)-1 {
		return
	}

	(*tasks)[taskIndex].Description = u.newDescription
	(*tasks)[taskIndex].UpdateTime = time.Now()

}
