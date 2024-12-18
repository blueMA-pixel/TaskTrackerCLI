package main

import (
	"fmt"
	"time"
)

type Status int

const (
	ToDo Status = iota
	InProgress
	Done
)

type Task struct {
	Id           int       `json:"ID"`
	Description  string    `json:"Description"`
	Status       Status    `json:"Status"`
	CreationTime time.Time `json:"CreationTime"`
	UpdateTime   time.Time `json:"UpdateTime"`
}

func (t Task) String() string {
	var statusString string

	switch t.Status {
	case ToDo:
		statusString = "to do"
	case InProgress:
		statusString = "in progress"
	case Done:
		statusString = "done"
	}

	return fmt.Sprintf(
		"ID: %-10d\nDescription: %-30s\nStatus: %-15s\nCreated: %-20s\nUpdated: %-20s\n",
		t.Id,
		t.Description,
		statusString,
		t.CreationTime.Format("2006-01-02 15:04:05"),
		t.UpdateTime.Format("2006-01-02 15:04:05"),
	)
}

type Tasks []Task

func (tasks *Tasks) findTask(id int) (*Task, int, error) {

	for index, task := range *tasks {
		if task.Id == id {
			return &(*tasks)[index], index, nil
		}
	}

	return nil, -1, nil // Todo: add an error
}

func (tasks *Tasks) removeTask(id int) error {

	_, index, err := tasks.findTask(id)

	if err != nil {
		return err
	}

	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)

	return nil
}

func (tasks *Tasks) generateTaskId() int {

	var id int = 0
	for _, task := range *tasks {
		if task.Id >= id {
			id = task.Id + 1
		}
	}
	return id
}

func (tasks *Tasks) addTask(task Task) error {
	if task.Id != 0 {
		return nil
	}

	task.Id = tasks.generateTaskId()

	*tasks = append(*tasks, task)

	return nil
}
