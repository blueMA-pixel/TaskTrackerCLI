package main

import (
	"fmt"
	"strconv"
	"time"
)

type MarkAsCommand struct {
	id        int
	newStatus Status
}

func (m *MarkAsCommand) initialize(commandLineArguments []string) error {
	if len(commandLineArguments) != 2 {
		return fmt.Errorf("mark command requires an id only")
	}

	switch commandLineArguments[0] {
	case "mark-as-done":
		m.newStatus = Done
	case "mark-in-progress":
		m.newStatus = InProgress
	default:
		return fmt.Errorf("no such command")
	}

	m.id, _ = strconv.Atoi(commandLineArguments[1])

	return nil
}

func (m *MarkAsCommand) execute(tasks *Tasks) error {

	task, _, err := tasks.findTask(m.id)

	if err != nil {
		return err
	}

	task.Status = m.newStatus
	task.UpdateTime = time.Now()

	return err
}
