package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"
)

const nonAlphabeticNumericalCharactersRE = "[^a-zA-Z0-9]"

type mockWriter struct {
	buffer *bytes.Buffer
}

func (cw *mockWriter) Write(p []byte) (n int, err error) {
	bytesWritten, err := cw.buffer.Write(p)

	if bytesWritten != len(p) {
		return bytesWritten, fmt.Errorf("capture writer has not written all bytes")
	}

	return bytesWritten, err

}

type MockTaskDataHandler struct {
	tasks *Tasks
}

func (handler MockTaskDataHandler) readTasks(tasks *Tasks) error {
	*tasks = *handler.tasks
	return nil
}

func (loader MockTaskDataHandler) writeTasks(tasks *Tasks) error {
	return nil
}

var (
	tasks Tasks
)

func SetupThreeTasks(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	mockTime, err := time.Parse(layout, "2024-12-28 16:24:40")

	if err != nil {
		t.Fatalf("time could not be parsed, %v", err)
	}

	err = tasks.addTask(Task{
		Description:  "Task 1 description",
		Status:       ToDo,
		CreationTime: mockTime,
		UpdateTime:   mockTime,
	})
	if err != nil {
		t.Fatalf("Error adding task 1: %v", err)
	}

	err = tasks.addTask(Task{
		Description:  "Task 2 description",
		Status:       InProgress,
		CreationTime: mockTime,
		UpdateTime:   mockTime,
	})
	if err != nil {
		t.Fatalf("Error adding task 2: %v", err)
	}

	err = tasks.addTask(Task{
		Description:  "Task 3 description",
		Status:       Done,
		CreationTime: mockTime,
		UpdateTime:   mockTime,
	})
	if err != nil {
		t.Fatalf("Error adding task 3: %v", err)
	}

	t.Cleanup(func() {
		tasks = Tasks{}
	})
}

func TestListCommand(t *testing.T) {

	SetupThreeTasks(t)

	var app Application = Application{}
	var cw mockWriter = mockWriter{buffer: &bytes.Buffer{}}
	var dataHandler MockTaskDataHandler = MockTaskDataHandler{tasks: &tasks}
	commandLineArguments := []string{"list"}
	app.initialize(commandLineArguments, &cw, dataHandler)
	app.run()

	expected :=
		`ID: 0
Description: Task 1 description
Status: to do
Created: 2024-12-28 16:24:40
Updated: 2024-12-28 16:24:40
ID: 1
Description: Task 2 description
Status: in progress
Created: 2024-12-28 16:24:40
Updated: 2024-12-28 16:24:40
ID: 2
Description: Task 3 description
Status: done
Created: 2024-12-28 16:24:40
Updated: 2024-12-28 16:24:40`

	re := regexp.MustCompile(nonAlphabeticNumericalCharactersRE)
	expectedTrimmed := re.ReplaceAllString(expected, "")
	actualTrimmed := re.ReplaceAllString(cw.buffer.String(), "")

	if expectedTrimmed != actualTrimmed {
		t.Errorf("expected\n%s\n\nreached\n%s", expected, cw.buffer.String())
	}
}

func TestAddCommand(t *testing.T) {
	SetupThreeTasks(t)
	var app Application = Application{}
	var dataHandler MockTaskDataHandler = MockTaskDataHandler{tasks: &tasks}
	commandLineArguments := []string{"add", "this is a new task"}
	app.initialize(commandLineArguments, os.Stdout, dataHandler)
	app.run()

	actualAddedTask := app.tasks[len(app.tasks)-1]

	expectedAddedTask := Task{
		Id:           3,
		Description:  "this is a new task",
		Status:       ToDo,
		CreationTime: time.Now(),
		UpdateTime:   time.Now(),
	}

	if expectedAddedTask != actualAddedTask {
		t.Errorf("expected\n%s\n\nreached\n%s", expectedAddedTask, actualAddedTask)
	}
}

func TestDeleteCommand(t *testing.T) {
	SetupThreeTasks(t)

	var app Application = Application{}
	var dataHandler MockTaskDataHandler = MockTaskDataHandler{tasks: &tasks}
	commandLineArguments := []string{"delete", "1"}
	app.initialize(commandLineArguments, os.Stdout, dataHandler)
	app.run()

	if len(app.tasks) != 2 {
		t.Errorf("Expected a task to be delete")
	}
}

func TestUpdateCommand(t *testing.T) {
	SetupThreeTasks(t)

	var app Application = Application{}
	var dataHandler MockTaskDataHandler = MockTaskDataHandler{tasks: &tasks}
	commandLineArguments := []string{"update", "1", "this is an updated task"}
	app.initialize(commandLineArguments, os.Stdout, dataHandler)
	app.tasks = tasks
	app.command.execute(&app.tasks)

	actualUpdatedTask := app.tasks[1]

	if actualUpdatedTask.Description != "this is an updated task" || actualUpdatedTask.UpdateTime != time.Now() {
		t.Errorf("The task was not updated\n%s", actualUpdatedTask)
	}
}

func TestMarkInProgressCommand(t *testing.T) {
	SetupThreeTasks(t)

	var app Application = Application{}
	var dataHandler MockTaskDataHandler = MockTaskDataHandler{tasks: &tasks}
	commandLineArguments := []string{"mark-in-progress", "1"}
	app.initialize(commandLineArguments, os.Stdout, dataHandler)
	app.tasks = tasks
	app.command.execute(&app.tasks)

	actualUpdatedTask := app.tasks[1]

	if actualUpdatedTask.Status != InProgress || actualUpdatedTask.UpdateTime != time.Now() {
		t.Errorf("The task was not marked in progress\n%s", actualUpdatedTask)
	}
}

func TestMarkAsDoneCommand(t *testing.T) {
	SetupThreeTasks(t)

	var app Application = Application{}
	var dataHandler MockTaskDataHandler = MockTaskDataHandler{tasks: &tasks}
	commandLineArguments := []string{"mark-as-done", "1"}
	app.initialize(commandLineArguments, os.Stdout, dataHandler)
	app.tasks = tasks
	app.command.execute(&app.tasks)

	actualUpdatedTask := app.tasks[1]

	if actualUpdatedTask.Status != Done || actualUpdatedTask.UpdateTime != time.Now() {
		t.Errorf("The task was not marked as done\n%s", actualUpdatedTask)
	}
}
