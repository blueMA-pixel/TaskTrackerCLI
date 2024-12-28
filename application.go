package main

import (
	"fmt"
	"io"
)

const listCommand string = "list"
const addCommand string = "add"
const deleteCommand string = "delete"
const updateCommand string = "update"
const markasDoneCommand string = "mark-as-done"
const markinProgressCommand string = "mark-in-progress"

type Application struct {
	tasks            Tasks
	command          ICommand
	writer           io.Writer
	tasksDataHandler ITaskDataHandler
}

func (app *Application) initialize(commandLineArguments []string, writer io.Writer, tasksDataHandler ITaskDataHandler) error {

	if len(commandLineArguments) < 1 {
		return fmt.Errorf("include a command")
	}

	if tasksDataHandler == nil {
		return fmt.Errorf("task data handler is not initialized")
	}

	app.tasksDataHandler = tasksDataHandler

	switch commandLineArguments[0] {
	case listCommand:
		app.command = &ListCommand{}
	case addCommand:
		app.command = &AddCommand{}
	case deleteCommand:
		app.command = &DeleteCommand{}
	case updateCommand:
		app.command = &UpdateCommand{}
	case markasDoneCommand, markinProgressCommand:
		app.command = &MarkAsCommand{}
	}

	return app.command.initialize(commandLineArguments, writer)
}

func (app *Application) run() error {
	err := app.tasksDataHandler.readTasks(&app.tasks)

	if err != nil {
		fmt.Fprintln(app.writer, err)
		return err
	}

	err = app.command.execute(&app.tasks)

	if err != nil {
		fmt.Fprintln(app.writer, err)
		return err
	}

	err = app.tasksDataHandler.writeTasks(&app.tasks)

	if err != nil {
		fmt.Fprintln(app.writer, err)
		return err
	}

	return nil
}
