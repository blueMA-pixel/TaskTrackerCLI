package main

import (
	"fmt"
	"strconv"
)

type DeleteCommand struct {
	taskId int
}

func (d *DeleteCommand) CheckCommandLineArguments(commandLineArguments []string) {
	if len(commandLineArguments) != 2 {
		return
	}

	d.taskId, _ = strconv.Atoi(commandLineArguments[1])
}

func findSliceIndex(tasks []Task, id int) int {
	for index, task := range tasks {
		if task.Id == id {
			return index
		}
	}
	return -1
}

func (d *DeleteCommand) executeCommand(tasks *[]Task) {
	var index int = findSliceIndex(*tasks, d.taskId)

	if index < 0 || index >= len(*tasks) {
		return
	}

	fmt.Println(index)

	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)

	fmt.Println(tasks)
}
