package main

import (
	"strconv"
	"time"
)

type MarkAsCommand struct {
	id        int
	newStatus Status
}

func (m *MarkAsCommand) CheckCommandLineArguments(commandLineArguments []string) {
	if len(commandLineArguments) != 2 {
		return
	}

	switch commandLineArguments[0] {
	case "mark-as-done":
		m.newStatus = Done
	case "mark-in-progress":
		m.newStatus = InProgress
	default:
		return
	}

	m.id, _ = strconv.Atoi(commandLineArguments[1])

}

func (m *MarkAsCommand) executeCommand(tasks *[]Task) {
	var taskIndex int = findSliceIndex(*tasks, m.id)

	if taskIndex < 0 || taskIndex > len(*tasks)-1 {
		return
	}

	(*tasks)[taskIndex].Status = m.newStatus
	(*tasks)[taskIndex].UpdateTime = time.Now()
}
